package initialize

import (
	"bwcxgdz/v2/user_srv/global"
	"fmt"

	"github.com/spf13/viper"
)

func InitConfig() {
	//从配置文件中读取出对应的配置
	configFilePrefix := "config"
	configFileName := fmt.Sprintf("user_srv/%s-debug.yaml", configFilePrefix)

	v := viper.New()
	//文件的路径如何设置
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	//这个对象如何在其他文件中使用 - 全局变量
	if err := v.Unmarshal(&global.ServerConfig); err != nil {
		panic(err)
	}

}
