package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/go-mail/mail"
	"github.com/spf13/viper"
)

var logger = log.New(os.Stdout, "sparrow: ", log.Ldate|log.Ltime|log.Lshortfile)

type mailSettings struct {
	hostname    string
	port        int
	user        string
	password    string
	message     string
	destination string
}

func processPassword(config *viper.Viper) string {
	logger.Println("Decoding password...")
	raw_password := config.GetString("MAIL_PASSWORD")
	decoded_password, err := base64.StdEncoding.DecodeString(raw_password)
	if err != nil {
		log.Println("Couldn't decode password!")
		log.Fatal(err)
	}
	trimmed_password := strings.TrimSpace(string(decoded_password))
	return string(trimmed_password)
}

func main() {
	logger.Println("Welcome to Sparrow!")
	logger.Println("Parsing options...")

	v := viper.New()

	v.SetEnvPrefix("SPARROW")
	v.AutomaticEnv()

	mailSettings := mailSettings{
		hostname:    v.GetString("MAIL_HOSTNAME"),
		port:        v.GetInt("MAIL_PORT"),
		user:        v.GetString("MAIL_USERNAME"),
		password:    processPassword(v),
		destination: v.GetString("MAIL_DESTINATION"),
		message:     "",
	}
	website := v.GetString("WEBSITE")
	expected_status_code := v.GetInt("STATUS_CODE")

	logger.Println("Preparing mailer & message")
	d := mail.NewDialer(mailSettings.hostname, mailSettings.port, mailSettings.user, mailSettings.password)
	m := mail.NewMessage()
	m.SetHeader("From", mailSettings.user)
	m.SetHeader("To", mailSettings.destination)

	if test := v.GetBool("TEST"); test {
		logger.Println("Sending test email...")
		m.SetHeader("Subject", "Test Email from Sparrow!")
		m.SetBody("text/html", "Test body from your Sparrow System!")
		d.DialAndSend(m)
		return
	}

	m.SetHeader("Subject", fmt.Sprintf("%s healthcheck", website))

	logger.Printf("Checking website response: %s", website)
	resp, err := http.Get(website)
	if err != nil {
		mailSettings.message = fmt.Sprintf("CIRITAL ERROR DURING STATUS CHECK: %s", err)
	} else {
		mailSettings.message = strconv.Itoa(resp.StatusCode)
		if resp.StatusCode == expected_status_code {
			logger.Println("All good, no email!")
			return
		} else {
			logger.Println("Something is no yes, sending email!")
		}
	}
	logger.Println("Status: ", mailSettings.message)
	m.SetBody("text/html", fmt.Sprintf("Status: %s", mailSettings.message))
	if err := d.DialAndSend(m); err != nil {
		logger.Println(err)
	}
}
