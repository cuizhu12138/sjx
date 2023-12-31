/*
Copyright © 2023 sjx <540643428@qq.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var OccurSimultaneously string = ""

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "PortScan",
	Short: "可以用这个工具进行端口扫描和主机发现",
	Long:  `你可以使用本工具对整个网络进行各种形式的端口扫描和主机发现，也可以仅对特定主机进行端口扫描,同时也支持扫描特定端口`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
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

	rootCmd.PersistentFlags().StringVarP(&OccurSimultaneously, "OccurSimultaneously", "O", "", "Set Max Goroutine")
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.PortScan.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
