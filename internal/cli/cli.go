package cli

import (
	"fmt"
	"github.com/EuricoCruz/stress_test_challenge/internal/tester"
	"github.com/spf13/cobra"
)

var url string
var totalRequests int
var concurrency int

var rootCmd = &cobra.Command{
	Use:   "stress_test_challenge",
	Short: "Stress test CLI em Go",
	RunE: func(cmd *cobra.Command, args []string) error {
		svc := tester.NewService()
		report := svc.Run(url, totalRequests, concurrency)
		fmt.Println(report.String())
		return nil
	},
}

func Execute() {
	rootCmd.Flags().StringVar(&url, "url", "", "URL do serviço a ser testado")
	rootCmd.Flags().IntVar(&totalRequests, "requests", 1, "Número total de requests")
	rootCmd.Flags().IntVar(&concurrency, "concurrency", 1, "Número de chamadas simultâneas")

	rootCmd.MarkFlagRequired("url")

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
