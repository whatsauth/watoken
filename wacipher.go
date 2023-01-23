package watoken

import (
	"encoding/base64"
	"math/rand"
	"strings"
	"time"
)

func GetAppUrl(wuid string) string {
	ss := strings.Split(wuid, ".")
	s := ss[len(ss)-1]
	base64Text := make([]byte, base64.StdEncoding.DecodedLen(len(s)))
	l, _ := base64.StdEncoding.Decode(base64Text, []byte(s))
	return string(base64Text[:l])
}

func GetAppInfo(wuid string) (protocol, hostname, pathname string) {
	appurl := GetAppUrl(wuid)
	var hostandpath string
	protocol, hostandpath, _ = strings.Cut(appurl, "://")
	hostname, pathname, _ = strings.Cut(hostandpath, "/")
	return
}

func GetAppSubDomain(wuid string) (subdomain string) {
	_, hostname, _ := GetAppInfo(wuid)
	subdomain, _, _ = strings.Cut(hostname, ".")
	return
}

func RandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return b.String()
}
