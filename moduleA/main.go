package main

import (
	"fmt"

	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use: "moduleA",
		Run: func(cmd *cobra.Command, args []string) {
			items := lo.Filter([]int{1, 2, 3, 4, 5}, func(x int, _ int) bool {
				return x%2 == 0
			})
			fmt.Println("Even numbers:", items)
		},
	}
	rootCmd.Execute()
}