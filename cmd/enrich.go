package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/zeshi09/ipenrich/internal/enrich"
	"github.com/zeshi09/ipenrich/internal/parser"
	"github.com/zeshi09/ipenrich/model"
)

// ip := m.ips[m.index]
// m.geo[ip] = enrich.FetchGeoInfo(ip)
// m.abuse[ip] = enrich.FetchAbuseScore(ip)
// m.vt[ip] = enrich.FetchVTStats(ip)
// m.index++
// m.percent = float64(m.index) / float64(len(m.ips))

var logFile string
var check = make(map[string]string)

var EnrichCommand = &cobra.Command{
	Use:   "enrich",
	Short: "Enrich IPs from log file and output JSON",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			fmt.Println("Arguments parsing error, please see -h")
			os.Exit(1)
		}

		logFile := args[0]
		ips, err := parser.ReadingFile(logFile)
		if err != nil {
			return err
		}

		var results []model.EnrichedIP
		for _, ip := range ips {
			results = append(results, model.EnrichedIP{
				LogFile: logFile,
				Ip:      ip,
				Geo:     enrich.FetchGeoInfo(ip),
				Abuse:   enrich.FetchAbuseScore(ip),
				Vt:      enrich.FetchVTStats(ip),
			})
		}

		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		return enc.Encode(results)
	},
}

func init() {
	// Command.Flags().StringVarP(&logFile, "log", "l", "", "Path to log file")
	// Command.MarkFlagRequired("log")
	rootCmd.AddCommand(EnrichCommand)
}
