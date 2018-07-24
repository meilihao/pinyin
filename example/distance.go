// 计算词间的相识度
package main

import (
	"fmt"

	"github.com/meilihao/pinyin"
)

func main() {
	s := []string{"张山", "张三", "张散", "张丹", "张成", "李四", "李奎"}
	list1 := pinyin.NewDistance("张山", 10)
	fmt.Println(list1.Search(s...))

	list2 := pinyin.NewDistance("李四", 10)
	fmt.Println(list2.Search(s...))
}
