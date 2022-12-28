package command

import (
	"fmt"
	"ginFrame"
	"ginFrame/config"
	"github.com/spf13/cobra"
	"os"
)

func Init() {
	var commands []*cobra.Command

	var cmd *cobra.Command

	cmd = &cobra.Command{
		Use:   "server",
		Short: "Server run",
		Run: func(cmd *cobra.Command, args []string) {
			ginFrame.New()
			//fmt.Println("Pull: " + strings.Join(args, " "))
		},
	}

	cmd.Flags().StringVarP(&config.Flag.Path, "path", "p", "", "config path")

	commands = append(commands, cmd)

	commands = append(commands, &cobra.Command{
		Use:   "version",
		Short: "Push an image or a repository from a registry",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("version: 12222")
			os.Exit(0)
		},
	})

	var rootCmd = &cobra.Command{Use: "frame"}
	rootCmd.AddCommand(commands...)
	if err := rootCmd.Execute(); err != nil {
		//fmt.Println("err = ", err)
		os.Exit(0)
	}
}
