# awsdyndns

Awsdyndns is a dynamic dns tool that uses Route53. The way I use this is by setting it up on a RaspberryPi running and running this program on a 10 minute cron.

The program updates a target dns hostname `my.host.com.` with the current public IP address of the computer that is calling it.

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
$ awsdyndns -d my.host.com. -z ZONEID
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
| Variable |
|----------|
| AWS_ACCESS_KEY |
| AWS_SECRET_KEY |
```

