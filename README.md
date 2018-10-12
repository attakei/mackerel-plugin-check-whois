# Mackerel check whois plugin

Mackerel agent plugin to check expired date and alert


## Installation 

This plugin can install by `mkr plugin install` command.

```
$ mkr plugin install attakei/mackerel-plugin-check-whois
```


## Usage

### Setting into config

```
[plugin.checks."whois_example.com"]
command = "/path/to/mackerel-plugin-check-whois --warning=30 --critical=7 --domain=example.com"
```

### Options

* `--domain`
    * Monitoring target domain
* `--warning`
    * Threshold of WARNING status
* `--critical`
    * Threshold of CRITICAL status

