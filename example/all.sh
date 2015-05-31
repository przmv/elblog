#!/bin/bash

# create fake log data
fake_log() {
	for m in {0..2880}; do
		TIMESTAMP=$(date -u -d "2 days ago + $m minutes" +"%FT%T.%NZ")
		echo "$TIMESTAMP my-loadbalancer 192.168.131.39:2817 10.0.0.1:80 0.000073 0.001048 0.000057 200 200 0 29 \"GET http://www.example.com:80/script.php?foo=bar HTTP/1.1\""
	done
}

# filter data to the previous day only
prev_day() {
	elblog interval --day
}

# fileter data to the previous hour only
prev_hour() {
	elblog interval --hour
}

# fileter data to the previous ten minutes only
prev_ten_mins() {
	elblog interval --duration 10m
}

# save status codes report as CSV
status_csv() {
	elblog --csv status > status.csv
}

# save ELB status codes report as CSV
elb_status_csv() {
	elblog --csv status --elb > status.csv
}

# output request parameter foo report
request_param_foo() {
	elblog request-param --param foo
}

# save client IP report as text
client_ip_txt() {
	elblog client-ip > client-ip.txt
}

# generate latency report from template
latency_tpl() {
	elblog --template "p50: {{.P50}}, p75: {{.P75}}, p99: {{.P99}}" latency
}

# mail the report
mail_snd() {
	mail -s "Report" john.smith@example.com
}

fake_log | prev_day | tee >(status_csv) >(elb_status_csv) >(prev_hour | request_param_foo) >(client_ip_txt) >(prev_ten_mins | latency_tpl | mail_snd) > /dev/null
