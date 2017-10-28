# gomanuf
通过mac地址计算设备厂家信息

## 安装[Installation]
```
go get -u github.com/timest/gomanuf
```

## 使用[Usage]
```
import (
    manuf "github.com/timest/gomanuf"
    "log"
)

m1 := manuf.Search("18:65:90:dc:c0:cb")
log.Info(m1) //Apple, Inc.

m2 := manuf.Search("00:ec:0a:ff:b7:27")
log.Info(m2) // Xiaomi Communications Co Ltd

```

