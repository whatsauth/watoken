package watoken

import (
	"testing"
	"time"
)

type sampleData struct {
	Role string `json:"role"`
}

// encoders memetakan setiap fungsi Encode* ke bentuk yang seragam untuk diuji.
func encoders() map[string]func(id, key string) (string, error) {
	data := &sampleData{Role: "admin"}
	return map[string]func(id, key string) (string, error){
		"Encode":                   func(id, key string) (string, error) { return Encode(id, key) },
		"EncodeWithStruct":         func(id, key string) (string, error) { return EncodeWithStruct(id, data, key) },
		"EncodeWithStructDuration": func(id, key string) (string, error) { return EncodeWithStructDuration(id, data, key, time.Hour) },
		"EncodeforHours":           func(id, key string) (string, error) { return EncodeforHours(id, "alias", key, 1) },
		"EncodeforMinutes":         func(id, key string) (string, error) { return EncodeforMinutes(id, key, 5) },
		"EncodeforSeconds":         func(id, key string) (string, error) { return EncodeforSeconds(id, key, 30) },
	}
}

func TestEncodeInvalidKeyReturnsNoToken(t *testing.T) {
	for name, enc := range encoders() {
		for _, key := range []string{"", "bukan-hex", "abcd"} {
			tok, err := enc("awangga", key)
			if err == nil {
				t.Errorf("%s(key=%q): expected error", name, key)
			}
			if tok != "" {
				t.Errorf("%s(key=%q): token harus kosong saat key tidak valid, dapat %q", name, key, tok)
			}
		}
	}
}

func TestEncodeDecodeRoundTrip(t *testing.T) {
	priv, pub := GenerateKey()
	for name, enc := range encoders() {
		tok, err := enc("awangga", priv)
		if err != nil || tok == "" {
			t.Fatalf("%s: err=%v token=%q", name, err, tok)
		}
		id, err := DecodeGetId(pub, tok)
		if err != nil || id != "awangga" {
			t.Errorf("%s: DecodeGetId = %q, %v", name, id, err)
		}
	}
}

func TestDecodeWithStruct(t *testing.T) {
	priv, pub := GenerateKey()
	tok, err := EncodeWithStruct("awangga", &sampleData{Role: "admin"}, priv)
	if err != nil {
		t.Fatal(err)
	}
	payload, err := DecodeWithStruct[sampleData](pub, tok)
	if err != nil || payload.Data.Role != "admin" || payload.Id != "awangga" {
		t.Fatalf("payload=%+v err=%v", payload, err)
	}
}

func TestDecodeRejectsBadTokens(t *testing.T) {
	priv, pub := GenerateKey()
	_, otherPub := GenerateKey()
	tok, _ := Encode("awangga", priv)

	if _, err := DecodeGetId(otherPub, tok); err == nil {
		t.Error("token harus ditolak oleh public key lain")
	}
	tampered := tok[:len(tok)-2] + "AA"
	if tampered == tok {
		tampered = tok[:len(tok)-2] + "BB"
	}
	if _, err := DecodeGetId(pub, tampered); err == nil {
		t.Error("token yang diubah harus ditolak")
	}
	if _, err := DecodeGetId("bukan-hex", tok); err == nil {
		t.Error("public key tidak valid harus ditolak")
	}
}

func TestDecodeRejectsExpiredToken(t *testing.T) {
	priv, pub := GenerateKey()
	tok, err := EncodeforSeconds("awangga", priv, 1)
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(1100 * time.Millisecond)
	if _, err := DecodeGetId(pub, tok); err == nil {
		t.Error("token kedaluwarsa harus ditolak")
	}
}
