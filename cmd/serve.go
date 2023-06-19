package cmd

import (
	"github.com/spf13/cobra"
	"mythogen/server"
)

var serve_cmd = &cobra.Command{
	Use: "serve",
	Run: func(cmd *cobra.Command, args []string) {
		port := "5051"
		if len(args)>0{
			port = args[0]
		}
		server.serve()
	},
}

func init(){
	root_cmd.AddCommand(serve_cmd)	
}