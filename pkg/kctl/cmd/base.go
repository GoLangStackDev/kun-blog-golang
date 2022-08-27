package cmd

import (
	"github.com/spf13/cobra"
	cmdDelete "kun-blog-golang/pkg/kctl/cmd/delete"
	cmdGet "kun-blog-golang/pkg/kctl/cmd/get"
	"log"
)

func RunCmd() {
	cmd := &cobra.Command{
		Use:          "kctl",
		Short:        "端博客",
		Example:      "kctl",
		SilenceUsage: true,
	}

	// 加入子模块
	cmd.AddCommand(
		VersionCMD,
		ApplyCMD,
		cmdGet.GetCMD,
		cmdDelete.DeleteCMD,
		configCMD)

	// 执行
	if err := cmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
