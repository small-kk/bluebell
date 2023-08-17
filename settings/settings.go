package settings

import (
	"flag"
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(AppConfig)

type AppConfig struct {
	Name         string `mapstructure:"name"`
	Mode         string `mapstructure:"mode"`
	Version      string `mapstructure:"version"`
	StartTime    string `mapstructure:"start_time"`
	Port         int    `mapstructure:"port"`
	MachineID    int64  `mapstructure:"machine_id"`
	*LogConfig   `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	FileName   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DbName       string `mapstructure:"dbname"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_openConns"`
	MaxIdleConns int    `mapstructure:"max_idleConns"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

func Init() (err error) {
	//设置配置文件常用方法：
	//第一种，直接指定配置文件的路径（相对路径或绝对路径）
	//viper.SetConfigFile("./conf/config.yaml")

	//第二种方式，指定配置文件名和配置文件位置，viper自行查找可用的配置文件
	//viper.SetConfigName("config")  配置文件名不需要带后缀
	//viper.AddConfigPath("./conf/") 配置文件位置可以配置多个，可以看到这里是Add

	path := flag.String("config_path", "./conf/config.yaml", "配置文件路径")
	flag.Parse()
	viper.SetConfigFile(*path)

	//viper.SetConfigName("config")  //指定配置文件名称（不需要带后缀）
	//viper.SetConfigType("yaml")    //指定配置文件类型（专门用于从远程获取配置信息时指定配置文件类型）
	//viper.AddConfigPath("./conf/") //指定查找配置文件的路径（这里使用相对路径）

	err = viper.ReadInConfig() //读取配置信息
	if err != nil {
		fmt.Printf("viper.ReadInConfig() failed,err:%v\n", err)
		return
	}
	//将读取到的配置信息反序列化到Conf变量中
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal failed,err:%v\n", err)
	}
	//自动监视配置文件变化
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了...")
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper.Unmarshal failed,err:%v\n", err)
		}
	})
	return
}
