/*
Copyright (c) 2020 TriggerMesh Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package inmemory

import (
	"github.com/triggermesh/eventstore/pkg/eventstore/sharedmain"
)

// EnvAccessorCtor for configuration parameters
func EnvAccessorCtor() sharedmain.EnvConfigAccessor {
	return &envAccessor{}
}

type envAccessor struct {
	sharedmain.EnvConfig

	DefaultGlobalTTL   int32 `envconfig:"EVENTSTORE_DEFAULT_GLOBAL_TTL" default:"604800"`
	DefaultBridgeTTL   int32 `envconfig:"EVENTSTORE_DEFAULT_BRIDGE_TTL" default:"604800"`
	DefaultInstanceTTL int32 `envconfig:"EVENTSTORE_DEFAULT_INSTANCE_TTL" default:"20"`

	DefaultExpiredGCPeriod int32 `envconfig:"EVENTSTORE_DEFAULT_EXPIRED_GC_PERIOD" default:"30"`
}