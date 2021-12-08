package entity

import (
	"errors"
	"regexp"
	"time"
)

type CreditCard struct {
	number          string
	name            string
	expirationMonth int
	expirtationYear int
	cvv             int
}

func NewCreditCard(number string, name string, expirationMonth int, expirationYear int, cvv int) (*CreditCard, error) {
	creditCard := &CreditCard{
		number:          number,
		name:            name,
		expirationMonth: expirationMonth,
		expirtationYear: expirationYear,
		cvv:             cvv,
	}

	err := creditCard.IsValid()
	if err != nil {
		return nil, err
	}

	return creditCard, nil
}

func (cc *CreditCard) IsValid() error {
	err := cc.validateNumber()
	if err != nil {
		return err
	}

	err = cc.validateMonth()
	if err != nil {
		return err
	}

	err = cc.validateYear()
	if err != nil {
		return err
	}

	return nil
}

func (cc *CreditCard) validateNumber() error {
	reg := regexp.MustCompile(`^(?:4[0-9]{12}(?:[0-9]{3})?|[25][1-7][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\d{3})\d{11})$`)
	if !reg.MatchString(cc.number) {
		return errors.New("invalid credit card number")
	}
	return nil
}

func (cc *CreditCard) validateMonth() error {
	if cc.expirationMonth < 1 || cc.expirationMonth > 12 {
		return errors.New("invalid credit card expiration month")
	}
	return nil
}

func (cc *CreditCard) validateYear() error {
	if cc.expirtationYear < time.Now().Year() {
		return errors.New("invalid credit card expiration year")
	}
	return nil
}
