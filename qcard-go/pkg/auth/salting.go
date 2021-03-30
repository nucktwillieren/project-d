package auth

import (
	"bytes"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"golang.org/x/crypto/pbkdf2"
)

func GetRandomString(len int) string {
	var container string
	var str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := bytes.NewBufferString(str)
	length := b.Len()
	bigInt := big.NewInt(int64(length))
	for i := 0; i < len; i++ {
		randomInt, _ := rand.Int(rand.Reader, bigInt)
		container += string(str[randomInt.Int64()])
	}
	return container
}

func Salting(password string, iterations int, salt string) (string, error) {
	if strings.TrimSpace(salt) == "" {
		salt = GetRandomString(12)
	}

	if strings.Contains(salt, "$") {
		return "", errors.New("salt contains dollar sign ($)")
	}

	if iterations <= 0 {
		iterations = 180000
	}

	hash := pbkdf2.Key([]byte(password), []byte(salt), iterations, sha256.Size, sha256.New)

	b64Hash := base64.StdEncoding.EncodeToString(hash)

	return fmt.Sprintf("%s$%d$%s$%s", "pbkdf2_sha256", iterations, salt, b64Hash), nil
}

func SaltingVerify(password string, encoded string) (bool, error) {
	s := strings.Split(encoded, "$")

	if len(s) != 4 {
		return false, errors.New("hashed password components mismatch")
	}

	algorithm, iterations, salt := s[0], s[1], s[2]

	if algorithm != "pbkdf2_sha256" {
		return false, errors.New("algorithm mismatch")
	}

	i, err := strconv.Atoi(iterations)
	if err != nil {
		return false, errors.New("unreadable component in hashed password")
	}

	newEncoded, err := Salting(password, i, salt)
	if err != nil {
		return false, err
	}

	return hmac.Equal([]byte(newEncoded), []byte(encoded)), nil
}
