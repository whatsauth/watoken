# watoken

<img width="1920" height="1043" alt="image" src="https://github.com/user-attachments/assets/60f8236c-3ec3-46f6-bd27-281dfdb04b1f" />

Whatsauth Token using PASETO, How to Use :

```go
package controller

import (
	"testing"

	"github.com/whatsauth/watoken"
)

func TestWatoken(t *testing.T) {
	privateKey, publicKey := watoken.GenerateKey()

	// generate token for user awangga (berlaku 2 jam)
	tokenstring, err := watoken.Encode("awangga", privateKey)
	if err != nil {
		t.Fatal(err) // key tidak valid: token selalu kosong
	}

	// decode token to get userid
	userid, err := watoken.DecodeGetId(publicKey, tokenstring)
	if err != nil {
		t.Fatal(err) // token palsu, diubah, atau kedaluwarsa
	}
	t.Log("id:", userid)
}
```

Semua fungsi `Encode*` mengembalikan string kosong dan error jika private key tidak valid, jadi selalu cek `err` sebelum memakai token.

## Generate Key

Untuk membuat pasangan `PRIVATEKEY` dan `PUBLICKEY`, clone repo ini lalu jalankan:

```sh
go test -run TestWatoken
```

Key dicetak ke output. Simpan private key sebagai secret dan jangan di-commit.

## Tagging

```sh
git tag v0.1.9
git push origin --tags
go list -m github.com/whatsauth/watoken@v0.1.9
```
