package check_whois

import (
    "os"

    "github.com/mackerelio/checkers"
)

// Do the plugin
func Do() {
    ckr := run(os.Args[1:])
    ckr.Name = "Whois"
    ckr.Exit()
}


func run(args []string) * checkers.Checker {
	return checkers.NewChecker(checkers.OK, "OK")
}
