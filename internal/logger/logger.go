package logger

import (
	"fmt"

	"github.com/fatih/color"
)

var (
	infoColor    = color.New(color.FgCyan).SprintFunc()
	successColor = color.New(color.FgGreen).SprintFunc()
	warnColor    = color.New(color.FgYellow).SprintFunc()
	errorColor   = color.New(color.FgRed).SprintFunc()
)

// Info logs general messages
func Info(msg string, args ...any) {
	fmt.Println(infoColor("[INFO]"), fmt.Sprintf(msg, args...))
}

// Success logs successful operations
func Success(msg string, args ...any) {
	fmt.Println(successColor("[SUCCESS]"), fmt.Sprintf(msg, args...))
}

// Warn logs warnings
func Warn(msg string, args ...any) {
	fmt.Println(warnColor("[WARN]"), fmt.Sprintf(msg, args...))
}

// Error logs errors
func Error(msg string, args ...any) {
	fmt.Println(errorColor("[ERROR]"), fmt.Sprintf(msg, args...))
}
