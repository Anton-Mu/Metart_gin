package dao

import (
	"Metart/config"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DRIVER 指定驱动
const DRIVER = "mysql"

var SqlSession *gorm.DB

// 配置参数映射结构体
type conf struct {
	Url      string `yaml:"url"`
	UserName string `yaml:"userName"`
	Password string `yaml:"password"`
	DbName   string `yaml:"dbname"`
	Port     string `yaml:"post"`
}

// InitMySql 初始化连接数据库，生成可操作基本增删改查结构的变量
func InitMySql() (err error) {
	//获取yaml配置参数
	databaseConfig, err := config.GetDatabaseConfig()
	//将yaml配置参数拼接成连接数据库的url
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		databaseConfig.Username,
		databaseConfig.Password,
		databaseConfig.Host,
		databaseConfig.Port,
		databaseConfig.DBName,
	)

	fmt.Print(dsn)
	//连接数据库
	SqlSession, err = gorm.Open(DRIVER, dsn)
	if err != nil {
		panic(err)
	}
	//验证数据库连接是否成功，若成功，则无异常
	return SqlSession.DB().Ping()
}

// Close 关闭数据库连接
func Close() {
	SqlSession.Close()
}
