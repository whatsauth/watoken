package watoken

import (
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWatoken(t *testing.T) {
	privateKey, publicKey := GenerateKey()
	fmt.Println("privateKey : ", privateKey)
	fmt.Println("publicKey : ", publicKey)
	userid := "awangga"

	tokenstring, err := Encode(userid, privateKey)
	require.NoError(t, err)
	body, err := Decode(publicKey, tokenstring)
	fmt.Println("signed : ", tokenstring)
	fmt.Println("isi : ", body)
	require.NoError(t, err)
}

func TestWatokenStruct(t *testing.T) {
	privateKey, publicKey := GenerateKey()
	fmt.Println("privateKey : ", privateKey)
	fmt.Println("publicKey : ", publicKey)
	userid := "awangga"

	type dataA struct {
		Name string `json:"name"`
	}

	data := dataA{
		Name: "awangga",
	}

	tokenstring, err := EncodeWithStruct(userid, &data, privateKey)
	require.NoError(t, err)
	body, err := DecodeWithStruct[dataA](publicKey, tokenstring)
	fmt.Println("signed : ", tokenstring)
	fmt.Println("isi : ", body)
	fmt.Println("data di Payload : ", body.Data)
	require.NoError(t, err)
}

func TestWatokenDuration(t *testing.T) {
	privateKey, publicKey := GenerateKey()
	fmt.Println("privateKey : ", privateKey)
	fmt.Println("publicKey : ", publicKey)
	userid := "awangga"

	tokenstring, err := EncodeforSeconds(userid, privateKey, 60)
	require.NoError(t, err)
	body, err := Decode(publicKey, tokenstring)
	fmt.Println("tokenstring duration : ", tokenstring)
	fmt.Println("isi : ", body)

	idstring := DecodeGetId(publicKey, tokenstring)
	if idstring == "" {
		fmt.Println("expire token")
	}
	fmt.Println("DecodeGetId isi : ", idstring)

	require.NoError(t, err)
}

func TestWaTokenDecodewithStaticKey(t *testing.T) {
	//privateKey := "461ce0e87748fd656c518b870da217dc200fc8d3b6275dda8cf14943424bf8c49e2ece1954df1ea8b151fba59cc7cbd4fb810b69716149e1c26169227bd5b686"
	publicKey := "20d9c3fbb219640c9baf0846a28fc6d3971924ccdd0175ec76bef435b248ed83"
	//userid := "awangga"
	tokenstring := "v4.public.eyJleHAiOiIyMDIzLTAxLTIzVDA0OjI3OjQzKzA3OjAwIiwiaWF0IjoiMjAyMy0wMS0yM1QwNDoyNjo0MyswNzowMCIsImlkIjoiYXdhbmdnYSIsIm5iZiI6IjIwMjMtMDEtMjNUMDQ6MjY6NDMrMDc6MDAifevjpmzIHYoUu9ZW_B15FPcEVSglaAxmACxyw5UEX4mpPnnIo53DOfUFAspGt2QNG8fwgnAH_volax7ZW4NAZwU"
	idstring := DecodeGetId(publicKey, tokenstring)
	if idstring == "" {
		fmt.Println("expire token")
	}
	fmt.Println("TestWaTokenDecodewithStaticKey idstring : ", idstring)

}

func TestWacipher(t *testing.T) {
	n := 100
	rnd := RandomString(n)
	require.Len(t, rnd, n)
	fmt.Println("rnd : ", rnd)
	host := "https://www.w3schools.com/js/js_window_location.asp"
	sEnc := base64.StdEncoding.EncodeToString([]byte(host))
	wh := GetAppSubDomain(rnd + "." + sEnc)
	require.Equal(t, "www", wh)
	fmt.Println("wh : ", wh)
}

func TestWaBcrypt(t *testing.T) {
	hasil_hash := GetBcryptHash("akuapa")
	fmt.Println(hasil_hash)
	fmt.Println("Hasil Hash : ", hasil_hash)

}
