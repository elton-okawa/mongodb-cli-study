package user

import (
	"github.com/spf13/cobra"
)

var databaseName string
var collectionName string = "users"

var rootCmd = &cobra.Command{
	Use:   "users",
	Short: "CRUD for user data",
}

func GetCommands() *cobra.Command {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	rootCmd.PersistentFlags().StringVarP(&databaseName, "database", "d", "blog", "Database name")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// templateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	return rootCmd
}
