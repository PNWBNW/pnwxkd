package main

import (
    "os"

    "github.com/cosmos/cosmos-sdk/server"
    "github.com/cosmos/cosmos-sdk/server/cmd"
    svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
    "github.com/spf13/cobra"

    "github.com/pnwxkd/app" // adjust import path if needed
)

func main() {
    rootCmd, _ := NewRootCmd()
    if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
        os.Exit(1)
    }
}

// NewRootCmd creates the root command for pnwxkd
func NewRootCmd() (*cobra.Command, app.EncodingConfig) {
    encoding := app.MakeEncodingConfig()
    initRootCmd := &cobra.Command{
        Use:   "pnwxkd",
        Short: "Proven National Workers Xero-Knowledge Daemon",
        PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
            return server.InterceptConfigsPreRunHandler(cmd)
        },
    }

    appCreator := app.NewAppCreator(encoding)

    return cmd.NewRootCmd(
        initRootCmd,
        encoding,
        appCreator,
    ), encoding
}
