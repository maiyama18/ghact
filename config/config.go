package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	jsonPath string
	vars     map[string]string
}

func New(jsonPath string) (*Config, error) {
	c := &Config{jsonPath: jsonPath, vars: make(map[string]string)}

	if _, err := os.Stat(jsonPath); os.IsNotExist(err) {
		if err := c.initConfigJson(); err != nil {
			return nil, err
		}
	}

	b, err := ioutil.ReadFile(jsonPath)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(b, &c.vars); err != nil {
		return nil, err
	}
	return c, nil
}

func (c *Config) Get(key string) (string, error) {
	val, ok := c.vars[key]
	if !ok {
		return "", fmt.Errorf("%s not set", key)
	}
	return val, nil
}

func (c *Config) Set(key, val string) error {
	c.vars[key] = val
	if err := c.save(); err != nil {
		return err
	}
	return nil
}

func (c *Config) save() error {
	b, err := json.Marshal(c.vars)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(c.jsonPath, b, 0644); err != nil {
		return err
	}
	return nil
}

func (c *Config) initConfigJson() error {
	if _, err := os.Stat(c.jsonPath); os.IsNotExist(err) {
		f, err := os.Create(c.jsonPath)
		if err != nil {
			return err
		}
		_ = f.Close()

		if err := c.save(); err != nil {
			return err
		}
	}
	return nil
}
