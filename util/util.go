package util

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

func Md5(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str))) //将[]byte转成16进制
}

func Sha1(str string) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(str))) //将[]byte转成16进制
}

// Sha256 生成哈希值
func Sha256(message []byte) string {
	hash := sha256.New()
	hash.Write(message)
	bytes := hash.Sum(nil)
	hashCode := hex.EncodeToString(bytes)
	return hashCode
}

// UnescapeUnicode unicode转utf8，用于json中unicode转义
func UnescapeUnicode(raw []byte) ([]byte, error) {
	str, err := strconv.Unquote(strings.Replace(strconv.Quote(string(raw)), `\\u`, `\u`, -1))
	if err != nil {
		return nil, err
	}
	return []byte(str), nil
}
