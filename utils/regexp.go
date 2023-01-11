package utils

import "regexp"

func CheckRegex(reg, s string) bool {
	return regexp.MustCompile(reg).Match([]byte(s))
}

func RetrieveRegexp(reg, s string) string {
	r, err := regexp.Compile(reg)
	if err != nil {
		return "none"
	}
	v := r.FindString(s)
	return v
}

func RegexpArray(reg, s string) []string {
	arr := regexp.MustCompile(reg).Split(s, -1)
	return arr
}
