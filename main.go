package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func logSrv() {
	app := fiber.New(fiber.Config{DisableStartupMessage: true})
	app.Get("/loglevel/:level?", func(c *fiber.Ctx) error {
		lev := c.Params("level")
		levelNo, err := zerolog.ParseLevel(lev)
		if err != nil {
			return c.SendString("unknown loglevel")
		}
		zerolog.SetGlobalLevel(levelNo)
		return c.SendString(fmt.Sprintf("%s level set.", c.Params("level")))
	})
	app.Listen(":3000")
}

func main() {
	// UNIX Time is faster and smaller than most timestamps
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	// Set global log level to DEBUG
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	// Beautify output
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	go logSrv()

	// log stuff
	for {
		fmt.Println()
		log.Trace().Msg("Something very low level.")
		log.Debug().Msg("Useful debugging information.")
		log.Info().Msg("Something noteworthy happened!")
		log.Warn().Msg("You should probably take a look at this.")
		log.Error().Msg("Something failed but I'm not quitting.")
		fmt.Println()
		time.Sleep(time.Second * 1)
	}
}