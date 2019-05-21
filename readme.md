# awsdyndns

This is a dynamic dns tool for AWS.

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

The way I use this is by setting it up on a RaspberryPi running a cron.

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

