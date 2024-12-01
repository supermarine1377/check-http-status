/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/spf13/cobra"
	"github.com/supermarine1377/check-http-status/internal/http_status"
	"github.com/supermarine1377/check-http-status/internal/log_files"
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
		intervalSeconds, err := cmd.Flags().GetInt(INTERVAL_SECONDS)
		if err != nil {
			fmt.Fprintln(cmd.OutOrStderr(), err)
			os.Exit(1)
		}

		createLogFile, err := cmd.Flags().GetBool(CREATE_LOG_FILE)
		if err != nil {
			fmt.Fprintln(cmd.OutOrStderr(), err)
			os.Exit(1)
		}

		logFile, err := log_files.New(createLogFile)
		if err != nil {
			fmt.Fprintln(cmd.OutOrStderr(), err)
			os.Exit(1)
		}

		timeoutSeconds, err := cmd.Flags().GetInt(TIMEOUT_SECONDS)
		if err != nil {
			fmt.Fprintln(cmd.OutOrStderr(), err)
			os.Exit(1)
		}

		m := http_status.NewMonitorer(targetURL, intervalSeconds, logFile)
		ctx, stop := signal.NotifyContext(
			context.Background(),
			os.Interrupt,
		)
		defer stop()
		ctx, cancel := context.WithTimeout(ctx, time.Duration(timeoutSeconds)*time.Second)
		defer cancel()

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

const INTERVAL_SECONDS = "interval-seconds"
const INTERVAL_SECONDS_SHORTHAND = "i"
const DEFAULT_INTERVAL_SECONDS = 10

const CREATE_LOG_FILE = "create-log-file"
const CREATE_LOG_FILE_SHORTHAND = "c"
const DEFAULT_CREATE_LOG_FILE = false

const TIMEOUT_SECONDS = "timeout-seconds"
const TIMEOUT_SECONDS_SHORTHAND = "t"
const DEFAULT_TIMEOUT_SECONDS = 30

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.check-http-status.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().IntP(
		INTERVAL_SECONDS,
		INTERVAL_SECONDS_SHORTHAND,
		DEFAULT_INTERVAL_SECONDS,
		"interval_seconds are interval time between monitoring HTTP requests.",
	)

	rootCmd.Flags().BoolP(
		CREATE_LOG_FILE,
		CREATE_LOG_FILE_SHORTHAND,
		DEFAULT_CREATE_LOG_FILE,
		"create a file to log results. In default log file won't be created. Log file name format: check-http-status_<timestamp>.log",
	)

	rootCmd.Flags().IntP(
		TIMEOUT_SECONDS,
		TIMEOUT_SECONDS_SHORTHAND,
		DEFAULT_TIMEOUT_SECONDS,
		"timeout in seconds for each HTTP request. If a response is not received within the specified time, the request will be considered failed.",
	)
}
