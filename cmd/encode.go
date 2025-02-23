package cmd

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var usePercent20 bool

// encodeCmd handles URL encoding
var encodeCmd = &cobra.Command{
	Use:   "encode",
	Short: "URL encode input",
	Run: func(cmd *cobra.Command, args []string) {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			encoded := url.QueryEscape(scanner.Text())
			// Replace '+' with '%20' if flag is set
			if usePercent20 {
				encoded = strings.ReplaceAll(encoded, "+", "%20")
			}
			fmt.Println(encoded)
		}
	},
}

func init() {
	encodeCmd.Flags().BoolVar(&usePercent20, "use-percent-20", false, "Encode spaces as %20 instead of +")
}
