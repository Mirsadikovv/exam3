package check

import (
	"errors"
	"regexp"
	"time"
)

func ValidatePhone(phone string) error {
	phoneRegex := regexp.MustCompile(`^\+998[0-9]{9}$`)
	if !phoneRegex.MatchString(phone) {
		return errors.New("phone is not valid")
	}
	return nil
}

func ValidateMail(mail string) error {
	mailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@(gmail\.com|mail\.ru)$`)
	if mailRegex.MatchString(mail) {
		return errors.New("mail is not valid")
	}
	return nil
}

func ValidateBitrthday(birthday string, age int) error {
	layout := "2006-01-02"

	date, err := time.Parse(layout, birthday)
	if err != nil {
		return errors.New("wrong date format")
	}
	today := time.Now()

	duration := today.Sub(date)
	years := int(duration.Hours() / (24 * 365))
	if years != age {
		return errors.New("age and birthday not equal")
	}
	return nil
}
