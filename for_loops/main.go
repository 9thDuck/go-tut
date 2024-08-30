package main

import "fmt"

func main() {
	var userNames = [4]string{
		"asdf", "9thduck",
	}

	userNames[3] = "msai"

	for index, value := range userNames {
		fmt.Println(`index`, index)
		fmt.Println(`value`, value)
	}

	userNameEmailMap := map[string]string{
		"asdf":    "asdf@gmail.com",
		"msai":    "msai@gmail.com",
		"9thduck": "9thduck@gmail.com",
	}

	for key, value := range userNameEmailMap {
		fmt.Println("key", key)
		fmt.Println("value", value)
	}

	for index := range 4 {
		fmt.Println(index, userNames[index])
	}
}
