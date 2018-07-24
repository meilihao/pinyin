package pinyin

// 汉语拼音音节表 (新华字典第11版)
// https://wenku.baidu.com/view/fe8f9fc503d276a20029bd64783e0912a3167c51.html?from=search
var InitialsWithoutH = "bpmfdtnlgkhjqxzcsryw" // no: zh,ch,sh
var Finals = "a,ia,ua,o,uo,e,i,u,ü,ai,uai,ei,ui,ao,iao,ou,iu,ie,üe," +
	"er,an,ian,uan,üan,en,in,un,ün,ang,iang,uang,eng,ing,ong,iong"

const (
	Normal      = 0 // 普通风格，不带声调（默认风格）。如： zhong guo
	Tone        = 1 // 声调风格1，拼音声调在韵母第一个字母上。如： zhōng guó
	Tone2       = 2 // 声调风格2，即拼音声调在各个韵母之后，用数字 [1-4] 进行表示。如： zho1ng guo2
	Tone3       = 8 // 声调风格3，即拼音声调在各个拼音之后，用数字 [1-4] 进行表示。如： zhong1 guo2
	Initial     = 3 // 声母风格，只返回各个拼音的声母部分。如： zh g
	FirstLetter = 4 // 首字母风格，只返回拼音的首字母部分。如： z g
	Final       = 5 // 韵母风格，只返回各个拼音的韵母部分，不带声调。如： ong uo
	FinalsTone  = 6 // 韵母风格1，带声调，声调在韵母第一个字母上。如： ōng uó
	FinalsTone2 = 7 // 韵母风格2，带声调，声调在各个韵母之后，用数字 [1-4] 进行表示。如： o1ng uo2
	FinalsTone3 = 9 // 韵母风格3，带声调，声调在各个拼音之后，用数字 [1-4] 进行表示。如： ong1 uo2
)

// 带音标字符。
var phoneticSymbol = map[string]string{
	"ā": "a1",
	"á": "a2",
	"ǎ": "a3",
	"à": "a4",
	"ē": "e1",
	"é": "e2",
	"ě": "e3",
	"è": "e4",
	"ō": "o1",
	"ó": "o2",
	"ǒ": "o3",
	"ò": "o4",
	"ī": "i1",
	"í": "i2",
	"ǐ": "i3",
	"ì": "i4",
	"ū": "u1",
	"ú": "u2",
	"ǔ": "u3",
	"ù": "u4",
	"ü": "v",
	"ǘ": "v2",
	"ǚ": "v3", // 女
	"ǜ": "v4",
	"ń": "n2", // 唔
	"ň": "n3",
	"ǹ": "n4",
	"ḿ": "m2", // 呒, 新华字典(v11)有m4, 但在Unicode里没找到
}
