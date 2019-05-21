package src

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53"
)

const PublicIpEndpoint = "https://api.ipify.org"

func GetPublicIp() (string, error) {
	res, err := http.Get(PublicIpEndpoint)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)
	s := buf.String()

	return s, nil
}

func SetDns(ip string, dnsName string, zoneID string) error {
	fmt.Println("Updating dns")
	sess := session.Must(session.NewSession(&aws.Config{}))

	svc := route53.New(sess)

	var ttl int64 = 300

	recordSetQuery := &route53.ListResourceRecordSetsInput{
		HostedZoneId:    aws.String(zoneID),
		StartRecordName: aws.String(dnsName),
	}

	res, err := svc.ListResourceRecordSets(recordSetQuery)
	if err != nil {
		return err
	}

	var changes []*route53.Change
	for _, recSet := range res.ResourceRecordSets {
		if *recSet.Name != dnsName {
			continue
		}
		for _, record := range recSet.ResourceRecords {
			// Do I need to change the ip?
			if *record.Value == ip {
				fmt.Println("No change in IP moving along.")
				return nil
			}

			// Need to changes

			ttl = *recSet.TTL
			chg := &route53.Change{
				Action:            aws.String("DELETE"),
				ResourceRecordSet: recSet,
			}
			changes = append(changes, chg)

			break
		}
	}

	newChange := &route53.Change{
		Action: aws.String("CREATE"),
		ResourceRecordSet: &route53.ResourceRecordSet{
			Name: aws.String(dnsName),
			Type: aws.String("A"),
			TTL:  aws.Int64(ttl),
			ResourceRecords: []*route53.ResourceRecord{
				{
					Value: aws.String(ip),
				},
			},
		},
	}

	changes = append(changes, newChange)

	changeSet := &route53.ChangeResourceRecordSetsInput{
		ChangeBatch: &route53.ChangeBatch{
			Changes: changes,
		},
		HostedZoneId: aws.String(zoneID),
	}

	_, err = svc.ChangeResourceRecordSets(changeSet)
	if err != nil {
		return err
	}

	return nil
}
