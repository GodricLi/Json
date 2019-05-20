package main

import (
	"encoding/json"
	"fmt"
)

type Students struct {
	Id   int
	Name string
}

type Class struct {
	Name     string
	Students []*Students
}

func main() {
	c := &Class{
		Name: "101",
	}
	for i := 0; i < 10; i++ {
		stu := &Students{
			Id:   i,
			Name: fmt.Sprintf("stu%d", i),
		}
		c.Students = append(c.Students, stu)

	}
	fmt.Println(c)

	//使用json序列化，Marshal返回一个[]byte,和error，没有出错时error为nil
	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println("json marshal failed")
		return
	}
	//将[]byte类型转化成string
	str_data := string(data)
	fmt.Println(str_data)

	// 反序列化
	var deserilize = `{"Name":"101","Students":
	[{"Id":0,"Name":"stu0"},{"Id":1,"Name":"stu1"},{"Id":2,"Name":"stu2"},
	{"Id":3,"Name":"stu3"},{"Id":4,"Name":"stu4"},{"Id":5,"Name":"stu5"},
	{"Id":6,"Name":"stu6"},{"Id":7,"Name":"stu7"},{"Id":8,"Name":"stu8"},
	{"Id":9,"Name":"stu9"}]
	}`
	//获取一个结构体的地址，必须是地址，将需要反序列化的数据使用[]byte转化之后进行Unmarshal序列化
	var c1 *Class = &Class{}
	err = json.Unmarshal([]byte(deserilize), c1)
	if err != nil {
		fmt.Println("Umarshal failed")
	}
	fmt.Printf("%#v", c1)
	for _, v := range c1.Students {
		fmt.Printf("%#v\n", v)
	}
}
