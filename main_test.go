package main

import (
	"encoding/base64"
	"testing"
)

func TestEncodeStd(t *testing.T) {
	got := base64.StdEncoding.EncodeToString([]byte("hello"))
	if got != "aGVsbG8=" {
		t.Errorf("期望 aGVsbG8=, 得到 %q", got)
	}
}

func TestDecodeStd(t *testing.T) {
	raw, err := base64.StdEncoding.DecodeString("aGVsbG8=")
	if err != nil {
		t.Fatal(err)
	}
	if string(raw) != "hello" {
		t.Errorf("期望 hello, 得到 %q", raw)
	}
}

func TestURLEncode(t *testing.T) {
	got := base64.URLEncoding.EncodeToString([]byte("a/b+c"))
	if got != "YS9iK2M=" {
		t.Errorf("URL 安全期望 YS9iK2M=, 得到 %q", got)
	}
}
