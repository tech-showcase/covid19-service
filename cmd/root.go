package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "covid19-service",
	Short: "Covid19 Service is a service that provide data about Covid-19",
	Long: `A service that can be a portal for all information about Covid-19
				from all over the world`,
}

func Execute() error {
	return rootCmd.Execute()
}
