package conf

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
	"sync"
)

var conf *viper.Viper

var cfg = pflag.StringP("config", "c", "configs/conf.yaml", "Configuration file.")

var confOnce sync.Once

func Conf() *viper.Viper {

	confOnce.Do(func() {
		viper.SetConfigFile(*cfg)

		if err := viper.ReadInConfig(); err != nil {
			log.Fatalln("read conf err:", err)
		}
		conf = viper.GetViper()
	})

	return conf

}
