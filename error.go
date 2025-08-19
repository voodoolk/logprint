package logprint

import (
	"fmt"
	"time"
)

func LocalError(msg error) {
	t := time.Now()
	fmt.Printf("[Error] %s:%v\n", t.Format("2006-01-02 15:04:05"), msg.Error())
}
