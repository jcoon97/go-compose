package writer

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"os/user"
)

const (
	defaultDBName = "db"
	outFileName   = "docker-compose.yml"
)

type Configuration struct {
	Version  string             `yaml:"version"`
	Services map[string]Service `yaml:"services"`
}

type Service struct {
	Image       string            `yaml:"image"`
	Command     string            `yaml:"command,omitempty"`
	Environment map[string]string `yaml:"environment"`
	Ports       []string          `yaml:"ports"`
	Restart     string            `yaml:"restart,omitempty"`
}

func (c *Configuration) getHomeDir() string {
	u, err := user.Current()
	c.throwError(err)
	return u.HomeDir + "/Desktop"
}

func (c *Configuration) throwError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func (c *Configuration) WriteFile(path string) {
	p := ""

	if path != "" {
		p = fmt.Sprintf("%s/%s", path, outFileName)
	} else {
		p = fmt.Sprintf("%s/%s", c.getHomeDir(), outFileName)
	}

	f, err := os.Create(p)
	c.throwError(err)
	defer f.Close()

	b, err := yaml.Marshal(&c)
	_, err = f.Write(b)
	c.throwError(err)
}

func NewConfiguration(service Service) *Configuration {
	m := make(map[string]Service)
	m[defaultDBName] = service

	return &Configuration{
		Version:  "3",
		Services: m,
	}
}
