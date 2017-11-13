# shellexec [![GoDoc](https://godoc.org/github.com/mibk/shellexec?status.png)](https://godoc.org/github.com/mibk/shellexec)

Package shellexec provides cross-platform shell-like command execution. It
supports a small subset of the Shell Command Language related to executing
of commands. Package shellexec isn't a replacement for `os/exec`. It's rather
just a convenience package for use-cases that require/prefer having just a
single string to define the command to execute.

## Instalation

```bash
$ go get github.com/mibk/shellexec
```

## Example

```go
package main

import (
	"log"
	"os"

	"github.com/mibk/shellexec"
)

func main() {
	run(`echo  'Sup'ports  "\"quotes\","`)
	run(`VARIABLE=assignment, X=3 env`)
	run(`echo "and variable expansion, $USER."`)
	run(`echo Returns just '"os/exec".Cmd,' that can be`)
	run(`echo further adjusted before running.`)

	run(`echo The goal is to support a really \*small\* subset`)
	run(`echo of the Shell language while being 100\% compatible.`)
}

func run(command string) {
	cmd, err := shellexec.Command(command)
	if err != nil {
		log.Fatal(err)
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
```
