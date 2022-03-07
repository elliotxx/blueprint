package blueprint

import (
	"os"
	"reflect"
	"testing"

	"github.com/elliotxx/gulu/json"
)

var (
	TestVariableMap   map[string]*Variable
	TestVariableSlice []*Variable
)

func TestNewTemplate(t *testing.T) {
	type args struct {
		tc *TemplateConfiguration
	}

	tests := []struct {
		name    string
		args    args
		want    *Template
		wantErr bool
	}{
		{
			name: "t1",
			args: args{
				tc: TestTemplateConfiguration,
			},
			want: &Template{
				variableMap:   TestVariableMap,
				variableSlice: TestVariableSlice,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTemplate(tt.args.tc)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTemplate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTemplate() = %v, want %v", json.MustPrettyMarshalString(got), json.MustPrettyMarshalString(tt.want))
			}
		})
	}
}

func TestTemplate_GetVariable(t *testing.T) {
	type fields struct {
		variableMap   map[string]*Variable
		variableSlice []*Variable
	}

	type args struct {
		varName string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Variable
		wantErr bool
	}{
		{
			name: "t1",
			fields: fields{
				variableMap:   TestVariableMap,
				variableSlice: TestVariableSlice,
			},
			args: args{
				varName: "Env",
			},
			want:    *TestVariableMap["Env"],
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Template{
				variableMap:   tt.fields.variableMap,
				variableSlice: tt.fields.variableSlice,
			}
			got, err := tr.GetVariable(tt.args.varName)
			if (err != nil) != tt.wantErr {
				t.Errorf("Template.GetVariable() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Template.GetVariable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTemplate_Variables(t *testing.T) {
	type fields struct {
		variableMap   map[string]*Variable
		variableSlice []*Variable
	}

	tests := []struct {
		name    string
		fields  fields
		want    []Variable
		wantErr bool
	}{
		{
			name: "t1",
			fields: fields{
				variableMap:   TestVariableMap,
				variableSlice: TestVariableSlice,
			},
			want: []Variable{
				*TestVariableMap["ClusterName"],
				*TestVariableMap["Env"],
				*TestVariableMap["Image"],
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Template{
				variableMap:   tt.fields.variableMap,
				variableSlice: tt.fields.variableSlice,
			}
			got, err := tr.Variables()
			if (err != nil) != tt.wantErr {
				t.Errorf("Template.Variables() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Template.Variables() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewVariable(t *testing.T) {
	type args struct {
		name         string
		description  string
		defaultValue string
	}

	tests := []struct {
		name    string
		args    args
		want    *Variable
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				name:         "test-name",
				description:  "test-description",
				defaultValue: "test-default-value",
			},
			want: &Variable{
				Name:        "test-name",
				Description: "test-description",
				Default:     "test-default-value",
			},
			wantErr: false,
		},
		{
			name: "missing-name",
			args: args{
				description:  "test-description",
				defaultValue: "test-default-value",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "missing-description",
			args: args{
				name:         "test-name",
				defaultValue: "test-default-value",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "missing-default",
			args: args{
				name:        "test-name",
				description: "test-description",
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewVariable(tt.args.name, tt.args.description, tt.args.defaultValue)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewVariable() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewVariable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMain(m *testing.M) {
	TestTemplateConfiguration = &TemplateConfiguration{
		Variables: map[string]Variable{
			"ClusterName": {
				Description: "The Cluster Name.",
				Default:     "kubernetes-dev",
			},
			"Env": {
				Description: "The env name. One of dev,test,stable,pre,sim,gray,prod.",
				Default:     "dev",
			},
			"Image": {
				Description: "The Image Address.",
				Default:     "gcr.io/google-samples/gb-frontend:v4",
			},
		},
	}

	TestVariableMap = map[string]*Variable{
		"ClusterName": {
			Name:        "ClusterName",
			Description: "The Cluster Name.",
			Default:     "kubernetes-dev",
		},
		"Env": {
			Name:        "Env",
			Description: "The env name. One of dev,test,stable,pre,sim,gray,prod.",
			Default:     "dev",
		},
		"Image": {
			Name:        "Image",
			Description: "The Image Address.",
			Default:     "gcr.io/google-samples/gb-frontend:v4",
		},
	}

	TestVariableSlice = []*Variable{
		{
			Name:        "ClusterName",
			Description: "The Cluster Name.",
			Default:     "kubernetes-dev",
		},
		{
			Name:        "Env",
			Description: "The env name. One of dev,test,stable,pre,sim,gray,prod.",
			Default:     "dev",
		},
		{
			Name:        "Image",
			Description: "The Image Address.",
			Default:     "gcr.io/google-samples/gb-frontend:v4",
		},
	}

	os.Exit(m.Run())
}
