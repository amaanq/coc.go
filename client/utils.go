package client

import (
	"fmt"
	"regexp"
	"strings"
)

// Credit: https://github.com/mathsman5133/coc.py/blob/master/coc/utils.py
func CorrectTag(tag string) string {
	re := regexp.MustCompile("[^A-Z0-9]+")
	tag = "#" + strings.ReplaceAll(re.ReplaceAllString(strings.ToUpper(tag), ""), "O", "0")
	return tag
}

func parseArgs(args []map[string]string) string {
	if len(args) == 0 {
		return ""
	}
	params := ""

	if len(args) > 0 {
		for _, arg := range args {
			for key, val := range arg {
				params += fmt.Sprintf("%s=%s&", key, val)
			}
		}
	}

	params = params[:len(params)-1]
	if params != "" {
		params = "?" + params
	}
	return params
}
