package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
	instances, err := ListInstances(os.Args[1])
	if err != nil {
		panic(err)
	}
	for _, i := range instances {
		fmt.Println(i)
	}
}

func ListInstances(name string) ([]string, error) {
	sess := session.Must(session.NewSession())

	svc := ec2.New(sess, &aws.Config{Region: aws.String("ap-southeast-2")})

	params := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("tag:Name"),
				Values: []*string{
					aws.String(name),
				},
			},
		},
	}
	resp, err := svc.DescribeInstances(params)

	if err != nil {
		return nil, err
	}

	var instances []string

	for _, r := range resp.Reservations {
		for _, t := range r.Instances[0].Tags {
			if *t.Key == "Name" {
				instances = append(instances, *t.Value)
			}
		}
	}

	return instances, nil
}
