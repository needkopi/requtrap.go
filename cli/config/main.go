package config

import (
	"github.com/libidev/requtrap.go/cli/errors"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type ConfigService struct {
	Path     string `yaml:"path"`
	Upstream string `yaml:"upstream"`
}

type ConfigYaml struct {
	Name     string          `yaml:"name"`
	Host     string          `yaml:"host"`
	Port     int             `yaml:"port"`
	Services []ConfigService `yaml:"services"`
}

var (
	Default = `
    name: book-store
    host: 127.0.0.1
    port: 8080
    services:
      - path: /books
        upstream: http://127.0.0.1:8001
      - path: /authors
        upstream: http://127.0.0.1:8002
	`

	config = ConfigYaml{}
)

func Parse(confile string) (*ConfigYaml, error) {
	var err error
<<<<<<< HEAD
	defer cli.IsError(err)
=======
	defer errors.IsError(err)
>>>>>>> 4f479139fe33164a3d363d06d93c4910d263bfd8

	f, err := ioutil.ReadFile(confile)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal([]byte(f), &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
