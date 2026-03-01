# base

`base` 是 infrago 的基础类型包，定义 `Map`、`Any`、`Res` 及核心通用类型。

## 安装

```bash
go get github.com/infrago/base@latest
```

## 使用

```go
import . "github.com/infrago/base"

func Demo() Map {
    return Map{"ok": true}
}
```

## 公开 API（摘自源码）

- `func (vvv *Var) Nil() bool`
