package snyk

import (
	"fmt"

	"github.com/snyk/parlay/lib"

	"github.com/package-url/packageurl-go"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

func NewPackageCommand(logger zerolog.Logger) *cobra.Command {
	cmd := cobra.Command{
		Use:   "package <purl>",
		Short: "Return package vulnerabilities from Snyk",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			logger.Debug().Str("purl", args[0]).Msg("Looking up package vulnerabilities from Snyk")
			purl, err := packageurl.FromString(args[0])
			if err != nil {
				logger.Fatal().Err(err)
			}
			resp, err := lib.GetPackageVulnerabilities(purl)
			if err != nil {
				logger.Fatal().Err(err)
			}
			fmt.Print(string(resp.Body))
		},
	}
	return &cmd
}