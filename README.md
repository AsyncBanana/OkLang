# OkLang
A simple declarative, way to check if someone is ok and help them if they aren't.
## Installation
Clone the repository and run the Go compiler on `main.go` (if you don't already have it installed, get it here at https://golang.org/dl/). Sometime hopefully, program files will be provided in the releases to allow you to not need to download the Go compiler.
## Running
Put your file `main.okng` in the same directory as the program file from the Go compiler, and run the program file.
## Syntax
The basic syntax is this: `command-{command-argument}-argument`.
On each line, you can write the command (or function) that you are executing, then all of the arguments separated by dashes. If you want to nest a command as an argument, you put the command in curly braces.
**Note:** Currently, you can not nest commands multiple times
## Commands
**Note:** All commands are written in camelCase
* **ok:** This command will simply print "Ok!"
* **notOk:** This command will open a page in the default
browser on a website with a bunch of pictures
* **promptOk:** This takes three arguments. The first one is a string of what to say before prompting "Are you ok?". The second one is the command to execute on the event that the user says yes. The third one is the command for answering no.
* **printOk:** This takes one argument, the string to print, and is identical to `Println` in Go.
* **print:** This takes one argument, the string to print, and is identical to `Print` in Go.
## Types
There are currently only two types, string and command. You don't need to do anything to make something a string, just put it in the argument. To use the command type, you have to put the command in curly braces, as stated above.
## Coming soon!
This project is still in an early stage, so make sure to check for updates or put this repository in your notifications stream to know when a new version comes out!
