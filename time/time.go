package time

import (
	"time"
	"fmt"
)

func FormatTimeNow() string{
	return time.Now().Format("2006-01-02 15:04:05")
}

func TimeStampNow() int64{
	return time.Now().Unix()
}

