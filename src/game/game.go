package game

import (
	"io"

	"github.com/ADM87/ggame/src/sys"
	"github.com/ADM87/ggame/src/sys/logger"
	"github.com/ADM87/ggame/src/sys/types"
	"github.com/hajimehoshi/ebiten"
	"github.com/spf13/cobra"
)

type ggame struct {
	Loop
	Renderer
	Window
}

var (
	windowWidth  = types.NewCmdArg("window-width", "", "Width of the game window", 640, false)
	windowHeight = types.NewCmdArg("window-height", "", "Height of the game window", 480, false)
)

func Start(version string) error {
	gameCmd := &cobra.Command{
		Use:   "ggame",
		Short: "Starts ggame",
		PreRun: func(cmd *cobra.Command, args []string) {
			sys.SetVersion(version)
			sys.Logger().MapWriters(map[logger.LogLevel]io.Writer{
				logger.LevelError: cmd.OutOrStderr(),
				logger.LevelWarn:  cmd.OutOrStderr(),
				logger.LevelInfo:  cmd.OutOrStdout(),
				logger.LevelDebug: cmd.OutOrStdout(),
			})
			sys.Logger().SetVerboseLevel(logger.LevelAll)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ebiten.SetWindowSize(windowWidth.GetValue(), windowHeight.GetValue())
			ebiten.SetWindowTitle("ggame")
			return ebiten.RunGame(&ggame{
				Loop:     NewGameLoop(),
				Renderer: NewGameRenderer(),
				Window:   NewGameWindow(windowWidth.GetValue(), windowHeight.GetValue()),
			})
		},
	}

	windowWidth.RegisterWith(gameCmd.Flags().IntVarP)
	windowHeight.RegisterWith(gameCmd.Flags().IntVarP)

	return gameCmd.Execute()
}
