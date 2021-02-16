package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

//Go语言的大多数数据类型都可以转化为有效的JSON文本，但channel、 complex和函数这几种
//类型除外
//如果转化前的数据结构中出现指针，那么将会转化指针所指向的值，如果指针指向的是零值，
//那么null将作为转化后的结果输出。

//在Go中， JSON转化前后的数据类型映射如下。
// 布尔值转化为JSON后还是布尔类型。
// 浮点数和整型会被转化为JSON里边的常规数字。
// 字符串将以UTF-8编码转化输出为Unicode字符集的字符串，特殊字符比如<将会被转义为
//\u003c。
// 数组和切片会转化为JSON里边的数组，但[]byte类型的值将会被转化为 Base64 编码后
//的字符串， slice类型的零值会被转化为 null。
// 结构体会转化为JSON对象，并且只有结构体里边以大写字母开头的可被导出的字段才会
//被转化输出，而这些可导出的字段会作为JSON对象的字符串索引。
// 转化一个map类型的数据结构时，该数据的类型必须是 map[string]T（ T可以是
//encoding/json 包支持的任意数据类型）。

type Book struct {
	Title       string
	Authors     []string
	Publisher   string
	IsPublished bool
	Price       float32
}

func testJson() {
	book := Book{
		"testTitle", []string{"a", "b"}, "xx.com.cn", true, 9.99,
	}

	b, err := json.Marshal(book)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("json:", string(b))
	}

	equal := bytes.Equal(b, []byte(`{"Title":"testTitle","Authors":["a","b"],"Publisher":"xx.com.cn","IsPublished":true,"Price":9.99}`))
	fmt.Println("equal:", equal)

	var unBook Book
	//如果JSON中的字段在Go目标类型中不存在， json.Unmarshal()函数在解码过程中会丢弃该字段,
	//目标类型中不可被导出的私有字段（非首字母大写）将不会受到解码转化的影响
	err = json.Unmarshal(b, &unBook)
	if err == nil {
		fmt.Println("unBook:", unBook)
	}
}

//如果要解码一段未知结构的JSON，只需将这段JSON数据解码输出到一个空接口即可。在解码JSON
//数据的过程中， JSON数据里边的元素类型将做如下转换：
// JSON中的布尔值将会转换为Go中的bool类型；
// 数值会被转换为Go中的float64类型；
// 字符串转换后还是string类型；
// JSON数组会转换为[]interface{}类型；
// JSON对象会转换为map[string]interface{}类型；
// null值会转换为nil。
func testUnknownType() {
	book := Book{
		"testTitle", []string{"a", "b"}, "xx.com.cn", true, 9.99,
	}

	b, _ := json.Marshal(book)

	var r interface{}
	json.Unmarshal(b, &r)
	fmt.Println("r:", r)
	gobook, ok := r.(map[string]interface{}) //判断类型
	if ok {
		for k, v := range gobook {
			switch v2 := v.(type) {
			case string:
				fmt.Println(k, "is string", v2)
			case int:
				fmt.Println(k, "is int", v2)
			case bool:
				fmt.Println(k, "is bool", v2)
			case []interface{}:
				fmt.Println(k, "is an array:")
				for i, iv := range v2 {
					fmt.Println(i, iv)
				}
			default:
				fmt.Println(k, "is another type not handle yet")
			}
		}
	}
}

//控制台输入：{"Title":"testTitle","Authors":["a","b"],"Publisher":"xx.com.cn","IsPublished":true,"Price":9.99}
func testJsonStream() {
	dec := json.NewDecoder(os.Stdin)
	enc := json.NewEncoder(os.Stdout)

	for {
		var v map[string]interface{}
		if err := dec.Decode(&v); err != nil { //从控制台读取，转换成v
			fmt.Println(err)
			return
		}
		for k := range v {
			if k != "Title" {
				v[k] = nil
			}
		}
		if err := enc.Encode(&v); err != nil { //没错误则输出到控制台
			fmt.Println(err)
		}
	}
}

func main() {
	//testJson()
	//testUnknownType()
	testJsonStream()
}
