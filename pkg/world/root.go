package world

import (
	"github.com/gookit/color"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// This flag should be used if the player just wants to watch their human, not make decisions for them.
var watchOnly bool
var logLevel string

func Init() {
	cobra.MousetrapHelpText = ""
	RootCmd.Flags().BoolVarP(&watchOnly, "watch", "w", false, "Just watch your human; do not make any decisions")
	RootCmd.Flags().StringVarP(&logLevel, "log-level", "l", "error", "Log level")

}

var RootCmd = &cobra.Command{
	Use:   "run",
	Short: "Adopt-A-Human is a human adoption simulator",
	Long:  `Adopt-A-Human is a human adoption simulator. You find a random human and make decisions for them (or just watch them grow and see how long it takes them to die)`,
	Run: func(cmd *cobra.Command, args []string) {

		switch logLevel {
		case "trace":
			log.SetLevel(log.TraceLevel)
		case "debug":
			log.SetLevel(log.DebugLevel)
		case "info":
			log.SetLevel(log.InfoLevel)
		case "warn":
			log.SetLevel(log.WarnLevel)
		default:
			log.SetLevel(log.ErrorLevel)
		}
		start()
	},
}

func start() {
	world := New(watchOnly)

	color.BgGreen.Println("Your human has arrived!")
	state := world.human.Mind().StateReport()
	color.Yellow.Println(state)

	if !watchOnly {
		options := []string{"start", "find another human"}
		selection, err := promptSelection("Would you like to start, or find another human?", options)
		if err != nil {
			log.WithError(err).Fatal("Failed to make game start selection")
		}

		if selection == "start" {
			pickName(world)
			world.Run()
			return
		}
		start()
		return
	}
	pickName(world)
	world.Run()
}

func pickName(world *World) {
	name := "Bob"
	if !watchOnly {
		var err error
		name, err = promptString("Pick a name for your human:", nil)
		if err != nil {
			log.WithError(err).Fatalf("Failed to get human name")
		}
	}
	world.SetHumanName(name)
}
