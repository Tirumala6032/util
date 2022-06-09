package config

import (
	"encoding/json"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

var Default_Config_File_Type = "yaml"
var Default_Config_Folder = "config"
var Default_Config_File_Name = "dev"

var Default_Path = "config/dev.yaml"

func InitConfiguration(path string, fileType string) (*Config, error) {

	var config *Config
	var err error

	if fileType == "" {
		fileType = Default_Config_File_Type
	}

	if path != "" {
		Default_Path = fmt.Sprintf("%s.%s", path, fileType)
	}

	switch fileType {
	case "json":
		config, err = getJsonConfiguration(Default_Path)
	case "yaml":
		config, err = getYamlConfiguration(Default_Path)
	}

	if err != nil {
		return nil, err
	}

	return config, nil
}

func getJsonConfiguration(path string) (*Config, error) {

	var configuration Config

	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&configuration)

	if err != nil {
		return nil, err
	}

	return &configuration, nil

}

func getYamlConfiguration(path string) (*Config, error) {
	config := &Config{}

	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	d := yaml.NewDecoder(file)

	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}
