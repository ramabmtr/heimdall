package nexmo

import (
	"errors"

	"github.com/matryer/resync"
	"github.com/njern/gonexmo"
	"github.com/ramabmtr/heimdall/helper/environ"
)

var (
	client *nexmo.Client

	apiKey    string
	apiSecret string
	ttl       int
	sender    string

	errNexmoNotInitialized = errors.New("nexmo client is not properly initialized")

	once resync.Once
)

func init() {
	apiKey = environ.GetEnv("NEXMO_API_KEY").Default("").ToString()
	apiSecret = environ.GetEnv("NEXMO_API_SECRET").Default("").ToString()
	ttl = environ.GetEnv("NEXMO_TTL_DELIVERY").Default("300000").ToInt()
	sender = environ.GetEnv("NEXMO_SENDER").Default("HEIMDALL").ToString()
}

func initClient() {
	once.Do(func() {
		client, _ = nexmo.NewClient(apiKey, apiSecret)
	})
}
