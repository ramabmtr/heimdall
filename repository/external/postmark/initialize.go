package postmark

import (
	"errors"

	"github.com/keighl/postmark"
	"github.com/matryer/resync"
	"github.com/ramabmtr/heimdall/helper/environ"
)

var (
	client *postmark.Client

	serverToken  string
	accountToken string
	sender       string
	subject      string

	errPostmarkNotInitialized = errors.New("postmark client is not properly initialized")

	once resync.Once
)

func init() {
	serverToken = environ.GetEnv("POSTMARK_SERVER_TOKEN").Default("").ToString()
	accountToken = environ.GetEnv("POSTMARK_ACCOUNT_TOKEN").Default("").ToString()
	sender = environ.GetEnv("POSTMARK_SENDER").Default("").ToString()
	subject = environ.GetEnv("POSTMARK_SUBJECT").Default("One Time Password").ToString()
}

func initClient() {
	once.Do(func() {
		client = postmark.NewClient(serverToken, accountToken)
	})
}
