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
    conf *config.Config
	)

  BeforeEach(func() {
    path = "testdata"
  })

	Describe("Define config", func() {
    BeforeEach(func() {
      path += "/simple.json"

      conf = config.New(&config.Options{File: path})
    })

    AfterEach(func() {
      os.Remove(path)
    })

    It("should save config with correct name", func() {
      conf.Save()
      if _, err := os.Stat(path); os.IsNotExist(err) {
        Ω(false).To(Equal(true))
      } else {
        Ω(true).To(Equal(true))
      }
    })

    It("should save config with correct name", func() {
      conf.Save()
      result, err := conf.Read()

      if err != nil {
        Ω(false).To(Equal(err))
      } else {
        Ω("").To(Equal(result))
      }
    })
	})
})
