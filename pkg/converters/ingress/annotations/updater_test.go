/*
Copyright 2019 The HAProxy Ingress Controller Authors.

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

package annotations

import (
	"testing"

	ing_helper "github.com/jcmoraisjr/haproxy-ingress/pkg/converters/ingress/helper_test"
	"github.com/jcmoraisjr/haproxy-ingress/pkg/converters/ingress/types"
	"github.com/jcmoraisjr/haproxy-ingress/pkg/haproxy"
	hatypes "github.com/jcmoraisjr/haproxy-ingress/pkg/haproxy/types"
	types_helper "github.com/jcmoraisjr/haproxy-ingress/pkg/types/helper_test"
)

/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 *
 *  BUILDERS
 *
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */

type testConfig struct {
	t       *testing.T
	haproxy haproxy.Config
	cache   *ing_helper.CacheMock
	logger  *types_helper.LoggerMock
}

func setup(t *testing.T) *testConfig {
	logger := &types_helper.LoggerMock{T: t}
	return &testConfig{
		t:       t,
		haproxy: haproxy.CreateInstance(logger, haproxy.InstanceOptions{}).Config(),
		cache:   &ing_helper.CacheMock{},
		logger:  logger,
	}
}

func (c *testConfig) teardown() {
	c.logger.CompareLogging("")
}

func (c *testConfig) createUpdater() *updater {
	return &updater{
		haproxy: c.haproxy,
		cache:   c.cache,
		logger:  c.logger,
	}
}

func (c *testConfig) createBackendData(namespace, name string, ann *types.BackendAnnotations) *backData {
	ann.Source = types.Source{
		Namespace: namespace,
		Name:      name,
		Type:      "ingress",
	}
	return &backData{
		backend: &hatypes.Backend{},
		ann:     ann,
	}
}
