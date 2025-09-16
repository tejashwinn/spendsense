package config

import (
	"github.com/jinzhu/configor"
	"github.com/tejashwinn/sependsense/mode"
)

type Configuration struct {
	Server struct {
		KeepAlivePeriodSeconds int
		ListenAddr             string `default:""`
		Port                   int    `default:"8080"`

		SSL struct {
			Enabled         bool   `default:"false"`
			RedirectToHTTPS bool   `default:"true"`
			ListenAddr      string `default:""`
			Port            int    `default:"443"`
			CertFile        string `default:""`
			CertKey         string `default:""`
			LetsEncrypt     struct {
				Enabled   bool   `default:"false"`
				AcceptTOS bool   `default:"false"`
				Cache     string `default:"data/certs"`
				Hosts     []string
			}
		}
		Cors struct {
			AllowOrigins []string
			AllowMethods []string
			AllowHeaders []string
		}
	}
	Database struct {
		Connection string
	}
}

func configFiles() []string {
	if mode.Get() == mode.TestDev {
		return []string{"config.yml"}
	}
	return []string{"config.yml", "/etc/sependsense/config.yml"}
}

func Get() *Configuration {
	conf := new(Configuration)
	err := configor.New(
		&configor.Config{ENVPrefix: "SPENDSENSE", Silent: true}).Load(
		conf, configFiles()...,
	)
	if err != nil {
		panic(err)
	}
	return conf
}
