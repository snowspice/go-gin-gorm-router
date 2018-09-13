package config

import (
	"os"
	"strings"
	"github.com/stackcats/gosh"
	"github.com/lunny/log"
)

//配置中心服务类

// Config ...
type Config struct {
//	Mongo string   `consul:"mongo/cache"`
	Mysql string   `consul:"mysql/test"`
//	Nsq   []string `consul:"nsq"`
}

var (
	config            *Config
	// defaultConsulHost = "127.0.0.1"
	defaultConsulHost = "192.168.99.100"
)

func GetAppConf() *Config{
	return config
}

func init() {
	consul := strings.Trim(os.Getenv("MICRO_REGISTRY_ADDRESS"), " ")
	config = &Config{}
	if consul == "" {
		consul = defaultConsulHost

	}
	c := gosh.DefaultConfig()
	c.Address = consul + ":8500"
	g, err := gosh.NewClient(c)
	if err != nil {
		log.Fatal(err)
	}

	err = g.Unmarshal(config)
	if err != nil {
		log.Fatal(err)
	}
	log.Debug(config)
}

