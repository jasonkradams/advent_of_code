package log

import (
	"log"
	"os"
)

func New(prefix string) *log.Logger {
	return log.New(os.Stdout, prefix, 80)
}
