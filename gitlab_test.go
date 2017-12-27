package main

import (
	"errors"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExpiryWithDate(t *testing.T) {
	date, err := addExpiry(url.Values{}, "2016-11-03")
	assert.Nil(t, err)
	assert.Equal(t, 1, len(date))
	assert.Contains(t, date, "personal_access_token[expires_at]")
}

func TestExpiryWithEmptyDate(t *testing.T) {
	date, err := addExpiry(url.Values{}, "")
	assert.Nil(t, err)
	assert.Empty(t, date)
}

func TestExpiryWithWrongDateAtYear(t *testing.T) {
	date, err := addExpiry(url.Values{}, "201-11-01")
	assert.NotNil(t, err)
	assert.Equal(t, errors.New("Year is too short"), err)
	assert.Empty(t, date)
}

func TestExpiryWithWrongDateAtMonth(t *testing.T) {
	date, err := addExpiry(url.Values{}, "2016-1-01")
	assert.NotNil(t, err)
	assert.Equal(t, errors.New("Month is too short needs to be with zero digest"), err)
	assert.Empty(t, date)
}

func TestExpiryWithWrongDateAtDay(t *testing.T) {
	date, err := addExpiry(url.Values{}, "2016-11-1")
	assert.NotNil(t, err)
	assert.Equal(t, errors.New("Day is too short needs to be with zero digest"), err)
	assert.Empty(t, date)
}

func TestExpiryWithWrongDateFormat(t *testing.T) {
	date, err := addExpiry(url.Values{}, "2016/11/01")
	assert.NotNil(t, err)
	assert.Equal(t, errors.New("Date is too short it should be formatted like this 2017-12-03"), err)
	assert.Empty(t, date)
}

func TestAccessToken(t *testing.T) {
	request := GitLabTokenRequest{
		URL:       "http://localhost:10080",
		Username:  "root",
		Password:  "abcd1234",
		Scope:     Scope{API: true},
		Date:      "",
		TokenName: "gogpat-test",
	}
	token, err := CreateToken(request)
	assert.Nil(t, err)
	assert.NotEmpty(t, token)
	assert.Equal(t, 20, len(token))
}

func TestFullAccessToken(t *testing.T) {
	request := GitLabTokenRequest{
		URL:       "http://localhost:10080",
		Username:  "root",
		Password:  "abcd1234",
		Scope:     Scope{API: true, ReadRegistry: false, ReadUser: true, Sudo: true},
		Date:      "",
		TokenName: "gogpat-test",
	}
	token, err := CreateToken(request)
	assert.Nil(t, err)
	assert.NotEmpty(t, token)
	assert.Equal(t, 20, len(token))
}
