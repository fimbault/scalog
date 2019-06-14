package cmd

import (
	"github.com/scalog/scalogger/order"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// orderCmd represents the order command
var orderCmd = &cobra.Command{
	Use:   "order",
	Short: "The ordering layer",
	Long:  `The ordering layer`,
	Run: func(cmd *cobra.Command, args []string) {
		order.Start()
	},
}

func init() {
	RootCmd.AddCommand(orderCmd)
	orderCmd.PersistentFlags().IntP("id", "i", 0, "Replica index")
	viper.BindPFlag("id", orderCmd.PersistentFlags().Lookup("id"))
}