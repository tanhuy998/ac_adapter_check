package main

import (
	"fmt"
	"os"
	"os/exec"
)

/*
	program that check for status laptop ac adapter
	for linux based os
	prerequisit:
	linux distros that already install the acpi package
*/

func main() {

	for {

		run()
	}
}

func run() {

	acpi_cmd := exec.Command("acpi", "-a")
	ouput, err := acpi_cmd.Output()

	if err != nil {

		fmt.Println("error: ", err.Error())
		return
	}

	os.Stdout.Write(ouput)
}
