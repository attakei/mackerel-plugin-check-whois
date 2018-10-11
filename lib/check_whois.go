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
		raw, err = whois.Whois(domain+"/e", "whois.jprs.jp")
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

func run(args []string) *checkers.Checker {
	var (
		domain       = flag.String("domain", "example.com", "check target domain")
		daysWarning  = flag.Int("warning", 30, "Threshold of WARNING status")
		daysCritical = flag.Int("critical", 7, "Threshold of CRITICAL status")
	)
	flag.Parse()
	expired, err := fetchExpired(*domain)
	if err != nil {
		msg := fmt.Sprintf("%s cannot fetch expired (%s)", *domain, err)
		return checkers.Unknown(msg)
	}
	delta := int(expired.Sub(time.Now()).Hours()) / 24
	msg := fmt.Sprintf("%s is expired at %d days", *domain, delta)
	if delta < *daysCritical {
		return checkers.Critical(msg)
	}
	if delta < *daysWarning {
		return checkers.Warning(msg)
	}
	return checkers.Ok(msg)
}
