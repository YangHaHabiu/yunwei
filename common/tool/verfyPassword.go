package tool

import (
	"regexp"
)

// 密码校验规则: 必须包含数字、大写字母、小写字母、特殊字符(如.@$!%*#_~?&^)至少3种的组合且长度在8-18之间
func VerifyPassword(minLength, maxLength int, pwd string) bool {
	if len(pwd) < minLength || len(pwd) > maxLength {
		return false
	}
	// 过滤掉这四类字符以外的密码串,直接判断不合法
	re, err := regexp.Compile(`^[a-zA-Z0-9.@$!%*#_~?&^]{8,18}$`)
	if err != nil {
		return false
	}
	match := re.MatchString(pwd)
	if !match {
		return false
	}

	var level = 0
	patternList := []string{`[0-9]+`, `[a-z]+`, `[A-Z]+`, `[.@$!%*#_~?&^]+`}
	for _, pattern := range patternList {
		match, _ := regexp.MatchString(pattern, pwd)
		if match {
			level++
		}
	}
	if level < 3 {
		return false
	}

	return true
}

func EncodeGuildMessageID(gid, uid, seq, rand, time int64) {

}
