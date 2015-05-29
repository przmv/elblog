# elblog

## Description

Tool for analyzing access logs for Amazon Elastic Load Balancing.

## Usage

### Filter log entries by the specified interval

`elblog interval interval [command options] [arguments...]`

The command `interval` outputs the log entries for the specified duration before now.

If `--hour` flag was specified, the output will contain log entries for the previous hour.

Specifying the `--day` flag will display the log entries since the day before now.

To output the log entries for the custom duration, use `--duration` flag with the custom duration value.

The value for the `--duration` flug must conform the input for the [`time.ParseDuration()`](http://golang.org/pkg/time/#ParseDuration) Go function.

A duration string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as "300ms", "-1.5h" or "2h45m". Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".

**NOTE:** Global flags '--csv' and '--template' are ignored by the 'interval' command.

### Count request parameters

`elblog request-param [command options] [arguments...]`

The command `request-param` outputs the list of different
request parameter's values along with the total amount of these values in the
request path field across the log entries provided.

The required `--param` flag is used to specify the HTTP request parameter.

If global flag `--csv` was set, the first column of the output denotes the request
parameter value and the second one is the count.

#### Template data

If using global flag '--template', the following data type is sent to the template
to execute:

`[]map[string]string`

The possible map keys are `RequestParameterValue` and `Count`.

The example template is the following:

```
{{range $i, $r := .}}
  {{$i}}. {{$r.RequestParameterValue}} ({{$r.Count}})
{{end}}
```

See https://golang.org/pkg/text/template/ for the reference.

### Count different client IP addresses

`elblog client-ip [arguments...]`

The command `client-ip` outputs the list of different client
IP addresses along with the total amount of requests.

If global flag `--csv` was set, the first column of the output denotes the IP address
and the second one is the count.

#### Template data

If using global flag `--template`, the following data type is sent to the template
to execute:

`[]map[string]string`

The possible map keys are `ClientIp` and `Count`.

The example template is the following:

```
{{range $i, $r := .}}
  {{$i}}. {{$r.ClientIp}} ({{$r.Count}})
{{end}}
```

See https://golang.org/pkg/text/template/ for the reference.

### Count status codes

`elblog status [command options] [arguments...]`

The command `status` outputs the list of different status codes
along with the total amount of requests.

By default it displays the backend status codes.

If the `--elb` flag was added, status codes for Elastic Load Balancing
will be outputted.

If global flag `--csv` was set, the first column of the output denotes status code
and the second one is the count.

#### Template data

If using global flag `--template`, the following data type is sent to the template
to execute:

`[]map[string]string`

The possible map keys are `Status` and `Count`.

The example template is the following:

```
{{range $i, $r := .}}
  {{$i}}. {{$r.Status}} ({{$r.Count}})
{{end}}
```

See https://golang.org/pkg/text/template/ for the reference.

## Install

To install, use `go get`:

```bash
$ go get github.com/pshevtsov/elblog/cmd/elblog
```

## Contribution

1. Fork ([https://github.com/pshevtsov/elblog/fork](https://github.com/pshevtsov/elblog/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run `gofmt -s`
1. Create a new Pull Request

## Author

[Petr Shevtsov](https://github.com/pshevtsov)

## Licence

MIT
