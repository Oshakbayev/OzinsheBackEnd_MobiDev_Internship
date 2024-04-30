package helpers

import (
	"math/rand"
	"os"
	"ozinshe/pkg/entity"
	"strconv"
	"strings"
	"time"
)

func GenerateRandomKey(length int) string {
	// Use current time as the seed
	rand.New(rand.NewSource(time.Now().UnixNano()))
	key := make([]byte, length)
	for i := range key {
		key[i] = entity.Charset[rand.Intn(len(entity.Charset))]
	}
	return string(key)
}

func StrToIntArr(str string) []int {
	var intArr []int
	str = strings.Trim(str, "{}NULL")
	for _, val := range strings.Split(str, ",") {
		if id, err := strconv.Atoi(val); err == nil {
			intArr = append(intArr, id)
		}
	}
	return intArr
}

func StrToStrArr(str string) []string {
	var strArr []string
	str = strings.Trim(str, "{}NULL")
	if str != "" {
		strArr = strings.Split(str, ",")
	}
	return strArr
}

func DeleteFile(filePath string) error {
	err := os.Remove(filePath)
	if err != nil {
		return err
	}
	return nil
}

func GeneratePassword() string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	lowercaseLetters := "abcdefghijklmnopqrstuvwxyz"
	uppercaseLetters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits := "0123456789"
	specialChars := "@$!&"
	allChars := lowercaseLetters + uppercaseLetters + digits + specialChars
	password := ""
	password += string(uppercaseLetters[rand.Intn(len(uppercaseLetters))])
	for i := 0; i < 6; i++ {
		password += string(allChars[rand.Intn(len(allChars))])
	}
	password += string(specialChars[rand.Intn(len(specialChars))])
	return password
}
