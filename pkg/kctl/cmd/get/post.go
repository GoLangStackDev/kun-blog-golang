package get

import (
	"fmt"
	"github.com/spf13/cobra"
	"kun-blog-golang/pkg/kctl/clientset"
)

var inputSlug string

func init() {
	getPostsCMD.Flags().StringVar(&inputSlug, "slug", "", "--slug")
}

var getPostsCMD = &cobra.Command{
	Use:          "posts",
	Example:      "kctl get posts xxx",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if inputSlug != "" {
			getPostsBySlug(inputSlug)
		} else {
			getAllPosts()
		}
		return nil
	},
}

func getAllPosts() {
	rst, _ := clientset.DefClientSet.V1().Posts().List()
	if len(rst) == 0 {
		fmt.Println("你还没发布文章")
	}
	for _, post := range rst {
		fmt.Println(post.Slug, post.Title)
	}
}

func getPostsBySlug(slug string) {
	post, err := clientset.DefClientSet.V1().Posts().GetBySlug(slug)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(post.Slug, post.Title,
			post.CreatedTime.Format("2006-01-02 15:03:04"))
	}
}
