package get

import "github.com/spf13/cobra"

var GetCMD = &cobra.Command{
	Use:          "get",
	Example:      "kctl get posts -s xxx",
	SilenceUsage: true,
}

func init() {
	GetCMD.AddCommand(getPostsCMD)
}
