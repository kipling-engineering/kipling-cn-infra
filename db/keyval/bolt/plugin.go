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

package bolt

import (
	"os"
	"time"

	"go.ligato.io/cn-infra/v2/db/keyval"
	"go.ligato.io/cn-infra/v2/db/keyval/kvproto"
	"go.ligato.io/cn-infra/v2/infra"
)

// Config represents configuration for Bolt plugin.
type Config struct {
	DbPath          string        `json:"db-path"`
	FileMode        os.FileMode   `json:"file-mode"`
	LockTimeout     time.Duration `json:"lock-timeout"`
	FilterDupNotifs bool          `json:"filter-duplicate-notifications"`
}

// Plugin implements bolt plugin.
type Plugin struct {
	Deps
	*Config

	// Plugin is disabled if there is no config file available
	disabled bool
	// Bolt DB encapsulation
	boltClient *Client
	// Read/Write proto modelled data
	protoWrapper *kvproto.ProtoWrapper
}

// Deps lists dependencies of the Bolt plugin.
// If injected, Bolt plugin will use StatusCheck to signal the connection status.
type Deps struct {
	infra.PluginDeps
}

// Disabled returns *true* if the plugin is not in use due to missing configuration.
func (p *Plugin) Disabled() bool {
	return p.disabled
}

func (p *Plugin) OnStateChange(func(updown bool)) {
}

// OnConnect executes callback from datasync
func (p *Plugin) OnConnect(callback func() error) {
	if err := callback(); err != nil {
		p.Log.Error(err)
	}
}

// Init initializes Bolt plugin.
func (p *Plugin) Init() (err error) {
	if p.Config == nil {
		p.Config, err = p.getConfig()
		if err != nil || p.disabled {
			return err
		}
	}

	p.boltClient, err = NewClient(p.Config)
	if err != nil {
		p.Log.Errorf("Err: %v", err)
		return err
	}

	p.protoWrapper = kvproto.NewProtoWrapper(p.boltClient, &keyval.SerializerJSON{})

	p.Log.Infof("BoltDB started with: %v", p.Config.DbPath)

	return nil
}

// Close closes the Bolt client.
func (p *Plugin) Close() error {
	if p.boltClient != nil {
		return p.boltClient.Close()
	}
	return nil
}

// NewBroker creates new instance of prefixed broker that provides API with arguments of type proto.Message.
func (p *Plugin) NewBroker(keyPrefix string) keyval.ProtoBroker {
	return p.protoWrapper.NewBroker(keyPrefix)
}

// NewWatcher creates new instance of prefixed broker that provides API with arguments of type proto.Message.
func (p *Plugin) NewWatcher(keyPrefix string) keyval.ProtoWatcher {
	return p.protoWrapper.NewWatcher(keyPrefix)
}

func (p *Plugin) getConfig() (*Config, error) {
	var cfg Config
	found, err := p.Cfg.LoadValue(&cfg)
	if err != nil {
		return nil, err
	}
	if !found {
		p.Log.Info("Bolt config not found, skip loading this plugin")
		p.disabled = true
		return nil, nil
	}
	return &cfg, nil
}
