package logutil

import (
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"os"
	"sync"
	"time"
)

// Global variables for debug mode and logging cache
var (
	debugMode    bool                    // Determines if debug logs should be displayed
	loggedEvents = make(map[string]bool) // Cache to store logged events to prevent duplicates
	mutex        sync.Mutex
	// Mutex to synchronize access to loggedEvents
)

// Struct to hold additional fields

type LogFields struct {
	Event         string                 `json:"event,omitempty"`
	Error         error                  `json:"error,omitempty"`
	Additional    map[string]interface{} `json:"additional,omitempty"`
	StatusCode    int                    `json:"status_code,omitempty"`
	CorrelationID string                 `json:"correlationID,omitempty"`
}

// Init initializes the logrus logger with JSON formatting and INFO level
func Init() {
	env := os.Getenv("BART_ENV")

	if env == "local" {
		logrus.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:   true,
			ForceColors:     true,
			TimestampFormat: time.RFC3339,
			DisableQuote:    true,
		})
	}
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
	})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)
}

// SetDebugMode enables or disables debug logging
func SetDebugMode(debug bool) {
	debugMode = debug
}

// GenerateCorrelationID generates a unique ID for tracking logs and events
func GenerateCorrelationID() string {
	return uuid.New().String()
}

// LogRelationalStart logs the start of an event if debug mode is enabled using struct
func LogRelationalStart(fields *LogFields) *logrus.Entry {
	if !debugMode {
		return nil
	}

	entry := logrus.WithFields(logrus.Fields{
		"event":         fields.Event,
		"correlationID": fields.CorrelationID,
		"timestamp":     getTimeStamp(),
		"status_code":   fields.StatusCode,
	})
	mergeFields(entry, fields.Additional)

	entry.Info("Event started")
	return entry
}

// LogRelationalEnd logs the end of an event if debug mode is enabled using struct
func LogRelationalEnd(fields *LogFields) *logrus.Entry {
	if !debugMode {
		return nil
	}

	entry := logrus.WithFields(logrus.Fields{
		"event":         fields.Event,
		"correlationID": fields.CorrelationID,
		"timestamp":     getTimeStamp(),
		"status_code":   fields.StatusCode,
	})
	mergeFields(entry, fields.Additional)

	entry.Info("Event completed")
	return entry
}

// LogError logs an error event regardless of debug mode using struct
func LogError(fields *LogFields) {

	entry := logrus.WithFields(logrus.Fields{
		"event":         fields.Event,
		"correlationID": fields.CorrelationID,
		"timestamp":     getTimeStamp(),
		"error":         fields.Error.Error(),
		"status_code":   fields.StatusCode,
	})
	mergeFields(entry, fields.Additional)

	entry.Error("Error occurred")
}

// LogOnce logs an event only once to prevent duplicate logs using struct
func LogOnce(fields *LogFields) {
	mutex.Lock()
	defer mutex.Unlock()

	logKey := fields.Event
	if loggedEvents[logKey] {
		return
	}

	entry := logrus.WithFields(logrus.Fields{
		"event":     fields.Event,
		"timestamp": getTimeStamp(),
	})
	mergeFields(entry, fields.Additional)

	entry.Info("Event logged once")
	loggedEvents[logKey] = true
}

// LogSuccess logs a successful event only once to prevent duplicate logs using struct
func LogSuccess(fields *LogFields) {
	mutex.Lock()
	defer mutex.Unlock()

	logKey := fields.Event
	if loggedEvents[logKey] {
		return
	}

	entry := logrus.WithFields(logrus.Fields{
		"event":     fields.Event,
		"timestamp": getTimeStamp(),
	})
	mergeFields(entry, fields.Additional)

	entry.Info("Event logged successfully")
	loggedEvents[logKey] = true
}

// mergeFields merges additional fields into the base log fields (for LogFields struct)
func mergeFields(entry *logrus.Entry, additionalFields map[string]interface{}) {
	if additionalFields != nil {
		for k, v := range additionalFields {
			entry = entry.WithField(k, v)
		}
	}
}

func getTimeStamp() string {
	return time.Now().UTC().Format(time.RFC3339)
}
