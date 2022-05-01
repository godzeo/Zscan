package tracklog

import (
	"os"
	"sync"
)

type TrackLog struct {
	traceMutex *sync.Mutex
	f          *os.File
}

func New(filename string) (*TrackLog, error) {
	t := &TrackLog{
		traceMutex: &sync.Mutex{},
	}
	f, err := os.Create(filename)
	if err != nil {
		return nil, err
	}
	t.f = f
	return t, nil
}

func (t *TrackLog) Write(s string) {
	t.traceMutex.Lock()
	_, _ = t.f.WriteString(s + "\n")
	t.traceMutex.Unlock()
}
func (t *TrackLog) Close() {
	_ = t.f.Close()
}
