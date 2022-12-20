package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"strings"
)

type viperConfig struct {
	AppDebug bool
	ServerIp string
	Ports    []int
	Log      struct {
		Level string
	}
}

func InitConfigFile() {
	//viper.SetEnvPrefix(cmdRoot) // 前缀
	//viper.SetDefault("time", time.Now().Unix()) // 设置自定义参数
	//viper.AddConfigPath("./")  // 多次调用添加多个配置文件的路径

	replacer := strings.NewReplacer(".", "_")
	viper.AutomaticEnv() // 读入匹配的环境变量
	viper.SetEnvKeyReplacer(replacer)

	viper.SetConfigType("yaml")   // 支持的扩展名有 "json", "toml", "yaml", "yml", "properties", "props", "prop", "env", "dotenv"
	viper.SetConfigName("config") // 配置文件名字，注意没有扩展名
	viper.AddConfigPath("./")     // 配置文件的路径

	err := viper.ReadInConfig() // 搜索路径，并读取配置数据
	if err != nil {
		fmt.Println(fmt.Errorf("Fatal error when reading config file:%s", err))
		os.Exit(1)
	}

	Viper.AppDebug = viper.GetBool("app_debug")
	Viper.ServerIp = viper.GetString("server_ip")
	Viper.Ports = viper.GetIntSlice("ports")
	Viper.Log.Level = viper.GetString("log.level")

	fmt.Println("----------------", viper.GetString("redis.password"))

	//fmt.Println("viper.AllSettings", viper.AllSettings())
	//fmt.Println("app_debug", viper.GetString("server_ip"))
	//fmt.Println("ports", viper.GetIntSlice("ports"))

	// 配置文件监听
	//viper.WatchConfig()
	//viper.OnConfigChange(func(e fsnotify.Event) {
	//	fmt.Println("Config file changed:", e.Name)
	//})
}
