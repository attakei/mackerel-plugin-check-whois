package check_whois

import (
    "flag"
    "fmt"
    "os"
    "strings"
    "time"

    "github.com/joyt/godate"
    "github.com/likexian/whois-go"
    "github.com/likexian/whois-parser-go"
    "github.com/mackerelio/checkers"
)

// Do the plugin
func Do() {
    ckr := run(os.Args[1:])
    ckr.Name = "Whois"
    ckr.Exit()
}


/**
 * Fetch WHOIS string from server
 */
func fetchWhois(domain string) (raw string, err error) {
    // If query to JPRS, set option to return English forcely
    if strings.HasSuffix(domain, ".jp") {
        raw, err = whois.Whois(domain + "/e", "whois.jprs.jp")
    } else {
        raw, err = whois.Whois(domain)
    }
    return raw, err
}


func fetchExpired(domain string) (expired time.Time, err error) {
    raw, err := fetchWhois(domain)
    if err != nil {
        return time.Now(), err
    }
    record, err := whois_parser.Parse(raw)
    if err != nil {
        return time.Now(), err
    }
    return date.Parse(record.Registrar.ExpirationDate)
}


func run(args []string) * checkers.Checker {
    var (
         d = flag.String("domain", "example.com", "check target domain")
    )
    flag.Parse()
    expired, err := fetchExpired(*d)
    if err != nil {
        // TODO: handle more formal for golang
        fmt.Println(err)
        os.Exit(1)
    }
	return checkers.NewChecker(checkers.OK, *d + " is expired at " + expired.String())
}
