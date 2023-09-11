package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

var version = `
  Mythogen: Fake data generator server
  Author: TheWisePigeon <https://github.com/TheWisePigeon>
  Version: 0.1
`

func main() {
	switch len(os.Args) {
	case 1:
		println(version)
	case 2:
		switch os.Args[1] {
		case "version":
			println(version)
		case "serve":
			Mythogen("8081")
		default:
			println("Unrecognized command. See https://github.com/TheWisePigeon/mythogen#README to learn how to use mythogen")
		}
	case 3:
		if os.Args[1] == "serve" {
			port := os.Args[2]
			conv_result, err := strconv.Atoi(port)
			if err != nil {
				fmt.Printf("%s is not a valid port", port)
				os.Exit(1)
			}
			if conv_result < 1024 {
				println("The port should be higher than 1024")
				os.Exit(1)
			}
			Mythogen(port)
		}
		println("Unrecognized command. See https://github.com/TheWisePigeon/mythogen#README to learn how to use mythogen")
		os.Exit(1)
	}
}

func Mythogen(port string) {
	fmt.Printf("Launching mythogen server on port %s ", port)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		return
	})
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		println("Something went wrong when launching mythogen server. Please report bugs at https://github.com/TheWisePigeon/mythogen/issues")
		println(err.Error())
	}
}
