package config

import (
	flag "github.com/spf13/pflag"
)

type FlagConfig struct {
	Path string
}

//var (
//	KafkaBrokers = flag.StringArray("kb", []string{"192.168.0.0:9092", "192.168.0.1:9092"}, "kafka brokers")
//	Conf         = flag.String("c", "doria.toml", "specify the configuration file, default is doria.toml")
//	Num          = flag.Int("n", 6, "specify the number")
//)

func InitFlag() {
	flag.StringVarP(&Flag.Path, "path", "p", "./", "config path")
	////fmt.Println("-------Flag.Path--------", *Flag.Path)
	////Flag.Ports = pflag.IntSlice("ports", []int{8081}, "ports")
	////Flag.ConfigType = pflag.StringP("config_type", "t", "yaml", "config type")
	flag.Parse()
}
