package main

import (
	"fmt"

	"github.com/andrey-solomenniy-test/standarta/app"
	"github.com/spf13/cobra"
)

//"github.com/andrey-solomenniy-test/standarta/app"

// rootCmd is the entry point for this binary
// var (
// 	rootCmd = &cobra.Command{
// 		Use:   "standartacli",
// 		Short: "StandartACoin light-client",
// 	}
// )

func main() {
	// disable sorting
	cobra.EnableCommandSorting = false

	// get the codec
	cdc := app.MakeCodec()
	fmt.Println(cdc)

	// // TODO: setup keybase, viper object, etc. to be passed into
	// // the below functions and eliminate global vars, like we do
	// // with the cdc

	// // add standard rpc, and tx commands
	// rpc.AddCommands(rootCmd)
	// rootCmd.AddCommand(client.LineBreak)
	// tx.AddCommands(rootCmd, cdc)
	// rootCmd.AddCommand(client.LineBreak)

	// // add query/post commands (custom to binary)
	// rootCmd.AddCommand(
	// 	client.GetCommands(
	// 		authcmd.GetAccountCmd("main", cdc, types.GetAccountDecoder(cdc)),
	// 	)...)
	// rootCmd.AddCommand(
	// 	client.PostCommands(
	// 		bankcmd.SendTxCmd(cdc),
	// 	)...)
	// rootCmd.AddCommand(
	// 	client.PostCommands(
	// 		ibccmd.IBCTransferCmd(cdc),
	// 	)...)
	// rootCmd.AddCommand(
	// 	client.PostCommands(
	// 		ibccmd.IBCRelayCmd(cdc),
	// 		simplestakingcmd.BondTxCmd(cdc),
	// 	)...)
	// rootCmd.AddCommand(
	// 	client.PostCommands(
	// 		simplestakingcmd.UnbondTxCmd(cdc),
	// 	)...)

	// // add proxy, version and key info
	// rootCmd.AddCommand(
	// 	client.LineBreak,
	// 	lcd.ServeCommand(cdc),
	// 	keys.Commands(),
	// 	client.LineBreak,
	// 	version.VersionCmd,
	// )

	// // prepare and add flags
	// executor := cli.PrepareMainCmd(rootCmd, "BC", os.ExpandEnv("$HOME/.standartacli"))
	// executor.Execute()
}
