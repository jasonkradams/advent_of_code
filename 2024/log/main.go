package log

import (
	"log"
	"os"
)

// New creates a new logger instance that writes to standard output (stdout)
// with the specified prefix and a predefined log flag.
//
// Parameters:
//   - prefix: A string that will be prepended to each log message.
//
// Returns:
//   - A pointer to a *log.Logger instance configured to write logs to stdout.
//
// Example:
//
//	logger := log.New("[MyApp] ")
//	logger.Println("Application started")
//
// Notes:
//   - The log flag is set to 80, which is a bitmask of log.LstdFlags (date and time)
//     and other formatting options.
func New(prefix string) *log.Logger {
	return log.New(os.Stdout, prefix, 80)
}
