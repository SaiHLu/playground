package main

import "fmt"

type User struct {
	name string
	age  int
}

func main() {
	users := make([]User, 16)

	users[0] = User{
		name: "Sai Hlaing",
		age:  1,
	}

	users[1] = User{
		name: "Sai Hlaing2",
		age:  2,
	}

	for key, value := range users {
		fmt.Printf("key %d, and value %+v\n", key, value)
	}
}
