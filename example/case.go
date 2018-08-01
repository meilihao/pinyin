package main

import (
	"fmt"

	"github.com/meilihao/pinyin"
)

func main() {
	hans := "中国人,124,鿅" // "鿅": 日语用汉字
	fmt.Println("---Raw---")
	fmt.Println(hans, len([]rune(hans)))

	tk := &pinyin.Token{
		Heteronym: true,
		Separator: "-", // token.ParseToString
	}

	fmt.Println("---Parse---")
	for i := 0; i <= pinyin.FinalsTone3; i++ {
		tk.Style = byte(i)

		fmt.Println(tk.Parse(hans))
		fmt.Println(tk.ParseToString(hans))
	}

	fmt.Println("---ParseByDict---")
	pinyin.LoadWordPinyinDict(pinyin.WordPinyinDict)
	for i := 0; i <= pinyin.FinalsTone3; i++ {
		tk.Style = byte(i)

		fmt.Println(tk.ParseByDict(hans))
	}

	fmt.Println("---Parse  Only Chinese---")
	tk.ExcludeOther = true
	for i := 0; i <= pinyin.FinalsTone3; i++ {
		tk.Style = byte(i)

		fmt.Println(tk.Parse(hans))
		fmt.Println(tk.ParseToString(hans))
	}

	fmt.Println("---IsContainsChinese---")
	fmt.Println(pinyin.IsContainsChinese(hans))
	fmt.Println(pinyin.IsContainsChinese(hans + "1243adsfaadsfaq8234*&%^*"))
	fmt.Println(pinyin.IsContainsChinese("1243adsfaadsfaq8234*&%^*"))

	fmt.Println("---IsChinese---")
	fmt.Println(pinyin.IsChinese('A'))
	fmt.Println(pinyin.IsChinese('中'))
	fmt.Println(pinyin.IsChinese('鿅'))
	fmt.Println(pinyin.IsChineseReal('鿅'))

	fmt.Println("---IsMultiPinyin---")
	fmt.Println(pinyin.IsMultiPinyin('中'))
	fmt.Println(pinyin.IsMultiPinyin('国'))
}
