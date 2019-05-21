# awsdyndns

Awsdyndns is a [dynamic dns](https://en.wikipedia.org/wiki/Dynamic_DNS) tool that updates a Route53 hostname with the current public IP of the caller. The way I use awsdyndns is by installing it on a RaspberryPi and running the program on a 10 minute cron.

Public IP is attained via [ipify](https://www.ipify.org/).

### Building
```
git clone git@github.com:jsfour/awsdyndns.git
cd awsdyndns
go install
```

### Usage
Once you have configured the `AWS_ACCESS_KEY`, and `AWS_SECRET_KEY` environment variables just run `awsdyndns`.

```
$ awsdyndns -d my.host.com. -z ZONEID // don't forget the trailing `.` on the hostname.
Updating dns
```

### Help
```
$ awsdyndns -h
Usage of awsdyndns:
  -d, --dnshostname string   Hostname on the zone id to update
  -z, --zoneid string        Route 53 zone id
```


### Required environment variables:
```
| Variable |
|----------|
| AWS_ACCESS_KEY |
| AWS_SECRET_KEY |
```

