package main

import _ "encoding/json"
import _ "fmt"
import (
	"encoding/json"
	"fmt"
	"os"
	_ "os"
)

type Response1 struct {
	Page   int
	Fruits []string
}

type Response2 struct {
	Page   int      `json:"Response2_page"`
	Fruits []string `json:"Response2_fruits"`
}

func main() {

	res1D := &Response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}

	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	// 你可以使用tag来自定义编码后JSON键的名称
	res2D := &Response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	boolB, _ := json.Marshal(true)
	fmt.Println(boolB)
	fmt.Println(string(boolB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(fltB)
	fmt.Println(string(fltB))

	// 现在我们看看解码JSON数据为Go数值
	byt := []byte(`{"num":6.13, "strs":["a","b"]}`)

	// 我们需要提供一个变量来存储解码后的JSON数据，这里
	// 的`map[string]interface{}`将以Key-Value的方式
	// 保存解码后的数据，Value可以为任意数据类型
	var dat map[string]interface{}

	// 解码过程，并检测相关可能存在的错误
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	num := dat["num"].(float64)
	fmt.Println(num)

	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	str := `{"Response2_page": 1, "Response2_fruits": ["apple", "peach"]}`
	res := Response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0:2])

	// 上面的例子中，我们使用bytes和strings来进行原始数据和JSON数据
	// 之间的转换，我们也可以直接将JSON编码的数据流写入`os.Writer`
	// 或者是HTTP请求回复数据。
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)

}
