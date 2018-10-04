package check_whois

import (
    "flag"
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
    var (
         d = flag.String("domain", "example.com", "check target domain")
    )
    flag.Parse()
	return checkers.NewChecker(checkers.OK, *d)
}
