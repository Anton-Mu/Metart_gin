package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
}

func GetDatabaseConfig() (*DatabaseConfig, error) {
	var databaseConfig DatabaseConfig

	// 读取配置文件
	yamlFile, err := os.ReadFile("./config/db.yaml")
	if err != nil {
		return nil, fmt.Errorf("读取数据库配置文件失败：%v", err)
	}
	// 解析配置文件中数据
	err = yaml.Unmarshal(yamlFile, &databaseConfig)
	if err != nil {
		return nil, fmt.Errorf("解析数据库配置文件失败：%v", err)
	}

	return &databaseConfig, nil
}
