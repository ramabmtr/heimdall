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
	serverToken = environ.GetEnv("POSTMARK_SERVER_TOKEN").Default("456403d5-67a1-4c2d-830b-045e2492621f").ToString()
	accountToken = environ.GetEnv("POSTMARK_ACCOUNT_TOKEN").Default("c0d5a38c-c5a3-4a61-86eb-2b86a259e987").ToString()
	sender = environ.GetEnv("POSTMARK_SENDER").Default("rama@prismapp.io").ToString()
	subject = environ.GetEnv("POSTMARK_SUBJECT").Default("One Time Password").ToString()
}

func initClient() {
	once.Do(func() {
		client = postmark.NewClient(serverToken, accountToken)
	})
}
