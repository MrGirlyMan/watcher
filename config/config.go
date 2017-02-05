package config

import (
	"os"
	// "fmt"
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// Global settings
var (
  TEMPLATE_DIR = "test"
)

type Config struct {
	Foo 					string
	Bar 					[]string
	Test 					string
	Frontman_build_path 	string
	Redis 					string
}

func getConfigFile(path string) (file string, err error) {
	// TODO: Rework this function. It doesn't check to make sure the path is actually a file.

	// Check if path is valid file or dir
	pwd, err := os.Getwd()
	
	if err != nil {
		return "", err
	}

	return filepath.Join(pwd, path), nil
}

func Configs(c string) *Config {
	var configuration Config

	config_file, _ := getConfigFile(c)

	content, err := ioutil.ReadFile(config_file)

	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(content, &configuration)

    if err != nil {
        panic(err)
    }

	return &configuration
}