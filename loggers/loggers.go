package loggers

import (
	"io/ioutil"
	"testing"

	"github.com/rs/zerolog"
	"github.com/sirupsen/logrus"
)

func benchmarkZerolog(interNo int, b *testing.B) {
	zlog := zerolog.New(ioutil.Discard)
	zerolog.SetGlobalLevel(zerolog.TraceLevel)
	// log stuff
	for i := 0; i < b.N; i++ {
		for i := 0; i < interNo; i++ {
			zlog.Trace().Msg("Something very low level.")
			zlog.Debug().Msg("Useful debugging information.")
			zlog.Info().Msg("Something noteworthy happened!")
			zlog.Warn().Msg("You should probably take a look at this.")
			zlog.Error().Msg("Something failed but I'm not quitting.")
		}
	}
}

func benchmarkLogrus(interNo int, b *testing.B) {
	var loggy = &logrus.Logger{
		Out:       ioutil.Discard,
		Formatter: new(logrus.JSONFormatter),
		Level: logrus.TraceLevel,
	}
	// log stuff
	for i := 0; i < b.N; i++ {
		for i := 0; i < interNo; i++ {
			loggy.Trace("Something very low level.")
			loggy.Debug("Useful debugging information.")
			loggy.Info("Something noteworthy happened!")
			loggy.Warn("You should probably take a look at this.")
			loggy.Error("Something failed but I'm not quitting.")
		}
	}
}
