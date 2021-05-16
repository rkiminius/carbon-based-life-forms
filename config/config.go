package config

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"os"
)

type Config struct {
	ClientPort  string  `yaml:"clientPort"`
	ManagerPort string  `yaml:"managerPort"`
	MongoDb     MongoDb `yaml:"mongoDb"`
	LogPath     string  `yaml:"logPath"`
	Log         *log.Logger
}

type MongoDb struct {
	DriverName string `yaml:"driverName"`
	Host       string `yaml:"host"`
	Port       string `yaml:"port"`
	DbName     string `yaml:"dbName"`
}

var Conf *Config

func GetConfig(configFile string) *Config {
	Conf = &Config{}
	if configFile != "" {
		Conf.GetConfFromFile(configFile)
	}
	Conf.Log = log.New()

	logFilePath := Conf.LogPath

	logFile, err := os.OpenFile(logFilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		Conf.Log.Info("Failed to log to file, using default stderr")
		Conf.Log.Panic(err)
		panic(err)
	}
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)

	return Conf
}

func (c *Config) SetDefaultValue(val *string, defValue string) {
	if *val == "" {
		*val = defValue
	}
}

func (c *Config) GetConfFromFile(fileName string) {
	pwd, err := os.Getwd()
	if err != nil {
		log.WithFields(log.Fields{"function": "GetConfFromFile method",
			"paramsFileName": fileName}).Error(err)
		// unhandled error
	}
	yamlFile, err := ioutil.ReadFile(pwd + "/" + fileName)
	if err != nil {
		log.WithFields(log.Fields{"function": "GetConfFromFile method",
			"paramsFileName": fileName,
			"paramsPwd":      pwd}).Fatal(err)
	}
	c.GetConfFromString(string(yamlFile))
}

func (c *Config) GetConfFromString(yamlString string) {
	err := yaml.Unmarshal([]byte(yamlString), c)
	if err != nil {
		log.WithFields(log.Fields{"function": "GetConfFromString method",
			"paramsYamlString": yamlString}).Fatal(err)
	}
	return
}
