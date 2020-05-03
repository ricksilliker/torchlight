package main

import (
	"github.com/ricksilliker/torchlight"
	"os"
	"github.com/spf13/cobra"
)

var ScenePath string
var Verbose bool

var rootCmd = &cobra.Command{
	Use:   "torchlight",
	Short: "Maya scene parser.",
	Long: "Parse maya scenes.",
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&ScenePath, "scene", "s", "", "Absolute file path to the maya scene.")
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "Verbosity of logs.")
}

func Execute() {
	logger := torchlight.GetLogger(Verbose)

	if err := rootCmd.Execute(); err != nil {
		logger.WithError(err).Error("failed to execute torchlight")
		os.Exit(1)
	}

	logger.Debug("Do work.")
}

func main() {
	Execute()
}