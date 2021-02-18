package cmd

import (
	"fmt"
	"json/utils"
	"time"

	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "returns metadata in json format",
	Long:  `json generate [all | ip | ts | user | userdir]:  returns user's metadata`,
	Run: func(cmd *cobra.Command, args []string) {

		timestamp := time.Now().Format(time.RFC850)
		ip := utils.GetIPAdress().String()
		userName := utils.GetUserData().Username
		homeDir := utils.GetUserData().HomeDir

		allMetaData := utils.AllMetaData{
			UserName:  userName,
			HomeDir:   homeDir,
			LocalHost: "localhost",
			IP:        ip,
			Timestamp: timestamp}

		if len(args) == 1 {
			if args[0] == "all" || args[0] == "ip" || args[0] == "ts" || args[0] == "user" || args[0] == "userdir" {
				switch args[0] {
				case "all":
					utils.OutputJSONall(&allMetaData)

				case "ip":
					utils.OutputJSONip(&ip)

				case "ts":
					utils.OutputJSONts(&timestamp)

				case "user":
					utils.OutputJSONuserName(&userName)

				case "userdir":
					utils.OutputJSONhomeDir(&homeDir)

				}

			} else {
				fmt.Println("Argument does not exist, Please try one of the following: all | ip | ts | user | userdir")
			}
		} else {
			fmt.Println("Please choose from one of the following arguments: all | ip | ts | user | userdir")

		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
