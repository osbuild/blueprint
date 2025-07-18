package blueprint

import (
	"encoding/json"
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/stretchr/testify/assert"
)

func TestJSON(t *testing.T) {
	tests := []struct {
		name  string
		json  string
		field FirstbootScriptCustomization
	}{
		{
			name: "custom",
			json: `{"type":"custom","name":"test","contents":"echo hello"}`,
			field: FirstbootScriptCustomization{
				union: json.RawMessage(`{"type":"custom","name":"test","contents":"echo hello"}`),
			},
		},
		{
			name: "satellite",
			json: `{"type":"satellite","name":"test","command":"echo hello"}`,
			field: FirstbootScriptCustomization{
				union: json.RawMessage(`{"type":"satellite","name":"test","command":"echo hello"}`),
			},
		},
		{
			name: "aap",
			json: `{"type":"aap","name":"test","job_template_url":"https://aap.example.com/api/v2/job_templates/9/callback/"}`,
			field: FirstbootScriptCustomization{
				union: json.RawMessage(`{"type":"aap","name":"test","job_template_url":"https://aap.example.com/api/v2/job_templates/9/callback/"}`),
			},
		},
		{
			name: "unknown",
			json: `{"type":"unknown","name":"test"}`,
			field: FirstbootScriptCustomization{
				union: json.RawMessage(`{"type":"unknown","name":"test"}`),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var actual FirstbootScriptCustomization

			err := json.Unmarshal([]byte(tt.json), &actual)
			assert.NoError(t, err)
			assert.Equal(t, tt.field, actual)

			b, err := json.Marshal(tt.field)
			assert.NoError(t, err)
			assert.Equal(t, tt.json, string(b))
		})
	}
}

func TestTOML(t *testing.T) {
	tests := []struct {
		name  string
		toml  string
		field FirstbootScriptCustomization
	}{
		{
			name: "custom",
			toml: `type = "custom"
name = "test"
contents = "echo hello"`,
			field: FirstbootScriptCustomization{
				union: json.RawMessage(`{"type":"custom","name":"test","contents":"echo hello"}`),
			},
		},
		{
			name: "satellite",
			toml: `type = "satellite"
name = "test"
command = "echo hello"`,
			field: FirstbootScriptCustomization{
				union: json.RawMessage(`{"type":"satellite","name":"test","command":"echo hello"}`),
			},
		},
		{
			name: "aap",
			toml: `type = "aap"
name = "test"
job_template_url = "https://aap.example.com/api/v2/job_templates/9/callback/"`,
			field: FirstbootScriptCustomization{
				union: json.RawMessage(`{"type":"aap","name":"test","job_template_url":"https://aap.example.com/api/v2/job_templates/9/callback/"}`),
			},
		},
		{
			name: "unknown",
			toml: `type = "unknown"
name = "test"`,
			field: FirstbootScriptCustomization{
				union: json.RawMessage(`{"type":"unknown","name":"test"}`),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var actual FirstbootScriptCustomization

			err := toml.Unmarshal([]byte(tt.toml), &actual)
			assert.NoError(t, err)
			assert.JSONEq(t, string(tt.field.union), string(actual.union))

			b, err := toml.Marshal(tt.field)
			assert.NoError(t, err)

			ok, err := tomlEq([]byte(tt.toml), b)
			if err != nil {
				assert.Fail(t, "TOML equality check failed", "expected: %s, actual: %s, error: %v", tt.toml, string(b), err)
			}

			if !ok {
				assert.Fail(t, "TOML mismatch", "expected: %s, actual: %s", tt.toml, string(b))
			}
		})
	}
}

func TestSelectUnion(t *testing.T) {
	tests := []struct {
		name              string
		json              string
		toml              string
		expectedCustom    *CustomFirstbootCustomization
		expectedSatellite *SatelliteFirstbootCustomization
		expectedAAP       *AAPFirstbootCustomization
		err               string // errors are common for json and toml
	}{
		{
			name: "err-bad-type",
			json: `{"type":"xxx"}`,
			toml: `type = "xxx"`,
			err:  "unknown firstboot customization: missing or invalid type field",
		},
		{
			name: "err-missing-type",
			json: `{}`,
			toml: ``,
			err:  "unknown firstboot customization: missing or invalid type field",
		},
		{
			name: "err-custom-with-aap",
			json: `{"type":"custom","job_template_url":"https://aap.example.com/api/v2/job_templates/9/callback/"}`,
			toml: `type = "custom"
job_template_url = "https://aap.example.com/api/v2/job_templates/9/callback/"`,
			err: "json: unknown field \"job_template_url\"",
		},
		{
			name: "custom",
			json: `{"type":"custom","name":"test","contents":"echo hello"}`,
			toml: `type = "custom"
name = "test"
contents = "echo hello"`,
			expectedCustom: &CustomFirstbootCustomization{
				FirstbootCommonCustomization: FirstbootCommonCustomization{
					Type: "custom",
					Name: "test",
				},
				Contents: "echo hello",
			},
		},
		{
			name: "satellite",
			json: `{"type":"satellite","name":"test","command":"echo hello"}`,
			toml: `type = "satellite"
name = "test"
command = "echo hello"`,
			expectedSatellite: &SatelliteFirstbootCustomization{
				FirstbootCommonCustomization: FirstbootCommonCustomization{
					Type: "satellite",
					Name: "test",
				},
				Command: "echo hello",
			},
		},
		{
			name: "missing-satellite-command",
			json: `{"type":"satellite"}`,
			toml: `type = "satellite"`,
			err:  "missing command field for satellite firstboot customization",
		},
		{
			name: "aap",
			json: `{"type":"aap","host_config_key":"test","job_template_url":"https://aap.example.com/api/v2/job_templates/9/callback/"}`,
			toml: `type = "aap"
host_config_key = "test"
job_template_url = "https://aap.example.com/api/v2/job_templates/9/callback/"`,
			expectedAAP: &AAPFirstbootCustomization{
				FirstbootCommonCustomization: FirstbootCommonCustomization{
					Type: "aap",
				},
				JobTemplateURL: "https://aap.example.com/api/v2/job_templates/9/callback/",
				HostConfigKey:  "test",
			},
		},
		{
			name: "missing-aap-host-config-key",
			json: `{"type":"aap"}`,
			toml: `type = "aap"`,
			err:  "missing job_template_url or host_config_key field for aap firstboot customization",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var actual FirstbootScriptCustomization

			err := json.Unmarshal([]byte(tt.json), &actual)
			if err != nil {
				t.Fatalf("failed to unmarshal JSON: %v", err)
			}

			cust, sat, aap, err := actual.SelectUnion()
			if tt.err != "" {
				if assert.Error(t, err) {
					assert.Equal(t, tt.err, err.Error())
				}
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.expectedCustom, cust)
			assert.Equal(t, tt.expectedSatellite, sat)
			assert.Equal(t, tt.expectedAAP, aap)

			err = toml.Unmarshal([]byte(tt.toml), &actual)
			if err != nil {
				t.Fatalf("failed to unmarshal JSON: %v", err)
			}

			cust, sat, aap, err = actual.SelectUnion()
			if tt.err != "" {
				if assert.Error(t, err) {
					assert.Equal(t, tt.err, err.Error())
				}
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.expectedCustom, cust)
			assert.Equal(t, tt.expectedSatellite, sat)
			assert.Equal(t, tt.expectedAAP, aap)
		})
	}
}
