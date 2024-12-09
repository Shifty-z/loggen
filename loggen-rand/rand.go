package loggen_rand

import (
	"fmt"
	"loggen/fileIO"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// GenerateRandomLogLine - Constructs a mocked log line that includes a timestamp, log level, a "class" name
// a "method" name, and a fake message.
func GenerateRandomLogLine(sources *fileIO.DataSources) string {
	sb := strings.Builder{}

	sb.WriteString("[")
	sb.WriteString(createTimestamp())
	sb.WriteString("] ")
	sb.WriteString(getRandomLogLevel())
	sb.WriteString("- ")
	sb.WriteString(getRandomValueInSource(&sources.Classes))
	sb.WriteString(".")
	sb.WriteString(getRandomValueInSource(&sources.Methods))
	sb.WriteString(": ") // This isn't always included in log lines
	sb.WriteString(getRandomValueInSource(&sources.Messages))

	return sb.String()
}

// randInRange - Returns a random value within the provided range. This is min and max inclusive: [min, max]
func randInRange(min, max int) int {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	return random.Intn(max+1-min) + min
}

// createTimestamp - Creates a string of "[YYYY-MM-ddThh:mm:ss.zzzZ]"
func createTimestamp() string {
	sb := strings.Builder{}

	sb.WriteString(strconv.Itoa(randInRange(1950, 2024)))
	sb.WriteString("-")

	month := strconv.Itoa(randInRange(1, 12))
	sb.WriteString(fmt.Sprintf("%02s", month))
	sb.WriteString("-")

	day := strconv.Itoa(randInRange(1, 30))
	sb.WriteString(fmt.Sprintf("%02s", day))

	sb.WriteString("T")

	hours := strconv.Itoa(randInRange(0, 23))
	sb.WriteString(fmt.Sprintf("%02s", hours))
	sb.WriteString(":")

	minutes := strconv.Itoa(randInRange(0, 59))
	sb.WriteString(fmt.Sprintf("%02s", minutes))
	sb.WriteString(":")

	seconds := strconv.Itoa(randInRange(0, 59))
	sb.WriteString(fmt.Sprintf("%02s", seconds))
	sb.WriteString(".")

	milliseconds := strconv.Itoa(randInRange(0, 999))
	sb.WriteString(fmt.Sprintf("%03s", milliseconds))
	sb.WriteString("Z")

	return sb.String()
}

// getRandomValueInSource - Generates a random index and returns the value from `source` at the generated index.
func getRandomValueInSource(source *[]string) string {
	idx := randInRange(0, len(*source)-1)

	return (*source)[idx]
}

// getRandomLogLevel - Generates a random log level (e.g., ERROR).
func getRandomLogLevel() string {
	switch level := randInRange(0, 3); level {
	case 0:
		return "DEBUG "
	case 1:
		return "INFO "
	case 2:
		return "WARN "
	case 3:
		return "ERROR "
	default:
		fmt.Println("")
		panic(fmt.Sprintf("getRandomLogLevel - The default case was hit! Level value generated: %d\n", level))
	}
}
