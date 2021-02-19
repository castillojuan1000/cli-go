package cmd

import (
	"fmt"
	"json/utils"
	"time"

	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
//creates a 'struct' variable
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "returns metadata in json format",
	Long:  `json generate all:  returns user's metadata`,
	Run: func(cmd *cobra.Command, args []string) {

		//Metedata
		timestamp := time.Now().Format(time.RFC850)
		ip := utils.GetIPAdress().String()
		userName := utils.GetUserData().Username
		homeDir := utils.GetUserData().HomeDir

		//Contains a collection of fileds with user's metadata
		allMetaData := utils.AllMetaData{
			UserName:  userName,
			HomeDir:   homeDir,
			LocalHost: "localhost",
			IP:        ip,
			Timestamp: timestamp}

		// getting the flags value, their default value is false
		tsFlagStatus, _ := cmd.Flags().GetBool("timestamp")
		ipFlatStatus, _ := cmd.Flags().GetBool("ipaddress")
		userFlagStatus, _ := cmd.Flags().GetBool("user")
		homedirFlagStatus, _ := cmd.Flags().GetBool("homedir")

		//if argument is 'all' it will output all metadata in json format and create a json file
		if len(args) >= 1 {
			if args[0] == "all" {
				utils.OutputJSONall(&allMetaData)

			} else {
				fmt.Println("Argument does not exist, Please try : all ")
			}
		}

		//flag will output its value in json format
		switch true {

		case tsFlagStatus:
			utils.OutputJSONts(&timestamp)
		case ipFlatStatus:
			utils.OutputJSONip(&ip)
		case homedirFlagStatus:
			utils.OutputJSONhomeDir(&homeDir)
		case userFlagStatus:
			utils.OutputJSONuserName(&userName)

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
	generateCmd.Flags().BoolP("timestamp", "t", false, "outputs timestamp into JSON format")
	generateCmd.Flags().BoolP("ipaddress", "i", false, "outputs ip address into JSON format")
	generateCmd.Flags().BoolP("user", "u", false, "outputs user name into JSON format")
	generateCmd.Flags().BoolP("homedir", "d", false, "outputs homedir  into JSON format")
}
