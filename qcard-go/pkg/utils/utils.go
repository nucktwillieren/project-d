package utils

import (
	"math/rand"
	"reflect"
)

func GetRandomString(n int) string {
	letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)

	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func StrIn(s string, a []string) bool {
	for _, v := range a {
		if v == s {
			return true
		}
	}
	return false
}

func PointerCheck(v interface{}) bool {
	return reflect.ValueOf(v).Kind() == reflect.Ptr
}
