package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func main() {

	msg := `System Shutdown in 10 minutes
	\rSave your work`
	cmd := exec.Command("notify-send", "-u", "critical", "SYSTEM SHUTDOWN NOTICE", msg)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v", out.String())
}
