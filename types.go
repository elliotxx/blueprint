package blueprint

import "errors"

const TemplateConfigurationFilenameWithoutExt = "blueprint"

var ErrNotTemplateConfiguration = errors.New("not a template configuration file")

type TemplateConfiguration struct {
	Variables map[string]Variable `yaml:"variables" json:"variables"`
}

type Variable struct {
	Name        string `yaml:"name" json:"name"`
	Description string `yaml:"description" json:"description"`
	Default     string `yaml:"default" json:"default"`
}
