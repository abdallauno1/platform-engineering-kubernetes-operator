package events

import "log"

type Recorder interface{ Event(reason, message string) }
type LogRecorder struct{}

func (LogRecorder) Event(reason, message string) {
	log.Printf("event reason=%s message=%q", reason, message)
}
