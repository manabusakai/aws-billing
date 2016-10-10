package main

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
)

const (
	region         = "us-east-1"
	namespace      = "AWS/Billing"
	metricName     = "EstimatedCharges"
	dimensionName  = "Currency"
	dimensionValue = "USD"
)

func GetBilling() (float64, error) {
	sess, err := session.NewSession(&aws.Config{Region: aws.String(region)})
	if err != nil {
		return 0, err
	}

	svc := cloudwatch.New(sess)

	params := &cloudwatch.GetMetricStatisticsInput{
		Namespace:  aws.String(namespace),
		MetricName: aws.String(metricName),
		Period:     aws.Int64(21600),
		StartTime:  aws.Time(time.Now().Add(time.Duration(21600) * time.Second * -1)),
		EndTime:    aws.Time(time.Now()),
		Statistics: []*string{
			aws.String(cloudwatch.StatisticMaximum),
		},
		Dimensions: []*cloudwatch.Dimension{
			{
				Name:  aws.String(dimensionName),
				Value: aws.String(dimensionValue),
			},
		},
		Unit: aws.String(cloudwatch.StandardUnitNone),
	}

	resp, err := svc.GetMetricStatistics(params)
	if err != nil {
		return 0, err
	}

	return float64(*resp.Datapoints[0].Maximum), nil
}

func main() {
	billing, err := GetBilling()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Billing: %v\n", billing)
}
