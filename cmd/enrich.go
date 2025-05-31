package cmd

import (
	"encoding/json"
	"fmt"
	"net"
	"os"

	"github.com/spf13/cobra"
	"github.com/zeshi09/ipenrich/internal/db"
	"github.com/zeshi09/ipenrich/internal/enrich"
	"github.com/zeshi09/ipenrich/internal/parser"

	"github.com/zeshi09/ipenrich/model"
)

var EnrichCommand = &cobra.Command{
	Use:   "enrich",
	Short: "Enrich IPs from log file and output JSON",
	RunE: func(cmd *cobra.Command, args []string) error {
		jsonFlag, _ := cmd.Flags().GetBool("json")

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
			country, city, org := enrich.FetchGeoInfo(ip)

			dns, _ := net.LookupAddr(ip)
			harmless, malicious, suspicious, undetected := enrich.FetchVTStats(ip)
			results = append(results, model.EnrichedIP{
				LogFile:      logFile,
				Ip:           ip,
				Dns:          dns[0],
				Country:      country,
				City:         city,
				Org:          org,
				AbuseScore:   enrich.FetchAbuseScore(ip),
				VTMalicious:  malicious,
				VTSuspicious: suspicious,
				VTHarmless:   harmless,
				VTUndetected: undetected,
			})

		}

		db.SaveToPostgres(results)
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")

		if jsonFlag {
			f, _ := os.Create("output.json")
			defer f.Close()
			f_json, _ := json.MarshalIndent(results, "", "  ")
			f.Write(f_json)
		}

		return enc.Encode(results)
	},
}

func init() {
	// Command.Flags().StringVarP(&logFile, "log", "l", "", "Path to log file")
	// Command.MarkFlagRequired("log")
	EnrichCommand.Flags().Bool("json", false, "A show for full json")

	rootCmd.AddCommand(EnrichCommand)
}
