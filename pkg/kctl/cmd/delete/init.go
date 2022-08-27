package delete

import "github.com/spf13/cobra"

var DeleteCMD = &cobra.Command{
	Use:          "delete",
	Example:      "kctl delete posts xxx",
	SilenceUsage: true,
}

func init() {
	DeleteCMD.AddCommand(deletePostsCMD)
}
