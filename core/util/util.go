package util

import (
	"fmt"
	"time"
)

func Duration2String(t time.Duration) string {
	sceond := t.Seconds()
	if sceond >= 60 {
		return fmt.Sprintf("%.2f min", t.Minutes())
	} else {
		return fmt.Sprintf("%.2f s", sceond)
	}
}
