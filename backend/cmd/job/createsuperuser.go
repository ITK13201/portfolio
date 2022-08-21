package job

import (
	"github.com/ITK13201/portfolio/backend/cmd"

	"github.com/ITK13201/portfolio/backend/jobs/createsuperuser"
	"github.com/spf13/cobra"
)

// createsuperuserCmd represents the createsuperuser command
var createsuperuserCmd = &cobra.Command{
	Use:   "createsuperuser",
	Short: "Create Super User",
	Long:  "Create Super User",
	Run: func(cmd *cobra.Command, args []string) {
		createsuperuser.Run(cmd, args)
	},
}

func init() {
	cmd.JobCmd.AddCommand(createsuperuserCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createsuperuserCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createsuperuserCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
