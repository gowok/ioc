package main

import (
	"fmt"

	"github.com/gowok/ioc"
)

type User struct {
	Name string
}

const KeyUser ioc.Key = "KeyUser"

func main() {
	ioc.Bind(KeyUser, func() interface{} {
		return &User{
			Name: "Alex Under",
		}
	})

	user, ok := ioc.Use(KeyUser)
	if !ok {
		panic(KeyUser + " not found!")
	}

	fmt.Println(user.(*User).Name)
}
