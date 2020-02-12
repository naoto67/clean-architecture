package main

import (
	"fmt"

	"github.com/naoto67/clean-architecture/src/infrastructure"
)

func main() {
	userRepository := infrastructure.NewUserRepository()
	fmt.Println(userRepository.FindByName("sample"))
	fmt.Println(userRepository.FindByName("aaa"))
}
