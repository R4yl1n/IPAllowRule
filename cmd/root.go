/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"net"
	"os"

	"github.com/abdfnx/gosh"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "IPAlowRule",
	Short: "Based on the inputs, the script will ADD / DELETE an iptables rule",
	Long: `Command to insert an allow rule from the given IP adress to any connection:

examples of usage are:
./IPAllowRule ACTIVE_DIRECTORY 192.168.1.1 connect
./IPAllowRule ACTIVE_DIRECTORY 192.168.1.1 disconnect
./IPAllowRule <NAME> <IP_ADDRESS> <ACTION>


Actions should be "connect" if you want to add an Ip 
or "disconnect" if you want to delete it
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	Run: func(cmd *cobra.Command, args []string) {

		log.Println(args)

		if len(args) >= 3 && args[0] != "" && args[2] == "connect" {
			name := args[0]
			ipadress := args[1]

			if ip_is_valid(ipadress) {
				add_iptables_rule(name, ipadress)
			}
		}

		if len(args) >= 3 && args[0] != "" && args[2] == "disconnect" {
			name := args[0]
			ipadress := args[1]
			if ip_is_valid(ipadress) {
				delete_iptables_rule(name, ipadress)
			}
		} else {
			log.Print("please check your Input")
			log.Fatal("to get help run: ipallowrule -h")
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

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func add_iptables_rule(name string, ip_adress string) {
	log.Printf("started adding %v with the adress %v", name, ip_adress)

	cmd := "sudo iptables -A INPUT -s" + ip_adress + " -j ACCEPT -m comment --comment '" + name + "' && sudo iptables --list"
	gosh.ShellCommand(cmd)
	log.Println("Added Succesfully")

}

func delete_iptables_rule(name string, ip_adress string) {
	log.Printf("started deleting %v with the adress %v", name, ip_adress)

	cmd := "sudo iptables -D INPUT -s" + ip_adress + " -j ACCEPT -m comment --comment '" + name + "' && sudo iptables --list"
	gosh.ShellCommand(cmd)
	log.Println("Deleted Succesfully")

}

func ip_is_valid(ip_adress string) bool {
	ip := net.ParseIP(ip_adress)
	if ip == nil {
		log.Fatal("IP address is not valid")
		return false
	}
	if ip.To4() != nil {
		log.Println("IP address is v4")
		return true
	} else {
		log.Fatal("IP address is v6 please use a ip4 adress")
		return false
	}
}
