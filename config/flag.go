package config

import (
	flag "github.com/spf13/pflag"
)

type FlagConfig struct {
	Path  *string
	Ports *[]int
}

//var (
//	KafkaBrokers = flag.StringArray("kb", []string{"192.168.0.0:9092", "192.168.0.1:9092"}, "kafka brokers")
//	Conf         = flag.String("c", "doria.toml", "specify the configuration file, default is doria.toml")
//	Num          = flag.Int("n", 6, "specify the number")
//)

func InitFlag() {
	Flag.Path = flag.String("path", "./", "config path")
	Flag.Ports = flag.IntSlice("ports", []int{8081}, "ports")
	flag.Parse()
}
