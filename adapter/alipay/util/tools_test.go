package util

import (
	"testing"
	"time"
)

func TestParseBeijingDateTime(t *testing.T) {
	t1 := "2025-10-17 10:00:00"
	t2 := ParseBeijingDateTime(t1)
	t.Log(t2.Format(time.DateTime))
}
