package config

import (
	"fmt"
	"time"
)

// const dbType string = "postgres" //"mssql" //"mysql" //"postgres"
// const pgsqlConnString string = "host=localhost port=5432 user=postmasters_user dbname=postmasters password=tbm123 sslmode=disable"
// const mysqlConnString string = "root:root@/postmasters?charset=utf8&parseTime=True&loc=Local"
// const mssqlConnString string = "server=(local)\\SQLEXPRESS;user id=pmuser;password=tbm123;database=PostMasters" //"sqlserver://pmuser:tbm123@(local)\\sqlexpress?database=PostMasters"
// const connectionString string = pgsqlConnString                                                                 //pgsqlConnString //mysqlConnString

// Config struct
type Config struct {
	Web *WebConfig
	DB  *DBConfig
}

// WebConfig struct
type WebConfig struct {
	Scheme       string
	Host         string
	Port         int
	WriteTimeout time.Duration
	ReadTimeout  time.Duration
	IdleTimeout  time.Duration
}

// DBConfig struct
type DBConfig struct {
	Dialect  string
	Host     string
	Port     int
	Username string
	Password string
	Name     string
	Charset  string
	SSLMode  string
}

// New is a constructor
func New() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "postgres",
			Host:     "localhost",
			Port:     5432,
			Username: "postmasters_user",
			Password: "tbm123",
			Name:     "postmasters",
			Charset:  "",
			SSLMode:  "disable",
		},
		Web: &WebConfig{
			Scheme:       "http",
			Host:         "localhost",
			Port:         3001,
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
			IdleTimeout:  60 * time.Second,
		},
	}
}

// GetConnectionString function
func (dbConfig *DBConfig) GetConnectionString() string {
	switch dbConfig.Dialect {
	case "postgres":
		return fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=%s",
			dbConfig.Host, dbConfig.Port, dbConfig.Name, dbConfig.Username, dbConfig.Password, dbConfig.SSLMode)
	case "mysql":
		return fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
			dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name, dbConfig.Charset)
	case "mssql":
		return fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s",
			dbConfig.Host, dbConfig.Username, dbConfig.Password, dbConfig.Name)
	default:
		return fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=%s",
			dbConfig.Host, dbConfig.Port, dbConfig.Name, dbConfig.Username, dbConfig.Password, dbConfig.SSLMode)
	}
}
