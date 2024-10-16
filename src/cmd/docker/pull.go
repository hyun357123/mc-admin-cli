package docker

import (
	"fmt"

	"github.com/mc-admin-cli/mcc/src/common"
	"github.com/spf13/cobra"
)

// pullCmd represents the pull command
var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pull images of M-CMP System containers",
	Long:  `Pull images of M-CMP System containers`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("\n[Pull images of M-CMP System containers]")
		fmt.Println()

		if DockerFilePath == "" {
			fmt.Println("file is required")
		} else {

			cmdStr := fmt.Sprintf("COMPOSE_PROJECT_NAME=%s docker compose -f %s pull", ComposeProjectName, DockerFilePath)
			//fmt.Println(cmdStr)
			common.SysCall(cmdStr)
		}

	},
}

func init() {
	infraCmd.AddCommand(pullCmd)

	pf := pullCmd.PersistentFlags()
	pf.StringVarP(&DockerFilePath, "file", "f", DefaultDockerComposeConfig, "User-defined configuration file")
	//	cobra.MarkFlagRequired(pf, "file")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pullCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pullCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
