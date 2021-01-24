package cmd

import "github.com/spf13/cobra"

var root = &cobra.Command{
	Use:   "goshm",
	Short: "Utils for get idx stock price",
}

func Execute() {
	root.AddCommand(fetchCmd())

	if err := root.Execute(); err != nil {
		panic(err)
	}
}
