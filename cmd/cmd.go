/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/spf13/cobra"
	"github.com/supermarine1377/check-http-status/cmd/flags"
	"github.com/supermarine1377/check-http-status/internal/monitorer"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: `
	check-http-status <URL> [flags]
	`,
	Short: "Monitors the HTTP status code of a specified website at regular intervals.",
	Long:  `Monitors the HTTP status code of a specified website at regular intervals.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Fprintln(cmd.OutOrStderr(), "no arguments provided")
			fmt.Fprintf(cmd.OutOrStderr(), "usage: %s\n", cmd.UseLine())
			os.Exit(1)
		}
		targetURL := args[0]

		flags, err := flags.Parse(cmd)
		if err != nil {
			fmt.Fprintln(cmd.OutOrStderr(), err)
			os.Exit(1)
		}
		options, err := monitorer.NewOptions(flags)
		if err != nil {
			fmt.Fprintln(cmd.OutOrStderr(), err)
			os.Exit(1)
		}

		m := monitorer.New(http.DefaultClient, targetURL, options)
		ctx, stop := signal.NotifyContext(
			context.Background(),
			os.Interrupt,
		)
		defer stop()

		m.Do(ctx)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.check-http-status.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().IntP(
		flags.INTERVAL_SECONDS,
		flags.INTERVAL_SECONDS_SHORTHAND,
		flags.DEFAULT_INTERVAL_SECONDS,
		"interval_seconds are interval time between monitoring HTTP requests.",
	)

	rootCmd.Flags().BoolP(
		flags.CREATE_LOG_FILE,
		flags.CREATE_LOG_FILE_SHORTHAND,
		flags.DEFAULT_CREATE_LOG_FILE,
		"create a file to log results. In default log file won't be created. Log file name format: check-http-status_<timestamp>.log",
	)

	rootCmd.Flags().IntP(
		flags.TIMEOUT_SECONDS,
		flags.TIMEOUT_SECONDS_SHORTHAND,
		flags.DEFAULT_TIMEOUT_SECONDS,
		"timeout in seconds for each HTTP request. If a response is not received within the specified time, the request will be considered failed.",
	)
}
