package validation

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPhone(t *testing.T) {
	testTable := []struct {
		number   int64
		allowNil bool
		want     bool
	}{
		{
			number:   79282638428,
			allowNil: true,
			want:     true,
		},
		{
			number:   78282638428,
			allowNil: true,
			want:     false,
		},
		{
			number:   78282638428,
			allowNil: false,
			want:     false,
		},
		{
			number:   79282638428,
			allowNil: false,
			want:     true,
		},
		{
			number:   0,
			allowNil: false,
			want:     false,
		},
	}

	for index, testCase := range testTable {
		t.Run(strconv.Itoa(index+1), func(t *testing.T) {
			result := Phone(&testCase.number, testCase.allowNil)

			assert.Equal(t, testCase.want, result)
		})
	}

}

func TestEmail(t *testing.T) {
	testTable := []struct {
		email    string
		allowNil bool
		want     bool
		msg      string
	}{
		{
			email:    "TAR@mail.ru",
			allowNil: true,
			want:     true,
		},
		{
			email:    "TARmail.ru",
			allowNil: true,
			want:     false,
		},
		{
			email:    "TAR@gmailcom",
			allowNil: true,
			want:     false,
		},
		{
			email:    "TAR@mail.ru",
			allowNil: false,
			want:     true,
		},
		{
			email:    "TAR@gmailcom",
			allowNil: false,
			want:     false,
		},
		{
			email:    "",
			allowNil: false,
			want:     false,
		},
		{
			email:    "TAR@mail.ru.ru",
			allowNil: false,
			want:     true,
		},
	}

	for index, testCase := range testTable {
		t.Run(strconv.Itoa(index+1), func(t *testing.T) {
			result := Email(&testCase.email, testCase.allowNil)

			assert.Equal(t, testCase.want, result)
		})
	}

}

func TestPassword(t *testing.T) {
	testTable := []struct {
		password string
		want     bool
	}{
		{
			password: "123456a",
			want:     true,
		},
		{
			password: "123456",
			want:     false,
		},
		{
			password: "a1!",
			want:     false,
		},
		{
			password: "a1!12312312",
			want:     true,
		},
		{
			password: "aaaaaaa",
			want:     false,
		},
		{
			password: "11111111",
			want:     false,
		},
		{
			password: "111@1a1.1",
			want:     true,
		},
	}

	for index, testCase := range testTable {
		t.Run(strconv.Itoa(index+1), func(t *testing.T) {
			result := Password(&testCase.password)

			assert.Equal(t, testCase.want, result)
		})
	}

}
