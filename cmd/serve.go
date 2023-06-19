package cmd

import (
	"fmt"
	"os"
	"strconv"

	"mythogen/server"

	"github.com/spf13/cobra"
)

var serve_cmd = &cobra.Command{
	Use: "serve",
	Run: func(cmd *cobra.Command, args []string) {
        port := "5051"
		if len(args)>0{
            port = args[0]
            v, err := strconv.Atoi(port)
            if v<1024 {
                fmt.Println("Invalid port. Please use a port that is not below 1024")
                os.Exit(1)
            }
            if err != nil {
                fmt.Println("Invalid port")
                os.Exit(1)
            }
		}
		server.Serve( port )
	},
}

func init(){
	root_cmd.AddCommand(serve_cmd)	
}
