# Mackerel check whois plugin

Mackerel agent plugin to check expired date and alert

## Usage

### Setting into config

```
[plugin.checks.whois_EXAMPLE-COM]
comand = "/path/to/check-whois --warning=30 --critical=7 --domain=example.com"
```

### Options

* `--domain`
    * Monitoring target domain
* `--warning`
    * Threshold of WARNING status
* `--critical`
    * Threshold of CRITICAL status

