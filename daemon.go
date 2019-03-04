package main

import (
	//	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	args := os.Args
	input_time := args[1]
	pre_time := ""
	for {
		now_time := time.Now().Format("200601021504")
		if strings.HasSuffix(now_time, input_time) && now_time != pre_time {
			cmd := exec.Command("/bin/bash", "-c", "sh "+args[2]+" > log."+args[2]+" 2>&1 &")
			_ = cmd.Run()
		}
		pre_time = now_time
		time.Sleep(time.Second * 10)
	}
}
