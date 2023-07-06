package cli

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	scanErr "piiScanner/pkg/errors"
)

func NewDetectRootCmd() *cobra.Command {
	showVersion := false
	// conf := config.New()

	rootCmd := &cobra.Command{
		Use:   "piiScanner",
		Short: "`piiScanner`",
		Long:  "`piiScanner`",
		Run: func(cmd *cobra.Command, args []string) {
			if showVersion {
				log.Info().Msg("Welcome to piiScanner")
			} else {
				_ = cmd.Usage()
				cmd.SilenceErrors = false
			}
		},
	}

	// "detect"
	rootCmd.AddCommand(newDetectCmd())

	// "version"
	rootCmd.Flags().BoolVarP(&showVersion, "version", "v", false, "show the version and exit")

	return rootCmd
}

func newDetectCmd() *cobra.Command {
	var dbSource, dbType string

	// "detect"
	detectCmd := &cobra.Command{
		Use:     "detect <config>",
		Aliases: []string{"serve"},
		Short:   "`detect` scans databases for PIIs",
		Long:    "`detect` scans databases for PIIs",
		RunE: func(cmd *cobra.Command, args []string) error {
			// if len(args) > 0 {
			// 	if err := LoadConfiguration(conf, args[0]); err != nil {
			// 		panic(err)
			// 	}
			// }

			sourcePath, err := cmd.Flags().GetString("add-source-path")
			if err != nil {
				cmd.SilenceUsage = true

				return err
			}
			if sourcePath == "" {
				return scanErr.ErrNoSourcePath
			}

			sourceType, err := cmd.Flags().GetString("db-type")
			if err != nil {
				cmd.SilenceUsage = true

				return err
			}
			if sourceType == "" {
				return scanErr.ErrNoSourceType
			}


			return nil
		},
	}

	detectCmd.Flags().StringVarP(&dbSource, "add-source-path", "A", "",
		"add relative path to source")
	detectCmd.Flags().StringVarP(&dbType, "db-type", "T", "",
		"specify db type of the source")

	return detectCmd
}
