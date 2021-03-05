package main

import (
	"fmt"
	library "xiaodim.com/xiaodi/project/model"
)

func main() {
	user, err :=library.GetUserInfo(1)
	fmt.Println(user.Nickname)
	fmt.Println(err)
}