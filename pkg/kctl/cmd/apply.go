package cmd

import (
	"github.com/spf13/cobra"
	"kun-blog-golang/pkg/kctl/clientset"
	"kun-blog-golang/pkg/utils"
	"log"
)

// 是否覆盖
var isForce bool

func init() {
	ApplyCMD.Flags().BoolVar(&isForce, "force", false, "--force")
	ApplyCMD.Flags().StringP("file", "f", "", "-f xxx.md")
}

var ApplyCMD = &cobra.Command{
	Use:          "apply",
	Example:      "kctl apply -f xxx",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		localFilePath, err := cmd.Flags().GetString("file")
		if err != nil {
			log.Fatalln(err)
		}
		return uploadFile(localFilePath)
	},
}

// uploadFile 上传文件的逻辑
func uploadFile(path string) error {
	fInfo, err := utils.FileInfo(path)
	if err != nil {
		log.Fatalln(err)
	}
	fb := utils.MustLoadFile(path)
	rst, err := clientset.DefClientSet.V1().Posts().Apply(fInfo.Name(), fb, isForce)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println(rst.Msg)
	}
	return nil
}
