package unpacking

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnpackString(t *testing.T) {
	cases := []struct {
		Input  string
		Result string
		Error  string
	}{
		{
			Input:  "a4bc2d5e",
			Result: "aaaabccddddde",
			Error:  "",
		},
		{
			Input:  "abcd",
			Result: "abcd",
			Error:  "",
		},
		{
			Input:  "",
			Result: "",
			Error:  "",
		},
		{
			Input:  "qwe\\45",
			Result: "qwe44444",
			Error:  "",
		},
		{
			Input:  "qwe\\4\\5",
			Result: "qwe45",
			Error:  "",
		},
		{
			Input:  "🥸🤩🥳11\\9",
			Result: "🥸🤩🥳🥳🥳🥳🥳🥳🥳🥳🥳🥳🥳9",
			Error:  "",
		},
		{
			Input:  "a",
			Result: "a",
			Error:  "",
		},
		{
			Input:  "\\9",
			Result: "9",
			Error:  "",
		},
		{
			Input:  "a1",
			Result: "a",
			Error:  "",
		},
		{
			Input:  "a10",
			Result: "aaaaaaaaaa",
			Error:  "",
		},
		{
			Input:  " a10 b",
			Result: " aaaaaaaaaa b",
			Error:  "",
		},
		{
			Input:  "a1000",
			Result: strings.Repeat("a", 1000),
			Error:  "",
		},
		{
			Input:  "45",
			Result: "",
			Error:  "string must starts with a letter or escaped number",
		},
		{
			Input:  "\\45adf\\",
			Result: "",
			Error:  "trailing backslash",
		},
	}

	for _, item := range cases {
		s, err := UnpackString(item.Input)

		assert.Equal(t, item.Result, s)
		if err != nil {
			assert.Equal(t, item.Error, err.Error())
		}
	}
}
