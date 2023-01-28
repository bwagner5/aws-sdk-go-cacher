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
	"github.com/bwagner5/aws-sdk-go-cacher/pkg/cacher"
	"github.com/samber/lo"
)

func main() {

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	ec2svc := cacher.NewClient(ec2.New(sess))
	for i := 0; i < 3; i++ {
		start := time.Now().UTC()
		dio := lo.Must(ec2svc.DescribeInstances(&ec2.DescribeInstancesInput{}))
		duration := time.Since(start)
		instanceIDs := lo.Flatten(lo.Map(dio.Reservations, func(r *ec2.Reservation, _ int) []string {
			return lo.Map(r.Instances, func(i *ec2.Instance, _ int) string { return *i.InstanceId })
		}))
		fmt.Printf("[%d] Took %s to DescribeInstances\n\n", i, duration)
		fmt.Println(instanceIDs)
	}
}