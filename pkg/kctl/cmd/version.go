package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"kun-blog-golang/pkg/kctl/clientset"
	"log"
)

type VersionInfo struct {
	Version   string
	GoVersion string
}

func NewVersionInfo() *VersionInfo {
	return &VersionInfo{
		Version:   "v0.1",
		GoVersion: "go1.7",
	}
}

// String 实现fmt的格式输出方法
func (this *VersionInfo) String() string {
	return fmt.Sprintf(
		"{version: %s, goversion: %s}",
		this.Version,
		this.GoVersion)
}

var VersionCMD = &cobra.Command{
	Use:          "version",
	Short:        "v",
	Example:      "kctl version",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("ctl version: %s \n", NewVersionInfo())
		serVersion, err := clientset.DefClientSet.V1().Version().Get()
		if err != nil {
			log.Panicln(err.Error())
		}
		fmt.Printf("server version: {version: %s, goversion: %s} \n",
			serVersion.Version,
			serVersion.GoVersion)
		return nil
	},
}
