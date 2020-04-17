package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command(
		"cmd.exe",
		"/C",
		"start",
		"/B",
		"hidebat.vbs",
		"idea.bat")

	err := cmd.Start()
	if err != nil {
		log.Fatalln("Error running the program", err)
	}
	fmt.Println("SUCCESSFUL!!")
}
