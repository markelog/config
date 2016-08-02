// Package config deals configurations file
package config

import (
  "io/ioutil"

  "github.com/Jeffail/gabs"
)

type Options struct {
  Path string
}

type Config struct {
  path string
  content *gabs.Container
}

func New(options *Options) *Config {
  config := &Config{}

  config.path = options.Path

  content, _ := gabs.ParseJSON([]byte(`{}`))
  config.content = content

  return config
}

func (config *Config) Set(key string, value interface{}) (*gabs.Container, error) {
  return config.content.SetP(value, key)
}

func (config *Config) Get(key string) interface{} {
  return config.content.Path(key).Data()
}

func (config *Config) Remove(key string) error {
  return config.content.DeleteP(key)
}

func (config *Config) Save() error {
  content := []byte(config.content.String())

  return ioutil.WriteFile(config.path, content, 0666)
}

func (config *Config) Read() (*gabs.Container, error) {
  bytes, err := ioutil.ReadFile(config.path)

  if err != nil {
    return nil, err
  }

  parsed, err := gabs.ParseJSON(bytes)
  config.content = parsed

  if err != nil {
    return nil, err
  }

  return parsed, nil
}

