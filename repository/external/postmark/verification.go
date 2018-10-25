package postmark

import (
	"errors"
	"fmt"
	"strings"

	depsPostmark "github.com/keighl/postmark"
	"github.com/labstack/echo"
	"github.com/ramabmtr/heimdall/helper/log"
	"github.com/ramabmtr/heimdall/model"
)

type (
	nexmoAuthRepository struct {
		ctx echo.Context
		db  model.VerificationDatabaseRepository
	}
)

func NewVerificationRepository(ctx echo.Context, db model.VerificationDatabaseRepository) model.VerificationRepository {
	initClient()

	return &nexmoAuthRepository{
		ctx: ctx,
		db:  db,
	}
}

func (c *nexmoAuthRepository) SendVerificationCode(v *model.Verification) error {
	if client == nil {
		return errPostmarkNotInitialized
	}

	email, ok := v.SendTo.(string)
	if !ok {
		return errors.New("fail to parse send_to param")
	}

	// save to db first to avoid sms sent but fail to save in db
	if err := c.db.Store(email, v.VerificationCode); err != nil {
		log.GetLogger(c.ctx).
			WithError(err).
			WithField("to", email).
			Error("error save to database")
		return err
	}

	message := strings.Replace(v.TemplateMessage, "{code}", v.VerificationCode, 1)

	content := depsPostmark.Email{
		From:     sender,
		To:       email,
		Subject:  subject,
		TextBody: message,
		Tag:      "OTP",
	}

	resp, err := client.SendEmail(content)

	fmt.Println(resp)

	if err != nil {
		log.GetLogger(c.ctx).
			WithError(err).
			WithField("to", email).
			Error("error send verification code")
		c.db.Delete(email)
		return err
	}

	log.GetLogger(c.ctx).
		WithField("to", email).
		Debug("verification code sent")

	return nil
}

func (c *nexmoAuthRepository) CheckVerificationCode(check interface{}, code string) (bool, error) {
	email, ok := check.(string)
	if !ok {
		return false, errors.New("fail to parse send_to param")
	}

	val, err := c.db.Get(email)
	if err != nil {
		log.GetLogger(c.ctx).
			WithError(err).
			WithField("to", email).
			Error("error read from database")
		return false, err
	}

	if val != code {
		return false, nil
	}

	c.db.Delete(email)

	return true, nil
}
