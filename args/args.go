package args

import (
	"strings"
)

// ErrEndOfArgs is a error type for EOA (End of arguments) error
type errEndOfArgs struct{}

func (e *errEndOfArgs) Error() string {
	return "end of arguments"
}

// Arguments is a struct for wrapper arguments from input
type Arguments struct {
	bin  string
	args []string
}

// Binary is a func to get the binary name
func (a Arguments) Binary() string {
	return a.bin
}

// First is a func to get the first argument
func (a Arguments) First() (string, error) {
	if len(a.args) == 0 {
		return "", &errEndOfArgs{}
	}
	return a.args[0], nil
}

// Tail is a func that leaving the first and get the rest
func (a Arguments) Tail() Arguments {
	if len(a.args) <= 1 {
		return Arguments{bin: a.bin, args: []string{}}
	}
	return Arguments{bin: a.bin, args: a.args[1:]}
}

// Parse is a func to transform the arguments in a struct for easier manipulate the arguments
func Parse(args []string) Arguments {
	nargs := make([]string, 0)
	for _, arg := range args {
		if !strings.HasPrefix(arg, "-") {
			nargs = append(nargs, arg)
		}
	}
	return Arguments{bin: nargs[0], args: nargs[1:]}
}
