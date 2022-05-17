package config

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Global struct {
	Mysql struct {
		Host string `mapstructrue:"host"`
		Port string `mapstructrue:"port"`
		User string `mapstructrue:"user"`
		Pass string `mapstructrue:"pass"`
		Name string `mapstructrue:"name"`
	}
	Pgsql struct {
		Name string `mapstructrue:"name"`
		Pass string `mapstructrue:"pass"`
	}
}

var g *Global

func GetGlobal() *Global {

	conf := "./config/config.yml"
	viper.SetConfigFile(conf)
	content, err := ioutil.ReadFile(conf)
	if err != nil {
		log.Fatal(fmt.Sprintf("Read conf file fail: %s", err.Error()))
	}
	//Replace environment variables
	err = viper.ReadConfig(strings.NewReader(os.ExpandEnv(string(content))))
	if err != nil {
		log.Fatal(fmt.Sprintf("Parse conf file fail: %s", err.Error()))
	}

	if err = viper.Unmarshal(&g); err != nil {
		color.Red.Println("unable to decode into struct, %v", err)
	}

	return g
}
