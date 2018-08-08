// Copyright (c) 2018 Cisco and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package kvscheduler

import (
	"context"
	"strings"
	"testing"
	. "github.com/onsi/gomega"

	"github.com/ligato/cn-infra/kvscheduler/test"
	. "github.com/ligato/cn-infra/kvscheduler/api"
	"time"
)

const (
	descriptor1Name = "descriptor1"

	prefixA = "/prefixA/"
)

func prefixSelector(prefix string) func(key string) bool {
	return func(key string) bool {
		return strings.HasPrefix(key, prefix)
	}
}

func TestEmptyResync(t *testing.T) {
	RegisterTestingT(t)

	// prepare KV Scheduler
	scheduler := NewPlugin()
	err := scheduler.Init()
	Expect(err).To(BeNil())

	// prepare mocks
	mockSB := test.NewMockSouthbound()
	descriptor1 := test.NewMockDescriptor(&test.MockDescriptorArgs{
		Name:            descriptor1Name,
		KeySelector:     prefixSelector(prefixA),
		NBKeyPrefixes:   []string{prefixA},
		WithMetadata:    true,
		DumpIsSupported: true,
	}, mockSB)

	// register descriptor with the scheduler
	scheduler.RegisterKVDescriptor(descriptor1)

	// get metadata map created for the descriptor
	metadataMap := scheduler.GetMetadataMap(descriptor1.GetName())
	_, withMetadataMap := metadataMap.(test.NameToInteger)
	Expect(withMetadataMap).To(BeTrue())

	// transaction history should be initially empty
	Expect(scheduler.getTransactionHistory(time.Now())).To(BeEmpty())

	// run transaction with empty resync
	startTime := time.Now()
	kvErrors, txnError := scheduler.StartNBTransaction().Resync([]KeyValueDataPair{}).Commit(context.Background())
	stopTime := time.Now()
	Expect(txnError).ShouldNot(HaveOccurred())
	Expect(kvErrors).To(BeEmpty())

	// check the state of SB
	Expect(mockSB.GetKeysWithInvalidData()).To(BeEmpty())
	Expect(mockSB.GetValues(nil)).To(BeEmpty())

	// check executed operations
	opHistory := mockSB.PopHistoryOfOps()
	Expect(opHistory).To(HaveLen(1))
	Expect(opHistory[0].OpType).To(Equal(test.Dump))
	Expect(opHistory[0].CorrelateDump).To(BeEmpty())

	// transaction consisted of zero operations
	txnHistory := scheduler.getTransactionHistory(time.Now())
	Expect(txnHistory).To(HaveLen(1))
	txnOp := txnHistory[0]
	Expect(txnOp.preRecord).To(BeFalse())
	Expect(txnOp.start.After(startTime)).To(BeTrue())
	Expect(txnOp.start.Before(txnOp.stop)).To(BeTrue())
	Expect(txnOp.stop.Before(stopTime)).To(BeTrue())
	Expect(txnOp.seqNum).To(BeEquivalentTo(0))
	Expect(txnOp.txnType).To(BeEquivalentTo(nbTransaction))
	Expect(txnOp.isResync).To(BeTrue())
	Expect(txnOp.values).To(BeEmpty())
	Expect(txnOp.preErrors).To(BeEmpty())
	Expect(txnOp.planned).To(BeEmpty())
	Expect(txnOp.executed).To(BeEmpty())

	// close scheduler
	err = scheduler.Close()
	Expect(err).To(BeNil())
}
