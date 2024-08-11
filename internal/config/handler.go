// Package config handling the declaration and reading of the server config
package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

const (
	envKey = "SERVICE_ENV"
)

// ReadJSONConfig read config file with json format
func (c Config) ReadJSONConfig(name string, serviceName string) (ret Config) {
	env := os.Getenv(envKey)
	ret = c
	paths := []string{
		"config/server",
		fmt.Sprintf("cmd/%s/config/server", serviceName),
		fmt.Sprintf("/etc/%s/config/server", serviceName),
	}
	errs := []error{}
	for _, path := range paths {
		fullPath := fmt.Sprintf("%v/%v.%v.json", path, name, env)
		b, err := os.ReadFile(filepath.Clean(fullPath))
		if err != nil {
			errs = append(errs, err)
			continue
		}
		err = json.Unmarshal(b, &ret)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		return
	}
	log.Fatal(errs)
	return
}

// ReadYAMLConfig read config file with json format
func (c Config) ReadYAMLConfig(name string, serviceName string) (ret Config) {
	ret = c
	env := os.Getenv(envKey)
	paths := []string{
		"config/server",
		fmt.Sprintf("cmd/%s/config/server", serviceName),
		fmt.Sprintf("/etc/%s/config/server", serviceName),
	}
	errs := []error{}
	for _, path := range paths {
		fullPath := fmt.Sprintf("%v/%v.%v.yaml", path, name, env)
		b, err := os.ReadFile(filepath.Clean(fullPath))
		if err != nil {
			errs = append(errs, err)
			continue
		}
		err = yaml.Unmarshal(b, &ret)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		return
	}
	log.Fatal(errs)
	return
}
