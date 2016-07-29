// Package config deals configurations file
package config

import (
  "io/ioutil"

  "github.com/Jeffail/gabs"
)

type Options struct {
  File string
}

type Config struct {
  options *Options
  content *gabs.Container
}

func New(options *Options) *Config {
  config := &Config{}
  config.options = options

  content, _ := gabs.ParseJSON([]byte(`{}`))
  config.content = content

  return config
}

func (config *Config) Set(key string, value interface{}) (*gabs.Container, error) {
  return config.content.SetP(value, key)
}

func (config *Config) Save() error {
  content := []byte(config.content.String())

  return ioutil.WriteFile(config.options.File, content, 0666)
}

func (config *Config) Read() (string, error) {
  bytes, err := ioutil.ReadFile(config.options.File)

  if err != nil {
    return "", err
  }

  parsed, err := gabs.ParseJSON(bytes)

  if err != nil {
    return "", err
  }

  return parsed.String(), nil
}

