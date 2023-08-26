package app

import "fmt"

func App() {
	config := configure()
	fmt.Printf("%v", config)
}
