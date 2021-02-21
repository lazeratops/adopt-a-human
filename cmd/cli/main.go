package main

import (
	"aah/pkg/world"
	"fmt"
	"github.com/sirupsen/logrus"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	world.Init()
	if err := world.RootCmd.Execute(); err != nil {
		logrus.WithError(err).Error("failed to run simulation")
	}
	fmt.Println("Press Enter to exit")
	fmt.Scanln()
}
