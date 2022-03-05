package blueprint

import (
	"encoding/json"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// ParseTemplateConfiguration parses the template configuration file
func ParseTemplateConfiguration(path string) (*TemplateConfiguration, error) {
	// Determine whether it is a template configuration file
	if !IsTemplateConfiguration(path) {
		return nil, ErrNotTemplateConfiguration
	}

	// Load the template configuration file
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// Parse the template configuration file
	var configuration TemplateConfiguration

	ext := filepath.Ext(path)
	switch ext {
	case ".yaml", ".yml":
		err = yaml.Unmarshal(data, &configuration)
	case ".json":
		err = json.Unmarshal(data, &configuration)
	default:
		err = ErrNotTemplateConfiguration
	}

	if err != nil {
		return nil, err
	}

	return &configuration, nil
}

// IsTemplateConfiguration determine whether the given path is Template Configuration file
func IsTemplateConfiguration(path string) bool {
	f, err := os.Stat(path)
	return err == nil && !f.IsDir() && f.Mode().IsRegular() && FilenameWithoutExt(filepath.Base(path)) == TemplateConfigurationFilenameWithoutExt
}

// FilenameWithoutExt returns the filename without extension
func FilenameWithoutExt(filename string) string {
	filename = filepath.Base(filename)
	ext := filepath.Ext(filename)

	return filename[:len(filename)-len(ext)]
}
