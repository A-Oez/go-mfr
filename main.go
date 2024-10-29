package main

import (
	"github.com/A-Oez/go-mfr/cmd"
	_ "github.com/A-Oez/go-mfr/cmd/sreq"
	_ "github.com/A-Oez/go-mfr/cmd/tui"
)

func main() {
	cmd.Execute()
}
