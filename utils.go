package main

import (
	"fmt"
	"log"
)

type LogEvent struct {
	level   int
	err     error
	success func()
	failed  func()
}

var levels = []string{"low", "medium", "high", "critical"}

func logEvent(event LogEvent) {
	if event.err != nil {
		var template = `[%s] %s`
		var isValidLevel = event.level >= 0 && event.level < len(levels)

		if !isValidLevel {
			log.Printf(template, "medium", fmt.Errorf("cannot call logEvent with level %d", event.level))
			return
		}

		if event.failed != nil {
			event.failed()
		}

		var levelStr = levels[event.level]
		if levelStr == "critical" {
			panic(fmt.Sprintf(template, levelStr, event.err))
		}

		log.Printf(template, levelStr, event.err.Error())
	} else {
		if event.success != nil {
			event.success()
		}
	}
}
