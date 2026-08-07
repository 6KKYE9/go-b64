package main

import (
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	mode := "encode"
	args := os.Args[1:]
	var texts []string
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "-d", "--decode":
			mode = "decode"
		case "-u", "--url":
			// url 安全变体，后面用的时候再判断，放这记个标记
			args[i] = "-u"
		case "-h", "--help":
			fmt.Println("go-b64 Base64 编解码")
			fmt.Println("用法: go-b64 [-d] [-u] <文本|文件?>")
			fmt.Println("  -d  解码")
			fmt.Println("  -u  使用 URL 安全变体（含 - 和 _）")
			return
		default:
			texts = append(texts, args[i])
		}
	}

	urlSafe := false
	for _, a := range args {
		if a == "-u" {
			urlSafe = true
		}
	}

	if len(texts) == 0 {
		fmt.Fprintln(os.Stderr, "要给点东西编解码啊")
		os.Exit(1)
	}

	for _, t := range texts {
		if mode == "decode" {
			var dec *base64.Encoding
			if urlSafe {
				dec = base64.URLEncoding
			} else {
				dec = base64.StdEncoding
			}
			raw, err := dec.DecodeString(t)
			if err != nil {
				fmt.Fprintf(os.Stderr, "解码失败: %v\n", err)
				os.Exit(1)
			}
			fmt.Println(string(raw))
		} else {
			var enc *base64.Encoding
			if urlSafe {
				enc = base64.URLEncoding
			} else {
				enc = base64.StdEncoding
			}
			fmt.Println(enc.EncodeToString([]byte(t)))
		}
	}
}
