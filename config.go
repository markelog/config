// Package config deals configurations file
package config

import (
  "io/ioutil"
)

type Options struct {
  File string
}

type Config struct {
  options *Options
  content string
}

func New(options *Options) *Config {
  config := &Config{}
  config.options = options
  config.content = "test"

  return config
}

func (config *Config) Save() error {
  content := []byte(config.content)

  return ioutil.WriteFile(config.options.File, content, 0666)
}


// func (config *Config) Read() (string, error) {
//   bytes, err := ioutil.ReadFile(config.options.file)

//   if err != nil {
//     return "", err
//   }

//   return string(bytes), nil
// }

