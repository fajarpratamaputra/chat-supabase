/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"chitchat/app"
	"chitchat/app/api"
	"chitchat/config"
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run the service",
	Long:  `RUN RUN RUN`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := config.Load(); err != nil {
			log.Fatal(err)
		}
		fmt.Println("serve called")

		ctx := context.Background()

		c := app.NewContainer(ctx)

		api.Serve(ctx, c)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
