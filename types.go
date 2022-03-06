package blueprint

import (
	"errors"
)

const TemplateConfigurationFilenameWithoutExt = "blueprint"

var ErrNotTemplateConfiguration = errors.New("not a template configuration file")

type Template struct {
	variableMap   map[string]*Variable
	variableSlice []*Variable
}

func NewTemplate(tc *TemplateConfiguration) (*Template, error) {
	t := &Template{
		variableMap:   map[string]*Variable{},
		variableSlice: []*Variable{},
	}

	for k, v := range tc.Variables {
		nv, err := NewVariable(k, v.Description, v.Default)
		if err != nil {
			return nil, err
		}

		t.variableMap[k] = nv
		t.variableSlice = append(t.variableSlice, nv)
	}

	return t, nil
}

func (t *Template) GetVariable(varName string) (Variable, error) {
	if v, ok := t.variableMap[varName]; ok {
		return *v, nil
	}

	return Variable{}, errors.New("variable not found")
}

func (t *Template) Variables() ([]Variable, error) {
	result := []Variable{}

	for _, v := range t.variableSlice {
		nv, err := NewVariable(v.Name, v.Description, v.Default)
		if err != nil {
			return nil, err
		}

		result = append(result, *nv)
	}

	return result, nil
}

type TemplateConfiguration struct {
	Variables map[string]Variable `yaml:"variables" json:"variables"`
}

type Variable struct {
	Name        string `yaml:"name" json:"name"`
	Description string `yaml:"description" json:"description"`
	Default     string `yaml:"default" json:"default"`
}

func NewVariable(name, description, defaultValue string) (*Variable, error) {
	if name == "" {
		return nil, errors.New("name is empty")
	}

	if description == "" {
		return nil, errors.New("description is empty")
	}

	if defaultValue == "" {
		return nil, errors.New("defaultValue is empty")
	}

	return &Variable{
		Name:        name,
		Description: description,
		Default:     defaultValue,
	}, nil
}
