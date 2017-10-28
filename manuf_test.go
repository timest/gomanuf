package manuf

import "testing"

func TestSearch(t *testing.T) {
    d1 := Search("48:e2:44:45:0b:04")
    if len(d1) == 0 {
        t.Fatal("数据有问题")
    }
    
    d2 := Search("24:1f:a0:17:6d:9b")
    if len(d2) == 0 {
        t.Fatal("数据有问题")
    }
    
}
