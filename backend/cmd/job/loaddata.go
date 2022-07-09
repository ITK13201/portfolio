package job

import (
	"github.com/ITK13201/portfolio/backend/jobs"

	"github.com/ITK13201/portfolio/backend/cmd"
	"github.com/spf13/cobra"
)

// loaddataCmd represents the loaddata command
var loaddataCmd = &cobra.Command{
	Use:   "loaddata",
	Short: "Load initial data",
	Long: `Load initial data
Caution: All table contents in the specified domain/model will be deleted before execution.`,
	Run: func(cmd *cobra.Command, args []string) {
		jobs.Run(cmd, args)
	},
}

func init() {
	cmd.JobCmd.AddCommand(loaddataCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loaddataCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loaddataCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
