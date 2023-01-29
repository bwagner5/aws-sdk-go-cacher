/*
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/bwagner5/aws-sdk-go-cacher/pkg/ec2cacher"
	"github.com/samber/lo"
)

func main() {

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	ec2client := ec2.New(sess)
	ec2CachingClient := ec2cacher.New(ec2client)
	for i := 0; i < 3; i++ {
		describeInstances(i, ec2CachingClient)
		describeInstancesPages(i, ec2CachingClient)
	}
}

func describeInstances(i int, ec2svc ec2iface.EC2API) {
	start := time.Now().UTC()
	dio := lo.Must(ec2svc.DescribeInstances(&ec2.DescribeInstancesInput{}))
	duration := time.Since(start)
	instanceIDs := lo.Flatten(lo.Map(dio.Reservations, func(r *ec2.Reservation, _ int) []string {
		return lo.Map(r.Instances, func(i *ec2.Instance, _ int) string { return *i.InstanceId })
	}))
	fmt.Printf("[%d] Took %s to DescribeInstances\n\n", i, duration)
	fmt.Println(instanceIDs)
}

func describeInstancesPages(i int, ec2svc ec2iface.EC2API) {
	start := time.Now().UTC()
	var dios []*ec2.DescribeInstancesOutput
	lo.Must0(ec2svc.DescribeInstancesPages(&ec2.DescribeInstancesInput{}, func(dio *ec2.DescribeInstancesOutput, b bool) bool {
		dios = append(dios, dio)
		return true
	}))
	duration := time.Since(start)
	var instanceIDs []string
	for _, dio := range dios {
		instanceIDs = append(instanceIDs, lo.Flatten(lo.Map(dio.Reservations, func(r *ec2.Reservation, _ int) []string {
			return lo.Map(r.Instances, func(i *ec2.Instance, _ int) string { return *i.InstanceId })
		}))...)
	}
	fmt.Printf("[%d] Took %s to DescribeInstancesPages and received %d pages\n\n", i, duration, len(dios))
	fmt.Println(instanceIDs)
}
