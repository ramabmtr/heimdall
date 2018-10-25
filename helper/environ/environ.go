package environ

import (
	"os"
	"strconv"
)

type Environment struct {
	value string
}

func GetEnv(key string) Environment {
	e := Environment{}
	if value, ok := os.LookupEnv(key); ok {
		e.value = value
	}

	return e
}

func SetEnv(key, val string) {
	if err := os.Setenv(key, val); err != nil {
		return
	}
}

func (e Environment) Default(data string) Environment {
	if e.value != "" {
		return e
	}

	e.value = data
	return e
}

func (e Environment) ToString() string {
	return e.value
}

func (e Environment) ToInt() int {
	v, err := strconv.Atoi(e.value)
	if err == nil {
		return v
	}

	return int(0)
}

func (e Environment) ToFloat64() float64 {
	v, err := strconv.ParseFloat(e.value, 64)
	if err == nil {
		return v
	}

	return float64(0)
}

func (e Environment) ToBool() bool {
	v, err := strconv.ParseBool(e.value)
	if err == nil {
		return v
	}
	return true
}
