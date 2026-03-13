package serviceinit

import (
	"LocalAPIGateway/src/common"
	"fmt"
	"log"
	"os"

	"gopkg.in/natefinch/lumberjack.v2"
	"gopkg.in/yaml.v3"
)

var serviceDir string
var dataDir string
var logDir string

type ServiceConfig struct {
	ConfigDir string `yaml:"config-dir"`
	LogDir    string `yaml:"log-dir"`
	DataDir   string `yaml:"data-dir"`
}

// init 初始化服务
func init() {
	fmt.Println("enter init")
	readAll()
	logInit()
}

// readAll 读取配置文件 获取配置
func readAll() {
	data, err := os.ReadFile(common.ConfigFilePath)
	if err != nil {
		panic(err)
	}

	var serviceConfig ServiceConfig
	err = yaml.Unmarshal(data, &serviceConfig)
	if err != nil {
		panic(err)
	}

	serviceDir = serviceConfig.ConfigDir
	dataDir = serviceConfig.DataDir
	logDir = serviceConfig.LogDir
}

// logInit 初始化日志配置
func logInit() {
	logFileName := logDir + "/" + common.DefaultLogFileName
	fmt.Println("enter logInit: ", logFileName)

	lumberJackLogger := &lumberjack.Logger{
		Filename:   logFileName,
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   true,
		LocalTime:  true,
	}

	// 标准 log 输出重定向到 lumberJackLogger
	log.SetOutput(lumberJackLogger)
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
}
