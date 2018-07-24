package main

import (
	"fmt"

	"github.com/meilihao/pinyin"
)

func main() {
	hans := "中国人"

	fmt.Println(hans)
	tk := &pinyin.Token{
		Heteronym: true,
	}

	for i := 0; i <= pinyin.FinalsTone3; i++ {
		tk.Style = byte(i)

		fmt.Println(tk.Parse(hans))
	}
}
