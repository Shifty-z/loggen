package main

import (
	"flag"
	"fmt"
	"loggen/fileIO"
	rand "loggen/loggen-rand"
	"strings"
)

type flagDescription struct {
	Extension string
	Prefix    string
	Help      string
	Count     string
}

type commandFlag struct {
	extension string
	prefix    string
	count     int
}

func main() {
	// TODO: Check comments in fileIO
	descriptions := createFlagDescription()

	extension := flag.String("ext", ".log", descriptions.Extension)
	prefix := flag.String("prefix", "loggen-", descriptions.Prefix)
	help := flag.Bool("help", false, descriptions.Help)
	count := flag.Int("count", 100, descriptions.Count)
	flag.Parse()

	commandFlags := createFlagWith(extension, prefix, count)

	if *help {
		fmt.Println(descriptions.Help)
	}

	if !*help {
		flags := fileIO.CommandFlags{
			Extension: commandFlags.extension,
			Prefix:    commandFlags.prefix,
		}

		dataSources := fileIO.LoadSources()

		logs := make([]string, 0)
		for i := 0; i < commandFlags.count; i++ {
			logs = append(logs, rand.GenerateRandomLogLine(&dataSources))
		}

		fmt.Printf("Writing to the log file\n")
		file := fileIO.CreateLogFile(&flags)
		fileIO.WriteLog(&logs, file)

		file = fileIO.OpenFile(file.Name())
		lines := fileIO.ReadTextFileLines(file)
		for _, line := range lines {
			fmt.Println(line)
		}
	}
}

func createFlagDescription() *flagDescription {
	extension := "Generated log file's extension. Defaults to .log"
	prefix := "Generated log file's prefixed value, what comes before the timestamp. Defaults to loggen-"
	count := "The number of log lines to generate."

	sb := strings.Builder{}
	sb.WriteString("Loggen creates mock log files.\n\n")
	sb.WriteString("Optional Flags: -ext, -prefix, -count\n\n")
	sb.WriteString("-ext: ")
	sb.WriteString(extension)
	sb.WriteString("\n\n")
	sb.WriteString("-prefix: ")
	sb.WriteString(prefix)
	sb.WriteString("\n\n")
	sb.WriteString("-count: ")
	sb.WriteString(count)
	sb.WriteString("\n\n")
	sb.WriteString("Example extension usage: -prefix=FAKE_LOG\n\n")
	sb.WriteString("Example command: ./loggen -ext=txt -prefix=data-log -count=250\n")

	fd := flagDescription{
		Extension: extension,
		Prefix:    prefix,
		Help:      sb.String(),
		Count:     count,
	}

	return &fd
}

// createFlagWith - Creates commandFlag struct with defaults. If properties have certain rules, they are applied here.
// For example, -ext is checked for a leading (dot). If it doesn't have one, it is added.
func createFlagWith(extension, prefix *string, count *int) *commandFlag {
	if (*extension)[0] != '.' {
		*extension = "." + *extension
	}

	// The trailing delimiter doesn't matter, include one for the sake of making filenames easier to read
	if !strings.HasSuffix(*prefix, "-") || !strings.HasSuffix(*prefix, "_") {
		*prefix = *prefix + "-"
	}

	return &commandFlag{
		extension: *extension,
		prefix:    *prefix,
		count:     *count,
	}
}
