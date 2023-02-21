package config

import (
	"fmt"
	"github.com/num30/config"
	"time"
)

type Link struct {
	Text   *string
	Target string `validate:"required"`
}

type DBConfig struct {
	Addr     string `default:"localhost:6379"`
	Database int    `default:"0"`
	Password string `default:""`
}

type Config struct {
	Db   DBConfig   `validate:"required"`
	Page PageConfig `default:"{}"`
	Port int        `default:"9000"`
}

type FooterConfig struct {
	Header        string
	CopyYear      int
	CopyStatement string
	Content       string
	Links         []Link `default:"[]"`
	MoreLinks     []Link `default:"[]"`
}

type PageConfig struct {
	Title  string       `default:"GURL - Shortener"`
	Lang   string       `default:"en"`
	Name   string       `default:"GURL"`
	Footer FooterConfig `default:"{}"`
	Color  string       `default:"blue"`
}

var conf Config

func NewConfig() *Config {
	err := config.NewConfReader("config").Read(&conf)
	if err != nil {
		panic(fmt.Errorf("error reading config: %v", err))
	}
	conf.Page.Footer.CopyYear = time.Now().Year()
	return &conf
}
