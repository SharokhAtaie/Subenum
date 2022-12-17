package functions

import (
	"github.com/tidwall/gjson"
	"regexp"
	"strings"
	"github.com/SharokhAtaie/Subenum/commands"
	"github.com/SharokhAtaie/Subenum/requests"
)

// crt.sh
func CRtsh(url string) string {
	data := requests.CRTSH(url)
	value := gjson.Get(data, "#.common_name")
	crtCom := commands.Beauty(value.String())
	return crtCom
}

func Jldcme(url string) string {
	// jldc.me
	jldc := requests.JldcMe(url)
	jldcbeauty := commands.Beauty(jldc)
	return jldcbeauty
}

func AbuseBeauty(url string) []string {
	abuse := requests.AbuseDB(url)
	pattern, _ := regexp.Compile(`<li>\w.*</li>`)
	allmatch := pattern.FindAllString(abuse, -1)
	return allmatch
}

func AbuseDb(url string) string {
	data := AbuseBeauty(url)
	str := strings.Join(data, " ")
	pattern, _ := regexp.Compile(`</?li>`)
	allmatch := pattern.ReplaceAllString(str, "\n")
	space := strings.ReplaceAll(allmatch, "\n", "")
	space2 := strings.ReplaceAll(space, " ", "\n")
	return space2
}

func GetSubstring(s string, indices []int) string {
	return string(s[indices[0]:indices[1]])
}
