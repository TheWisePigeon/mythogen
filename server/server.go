package server

import (
	"fmt"
	"net/http"
	"os"
)

func mythogen( w http.ResponseWriter, r *http.Request ){

}

func Serve( port string ){
    mux := http.NewServeMux()
    mux.HandleFunc("/", mythogen )
    fmt.Printf("Mythogen server launched on port %s\n", port)
    err := http.ListenAndServe(fmt.Sprintf("%s", port), mux)
    if err!= nil {
        println("Mythogen server crashed")
        fmt.Printf("%s", err)
        os.Exit(1)
    }
}
