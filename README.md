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

### Running the command or application

- Run the command or application with the name of the command that you want to run in the background as the argument

`./background -command=commandName`

- Optionaly you can add multiple arguments to command:

`./background -command=commandName -commandArgs=exampleArguments`

The application will redirect the standard output and error of the command to files "/tmp/background-exec-stdout.log" and "/tmp/background-exec-stderr.log", respectively. 

While running the command will also run in background and you can track the process using the process id by running `ps -ef | grep commandName` or `pgrep -f commandName <arg1> <arg2>`.

## Testing

Run app with `test.sh` script:

`go run background.go -command="./test.sh" -commandArgs="--=this 'is' example usage"`

Check result:

`tail -f /tmp/background-exec-stdout.log`