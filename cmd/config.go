/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	//"fmt"

	//"github.com/spf13/cobra"
	ksUtils "github.com/stellar/kelp/support/utils"
)

type CuttlefishConfig struct{
	Users	[]UserInput	`valid:"-" toml:"USERS"`
}

type UserInput struct {
	UserIdCode	string			`valid:"-" toml:"USERID"`
	CbSetUps	[]SetupInput	`valid:"-" toml:"CBSETUPS"`
}

type SetupInput struct {
	CBId					int		`valid:"-" toml:"CBID"`
	CBSendMaxNo				int		`valid:"-" toml:"CB_SEND_MAX"`
	SourcePublicKey			string	`valid:"-" toml:"SOURCE_PKEY"`
	AssetCode				string	`valid:"-" toml:"ASSET_CODE"`
	IssuerPublicKey			string	`valid:"-" toml:"ISSUER_PKEY"`
	IssuerSecretKey			string	`valid:"-" toml:"ISSUER_SKEY"`
	CBAmountToSend			int		`valid:"-" toml:"AMOUNT"`
	DestinationPublicKey	string	`valid:"-" toml:"DEST_PKEY"`
	Enabled					bool    `valid:"-" toml:"ENABLED"`
	PublicNetwork			bool    `valid:"-" toml:"PUBLIC_NETWORK"`
}

// String impl.
func (c CuttlefishConfig) String() string {
	return ksUtils.StructString(c,1, nil)
}

/* // configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("config called")
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
} */
