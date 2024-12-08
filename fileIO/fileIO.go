package fileIO

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path"
	"strings"
	"time"
)

const (
	resourcesFolder = "resources"
	classNamesFile  = "class-names.csv"
	methodNamesFile = "method-names.csv"
	messagesFile    = "messages.txt"
)

type DataSources struct {
	Classes  []string
	Methods  []string
	Messages []string
}

type CommandFlags struct {
	Extension string
	Prefix    string
}

// LoadSources - Loads content from each of the data files.
func LoadSources() DataSources {
	classNamesFile := OpenFile(path.Join(resourcesFolder, classNamesFile))
	defer classNamesFile.Close()
	classes := ReadCsvAndGetLines(classNamesFile)

	methodNamesFile := OpenFile(path.Join(resourcesFolder, methodNamesFile))
	defer methodNamesFile.Close()
	methods := ReadCsvAndGetLines(methodNamesFile)

	messagesFile := OpenFile(path.Join(resourcesFolder, messagesFile))
	defer messagesFile.Close()
	messages := ReadTextFileLines(messagesFile)

	return DataSources{Classes: classes, Methods: methods, Messages: messages}
}

// OpenFile - Opens a file with the given `filename` and returns a handle to it. Does not close the file itself!
// PANIC NOTE: If no file with the provided name exists, or it cannot be interacted with for some reason, this
// panics.
func OpenFile(filename string) *os.File {
	file, err := os.Open(filename)

	if err != nil {
		e := fmt.Sprintf("OpenFile: Unable to interact with the file: %s\n Error: %s\n", filename, err)
		panic(e)
	}

	return file
}

// ReadCsvAndGetLines - Returns all entries in the CSV file.
func ReadCsvAndGetLines(file *os.File) []string {
	scanner := bufio.NewScanner(file)

	var content string
	for scanner.Scan() {
		content += scanner.Text()
	}
	if scanner.Err() != nil {
		fmt.Printf("Scanner encountered a problem while reading the CSV file: %s\n", file.Name())
	}

	return strings.Split(content, ",")
}

// ReadTextFileLines - Reads each row in the provided file. Rows are separated by a trailing newline
func ReadTextFileLines(file *os.File) []string {
	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func timestamp() string {
	now := time.Now()

	return now.Format(time.RFC3339)
}

func createFileName(args *[]string) string {
	var sb strings.Builder
	sb.Grow(len(*args))

	for _, arg := range *args {
		sb.WriteString(arg)
	}

	return sb.String()
}

func CreateLogFile(flags *CommandFlags) *os.File {
	fileName := createFileName(&[]string{flags.Prefix, timestamp(), flags.Extension})

	if _, err := os.Stat(fileName); err != nil && errors.Is(err, os.ErrExist) {
		fileName = flags.Prefix + timestamp() + flags.Extension
	}

	file, err := os.Create(fileName)

	if err != nil {
		fmt.Println("Unable to create the log file: ", fileName)
		panic("Log file creation failure")
	}

	return file
}

// WriteLog - Writes provided data to the file and returns the destination file's name.
func WriteLog(logs *[]string, file *os.File) {
	defer file.Close()

	fmt.Println("Writing lines")
	for _, line := range *logs {
		file.WriteString(line + "\n")
	}
}
