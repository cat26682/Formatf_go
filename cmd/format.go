/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// formatCmd represents the format command
var formatCmd = &cobra.Command{
	Use:   "format",
	Short: "Изменяет формат файла фото",
	Long: `Изменяет формат файла фото на jpg, jpeg, png, gif, bmp, tiff, webp, heif, heic, raw, dng, psd, avif, ico, cur, pcx, ppm, pgm, pbm, pgf, xcf, svg
var th = "."`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("format called")
	},
}

func init() {
	rootCmd.AddCommand(formatCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// formatCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// formatCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
