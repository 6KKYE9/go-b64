# go-b64

又在终端里为个编码解码去开浏览器搜网页？别了，一行命令的事。

Base64 编解码，零依赖，标准库 `encoding/base64`。

## 用法

```bash
go run . "hello"            # 编码
go run . -d "aGVsbG8="      # 解码
go run . -u "a/b+c"         # URL 安全变体（用 - 和 _ 代替 + 和 /）
```

解码失败会报错退出，不会吐出半截数据。
