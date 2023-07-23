package main

import (
	"fmt"
	"modules/types"
	"modules/util"
)

func main() {
	user := types.User{
		Username: "Sai Hlaing",
		Age:      util.GetNumber(),
	}

	fmt.Printf("User is %+v", user)
}
