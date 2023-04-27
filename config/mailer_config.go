package config

import (
	"crypto/tls"
	"log"
	"os"

	"gopkg.in/gomail.v2"
	"gopkg.in/yaml.v3"
)

type Mailer struct {
	Username string `yaml:"mailer_username"`
	Password string `yaml:"mailer_password"`
}

func InitMailer() *gomail.Dialer {
	file, err := os.ReadFile("config/mailer_config.yml")
	if err != nil {
		log.Println(err)
		panic(err)
	}

	var mailer Mailer
	err = yaml.Unmarshal(file, &mailer)
	if err != nil {
		return nil
	}

	m := gomail.NewDialer("smtp.gmail.com", 587, mailer.Username, mailer.Password)
	m.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	return m
}
