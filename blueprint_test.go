package blueprint

import (
	"reflect"
	"testing"

	"github.com/elliotxx/gulu/json"
)

var TestTemplateConfiguration = &TemplateConfiguration{
	Variables: map[string]Variable{
		"Env": {
			Description: "The env name. One of dev,test,stable,pre,sim,gray,prod.",
			Default:     "dev",
		},
		"ClusterName": {
			Description: "The Cluster Name.",
			Default:     "kubernetes-dev",
		},
		"Image": {
			Description: "The Image Address.",
			Default:     "gcr.io/google-samples/gb-frontend:v4",
		},
	},
}

func TestParseTemplateConfiguration(t *testing.T) {
	type args struct {
		path string
	}

	tests := []struct {
		name    string
		args    args
		want    *TemplateConfiguration
		wantErr bool
	}{
		{
			name: "success-parse-yaml-file",
			args: args{
				path: "./testdata/blueprint.yaml",
			},
			want:    TestTemplateConfiguration,
			wantErr: false,
		},
		{
			name: "success-parse-yml-file",
			args: args{
				path: "./testdata/blueprint.yml",
			},
			want:    TestTemplateConfiguration,
			wantErr: false,
		},
		{
			name: "success-parse-json-file",
			args: args{
				path: "./testdata/blueprint.json",
			},
			want:    TestTemplateConfiguration,
			wantErr: false,
		},
		{
			name: "invalid-file-name",
			args: args{
				path: "./testdata/blueprint-invalid.json",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "not-existed-file",
			args: args{
				path: "./testdata/blueprint-not-existed.json",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid-file-ext",
			args: args{
				path: "./testdata/blueprint.txt",
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseTemplateConfiguration(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseTemplateConfiguration() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseTemplateConfiguration() = %v, want %v", json.MustPrettyMarshalString(got), json.MustPrettyMarshalString(tt.want))
			}
		})
	}
}
