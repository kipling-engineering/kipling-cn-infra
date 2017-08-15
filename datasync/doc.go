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

// The datasync package defines the interfaces for the abstraction of data
// transport between app plugins and backend data sources (such
// as data stores, message buses, or gRPC-connected clients).
//
// These events are processed asynchronously.
// The app plugin that watches data changes gives callback for each event
// (e.g. successful configuration or an error).
//
// See the examples under the dedicated examples package.
package datasync
