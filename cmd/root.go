package cmd

import (
	"fmt"
	"os"

	"github.com/vinsensiussatya/bego-training/internal/server"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "BEGO training",
	Short: "BEGO training REST API",
	Run: func(_ *cobra.Command, _ []string) {
		start()
	},
}

func init() {
	cobra.OnInitialize()
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func start() {
	h := initInjections()
	s := server.NewServer(h)
	s.StartRestAPI()
}
