package tag

import (
	"net/url"
	"regexp"
	"strings"
)

type PlayerTag string

func ToPlayerTag(tag string) PlayerTag {
	return PlayerTag(tag).CorrectTag()
}

// Credit: https://github.com/mathsman5133/coc.py/blob/master/coc/utils.py
func (p PlayerTag) CorrectTag() PlayerTag {
	re := regexp.MustCompile("[^A-Z0-9]+")
	p = PlayerTag("#" + strings.ReplaceAll(re.ReplaceAllString(strings.ToUpper(string(p)), ""), "O", "0"))
	return p
}

func (p PlayerTag) OpenInGameURL() string {
	return "https://link.clashofclans.com/en?action=OpenPlayerProfile&tag=" + url.PathEscape(string(p))
}

type ClanTag string

func ToClanTag(tag string) ClanTag {
	return ClanTag(tag).CorrectTag()
}

// Credit: https://github.com/mathsman5133/coc.py/blob/master/coc/utils.py
func (c ClanTag) CorrectTag() ClanTag {
	re := regexp.MustCompile("[^A-Z0-9]+")
	c = ClanTag("#" + strings.ReplaceAll(re.ReplaceAllString(strings.ToUpper(string(c)), ""), "O", "0"))
	return c
}

func (c ClanTag) OpenInGameURL() string {
	return "https://link.clashofclans.com/en?action=OpenClanProfile&tag=" + url.PathEscape(string(c))
}
