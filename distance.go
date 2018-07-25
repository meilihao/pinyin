// from http://ai.baidu.com/docs/#/ASR-Tool-diff/top
package pinyin

import (
	"fmt"
	"sort"
)

type Distance struct {
	origin *Word
	limit  int
	tk     *Token
}

func NewDistance(raw string, limit int) *Distance {
	tk := &Token{
		Separator: ",",
	}

	return &Distance{
		limit:  limit,
		tk:     tk,
		origin: NewWord(raw, tk),
	}
}

type Word struct {
	raw     string
	pNormal string
	pTone3  string
}

func NewWord(raw string, tk *Token) *Word {
	w := &Word{
		raw: raw,
	}

	tk.Style = Normal
	w.pNormal = tk.ParseToString(raw)

	tk.Style = Tone3
	w.pTone3 = tk.ParseToString(raw)

	return w
}

func (w *Word) Compare(o *Word) int {
	num1 := caclDistance(w.pNormal, o.pNormal)
	num2 := caclDistance(w.pTone3, o.pTone3)

	return num1 + num2
}

type Score struct {
	Word *Word
	Num  int
}

func (s *Score) String() string {
	return fmt.Sprintf("{word=%s, num=%d}", s.Word.raw, s.Num)
}

func (d *Distance) Search(ss ...string) []*Score {
	if len(ss) == 0 {
		return nil
	}

	tmp := make([]*Score, len(ss))
	for i, v := range ss {
		s := &Score{
			Word: NewWord(v, d.tk),
		}
		s.Num = s.Word.Compare(d.origin)

		tmp[i] = s
	}

	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i].Num < tmp[j].Num
	})

	if len(tmp) > d.limit {
		tmp = tmp[:d.limit]
	}

	return tmp
}

// t is origin
// [edit distance度量方法O(M*N)](http://www.cnblogs.com/ivanyb/archive/2011/11/25/2263356.html)
// https://zh.wikipedia.org/wiki/%E8%90%8A%E6%96%87%E6%96%AF%E5%9D%A6%E8%B7%9D%E9%9B%A2
// other methods: [字符串相似性的几种度量方法](https://blog.csdn.net/shijing_0214/article/details/53100992)
func caclDistance(s, t string) int {
	n := len(s)
	m := len(t)
	if n == 0 {
		return m
	}
	if m == 0 {
		return n
	}

	if n == m && s == t {
		return 0
	}

	d := make([][]int, n+1)
	for i := range d {
		d[i] = make([]int, m+1)
	}

	i, j, cost := 0, 0, 0
	for ; i <= n; i++ {
		d[i][0] = i // left column
	}
	for ; j <= m; j++ {
		d[0][j] = j // top row
	}

	var s_i uint8 // ith character of s
	var t_j uint8 // jth character of t

	for i = 1; i <= n; i++ {
		s_i = s[i-1]

		for j = 1; j <= m; j++ {
			t_j = t[j-1]

			cost = 1
			if s_i == t_j {
				cost = 0
			}

			d[i][j] = min(d[i-1][j]+1, d[i][j-1]+1,
				d[i-1][j-1]+cost) //上方(del s[i]),左方(insert t[j]),左上(replace s[i] -> t[j])
		}
	}

	return d[n][m]
}

func min(a, b, c int) int {
	tmp := a
	if b < tmp {
		tmp = b
	}

	if c < tmp {
		tmp = c
	}

	return tmp
}
