package assert

import "fmt"

// PanicIf checks a condition and panics if true.
// Ensures critical conditions are met, preventing further execution in an invalid state.
func PanicIf(condition bool, message string) {
    if condition {
        panic(message)
    }
}

// PanicIfError checks for an error and panics if present.
// Ensures the server doesn't continue running when critical errors occur, maintaining reliability.
func PanicIfError(err error, message string) {
    if err != nil {
        panic(fmt.Sprintf("%s: %v", message, err))
    }
}
