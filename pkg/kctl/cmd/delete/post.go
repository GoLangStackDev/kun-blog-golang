package delete

import (
	"fmt"
	"github.com/spf13/cobra"
	"kun-blog-golang/pkg/kctl/clientset"
)

var inputSlug string

func init() {
	deletePostsCMD.Flags().StringVar(&inputSlug, "slug", "", "--slug")
}

var deletePostsCMD = &cobra.Command{
	Use:          "posts",
	Example:      "kctl delete posts -s xxx",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if inputSlug != "" {
			deletePostsBySlug(inputSlug)
		}
		return nil
	},
}

func deletePostsBySlug(slug string) {
	err := clientset.DefClientSet.V1().Posts().DeleteBySlug(slug)
	if err == nil {
		fmt.Println("删除成功！")
	} else {
		fmt.Println(err)
	}
}
