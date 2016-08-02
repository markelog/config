package config_test

import (
	"os"
  // "fmt"
  // "reflect"

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

    It("should read saved file", func() {
      conf.Save()
      result, err := conf.Read()

      if err != nil {
        Ω(false).To(Equal(err))
      } else {
        Ω(result).To(Equal("{}"))
      }
    })

    It("should set one level key with string value", func() {
      conf.Set("test", "1")
      conf.Save()
      result, err := conf.Read()

      if err != nil {
        Ω(false).To(Equal(err))
      } else {
        Ω(`{"test":"1"}`).To(Equal(result))
      }
    })

    It("should set one level key with int value", func() {
      conf.Set("test", 1)
      conf.Save()
      result, err := conf.Read()

      if err != nil {
        Ω(false).To(Equal(err))
      } else {
        Ω(`{"test":1}`).To(Equal(result))
      }
    })

    It("should set second level key with string value", func() {
      conf.Set("test.path", "tester")
      conf.Save()
      result, err := conf.Read()

      if err != nil {
        Ω(false).To(Equal(err))
      } else {
        Ω(`{"test":{"path":"tester"}}`).To(Equal(result))
      }
    })

    It("should set second level key with int value", func() {
      conf.Set("test.path", 1)
      conf.Save()
      result, err := conf.Read()

      if err != nil {
        Ω(false).To(Equal(err))
      } else {
        Ω(`{"test":{"path":1}}`).To(Equal(result))
      }
    })

    It("should get int value", func() {
      conf.Set("test.path", 1)
      conf.Save()

      value := conf.Get("test.path")

      Ω(value).To(Equal(1))
    })

    It("should get string value", func() {
      conf.Set("test.path", "test")
      conf.Save()

      value := conf.Get("test.path")

      Ω(value).To(Equal("test"))
    })

    It("should get boolean value", func() {
      conf.Set("test.path", true)
      conf.Save()

      value := conf.Get("test.path")

      Ω(value).To(Equal(true))
    })
	})
})
