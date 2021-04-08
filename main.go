package main

import (
	"Project/database"
	"Project/routes"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
)

type Config struct{
	Development struct {
		Database string `yaml:"database"`
		UserName string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"development"`
}

func ValidateConfigPath(path string) error {
	s, err := os.Stat(path)
    if err != nil {
        return err
    }
    if s.IsDir() {
        return fmt.Errorf("'%s' is a directory, not a normal file", path)
    }
    return nil
}

func NewConfig(configPath string) (*Config, error) {
    config := &Config{}
    file, err := os.Open(configPath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    d := yaml.NewDecoder(file)

    if err := d.Decode(&config); err != nil {
        return nil, err
    }

    return config, nil
}


func ParseFlags() (string, error) {
	var configPath string

	flag.StringVar(&configPath, "config", "./config.yml", "path to config file")
	flag.Parse()

	if err := ValidateConfigPath(configPath); err != nil {
        return "", err
    }
	return configPath, nil
}

func main() {

	cfgPath, err := ParseFlags()
	if err != nil {
        log.Fatal(err)
    }
	fmt.Println("cfgPath", cfgPath);
	cfg, err := NewConfig(cfgPath)
    if err != nil {
        log.Fatal(err)
    }

	//fmt.Println("cnfgValue", cfg.Development)
	cfgDev := cfg.Development
	database.Connect(cfgDev.UserName, cfgDev.Password, cfgDev.Database)
	app := gin.New()
	app.Use(CORSMiddleware())
	routes.Setup(app)
	app.Run(":8080")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
