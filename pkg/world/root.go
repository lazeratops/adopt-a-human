package world

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "run",
	Short: "Adopt-A-Human is a human adoption simulator",
	Long:  `That's pretty much it`,
	Run: func(cmd *cobra.Command, args []string) {
		world := New()
		name, err := promptString("Pick a name for your human:", nil)
		if err != nil {
			log.WithError(err).Fatalf("Failed to get human name")
		}
		world.SetHumanName(name)
		fmt.Println("Press any key to start growing your human")
		fmt.Scanln()
		world.Run()
	},
}

var nameHuman = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Hugo",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
	},
}
