package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/zeshi09/ipenrich/internal/parser"
)

var rootCmd = &cobra.Command{
	Use:   "ipenrich <path_to_log_file>",
	Short: "IPEnrich is a very cool auth.log parser and IP addresses enricher for DM TI Specialist and my bro\n\nPlease, don't forget export apikeys:\n\nexport ABUSEIPDB_API_KEY=your-abuse-key\nexport VT_API_KEY=your-virustotal-key",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Arguments parsing error, please see -h")
			os.Exit(1)
		}

		logFile := args[0]

		ips, err := parser.ReadingFile(logFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Parsing file error: %s\n", err)
			os.Exit(1)
		}

		fmt.Printf("%d IPs was found in %s\n", len(ips), logFile)
	},
}

func Execute() {
	// rootCmd.AddCommand(enrich.Command)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
