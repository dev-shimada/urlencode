package cmd

import (
	"bufio"
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// decodeCmd handles URL decoding
var decodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "URL decode input",
	Run: func(cmd *cobra.Command, args []string) {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			decoded, err := url.QueryUnescape(scanner.Text())
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error decoding: %v\n", err)
				os.Exit(1)
			}
			fmt.Println(decoded)
		}
	},
}
