# watoken

Whatsauth Token using PASETO

```go
 privateKey, publicKey := GenerateKey()
 fmt.Println("privateKey : ", privateKey)
 fmt.Println("publicKey : ", publicKey)
 userid := "awangga"

 signed, _ := Encode(userid, privateKey)
 body, _ := Decode(publicKey, signed)
 fmt.Println("signed : ", signed)
 fmt.Println("isi : ", body)
```

## Tagging

```sh
git tag v0.0.1
git push origin --tags
go list -m github.com/whatsauth/watoken@v0.0.1
```
