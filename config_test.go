package config_test

import (
	"os"

	"github.com/markelog/config"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Config", func() {
	var (
		path string
	)

  BeforeEach(func() {
    path = "testdata"
  })

	Describe("Define config", func() {
    It("should save config with correct name", func() {
      path += "/simple.json"

      conf := config.New(&config.Options{File: path})
      conf.Save()

      if _, err := os.Stat(path); os.IsNotExist(err) {
        Ω(false).To(Equal(true))
      } else {
        Ω(true).To(Equal(true))
      }

      os.Remove(path)
    })
	})
})
