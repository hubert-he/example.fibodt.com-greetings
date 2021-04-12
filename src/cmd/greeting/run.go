package greeting

import (
	"fmt"
	"github.com/hubert-he/translate/src"
	"github.com/spf13/cobra"
)

var StartCmd = &cobra.Command{
	Use:     "hello",
	Short:   "start",
	Example: "bin/greetings hello",
	RunE: func(cmd *cobra.Command, args []string) error {
		lang, err := cmd.Flags().GetString("lang")
		if err != nil {
			return err
		}
		return run(lang)
	},
}

func init() {
	StartCmd.PersistentFlags().StringP("lang", "l", "", "language")
}

func run(language string) error {
	// start greeting
	fmt.Println(lang.Tranlate(language))
	//	fmt.Println("你好，Go语言")
	return nil
}
