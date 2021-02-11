package main

import (
	"fmt"
	"time"

	"github.com/claudiocandio/gemini-api/logger"
	"github.com/urfave/cli/v2"
)

func parseConvertTimestamp(timestamp_str string) (*time.Time, error) {

	// checking timestamp input
	_, err := time.Parse("2006-01-02T15:04:05", timestamp_str)
	if err != nil {
		return nil, fmt.Errorf("timestamp not valid: %s", timestamp_str)
	}

	// adding timezone offset to timestamp_str
	// convert offset secs in +01:00
	t := time.Now()
	_, offset := t.Zone()
	timestamp_str = fmt.Sprintf("%s+%02d:%02d", timestamp_str, offset/3600, offset/60-(offset/3600*60))
	timestamp, err := time.Parse("2006-01-02T15:04:05Z07:00", timestamp_str)
	if err != nil {
		return nil, fmt.Errorf("timestamp not valid: %s", timestamp_str)
	}
	return &timestamp, nil

}

func parse_params(c *cli.Context) string {

	if c.Bool("trace") {
		logger.SetLevel(logger.Level(logger.TraceLevel))
		logger.Debug("Trace enabled")
	} else if c.Bool("debug") {
		logger.SetLevel(logger.Level(logger.DebugLevel))
		logger.Debug("Debug enabled")
	}

	if c.IsSet("config") {
		return c.String("config")
	}
	return ""
}

func get_timestampms(ts time.Time) int64 {
	timestampms := ts.UnixNano() / 1e6
	return timestampms
}
