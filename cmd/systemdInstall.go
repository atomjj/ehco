/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	SystemDFilePath = "/etc/systemd/system/ehco.service"
	SystemDTemplate = `# Ehco service
[Unit]
Description=ehco
After=network.target

[Service]
LimitNOFILE=65535
ExecStart=/root/ehco -c ""
Restart=always

[Install]
WantedBy=multi-user.target
`
)

// systemdInstallCmd represents the systemdInstall command
var systemdInstallCmd = &cobra.Command{
	Use:   "systemdInstall",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("systemdInstall called")
	},
}

func init() {

	// rootCmd.AddCommand(systemdInstallCmd)

}

// app.Commands = []*cli.Command{
// 	{
// 		Name:  "install",
// 		Usage: "install ehco systemd service",
// 		Action: func(c *cli.Context) error {
// 			fmt.Printf("Install ehco systemd file to `%s`\n", SystemFilePath)

// 			if _, err := os.Stat(SystemFilePath); err != nil && os.IsNotExist(err) {
// 				f, _ := os.OpenFile(SystemFilePath, os.O_CREATE|os.O_WRONLY, 0644)
// 				if _, err := f.WriteString(SystemDTMPL); err != nil {
// 					logger.Fatal(err)
// 				}
// 				f.Close()
// 			}

// 			command := exec.Command("vi", SystemFilePath)
// 			command.Stdin = os.Stdin
// 			command.Stdout = os.Stdout
// 			return command.Run()
// 		},
// 	},
// }
// return app
