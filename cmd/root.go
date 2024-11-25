/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "IPAlowRule",
	Short: "Based on the inputs, the script will ADD / DELETE an iptables rule",
	Long: `Command to insert an allow rule from the given IP adress to any connection:

examples of commands are:
./IPAllowRule <NAME> <IP_ADDRESS> <ACTION>
go run main.go <NAME> <IP_ADDRESS> <ACTION>

Actions should be connect if you want to add an Ip pr disconnect if you want to delete it
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 1 && args[0] != "" && args[2] == "connect" {
			name := args[0]
			ipadress := args[1]
			ipallowrule.add_iptables_rule(name, ipadress)

		}
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ControllsDatabase.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
