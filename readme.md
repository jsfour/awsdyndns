# awsdyndns

Awsdyndns is a [dynamic dns](https://en.wikipedia.org/wiki/Dynamic_DNS) tool that updates a Route53 hostname with the current public IP of the caller. The way I use awsdyndns is by installing it on a RaspberryPi and running the program on a 10 minute cron.

Public IP is obtained via [ipify](https://www.ipify.org/).

## Installing

Download the [latest release for your system](https://github.com/jsfour/awsdyndns/releases), rename it to awsdyndns and put it into your path.

#### macOs example:
```
cd ~/Downloads
curl -O https://github.com/jsfour/awsdyndns/releases/download/v1/awsdyndns-v1-darwin-amd64
chmod +x awsdyndns-v1-darwin-amd64
sudo mv ./awsdyndns-v1-darwin-amd64 /usr/local/bin/awsdyndns
awsdyndns -h
```

## Building from source
```
git clone git@github.com:jsfour/awsdyndns.git
cd awsdyndns
make all
```

## Usage
Setup your `AWS_ACCESS_KEY` and `AWS_SECRET_KEY` [from AWS](https://docs.aws.amazon.com/general/latest/gr/aws-security-credentials.html). Make sure that this account has permissions to the Route53 zone that you want to modify.

Once you have configured the `AWS_ACCESS_KEY`, and `AWS_SECRET_KEY` environment variables just run `awsdyndns`.

```
$ awsdyndns -d my.host.com. -z ZONEID // don't forget the trailing `.` on the hostname.
Updating dns
```

## Help
```
$ awsdyndns -h
Usage of awsdyndns:
  -d, --dnshostname string   Hostname on the zone id to update
  -z, --zoneid string        Route 53 zone id
```


## Required environment variables:
| Variable |
|----------|
| AWS_ACCESS_KEY |
| AWS_SECRET_KEY |

