package watoken

import (
	"fmt"
	"testing"
)

/* func TestEncode(t *testing.T) {
	privkey := os.Getenv("PRIVATEKEY")
	str, _ := EncodeforHours(os.Getenv("PHONENUMBER"), "Helpdesk Pamong Desa", privkey, 43830)
	println(str)
	//atr, _ := DecodeGetId("", str)
	//println(atr)

} */

func TestWatoken(t *testing.T) {
	privateKey, publicKey := GenerateKey()
	println("PRIVATEKEY:")
	println(privateKey)
	println("PUBLICKEY")
	println(publicKey)

	//generate token for user awangga
	userid := "awangga"
	tokenstring, _ := Encode(userid, privateKey)

	//decode token to get userid
	useridstring, _ := DecodeGetId(publicKey, tokenstring)
	if useridstring == "" {
		fmt.Println("expire token")
	} else {
		println("id:" + useridstring)
	}

}
