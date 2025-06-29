package src

import (
	"io"

	"github.com/ADM87/ggame/src/game"
	"github.com/ADM87/ggame/src/sys"
	"github.com/ADM87/ggame/src/sys/exceptions"
	"github.com/ADM87/ggame/src/sys/logger"
	"github.com/ADM87/ggame/src/sys/types"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/spf13/cobra"
)

var (
	windowed     = types.NewCmdArg("windowed", "", "Run the game in windowed mode", false, false)
	windowWidth  = types.NewCmdArg("window-width", "", "Width of the game window", 800, false)
	windowHeight = types.NewCmdArg("window-height", "", "Height of the game window", 600, false)
)

func Boot(version string) error {
	gameCmd := &cobra.Command{
		Use:     "ggame",
		Short:   "Starts ggame",
		Version: version,
		PreRun: func(cmd *cobra.Command, args []string) {
			sys.SetVersion(cmd.Version)

			sys.Logger().MapWriters(map[logger.LogLevel]io.Writer{
				logger.LevelError: cmd.OutOrStderr(),
				logger.LevelWarn:  cmd.OutOrStderr(),
				logger.LevelInfo:  cmd.OutOrStdout(),
				logger.LevelDebug: cmd.OutOrStdout(),
			})
			sys.Logger().SetVerboseLevel(logger.LevelAll)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			width, height := windowWidth.GetValue(), windowHeight.GetValue()
			if width <= 0 || height <= 0 {
				return exceptions.InvalidCommandWith("window width and height must be positive integers")
			}

			ebiten.SetFullscreen(!windowed.GetValue())
			ebiten.SetWindowSize(width, height)
			ebiten.SetWindowTitle("ggame")

			return game.NewGame(width, height).Start()
		},
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	windowed.RegisterWith(gameCmd.Flags().BoolVarP)
	windowWidth.RegisterWith(gameCmd.Flags().IntVarP)
	windowHeight.RegisterWith(gameCmd.Flags().IntVarP)

	return gameCmd.Execute()
}
