package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version = &cobra.Command{
	Use:   "version",
	Short: "версия программы",
	Long:  `версия программы`,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("0.2")
	},
}

func init() {
	rootCmd.AddCommand(version) // добавляем новую команду к корневой команде
}
