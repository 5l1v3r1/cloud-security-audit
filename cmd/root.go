package cmd

import (
	"fmt"
	"os"
	"tyr/configuration"
	"tyr/scanner"
	"tyr/tyrlogger"

	"github.com/spf13/cobra"
)

// var cfgFile string
var config configuration.Config
var logger = tyrlogger.GetInstance()

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tyr",
	Short: "Scan for vulnerabilities in your AWS Account.",
	Long:  `Scan for vulnerabilities in your AWS Account.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := scanner.Run(&config)
		if err != nil {
			logger.Fatalln(err)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {

	rootCmd.Flags().StringVarP(&config.Region, "region", "r", "", "specify aws region to scan your account,e.g. --region us-east-1")
	rootCmd.MarkFlagRequired("region")

	rootCmd.Flags().StringVarP(&config.Service, "service", "s", "", "specify aws service to scan in your account,e.g. --service [ec2:x,ec2:image]")
	rootCmd.MarkFlagRequired("service")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
