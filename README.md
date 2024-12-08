# Loggen
Generates fake log file data. If you want to add your own
content, put them in one of the included files. Use `-help`
to see more information about loggen.

## Installation
Clone the repository onto your local machine. If you do not
want loggen installed globally, you can use `go build .`.
If you want loggen installed globally,
use `go install loggen@latest`.

## Usage
If you installed loggen globally, you'll use `loggen`.
If you did not install it globally, you'll use `./loggen`.
Both versions of that command will execute loggen with the
default arguments. You can, however, provide arguments.

### Changing Defaults
Loggen's default settings can be changed by using
the following commands:
* `-prefix=` - Value prefixed to the generated log file's name.
* `-extension=` - Generated log file's extension.
* `-count=` - The number of lines you'd like created.

### Adding Content
More class names, method names, and log messages can be
added to loggen by updating the appropriate file.

* `class-names.csv` - Represents names of classes
* `method-names.csv` - Names of methods or functions that emitted the log line
* `messages.txt` - The logging statement. E.g., `log.info("Your message would be this.");`
