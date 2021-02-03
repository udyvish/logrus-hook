package logrus_hook

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
)

type Hook struct {
	InMemoryLogStorage map[string]string
}

func NewHook() *Hook {
	return &Hook{InMemoryLogStorage: make(map[string]string)}
}

func (h Hook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.DebugLevel,
		//logrus.ErrorLevel,
		//logrus.FatalLevel,
		//logrus.InfoLevel,
		//logrus.PanicLevel,
		//logrus.TraceLevel,
		//logrus.WarnLevel,
	}
}

func (h *Hook) Fire(entry *logrus.Entry) error {
	if len(h.InMemoryLogStorage) == 2 {
		for k, v := range h.InMemoryLogStorage {
			fmt.Printf("key:=%s value:=%s\n", k, v)
			delete(h.InMemoryLogStorage, k)
		}
	}
	h.InMemoryLogStorage[entry.Time.Format(time.RFC3339)] = entry.Message
	return nil
}
