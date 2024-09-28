package validator

import (
	"regexp"
	"unicode"

	"github.com/go-playground/validator/v10"
	"github.com/sail-host/cloud/internal/global"
)

func Init() {
	validator := validator.New()
	if err := validator.RegisterValidation("name", checkNamePattern); err != nil {
		panic(err)
	}
	if err := validator.RegisterValidation("ip", checkIpPattern); err != nil {
		panic(err)
	}
	if err := validator.RegisterValidation("password", checkPasswordPattern); err != nil {
		panic(err)
	}
	if err := validator.RegisterValidation("email", checkEmailPattern); err != nil {
		panic(err)
	}
	global.VALID = validator
	global.ECHO.Validator = &CustomValidator{validator: validator}
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func checkNamePattern(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	result, err := regexp.MatchString("^[a-zA-Z\u4e00-\u9fa5]{1}[a-zA-Z0-9_\u4e00-\u9fa5]{0,30}$", value)
	if err != nil {
		global.LOG.Errorf("regexp matchString failed, %v", err)
	}
	return result
}

func checkIpPattern(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	result, err := regexp.MatchString(`^((2(5[0-5]|[0-4]\d))|[0-1]?\d{1,2})(\.((2(5[0-5]|[0-4]\d))|[0-1]?\d{1,2})){3}$`, value)
	if err != nil {
		global.LOG.Errorf("regexp check ip matchString failed, %v", err)
	}
	return result
}

func checkPasswordPattern(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	if len(value) < 8 || len(value) > 30 {
		return false
	}

	hasNum := false
	hasLetter := false
	for _, r := range value {
		if unicode.IsLetter(r) && !hasLetter {
			hasLetter = true
		}
		if unicode.IsNumber(r) && !hasNum {
			hasNum = true
		}
		if hasLetter && hasNum {
			return true
		}
	}

	return false
}

func checkEmailPattern(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	result, err := regexp.MatchString(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, value)
	if err != nil {
		global.LOG.Errorf("regexp check email matchString failed, %v", err)
	}
	return result
}
