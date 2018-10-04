package check_whois

import (
    "flag"
    "fmt"
    "os"

    "github.com/domainr/whois"
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
    request, err := whois.NewRequest(*d)
    if err != nil {
        // TODO: Echo reason
        os.Exit(1)
    }
    response, err := whois.DefaultClient.Fetch(request)
    if err != nil {
        // TODO: Echo reason
        os.Exit(1)
    }
    record, err := response.Parse() // not implemented yet
    if err != nil {
        // TODO: Echo reason
        os.Exit(1)
    }
    fmt.Println(record)
	return checkers.NewChecker(checkers.OK, *d)
}
