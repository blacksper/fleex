package cmd

import (
	"log"

	"github.com/FleexSecurity/fleex/pkg/controller"
	"github.com/FleexSecurity/fleex/pkg/utils"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an existing fleet or even a single box",
	Run: func(cmd *cobra.Command, args []string) {
		var token string

		proxy, _ := rootCmd.PersistentFlags().GetString("proxy")
		utils.SetProxy(proxy)

		name, _ := cmd.Flags().GetString("name")
		providerFlag, _ := cmd.Flags().GetString("provider")

		if globalConfig.Settings.Provider != providerFlag && providerFlag == "" {
			providerFlag = globalConfig.Settings.Provider
		}

		provider := controller.GetProvider(providerFlag)
		if provider == -1 {
			log.Fatal("invalid provider")
		}

		controller.DeleteFleet(name, token, provider)

	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().StringP("name", "n", "pwn", "Fleet name. Boxes will be named [name]-[number]")
	deleteCmd.Flags().StringP("provider", "p", "", "Service provider (Supported: linode, digitalocean, vultr)")

}
