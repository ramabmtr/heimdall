package twilio

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/ramabmtr/heimdall/helper/formatter"
	"github.com/ramabmtr/heimdall/helper/log"
	"github.com/ramabmtr/heimdall/model"
)

type (
	twilioAuthRepository struct {
		ctx echo.Context
		db  model.VerificationDatabaseRepository
	}

	SendTo struct {
		CountryCode string `json:"country_code" validate:"required"`
		PhoneNumber string `json:"phone_number" validate:"required"`
	}

	CheckKey SendTo
)

func NewVerificationRepository(ctx echo.Context, db model.VerificationDatabaseRepository) model.VerificationRepository {
	initClient()

	return &twilioAuthRepository{
		ctx: ctx,
		db:  db,
	}
}

func (c *twilioAuthRepository) SendVerificationCode(v *model.Verification) error {
	if client == nil {
		return errTwilioNotInitialized
	}

	validate := validator.New()

	sendTo := new(SendTo)

	b, _ := json.Marshal(v.SendTo)
	if err := json.Unmarshal(b, &sendTo); err != nil {
		return err
	}

	if err := validate.Struct(sendTo); err != nil {
		return err
	}

	phone := formatter.Phone(sendTo.CountryCode, sendTo.PhoneNumber)

	// save to db first to avoid sms sent but fail to save in db
	if err := c.db.Store(phone, v.VerificationCode); err != nil {
		log.GetLogger(c.ctx).
			WithError(err).
			WithField("to", phone).
			Error("error save to database")
		return err
	}

	from := twilioNumber
	to := phone
	message := strings.Replace(v.TemplateMessage, "{code}", v.VerificationCode, 1)
	_, e, err := client.SendSMS(from, to, message, "", "")
	if err != nil {
		log.GetLogger(c.ctx).
			WithError(err).
			WithField("to", phone).
			Error("error send verification code")
		c.db.Delete(phone)
		return err
	}

	if e != nil {
		err = errors.New(e.Message)
		log.GetLogger(c.ctx).
			WithError(err).
			WithField("to", phone).
			Error("fail send verification code")
		c.db.Delete(phone)
		return err
	}

	log.GetLogger(c.ctx).
		WithField("to", phone).
		Debug("verification code sent")

	return nil
}

func (c *twilioAuthRepository) CheckVerificationCode(check interface{}, code string) (bool, error) {
	validate := validator.New()

	checkKey := new(CheckKey)

	b, _ := json.Marshal(check)
	if err := json.Unmarshal(b, &checkKey); err != nil {
		return false, err
	}

	if err := validate.Struct(checkKey); err != nil {
		return false, err
	}

	phone := formatter.Phone(checkKey.CountryCode, checkKey.PhoneNumber)

	val, err := c.db.Get(phone)
	if err != nil {
		log.GetLogger(c.ctx).
			WithError(err).
			WithField("to", phone).
			Error("error read from database")
		return false, err
	}

	if val != code {
		return false, nil
	}

	c.db.Delete(phone)

	return true, nil
}
