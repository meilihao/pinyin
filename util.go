package pinyin

import (
	"regexp"
	"strings"
)

var initialArray = make([]string, 128)

func init() {
	initials := []byte(InitialsWithoutH)
	for _, v := range initials {
		initialArray[v] = string(v)
	}
}

// 获取单个拼音中的声母
func initial(p string) string {
	tmp := initialArray[p[0]]

	switch tmp {
	case "z", "c", "s":
		if len(p) >= 2 && p[1] == 'h' {
			tmp += "h"
		}
	}

	return tmp
}

// 获取单个拼音中的韵母
func final(p string) string {
	i := initial(p)
	if i == "" {
		return handleYW(p)
	}

	return handleJQXU(p, i)
}

func handleJQXU(p, i string) string {
	if len(p) >= 2 && (p[0] == 'j' || p[0] == 'q' || p[0] == 'x') && p[1] == 'u' {
		return "v" + p[2:]
	}

	return strings.TrimPrefix(p, i)
}

// 处理 y, w
func handleYW(p string) string {
	// 特例 y/w
	if strings.HasPrefix(p, "yu") {
		p = "v" + p[2:] // yu -> v
	} else if strings.HasPrefix(p, "y") {
		p = p[1:]
	} else if strings.HasPrefix(p, "w") {
		p = p[1:]
	}
	return p
}

// 所有带声调的字符
var rePhoneticSymbolKeys = func(m map[string]string) string {
	s := ""

	for k := range m {
		s += k
	}

	return s
}(phoneticSymbol)

var rePhoneticSymbol = regexp.MustCompile("[" + rePhoneticSymbolKeys + "]")

//var reTone2 = regexp.MustCompile("([aeoiuvnm])([1-4])$")

// 去掉声调: a1 -> a
// TrimTone(s) == reTone2.ReplaceAllString(s, "$1")
func TrimTone(raw string) string {
	n := len(raw)

	if raw[n-1] >= '1' && raw[n-1] <= '4' {
		raw = raw[:n-1]
	}

	return raw
}

// 匹配 Tone2 中标识韵母声调的正则表达式
//var reTone3 = regexp.MustCompile("^([a-z]+)([1-4])([a-z]*)$")

// RepositionTone(s) == reTone3.ReplaceAllString(s, "$1$3$2")
func RepositionTone(raw string) string {
	n := len(raw)

	index := -1
	for i, v := range raw {
		if v >= '1' && v <= '4' {
			index = i

			break
		}
	}

	if index == -1 || index == n-1 {
		return raw
	}

	return raw[:index] + raw[index+1:] + string(raw[index])
}

// 替换拼音中的带声调字符
func ReplacePhoneticSymbol(s string, style byte) string {
	var symbol string
	var ok bool
	isNeed := true // for FinalsTone

	for i, v := range s {
		symbol, ok = phoneticSymbol[string(v)]
		if ok {
			switch style {
			case Normal, FirstLetter, Final: // 不包含声调
				symbol = TrimTone(symbol)
			case Tone2, FinalsTone2, Tone3, FinalsTone3: // 使用数字标识声调的字符
			default: // 声调在头上
				isNeed = false
			}

			if isNeed {
				return s[:i] + symbol + s[i+2:] // symbol(2B) in s
			}

			return s
		}
	}

	return s
}

// 是否是多音字
func IsMultiPinyin(r rune) bool {
	v := pinyinDictPool[r]

	return len(v) > 1
}

// 是不是中文,不是很准确(因为范围包含日文),但够用.
func IsChinese(r rune) bool {
	if r >= 0x4E00 && r <= 0x9FEA {
		return true
	}

	return false
}

// 是不是中文, 准确
func IsChineseReal(r rune) bool {
	return pinyinDictPool[r] != nil
}

// 判断字符串中是否包含中文
func IsContainsChinese(s string) bool {
	tmp := []rune(s)

	for _, v := range tmp {
		if v >= 0x4E00 && v <= 0x9FEA {
			return true
		}
	}

	return false
}
