package kvlog

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// Log take a type and map of key/values
// and prints a timestamped log entry to stdout.
func Log(t string, m map[string]string) {
	out, _ := json.Marshal(m)
	fmt.Fprintf(os.Stdout, `ts='%s' type='%s' msg='%s'`+"\n",
		time.Now().Format("2006-01-02T15:04:05-0700"), t, out)
}

// LogInfo takes a type and message string
// and prints timestamped info log entry to stdout.
func LogInfo(t, m string) {
	fmt.Fprintf(os.Stdout, `ts='%s' type='%s' msg='{"info": "%s"}'`+"\n",
		time.Now().Format("2006-01-02T15:04:05-0700"), t, m)
}

// LogErr takes a type and message string
// and prints timestamped error log entry to stdout.
func LogErr(t, m string) {
	fmt.Fprintf(os.Stdout, `ts='%s' type='%s' msg='{"error": "%s"}'`+"\n",
		time.Now().Format("2006-01-02T15:04:05-0700"), t, m)
}