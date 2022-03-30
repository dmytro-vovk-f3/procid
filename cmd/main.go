package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/dmytro-vovk-f3/procid/internal/detector"
	"github.com/dmytro-vovk-f3/procid/internal/finder"
)

func main() {
	if len(os.Args) != 2 {
		_, _ = fmt.Fprintf(os.Stderr, "Usage: %s file_name", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	arg := os.Args[1]

	path, err := finder.Resolve(arg)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error finding the file: %s", err)
		os.Exit(1)
	}

	pid, err := detector.FindPID(path)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error finding pid for file: %s", err)
		os.Exit(1)
	}

	fmt.Print(pid + "\n")
}
