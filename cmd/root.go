/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/Ehco1996/ehco/internal/constant"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "ehco",
	Version: constant.Version,
	Short:   "ehco is a network relay tool and a typo :)",
	Run:     start,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("local", "l", "", "监听地址，例如 0.0.0.0:1234")

	rootCmd.Flags().StringP("remote", "r", "", "TCP 转发地址，例如 0.0.0.0:5201，通过 ws 隧道转发时应为 ws://0.0.0.0:2443")
	rootCmd.Flags().StringP("udp_remote", "u", "", "UDP 转发地址，例如 0.0.0.0:1234")

	rootCmd.Flags().StringP("listen_type", "lt", "raw", "监听类型，可选项有 raw,ws,wss,mwss")
	rootCmd.Flags().StringP("transport_type", "tt", "raw", "传输类型，可选选有 raw,ws,wss,mwss")

}

func start(cmd *cobra.Command, args []string) {
	println("haha")

	r, _ := cmd.Flags().GetString("local")

	println(r)

}
