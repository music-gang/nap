// Command nap is scale-to-zero for HashiCorp Nomad: see docs/README.md for
// what it is and docs/architecture.md for how the pieces fit together.
package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/music-gang/nap/internal/version"
)

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}

// run is main without the process: it returns the exit code.
func run(args []string, stdout, stderr io.Writer) int {
	fs := flag.NewFlagSet("nap", flag.ContinueOnError)
	fs.SetOutput(stderr)
	showVersion := fs.Bool("version", false, "print the version and exit")
	if err := fs.Parse(args); err != nil {
		return 2
	}
	if *showVersion {
		fmt.Fprintln(stdout, version.String())
		return 0
	}
	fmt.Fprintln(stderr, "nap: nothing to run yet, see -version")
	return 2
}
