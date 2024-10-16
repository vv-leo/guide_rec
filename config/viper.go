package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func ViperInit() {

	env := os.Getenv("MY_ENV")
	if len(env) == 0 {
		env = "dev"
	}
	//environment := "dev"
	//args := os.Args
	////默认走dev环境，如果启动带参数则取环境参数
	//if len(args) == 2 {
	//	environment = args[1]
	//}

	viper.SetConfigName(env)        // 配置文件名称(无扩展名)
	viper.SetConfigType("yaml")     // 或viper.SetConfigType("YAML")
	viper.AddConfigPath("./config") // 配置文件路径
	err := viper.ReadInConfig()     // 查找并读取配置文件
	if err != nil {                 // 处理读取配置文件的错误
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

}
