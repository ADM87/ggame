package main

import (
	"io"

	resCmd "github.com/ADM87/ggame/cmd/resources"

	"github.com/ADM87/ggame/resources"
	"github.com/ADM87/ggame/src/game"
	"github.com/ADM87/ggame/src/sys"
	"github.com/ADM87/ggame/src/sys/exceptions"
	"github.com/ADM87/ggame/src/sys/logger"
	"github.com/ADM87/ggame/src/sys/types"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/spf13/cobra"
)

var version = "0.0.0-unreleased"

var (
	rootDir      = types.NewCmdArg("root-dir", "", "Root directory for the game resources", ".", false)
	windowed     = types.NewCmdArg("windowed", "", "Run the game in windowed mode", false, false)
	windowWidth  = types.NewCmdArg("window-width", "", "Width of the game window", 800, false)
	windowHeight = types.NewCmdArg("window-height", "", "Height of the game window", 600, false)
)

// bootstrap
func main() {
	gameCmd := &cobra.Command{
		Use:     "ggame",
		Short:   "Starts ggame",
		Version: version,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			sys.SetVersion(cmd.Version)
			sys.Logger().MapWriters(map[logger.LogLevel]io.Writer{
				logger.LevelError: cmd.OutOrStderr(),
				logger.LevelWarn:  cmd.OutOrStderr(),
				logger.LevelInfo:  cmd.OutOrStdout(),
				logger.LevelDebug: cmd.OutOrStdout(),
			})
			sys.Logger().SetVerboseLevel(logger.LevelAll)

			if err := resources.Initialize(rootDir.GetValue()); err != nil {
				return err
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			width, height := windowWidth.GetValue(), windowHeight.GetValue()
			if width <= 0 || height <= 0 {
				return exceptions.InvalidCommandWith("window width and height must be positive integers")
			}

			fullscreen := !windowed.GetValue()

			ebiten.SetFullscreen(fullscreen)
			ebiten.SetWindowSize(width, height)
			ebiten.SetWindowTitle("ggame")

			return game.NewGame().Start()
		},
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	rootDir.RegisterWith(gameCmd.Flags().StringVarP)
	windowed.RegisterWith(gameCmd.Flags().BoolVarP)
	windowWidth.RegisterWith(gameCmd.Flags().IntVarP)
	windowHeight.RegisterWith(gameCmd.Flags().IntVarP)

	gameCmd.AddCommand(
		resCmd.GenerateResourcesManifest(),
	)

	if err := gameCmd.Execute(); err != nil {
		sys.Logger().Error(err)
		sys.ShutdownWith(1)
		return
	}
}
