# Everything SDK Go Bindings

基于 [Everything SDK](https://www.voidtools.com/support/sdk/) 的 Go 语言绑定，**纯 Go 实现，无 CGO 依赖**。

## 特性

- **🚀 无 CGO**: 纯 Go 实现
- **DLL 嵌入**: `Everything64.dll` 嵌入到 Go 二进制
- **文件搜索**: Windows 快速文件搜索

## 使用方法

```go
package main

import "github.com/ddkwork/everything"

func main() {
    e := &everything.Everything{}
    e.SetSearchW("*.go")
    if e.QueryW(1) == 1 {
        count := e.GetNumResults()
        fmt.Printf("Found %d files\n", count)
    }
}
```

## 测试

```bash
go test -v
```

## 许可证

MIT License
