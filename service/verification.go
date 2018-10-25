package service

import (
	"math/rand"

	"github.com/labstack/echo"
	"github.com/ramabmtr/heimdall/model"
)

const codeNumber = "1234567890"

type (
	VerificationService struct {
		ctx    echo.Context
		verify model.VerificationRepository
	}
)

func NewVerificationService(
	ctx echo.Context,
	verify model.VerificationRepository,
) *VerificationService {
	return &VerificationService{
		ctx:    ctx,
		verify: verify,
	}
}

func randCode(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = codeNumber[rand.Int63()%int64(len(codeNumber))]
	}
	return string(b)
}

func (c *VerificationService) SendVerificationCode(sendTo interface{}) error {
	verificationCode := randCode(4)
	v := model.Verification{
		SendTo:           sendTo,
		TemplateMessage:  "Hi, this is your verification code: {code}",
		VerificationCode: verificationCode,
	}

	if err := c.verify.SendVerificationCode(&v); err != nil {
		return err
	}

	return nil
}

func (c *VerificationService) CheckVerificationCode(checkKey interface{}, code string) (bool, error) {
	return c.verify.CheckVerificationCode(checkKey, code)
}
