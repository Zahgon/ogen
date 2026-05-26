package conv

import (
	"time"
)

const (
	dateLayout     = "2006-01-02"
	timeLayout     = "15:04:05"
	httpDateLayout = "Mon, 02 Jan 2006 15:04:05 GMT"
)

func Date(t time.Time) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func Time(t time.Time) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func DateTime(t time.Time) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func HTTPDate(t time.Time) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }
