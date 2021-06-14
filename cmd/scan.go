package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/sw33tLie/fleex/pkg/controller"
	scan "github.com/sw33tLie/fleex/pkg/scan"
)

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Distributed scanning",
	Run: func(cmd *cobra.Command, args []string) {
		var token string

		command, _ := cmd.Flags().GetString("command")
		delete, _ := cmd.Flags().GetBool("delete")
		fleetName, _ := cmd.Flags().GetString("name")
		input, _ := cmd.Flags().GetString("input")
		output, _ := cmd.Flags().GetString("output")

		port, _ := cmd.Flags().GetInt("port")
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")

		provider := controller.GetProvider(viper.GetString("provider"))

		switch provider {
		case controller.PROVIDER_LINODE:
			token = viper.GetString("linode.token")

		case controller.PROVIDER_DIGITALOCEAN:
			token = viper.GetString("digitalocean.token")
		}

		scan.Start(fleetName, command, delete, input, output, token, port, username, password, provider)

	},
}

func init() {
	rootCmd.AddCommand(scanCmd)
	scanCmd.Flags().StringP("name", "n", "pwn", "Fleet name")
	scanCmd.Flags().StringP("command", "c", "whoami", "Command to send. Supports {{INPUT}} and {{OUTPUT}}")
	scanCmd.Flags().StringP("input", "i", "", "Input file")
	scanCmd.Flags().StringP("output", "o", "scan-results.txt", "Output file path")
	scanCmd.Flags().IntP("port", "P", 2266, "SSH port")
	scanCmd.Flags().StringP("username", "u", "op", "SSH username")
	scanCmd.Flags().StringP("password", "p", "1337superPass", "SSH password")
	scanCmd.Flags().BoolP("delete", "d", false, "Delete boxes as soon as they finish their job")
}
