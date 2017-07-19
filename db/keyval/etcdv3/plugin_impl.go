// Copyright (c) 2017 Cisco and/or its affiliates.
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

package etcdv3

import (
	"github.com/coreos/etcd/clientv3"
	"github.com/ligato/cn-infra/core"
	"github.com/ligato/cn-infra/db/keyval/plugin"
	"github.com/ligato/cn-infra/logging"
)

// PluginID used in the Agent Core flavors
const PluginID core.PluginName = "ETCD"

// Plugin implements Plugin interface therefore can be loaded with other plugins
type Plugin struct {
	*plugin.Skeleton
}

// NewEtcdPlugin creates a new instance of Plugin. Configuration of etcd connection is loaded from file.
func NewEtcdPlugin(cfg *Config) *Plugin {

	skeleton := plugin.NewSkeleton(string(PluginID),
		func(log logging.Logger) (plugin.Connection, error) {
			etcdConfig, err := ConfigToClientv3(cfg)
			if err != nil {
				return nil, err
			}
			return NewEtcdConnectionWithBytes(*etcdConfig, log)
		},
	)
	return &Plugin{Skeleton: skeleton}
}

// NewEtcdPluginUsingClient creates a new instance of Plugin using given etcd client
func NewEtcdPluginUsingClient(client *clientv3.Client) *Plugin {
	skeleton := plugin.NewSkeleton(string("etcdv3"),
		func(log logging.Logger) (plugin.Connection, error) {
			return NewEtcdConnectionUsingClient(client, log)
		},
	)
	return &Plugin{Skeleton: skeleton}
}
