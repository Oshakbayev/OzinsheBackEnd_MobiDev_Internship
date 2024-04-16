package helpers

import (
	"math/rand"
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
