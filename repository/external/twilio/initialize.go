package twilio

import (
	"errors"

	"github.com/matryer/resync"
	"github.com/ramabmtr/heimdall/helper/environ"
	"github.com/sfreiberg/gotwilio"
)

var (
	client *gotwilio.Twilio

	accountSid   string
	authToken    string
	twilioNumber string

	errTwilioNotInitialized = errors.New("twilio client is not properly initialized")

	once resync.Once
)

func init() {
	accountSid = environ.GetEnv("TWILIO_ACCOUNT_SID").Default("").ToString()
	authToken = environ.GetEnv("TWILIO_AUTH_TOKEN").Default("").ToString()
	twilioNumber = environ.GetEnv("TWILIO_NUMBER").Default("").ToString()
}

func initClient() {
	once.Do(func() {
		client = gotwilio.NewTwilioClient(accountSid, authToken)
	})
}
