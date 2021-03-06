package pinyin

import (
	"log"
	"strings"
)

// what to convert
type Token struct {
	Style        byte
	Heteronym    bool
	Segment      bool
	Separator    string
	ExcludeOther bool
}

func (tk *Token) Parse(str string) [][]string {
	result := make([][]string, 0, len(str))

	for _, r := range str {
		if IsChinese(r) {
			result = append(result, tk.parse(r))
		} else if tk.ExcludeOther {
		} else {
			result = append(result, []string{string(r)})
		}
	}

	return result
}

func (tk *Token) ParseToString(str string) string {
	result := tk.Parse(str)

	var n int
	tmp := make([]string, 0, len(result))
	for _, v := range result {
		n = len(v)

		if n == 0 {
			continue
		}

		tmp = append(tmp, v[0])
	}

	return strings.Join(tmp, tk.Separator)
}

func (tk *Token) parse(r rune) []string {
	v := pinyinDictPool[r]

	if len(v) > 0 {
		if !tk.Heteronym {
			v = v[:1]
		}

		if tk.Style == Tone {
			return v
		}

		return handleStyle(v, tk.Style)
	}

	if !tk.ExcludeOther {
		log.Printf("unknown char: %s\n", string(r))
	}

	return nil
}

func handleStyle(ss []string, style byte) []string {
	ns := make([]string, len(ss))

	for i, v := range ss {
		ns[i] = handleStyleSingle(v, style)
	}

	return ns
}

func handleStyleSingle(s string, style byte) string {
	if style == Initial {
		return initial(s)
	}

	originS := s

	// 替换拼音中的带声调字符
	// s = rePhoneticSymbol.ReplaceAllStringFunc(s, func(m string) string {
	// 	symbol, _ := phoneticSymbol[m]

	// 	switch style {
	// 	case Normal, FirstLetter, Final: // 不包含声调
	// 		// m = reTone2.ReplaceAllString(symbol, "$1")
	// 		m = TrimTone(symbol)
	// 	case Tone2, FinalsTone2, Tone3, FinalsTone3: // 返回使用数字标识声调的字符
	// 		m = symbol
	// 	}

	// 	return m
	// })
	s = ReplacePhoneticSymbol(s, style)

	switch style {
	case Tone3, FinalsTone3: // 将声调移动到最后
		// s = reTone3.ReplaceAllString(s, "$1$3$2")
		s = RepositionTone(s)
	}

	switch style {
	case FirstLetter: // 首字母
		return s[:1]
	case Final, FinalsTone, FinalsTone2, FinalsTone3: // 韵母
		// 转换为 []rune unicode 编码用于获取第一个拼音字符
		// 因为 string 是 utf-8 编码不方便获取第一个拼音字符
		rs := []rune(originS)
		switch string(rs[0]) {
		// 因为鼻音没有声母所以不需要去掉声母部分
		case "ḿ", "ń", "ň", "ǹ":
		default:
			s = final(s)
		}
	}
	return s
}

// from github.com/stuxuhai/jpinyin
func (tk *Token) ParseByDict(str string) []string {
	tmp := []rune(str)
	i, strLen := 0, len(tmp)
	result := make([]string, 0, len(tmp))

	for i < strLen {
		substr := tmp[i:]

		ls := wordPinyinTrie.PrefixPredict([]byte(string(substr)), 0)
		if len(ls) == 0 {
			if IsChinese(tmp[i]) {
				pTmp := tk.parse(tmp[i])

				if len(pTmp) > 0 {
					result = append(result, pTmp[0])
				} else {
					result = append(result, "")
				}
			} else if tk.ExcludeOther {
			} else {
				result = append(result, string(tmp[i]))
			}

			i++
		} else {
			tmp := wordPinyinPool[ls[len(ls)-1]].pinyin

			for _, v := range tmp {
				result[i] = handleStyleSingle(v, tk.Style)
				i++
			}
		}
	}

	return result
}
