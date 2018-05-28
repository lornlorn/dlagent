package utils

import (
	"log"

	"github.com/robfig/config"
)

/*
Conf *config.Config
*/
var Conf *config.Config

/*
ReadConf func(string,string) (string,error)
*/
func ReadConf(section string, option string) (string, error) {
	Conf, err := config.ReadDefault("conf/app.conf")
	if err != nil {
		log.Printf("Read Config File Fail : %v\n", err)
		return "", err
	}
	return Conf.String(section, option)
	// result is string "http://www.example.com/some/path"

	// c.Int("service-1", "maxclients")
	// // result is int 200

	// c.Bool("service-1", "delegation")
	// // result is bool true

	// c.String("service-1", "comments")
	// // result is string "This is a multi-line\nentry"

}
