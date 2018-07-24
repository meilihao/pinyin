package main

import (
	"fmt"

	"github.com/meilihao/pinyin"
)

func main() {
	hans := "中国人"

	fmt.Println(hans)
	tk := &pinyin.Token{}

	fmt.Println(tk.Parse(hans))
}
