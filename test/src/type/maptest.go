package main

import "fmt"

//PersonInfo是包含个人详细信息的类型
type PersonInfo struct {
	ID      string
	Name    string
	Address string
}

func testCreate() {

	var myMap1 = make(map[string]PersonInfo)
	//创建并指定存储能力
	var myMap2 = make(map[string]PersonInfo, 100)
	//创建并初始化
	var myMap3 = map[string]PersonInfo{
		"1234": PersonInfo{"1", "jack", "Room 101,..."},
	}

	fmt.Println(myMap1, myMap2, myMap3)
}

func testDelete() {

	var myMap = make(map[string]PersonInfo)
	//若不存在则什么不发生，也不会有副作用，若myMap为nil则异常(panic)
	delete(myMap, "1234")
}

func main() {
	var personDB map[string]PersonInfo //key:string,value:PersonInfo
	personDB = make(map[string]PersonInfo)

	personDB["12345"] = PersonInfo{"12345", "Tom", "Room 203,..."}
	personDB["1"] = PersonInfo{"1", "jack", "Room 101,..."}

	//判断是否成功找到特定的键，不需要检查取到的值是否为nil，只需查看第二个返回值ok，这让表意清晰很多
	person, ok := personDB["1234"]
	if ok {
		fmt.Println("Found person", person.Name, "with ID 1234.")
	} else {
		fmt.Println("Did not find person with ID 1234.", )
	}
}
