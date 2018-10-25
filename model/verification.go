package model

type (
	Verification struct {
		SendTo           interface{}
		TemplateMessage  string
		VerificationCode string
	}

	VerificationRepository interface {
		SendVerificationCode(verification *Verification) (err error)
		CheckVerificationCode(check interface{}, code string) (verified bool, err error)
	}

	VerificationDatabaseRepository interface {
		Get(key string) (value string, err error)
		Store(key, value string) (err error)
		Delete(key string)
	}
)
