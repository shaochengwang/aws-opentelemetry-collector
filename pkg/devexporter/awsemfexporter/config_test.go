// Copyright 2019, OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package awsemfexporter

import (
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/config/configmodels"
)

func TestLoadConfig(t *testing.T) {
	factories, err := config.ExampleComponents()
	assert.Nil(t, err)

	factory := &Factory{}
	factories.Exporters[configmodels.Type(typeStr)] = factory
	cfg, err := config.LoadConfigFile(
		t, path.Join(".", "testdata", "config.yaml"), factories,
	)

	require.NoError(t, err)
	require.NotNil(t, cfg)

	assert.Equal(t, len(cfg.Exporters), 2)

	r0 := cfg.Exporters["awsemf"]
	assert.Equal(t, r0, factory.CreateDefaultConfig())

	r1 := cfg.Exporters["awsemf/1"].(*Config)
	assert.Equal(t, r1,
		&Config{
			ExporterSettings:      configmodels.ExporterSettings{TypeVal: configmodels.Type(typeStr), NameVal: "awsemf/1"},
			LogGroupName:          "",
			LogStreamName:         "",
			Endpoint:              "",
			RequestTimeoutSeconds: 30,
			MaxRetries:            5,
			NoVerifySSL:           false,
			ProxyAddress:          "",
			Region:                "us-west-2",
			LocalMode:             true,
			ResourceARN:           "arn:aws:ec2:us-east1:123456789:instance/i-293hiuhe0u",
			RoleARN:               "arn:aws:iam::123456789:role/monitoring-EKS-NodeInstanceRole",
			ForceFlushInterval:    30,
		})
}
