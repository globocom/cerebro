package modules

import (
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
)

//LogLevelDecoder represents the current log level for this application
type LogLevelDecoder log.Level

var (
	mapLogLevel = map[string]log.Level{
		"DEBUG": log.DebugLevel,
		"INFO":  log.InfoLevel,
		"WARN":  log.WarnLevel,
		"ERROR": log.ErrorLevel,
	}
)

//Decode decodes strings to the current log level
func (lld *LogLevelDecoder) Decode(value string) error {
	var upper = strings.ToUpper(value)
	if val, ok := mapLogLevel[upper]; ok {
		*lld = LogLevelDecoder(val)
	} else {
		return fmt.Errorf("log Level %s is not valid", value)
	}

	return nil
}
