package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"regexp"
	"sync"
)

type Config struct {
	Buckets map[string]Bucket `yaml:"buckets"`
	Headers []HeaderYaml      `yaml:"headers"`
}

var instance *Config
var once sync.Once

func GetInstance() *Config {
	once.Do(func() {
		instance = &Config{}
	})
	return instance
}

func (self *Config) Load(filePath string) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	errYaml := yaml.Unmarshal([]byte(data), self)

	for name, bucket := range self.Buckets {
		if bucket.Transform != nil && bucket.Transform.Path != "" {
			bucket.Transform.PathRegexp = regexp.MustCompile(bucket.Transform.Path)
			self.Buckets[name] = bucket
		}
	}

	if errYaml != nil {
		panic(errYaml)
	}

}