package formatter

import (
	"bytes"
	"strings"
)

func CacheKey(args ...string) string {
	const redisPrefix string = "heimdall"

	var buffer bytes.Buffer

	buffer.WriteString(redisPrefix)

	for _, arg := range args {
		buffer.WriteString("::") // separator
		buffer.WriteString(arg)
	}

	return buffer.String()
}

func Phone(countryCode, phoneNumber string) string {
	var buffer bytes.Buffer

	if !strings.HasPrefix(countryCode, "+") {
		buffer.WriteString("+")
	}

	buffer.WriteString(countryCode)

	if strings.HasPrefix(phoneNumber, "0") {
		phoneNumber = phoneNumber[1:]
	}

	buffer.WriteString(phoneNumber)

	return buffer.String()
}
