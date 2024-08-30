package main

import "fmt"

func main() {
	websites:= map[string]string {
		"google": "https://www.google.com",
		"aws": "https://www.aws.com",
	}

	websites["msai"] = "https://msai.me"

	fmt.Println(websites)

	delete(websites, "aws")

	fmt.Println(websites)
}