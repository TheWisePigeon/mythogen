package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func print_help() {
	println("Mythogen v 0.1 by TheWisePigeon")
	println("Documentation at https://github.com/TheWisePigeon/mythogen#README")
	return
}

var DEFAULT_PORT = 6061

func ServeMythogen(port string) {
  err := http.ListenAndServe(port, nil)
  if err!= nil {
    println("Something went wrong while launching the Mythogen server: ", err.Error())
    println("If this issue persist please report it at https://github.com/TheWisePigeon/mythogen/issues")
    os.Exit(1)
  }
}

func main() {
	switch len(os.Args) {
	case 2:
		if os.Args[1] == "serve" {
			println("Serving Mythogen server on default port:", DEFAULT_PORT)
      ServeMythogen(fmt.Sprintf(":%d", DEFAULT_PORT))
		}
		print_help()
	case 3:
		if os.Args[1] == "serve" {
			port := os.Args[2]
			if port_to_int, err := strconv.Atoi(port); err != nil {
				println("Invalid port")
				return
			} else {
				if port_to_int <= 1024 {
					println("Invalid port! Port number must be higher than 1024")
					return
				}
			}
			println("Serving Mythogen server on port", port)
      ServeMythogen(fmt.Sprintf(":%s", port))
		}
		print_help()
	default:
		print_help()
	}

}
