package util

import (
	"regexp"
)

func checkMobileReg(mobile string, regStr string) bool {
	if len(mobile) == 0 {
		return false
	}

	reg := regexp.MustCompile(regStr)
	return reg.MatchString(mobile)
}

// CheckMobile 手机号格式验证，正确 true, 失败 false
func CheckMobile(mobile string) bool {

	// 中国大陆
	if checkMobileReg(mobile, `^(1(3|4|5|6|7|8)\d{9})$`) {
		return true
	}

	// 中国香港
	if checkMobileReg(mobile, `^(852(5|6|8|9)\d{7})$`) {
		return true
	}

	// 中国澳门
	if checkMobileReg(mobile, `^(853(5|6|8|9)\d{7})$`) {
		return true
	}

	// 中国澳门
	if checkMobileReg(mobile, `^(853(5|6|8|9)\d{7})$`) {
		return true
	}

	// 中国台湾
	if checkMobileReg(mobile, `^(88609\d{8})$`) {
		return true
	}

	// 美国/加拿大
	if checkMobileReg(mobile, `^(001\d{10})$`) {
		return true
	}

	// 英国
	if checkMobileReg(mobile, `^(447\d{9})$`) {
		return true
	}

	// 新加坡
	if checkMobileReg(mobile, `^(65(8|9)\d{7})$`) {
		return true
	}

	return false
}
