package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
)

func main() {
	var command string
	flag.StringVar(&command, "command", "", "Command to execute.")
	var argsCommand string
	flag.StringVar(&argsCommand, "commandArgs", "", "Arguments to command.")
	flag.Parse()
	// Opening the stdout stderr files in append mode, if the files not present it will create it.
	stdoutFile, err := os.OpenFile("/tmp/background-exec-stdout.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer stdoutFile.Close()

	stderrFile, err := os.OpenFile("/tmp/background-exec-stderr.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer stderrFile.Close()
	// Creating a command to run the provided program
	cmd := exec.Command(command, argsCommand)
	// Assigning the opened stdout and stderr files to cmd's stdout and stderr respectively
	cmd.Stdout = stdoutFile
	cmd.Stderr = stderrFile
	// Starting the command in the background and storing the process
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	log.Printf("Started command with PID: %d", cmd.Process.Pid)
	// Spawning a goroutine to wait for the command to finish in background
	go func() {
		if err := cmd.Wait(); err != nil {
			log.Printf("Command finished with error: %v", err)
		} else {
			log.Printf("Command finished successfully")
		}
	}()
	log.Print("Command started in background.")
}
