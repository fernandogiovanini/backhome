package printer

import (
	"fmt"
)

func Error(message string, args ...any) {
	fmt.Printf("\nERROR! %s\n", fmt.Sprintf(message, args...))
}
