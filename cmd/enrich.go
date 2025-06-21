package cmd

import (
	"encoding/json"
	"fmt"
	"net"
	"os"

	"github.com/spf13/cobra"
	"github.com/zeshi09/ipenrich/internal/enrich"
	"github.com/zeshi09/ipenrich/internal/parser"

	"github.com/zeshi09/ipenrich/model"
)

var EnrichCommand = &cobra.Command{
	Use:   "enrich",
	Short: "Enrich IPs from log file and output JSON",
	RunE: func(cmd *cobra.Command, args []string) error {
		ipFlag, _ := cmd.Flags().GetBool("ip")
		hashFlag, _ := cmd.Flags().GetBool("hash")

		if len(args) == 0 {
			fmt.Println("Arguments parsing error, please see -h")
			os.Exit(1)
		}

		logFile := args[0]

		if ipFlag {
			ips, err := parser.ReadingFileForIP(logFile)
			if err != nil {
				return err
			}

			var results []model.EnrichedIP
			for _, ip := range ips {
				country, city, org := enrich.FetchGeoInfo(ip)

				var dnsname string
				dns, err := net.LookupAddr(ip)
				if err != nil {
					dnsname = "N/A"
				} else {
					dnsname = dns[0]
				}

				harmless, malicious, suspicious, undetected := enrich.FetchVTStatsIP(ip)
				results = append(results, model.EnrichedIP{
					LogFile:      logFile,
					Ip:           ip,
					Dns:          dnsname,
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

			// db.SaveToPostgres(results)
			enc := json.NewEncoder(os.Stdout)
			enc.SetIndent("", "  ")

			// Save outputjson
			outputfilename := fmt.Sprintf("%s_output.json", logFile)
			f, _ := os.Create(outputfilename)
			defer f.Close()
			f_json, _ := json.MarshalIndent(results, "", "  ")
			f.Write(f_json)

			defer fmt.Printf("Output was saved to %s", outputfilename)
			return enc.Encode(results)
		} else if hashFlag {
			return nil
		} else {
			return nil
		}

	},
}

func init() {
	// Command.Flags().StringVarP(&logFile, "log", "l", "", "Path to log file")
	// Command.MarkFlagRequired("log")
	EnrichCommand.Flags().Bool("ip", false, "Parse IOC file for IP's")
	EnrichCommand.Flags().Bool("hash", false, "Parse IOC file for hashes")

	rootCmd.AddCommand(EnrichCommand)
}
