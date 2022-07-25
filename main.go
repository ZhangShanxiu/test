package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

type User struct {
	Name  string
	Sex   int
	Price float64
	Addr  string
	Age   int32
	Arr   []string
	Arr1  []string
	Arr2  []string
	Arr3  []string
	Ctime time.Time
}

func CopyUser(src, dst *User) {
	dst.Name = src.Name
	dst.Sex = src.Sex
	dst.Arr = src.Arr
	dst.Price = src.Price
	dst.Ctime = src.Ctime
	dst.Arr1 = src.Arr1
	dst.Arr2 = src.Arr2
	dst.Arr3 = src.Arr3
	dst.Age = src.Age
}

func CopyStruct(src, dst interface{}) {
	s := reflect.ValueOf(src).Elem()
	d := reflect.ValueOf(dst).Elem()

	for i := 0; i < s.NumField(); i++ {
		sField := s.Type().Field(i)
		sValue := s.Field(i)
		sType := sField.Type.String()

		dValue := d.FieldByName(sField.Name)
		dType := dValue.Type().String()
		if !dValue.IsValid() {
			continue
		}

		if sType == dType {
			dValue.Set(sValue)
		}
	}
}

func CopyObj(src, dst interface{}) {
	data, _ := json.Marshal(src)
	json.Unmarshal(data, dst)
}

func main() {
	user1 := &User{
		Name:  "xiaoming",
		Sex:   10,
		Arr:   []string{"a", "b"},
		Price: 1.1,
		Ctime: time.Now(),
	}

	user2 := &User{}
	beg := time.Now()
	for i := 0; i < 1000000; i++ {
		//CopyUser(user1, user2)
		//CopyStruct(user1, user2)
		CopyObj(user1, user2)
		time.Sleep(1 * time.Second)

		fmt.Println("Hello world")

		fmt.Println("HSX change again")
		fmt.Println("ZSX change again loaclly")

		fmt.Println("ZSX change again by goland")

		fmt.Println("add an account.")

		fmt.Println("online.")
		fmt.Println("local.")

		fmt.Println("modyfied 20220725.")
	}
	fmt.Printf("cost: %dms", time.Since(beg).Milliseconds())
}
