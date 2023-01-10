# Background Process Logger

This application is used to run a program in the background and write all logs to files.

## Getting Started

These instructions will help you to run this application on your local machine.

### Prerequisites

- Go version 1.14 or above
- A program that you want to run in the background

### Installing

- Clone the repository
- Run `go build` command to build the application

### Running the application

- Run the application with the name of the command that you want to run in the background as the argument

`./background -appName=commandName`

The application will redirect the standard output and error of the command to files "stdout.log" and "stderr.log", respectively. 

While running the command will also run in background and you can track the process using the process id by running `ps -ef | grep commandName` or `pgrep -f commandName <arg1> <arg2>`.