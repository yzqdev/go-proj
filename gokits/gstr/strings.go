package gstr

import "regexp"

const regexStrictEmailPattern = `(?i)[A-Z0-9!#$%&'*+/=?^_{|}~-]+` +
	`(?:\.[A-Z0-9!#$%&'*+/=?^_{|}~-]+)*` +
	`@(?:[A-Z0-9](?:[A-Z0-9-]*[A-Z0-9])?\.)+` +
	`[A-Z0-9](?:[A-Z0-9-]*[A-Z0-9])?`

var (
	digitRegexp             = regexp.MustCompile(`^[0-9]+$`)
	chLetterNumUnlineRegexp = regexp.MustCompile(`^[\x{4e00}-\x{9fa5}_a-zA-Z0-9]+$`)
	letterNumUnlineRegexp   = regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
	chineseRegexp           = regexp.MustCompile(`^[\x{4e00}-\x{9fa5}]+$`)
	emailRegexp             = regexp.MustCompile(`(?i)[A-Z0-9._%+-]+@(?:[A-Z0-9-]+\.)+[A-Z]{2,6}`)
	strictEmailRegexp       = regexp.MustCompile(regexStrictEmailPattern)
	urlRegexp               = regexp.MustCompile(`(ftp|http|https):\/\/(\w+:{0,1}\w*@)?(\S+)(:[0-9]+)?(\/|\/([\w#!:.?+=&%@!\-\/]))?`)
)

// IsDigit ...
func IsDigit(str string) bool {
	return digitRegexp.MatchString(str)
}

// IsChineseLetterNumUnline ...
func IsChineseLetterNumUnline(str string) bool {
	return chLetterNumUnlineRegexp.MatchString(str)
}

// IsLetterNumUnline ...
func IsLetterNumUnline(str string) bool {
	return letterNumUnlineRegexp.MatchString(str)
}

// IsChinese ...
func IsChinese(str string) bool {
	return chineseRegexp.MatchString(str)
}

// IsEmail ...
func IsEmail(str string) bool {
	return emailRegexp.MatchString(str)
}

// IsEmailRFC ...
func IsEmailRFC(str string) bool {
	return strictEmailRegexp.MatchString(str)
}

// IsURL ...
func IsURL(str string) bool {
	return urlRegexp.MatchString(str)
}

// IfEmpty ...
func IfEmpty(a, b string) string {
	if a == "" {
		return b
	}
	return a
}
