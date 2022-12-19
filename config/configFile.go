package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"strings"
)

func InitConfigFile() {
	//viper.SetEnvPrefix(cmdRoot)
	viper.AutomaticEnv() // 读入匹配的环境变量
	//viper.SetDefault("time", time.Now().Unix()) // 设置自定义参数
	viper.SetConfigType("yaml") // 支持的扩展名有 "json", "toml", "yaml", "yml", "properties", "props", "prop", "env", "dotenv"
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetConfigName("core") // 配置文件名字，注意没有扩展名
	viper.AddConfigPath("./")   // 配置文件的路径
	//viper.AddConfigPath("./")  // 多次调用添加多个配置文件的路径

	err := viper.ReadInConfig() // 搜索路径，并读取配置数据
	if err != nil {
		fmt.Println(fmt.Errorf("Fatal error when reading config file:%s", err))
		os.Exit(1)
	}

	// 配置文件监听
	//viper.WatchConfig()
	//viper.OnConfigChange(func(e fsnotify.Event) {
	//	fmt.Println("Config file changed:", e.Name)
	//})
}
