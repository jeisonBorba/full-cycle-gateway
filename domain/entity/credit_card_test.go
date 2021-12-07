package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreditCardNumber(t *testing.T) {
	_, err := NewCreditCard("400000000000000", "Jose da Silva", 12, 2022, 123)
	assert.Equal(t, "invalid credit card number", err.Error())

	_, err = NewCreditCard("4193523830170205", "Jose da Silva", 12, 2022, 123)
	assert.Nil(t, err)
}

func TestCreditCardExpirationMonth(t *testing.T) {
	_, err := NewCreditCard("4193523830170205", "Jose da Silva", 13, 2022, 123)
	assert.Equal(t, "invalid credit card expiration month", err.Error())

	_, err = NewCreditCard("4193523830170205", "Jose da Silva", 0, 2022, 123)
	assert.Equal(t, "invalid credit card expiration month", err.Error())

	_, err = NewCreditCard("4193523830170205", "Jose da Silva", 1, 2022, 123)
	assert.Nil(t, err)
}

func TestCreditCardExpirationYear(t *testing.T) {
	lastYear := time.Now().AddDate(-1, 0, 0)
	_, err := NewCreditCard("4193523830170205", "Jose da Silva", 12, lastYear.Year(), 123)
	assert.Equal(t, "invalid credit card expiration year", err.Error())

	lastYear = time.Now().AddDate(1, 0, 0)
	_, err = NewCreditCard("4193523830170205", "Jose da Silva", 12, lastYear.Year(), 123)
	assert.Nil(t, err)
}
