package utils

import (
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
)

func RandomString(len int) string {
	b := make([]byte, len)       //随机生成字符数组
	rand.Read(b)                 //整合
	str := hex.EncodeToString(b) //转换为string
	return fmt.Sprintf("%s", str)
}

func StringPort2Int(port string) int {
	arr := strings.Split(port, ":")
	PortInt, err := strconv.Atoi(arr[len(arr)-1])
	if err != nil {
		log.Println("[Port To Int Err]", err)
	}
	return PortInt
}

// RandomStringUsingChars 从chars中选取字符生成长度为length的字符串，不适用于含中文等特殊字符
func RandomStringUsingChars(length int, chars string) string {
	if chars == "" {
		return RandomString(length)
	}
	res := make([]uint8, length)
	for i := 0; i < length; i++ {
		randIdx := rand.Intn(len(chars))
		res[i] = chars[randIdx]
	}
	return string(res)
}

func TriPrefixAndSuffix(origin string, target string) string {
	origin = strings.TrimPrefix(origin, target)
	origin = strings.TrimRight(origin, target)
	return origin
}
