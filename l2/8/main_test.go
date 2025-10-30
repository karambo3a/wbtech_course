package main

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Address string
	Error   error
}

func TestGetCurrTime(t *testing.T) {
	cases := []TestCase{
		{
			Address: "0.beevik-ntp.pool.ntp.org",
			Error:   nil,
		},
		{
			Address: "",
			Error:   errors.New("address string is empty"),
		},
		{
			Address: "abc",
			Error:   errors.New("lookup abc: no such host"),
		},
	}

	for _, item := range cases {
		startTime := time.Now()
		currTime, err := GetCurrTime(item.Address)

		assert.False(t, currTime.IsZero())
		assert.True(t, currTime.Before(startTime.Add(time.Second)))
		if err != nil {
			assert.Equal(t, err.Error(), item.Error.Error())
		}
	}
}
