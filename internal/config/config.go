// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf

	Database struct {
		Host         string
		Port         int
		User         string
		Password     string
		Name         string
		SSLMode      string
		MaxOpenConns int
		MaxIdleConns int
	}

	Redis struct {
		Host string
		Pass string
		DB   int
	}
}
