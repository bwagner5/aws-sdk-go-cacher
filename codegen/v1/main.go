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

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/service/accessanalyzer/accessanalyzeriface"
	"github.com/aws/aws-sdk-go/service/account/accountiface"
	"github.com/aws/aws-sdk-go/service/acm/acmiface"
	"github.com/aws/aws-sdk-go/service/acmpca/acmpcaiface"
	"github.com/aws/aws-sdk-go/service/alexaforbusiness/alexaforbusinessiface"
	"github.com/aws/aws-sdk-go/service/amplify/amplifyiface"
	"github.com/aws/aws-sdk-go/service/amplifybackend/amplifybackendiface"
	"github.com/aws/aws-sdk-go/service/amplifyuibuilder/amplifyuibuilderiface"
	"github.com/aws/aws-sdk-go/service/apigateway/apigatewayiface"
	"github.com/aws/aws-sdk-go/service/apigatewaymanagementapi/apigatewaymanagementapiiface"
	"github.com/aws/aws-sdk-go/service/apigatewayv2/apigatewayv2iface"
	"github.com/aws/aws-sdk-go/service/appconfig/appconfigiface"
	"github.com/aws/aws-sdk-go/service/appconfigdata/appconfigdataiface"
	"github.com/aws/aws-sdk-go/service/appflow/appflowiface"
	"github.com/aws/aws-sdk-go/service/appintegrationsservice/appintegrationsserviceiface"
	"github.com/aws/aws-sdk-go/service/applicationautoscaling/applicationautoscalingiface"
	"github.com/aws/aws-sdk-go/service/applicationcostprofiler/applicationcostprofileriface"
	"github.com/aws/aws-sdk-go/service/applicationdiscoveryservice/applicationdiscoveryserviceiface"
	"github.com/aws/aws-sdk-go/service/applicationinsights/applicationinsightsiface"
	"github.com/aws/aws-sdk-go/service/appmesh/appmeshiface"
	"github.com/aws/aws-sdk-go/service/appregistry/appregistryiface"
	"github.com/aws/aws-sdk-go/service/apprunner/apprunneriface"
	"github.com/aws/aws-sdk-go/service/appstream/appstreamiface"
	"github.com/aws/aws-sdk-go/service/appsync/appsynciface"
	"github.com/aws/aws-sdk-go/service/arczonalshift/arczonalshiftiface"
	"github.com/aws/aws-sdk-go/service/athena/athenaiface"
	"github.com/aws/aws-sdk-go/service/auditmanager/auditmanageriface"
	"github.com/aws/aws-sdk-go/service/augmentedairuntime/augmentedairuntimeiface"
	"github.com/aws/aws-sdk-go/service/autoscaling/autoscalingiface"
	"github.com/aws/aws-sdk-go/service/autoscalingplans/autoscalingplansiface"
	"github.com/aws/aws-sdk-go/service/backup/backupiface"
	"github.com/aws/aws-sdk-go/service/backupgateway/backupgatewayiface"
	"github.com/aws/aws-sdk-go/service/backupstorage/backupstorageiface"
	"github.com/aws/aws-sdk-go/service/batch/batchiface"
	"github.com/aws/aws-sdk-go/service/billingconductor/billingconductoriface"
	"github.com/aws/aws-sdk-go/service/braket/braketiface"
	"github.com/aws/aws-sdk-go/service/budgets/budgetsiface"
	"github.com/aws/aws-sdk-go/service/chime/chimeiface"
	"github.com/aws/aws-sdk-go/service/chimesdkidentity/chimesdkidentityiface"
	"github.com/aws/aws-sdk-go/service/chimesdkmediapipelines/chimesdkmediapipelinesiface"
	"github.com/aws/aws-sdk-go/service/chimesdkmeetings/chimesdkmeetingsiface"
	"github.com/aws/aws-sdk-go/service/chimesdkmessaging/chimesdkmessagingiface"
	"github.com/aws/aws-sdk-go/service/chimesdkvoice/chimesdkvoiceiface"
	"github.com/aws/aws-sdk-go/service/cleanrooms/cleanroomsiface"
	"github.com/aws/aws-sdk-go/service/cloud9/cloud9iface"
	"github.com/aws/aws-sdk-go/service/cloudcontrolapi/cloudcontrolapiiface"
	"github.com/aws/aws-sdk-go/service/clouddirectory/clouddirectoryiface"
	"github.com/aws/aws-sdk-go/service/cloudformation/cloudformationiface"
	"github.com/aws/aws-sdk-go/service/cloudfront/cloudfrontiface"
	"github.com/aws/aws-sdk-go/service/cloudhsm/cloudhsmiface"
	"github.com/aws/aws-sdk-go/service/cloudhsmv2/cloudhsmv2iface"
	"github.com/aws/aws-sdk-go/service/cloudsearch/cloudsearchiface"
	"github.com/aws/aws-sdk-go/service/cloudsearchdomain/cloudsearchdomainiface"
	"github.com/aws/aws-sdk-go/service/cloudtrail/cloudtrailiface"
	"github.com/aws/aws-sdk-go/service/cloudwatch/cloudwatchiface"
	"github.com/aws/aws-sdk-go/service/cloudwatchevents/cloudwatcheventsiface"
	"github.com/aws/aws-sdk-go/service/cloudwatchevidently/cloudwatchevidentlyiface"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs/cloudwatchlogsiface"
	"github.com/aws/aws-sdk-go/service/cloudwatchrum/cloudwatchrumiface"
	"github.com/aws/aws-sdk-go/service/codeartifact/codeartifactiface"
	"github.com/aws/aws-sdk-go/service/codebuild/codebuildiface"
	"github.com/aws/aws-sdk-go/service/codecommit/codecommitiface"
	"github.com/aws/aws-sdk-go/service/codedeploy/codedeployiface"
	"github.com/aws/aws-sdk-go/service/codeguruprofiler/codeguruprofileriface"
	"github.com/aws/aws-sdk-go/service/codegurureviewer/codegururevieweriface"
	"github.com/aws/aws-sdk-go/service/codepipeline/codepipelineiface"
	"github.com/aws/aws-sdk-go/service/codestar/codestariface"
	"github.com/aws/aws-sdk-go/service/codestarconnections/codestarconnectionsiface"
	"github.com/aws/aws-sdk-go/service/codestarnotifications/codestarnotificationsiface"
	"github.com/aws/aws-sdk-go/service/cognitoidentity/cognitoidentityiface"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider/cognitoidentityprovideriface"
	"github.com/aws/aws-sdk-go/service/cognitosync/cognitosynciface"
	"github.com/aws/aws-sdk-go/service/comprehend/comprehendiface"
	"github.com/aws/aws-sdk-go/service/comprehendmedical/comprehendmedicaliface"
	"github.com/aws/aws-sdk-go/service/computeoptimizer/computeoptimizeriface"
	"github.com/aws/aws-sdk-go/service/configservice/configserviceiface"
	"github.com/aws/aws-sdk-go/service/connect/connectiface"
	"github.com/aws/aws-sdk-go/service/connectcampaigns/connectcampaignsiface"
	"github.com/aws/aws-sdk-go/service/connectcases/connectcasesiface"
	"github.com/aws/aws-sdk-go/service/connectcontactlens/connectcontactlensiface"
	"github.com/aws/aws-sdk-go/service/connectparticipant/connectparticipantiface"
	"github.com/aws/aws-sdk-go/service/connectwisdomservice/connectwisdomserviceiface"
	"github.com/aws/aws-sdk-go/service/controltower/controltoweriface"
	"github.com/aws/aws-sdk-go/service/costandusagereportservice/costandusagereportserviceiface"
	"github.com/aws/aws-sdk-go/service/costexplorer/costexploreriface"
	"github.com/aws/aws-sdk-go/service/customerprofiles/customerprofilesiface"
	"github.com/aws/aws-sdk-go/service/databasemigrationservice/databasemigrationserviceiface"
	"github.com/aws/aws-sdk-go/service/dataexchange/dataexchangeiface"
	"github.com/aws/aws-sdk-go/service/datapipeline/datapipelineiface"
	"github.com/aws/aws-sdk-go/service/datasync/datasynciface"
	"github.com/aws/aws-sdk-go/service/dax/daxiface"
	"github.com/aws/aws-sdk-go/service/detective/detectiveiface"
	"github.com/aws/aws-sdk-go/service/devicefarm/devicefarmiface"
	"github.com/aws/aws-sdk-go/service/devopsguru/devopsguruiface"
	"github.com/aws/aws-sdk-go/service/directconnect/directconnectiface"
	"github.com/aws/aws-sdk-go/service/directoryservice/directoryserviceiface"
	"github.com/aws/aws-sdk-go/service/dlm/dlmiface"
	"github.com/aws/aws-sdk-go/service/docdb/docdbiface"
	"github.com/aws/aws-sdk-go/service/docdbelastic/docdbelasticiface"
	"github.com/aws/aws-sdk-go/service/drs/drsiface"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/aws/aws-sdk-go/service/dynamodbstreams/dynamodbstreamsiface"
	"github.com/aws/aws-sdk-go/service/ebs/ebsiface"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/aws/aws-sdk-go/service/ec2instanceconnect/ec2instanceconnectiface"
	"github.com/aws/aws-sdk-go/service/ecr/ecriface"
	"github.com/aws/aws-sdk-go/service/ecrpublic/ecrpubliciface"
	"github.com/aws/aws-sdk-go/service/ecs/ecsiface"
	"github.com/aws/aws-sdk-go/service/efs/efsiface"
	"github.com/aws/aws-sdk-go/service/eks/eksiface"
	"github.com/aws/aws-sdk-go/service/elasticache/elasticacheiface"
	"github.com/aws/aws-sdk-go/service/elasticbeanstalk/elasticbeanstalkiface"
	"github.com/aws/aws-sdk-go/service/elasticinference/elasticinferenceiface"
	"github.com/aws/aws-sdk-go/service/elasticsearchservice/elasticsearchserviceiface"
	"github.com/aws/aws-sdk-go/service/elastictranscoder/elastictranscoderiface"
	"github.com/aws/aws-sdk-go/service/elb/elbiface"
	"github.com/aws/aws-sdk-go/service/elbv2/elbv2iface"
	"github.com/aws/aws-sdk-go/service/emr/emriface"
	"github.com/aws/aws-sdk-go/service/emrcontainers/emrcontainersiface"
	"github.com/aws/aws-sdk-go/service/emrserverless/emrserverlessiface"
	"github.com/aws/aws-sdk-go/service/eventbridge/eventbridgeiface"
	"github.com/aws/aws-sdk-go/service/finspace/finspaceiface"
	"github.com/aws/aws-sdk-go/service/finspacedata/finspacedataiface"
	"github.com/aws/aws-sdk-go/service/firehose/firehoseiface"
	"github.com/aws/aws-sdk-go/service/fis/fisiface"
	"github.com/aws/aws-sdk-go/service/fms/fmsiface"
	"github.com/aws/aws-sdk-go/service/forecastqueryservice/forecastqueryserviceiface"
	"github.com/aws/aws-sdk-go/service/forecastservice/forecastserviceiface"
	"github.com/aws/aws-sdk-go/service/frauddetector/frauddetectoriface"
	"github.com/aws/aws-sdk-go/service/fsx/fsxiface"
	"github.com/aws/aws-sdk-go/service/gamelift/gameliftiface"
	"github.com/aws/aws-sdk-go/service/gamesparks/gamesparksiface"
	"github.com/aws/aws-sdk-go/service/glacier/glacieriface"
	"github.com/aws/aws-sdk-go/service/globalaccelerator/globalacceleratoriface"
	"github.com/aws/aws-sdk-go/service/glue/glueiface"
	"github.com/aws/aws-sdk-go/service/gluedatabrew/gluedatabrewiface"
	"github.com/aws/aws-sdk-go/service/greengrass/greengrassiface"
	"github.com/aws/aws-sdk-go/service/greengrassv2/greengrassv2iface"
	"github.com/aws/aws-sdk-go/service/groundstation/groundstationiface"
	"github.com/aws/aws-sdk-go/service/guardduty/guarddutyiface"
	"github.com/aws/aws-sdk-go/service/health/healthiface"
	"github.com/aws/aws-sdk-go/service/healthlake/healthlakeiface"
	"github.com/aws/aws-sdk-go/service/honeycode/honeycodeiface"
	"github.com/aws/aws-sdk-go/service/iam/iamiface"
	"github.com/aws/aws-sdk-go/service/identitystore/identitystoreiface"
	"github.com/aws/aws-sdk-go/service/imagebuilder/imagebuilderiface"
	"github.com/aws/aws-sdk-go/service/inspector/inspectoriface"
	"github.com/aws/aws-sdk-go/service/inspector2/inspector2iface"
	"github.com/aws/aws-sdk-go/service/iot/iotiface"
	"github.com/aws/aws-sdk-go/service/iot1clickdevicesservice/iot1clickdevicesserviceiface"
	"github.com/aws/aws-sdk-go/service/iot1clickprojects/iot1clickprojectsiface"
	"github.com/aws/aws-sdk-go/service/iotanalytics/iotanalyticsiface"
	"github.com/aws/aws-sdk-go/service/iotdataplane/iotdataplaneiface"
	"github.com/aws/aws-sdk-go/service/iotdeviceadvisor/iotdeviceadvisoriface"
	"github.com/aws/aws-sdk-go/service/iotevents/ioteventsiface"
	"github.com/aws/aws-sdk-go/service/ioteventsdata/ioteventsdataiface"
	"github.com/aws/aws-sdk-go/service/iotfleethub/iotfleethubiface"
	"github.com/aws/aws-sdk-go/service/iotfleetwise/iotfleetwiseiface"
	"github.com/aws/aws-sdk-go/service/iotjobsdataplane/iotjobsdataplaneiface"
	"github.com/aws/aws-sdk-go/service/iotroborunner/iotroborunneriface"
	"github.com/aws/aws-sdk-go/service/iotsecuretunneling/iotsecuretunnelingiface"
	"github.com/aws/aws-sdk-go/service/iotsitewise/iotsitewiseiface"
	"github.com/aws/aws-sdk-go/service/iotthingsgraph/iotthingsgraphiface"
	"github.com/aws/aws-sdk-go/service/iottwinmaker/iottwinmakeriface"
	"github.com/aws/aws-sdk-go/service/iotwireless/iotwirelessiface"
	"github.com/aws/aws-sdk-go/service/ivs/ivsiface"
	"github.com/aws/aws-sdk-go/service/ivschat/ivschatiface"
	"github.com/aws/aws-sdk-go/service/kafka/kafkaiface"
	"github.com/aws/aws-sdk-go/service/kafkaconnect/kafkaconnectiface"
	"github.com/aws/aws-sdk-go/service/kendra/kendraiface"
	"github.com/aws/aws-sdk-go/service/kendraranking/kendrarankingiface"
	"github.com/aws/aws-sdk-go/service/keyspaces/keyspacesiface"
	"github.com/aws/aws-sdk-go/service/kinesis/kinesisiface"
	"github.com/aws/aws-sdk-go/service/kinesisanalytics/kinesisanalyticsiface"
	"github.com/aws/aws-sdk-go/service/kinesisanalyticsv2/kinesisanalyticsv2iface"
	"github.com/aws/aws-sdk-go/service/kinesisvideo/kinesisvideoiface"
	"github.com/aws/aws-sdk-go/service/kinesisvideoarchivedmedia/kinesisvideoarchivedmediaiface"
	"github.com/aws/aws-sdk-go/service/kinesisvideomedia/kinesisvideomediaiface"
	"github.com/aws/aws-sdk-go/service/kinesisvideosignalingchannels/kinesisvideosignalingchannelsiface"
	"github.com/aws/aws-sdk-go/service/kinesisvideowebrtcstorage/kinesisvideowebrtcstorageiface"
	"github.com/aws/aws-sdk-go/service/kms/kmsiface"
	"github.com/aws/aws-sdk-go/service/lakeformation/lakeformationiface"
	"github.com/aws/aws-sdk-go/service/lambda/lambdaiface"
	"github.com/aws/aws-sdk-go/service/lexmodelbuildingservice/lexmodelbuildingserviceiface"
	"github.com/aws/aws-sdk-go/service/lexmodelsv2/lexmodelsv2iface"
	"github.com/aws/aws-sdk-go/service/lexruntimeservice/lexruntimeserviceiface"
	"github.com/aws/aws-sdk-go/service/lexruntimev2/lexruntimev2iface"
	"github.com/aws/aws-sdk-go/service/licensemanager/licensemanageriface"
	"github.com/aws/aws-sdk-go/service/licensemanagerlinuxsubscriptions/licensemanagerlinuxsubscriptionsiface"
	"github.com/aws/aws-sdk-go/service/licensemanagerusersubscriptions/licensemanagerusersubscriptionsiface"
	"github.com/aws/aws-sdk-go/service/lightsail/lightsailiface"
	"github.com/aws/aws-sdk-go/service/locationservice/locationserviceiface"
	"github.com/aws/aws-sdk-go/service/lookoutequipment/lookoutequipmentiface"
	"github.com/aws/aws-sdk-go/service/lookoutforvision/lookoutforvisioniface"
	"github.com/aws/aws-sdk-go/service/lookoutmetrics/lookoutmetricsiface"
	"github.com/aws/aws-sdk-go/service/m2/m2iface"
	"github.com/aws/aws-sdk-go/service/machinelearning/machinelearningiface"
	"github.com/aws/aws-sdk-go/service/macie/macieiface"
	"github.com/aws/aws-sdk-go/service/macie2/macie2iface"
	"github.com/aws/aws-sdk-go/service/managedblockchain/managedblockchainiface"
	"github.com/aws/aws-sdk-go/service/managedgrafana/managedgrafanaiface"
	"github.com/aws/aws-sdk-go/service/marketplacecatalog/marketplacecatalogiface"
	"github.com/aws/aws-sdk-go/service/marketplacecommerceanalytics/marketplacecommerceanalyticsiface"
	"github.com/aws/aws-sdk-go/service/marketplaceentitlementservice/marketplaceentitlementserviceiface"
	"github.com/aws/aws-sdk-go/service/marketplacemetering/marketplacemeteringiface"
	"github.com/aws/aws-sdk-go/service/mediaconnect/mediaconnectiface"
	"github.com/aws/aws-sdk-go/service/mediaconvert/mediaconvertiface"
	"github.com/aws/aws-sdk-go/service/medialive/medialiveiface"
	"github.com/aws/aws-sdk-go/service/mediapackage/mediapackageiface"
	"github.com/aws/aws-sdk-go/service/mediapackagevod/mediapackagevodiface"
	"github.com/aws/aws-sdk-go/service/mediastore/mediastoreiface"
	"github.com/aws/aws-sdk-go/service/mediastoredata/mediastoredataiface"
	"github.com/aws/aws-sdk-go/service/mediatailor/mediatailoriface"
	"github.com/aws/aws-sdk-go/service/memorydb/memorydbiface"
	"github.com/aws/aws-sdk-go/service/mgn/mgniface"
	"github.com/aws/aws-sdk-go/service/migrationhub/migrationhubiface"
	"github.com/aws/aws-sdk-go/service/migrationhubconfig/migrationhubconfigiface"
	"github.com/aws/aws-sdk-go/service/migrationhuborchestrator/migrationhuborchestratoriface"
	"github.com/aws/aws-sdk-go/service/migrationhubrefactorspaces/migrationhubrefactorspacesiface"
	"github.com/aws/aws-sdk-go/service/migrationhubstrategyrecommendations/migrationhubstrategyrecommendationsiface"
	"github.com/aws/aws-sdk-go/service/mobile/mobileiface"
	"github.com/aws/aws-sdk-go/service/mobileanalytics/mobileanalyticsiface"
	"github.com/aws/aws-sdk-go/service/mq/mqiface"
	"github.com/aws/aws-sdk-go/service/mturk/mturkiface"
	"github.com/aws/aws-sdk-go/service/mwaa/mwaaiface"
	"github.com/aws/aws-sdk-go/service/neptune/neptuneiface"
	"github.com/aws/aws-sdk-go/service/networkfirewall/networkfirewalliface"
	"github.com/aws/aws-sdk-go/service/networkmanager/networkmanageriface"
	"github.com/aws/aws-sdk-go/service/nimblestudio/nimblestudioiface"
	"github.com/aws/aws-sdk-go/service/oam/oamiface"
	"github.com/aws/aws-sdk-go/service/omics/omicsiface"
	"github.com/aws/aws-sdk-go/service/opensearchserverless/opensearchserverlessiface"
	"github.com/aws/aws-sdk-go/service/opensearchservice/opensearchserviceiface"
	"github.com/aws/aws-sdk-go/service/opsworks/opsworksiface"
	"github.com/aws/aws-sdk-go/service/opsworkscm/opsworkscmiface"
	"github.com/aws/aws-sdk-go/service/organizations/organizationsiface"
	"github.com/aws/aws-sdk-go/service/outposts/outpostsiface"
	"github.com/aws/aws-sdk-go/service/panorama/panoramaiface"
	"github.com/aws/aws-sdk-go/service/personalize/personalizeiface"
	"github.com/aws/aws-sdk-go/service/personalizeevents/personalizeeventsiface"
	"github.com/aws/aws-sdk-go/service/personalizeruntime/personalizeruntimeiface"
	"github.com/aws/aws-sdk-go/service/pi/piiface"
	"github.com/aws/aws-sdk-go/service/pinpoint/pinpointiface"
	"github.com/aws/aws-sdk-go/service/pinpointemail/pinpointemailiface"
	"github.com/aws/aws-sdk-go/service/pinpointsmsvoice/pinpointsmsvoiceiface"
	"github.com/aws/aws-sdk-go/service/pinpointsmsvoicev2/pinpointsmsvoicev2iface"
	"github.com/aws/aws-sdk-go/service/pipes/pipesiface"
	"github.com/aws/aws-sdk-go/service/polly/pollyiface"
	"github.com/aws/aws-sdk-go/service/pricing/pricingiface"
	"github.com/aws/aws-sdk-go/service/privatenetworks/privatenetworksiface"
	"github.com/aws/aws-sdk-go/service/prometheusservice/prometheusserviceiface"
	"github.com/aws/aws-sdk-go/service/proton/protoniface"
	"github.com/aws/aws-sdk-go/service/qldb/qldbiface"
	"github.com/aws/aws-sdk-go/service/qldbsession/qldbsessioniface"
	"github.com/aws/aws-sdk-go/service/quicksight/quicksightiface"
	"github.com/aws/aws-sdk-go/service/ram/ramiface"
	"github.com/aws/aws-sdk-go/service/rds/rdsiface"
	"github.com/aws/aws-sdk-go/service/rdsdataservice/rdsdataserviceiface"
	"github.com/aws/aws-sdk-go/service/recyclebin/recyclebiniface"
	"github.com/aws/aws-sdk-go/service/redshift/redshiftiface"
	"github.com/aws/aws-sdk-go/service/redshiftdataapiservice/redshiftdataapiserviceiface"
	"github.com/aws/aws-sdk-go/service/redshiftserverless/redshiftserverlessiface"
	"github.com/aws/aws-sdk-go/service/rekognition/rekognitioniface"
	"github.com/aws/aws-sdk-go/service/resiliencehub/resiliencehubiface"
	"github.com/aws/aws-sdk-go/service/resourceexplorer2/resourceexplorer2iface"
	"github.com/aws/aws-sdk-go/service/resourcegroups/resourcegroupsiface"
	"github.com/aws/aws-sdk-go/service/resourcegroupstaggingapi/resourcegroupstaggingapiiface"
	"github.com/aws/aws-sdk-go/service/robomaker/robomakeriface"
	"github.com/aws/aws-sdk-go/service/rolesanywhere/rolesanywhereiface"
	"github.com/aws/aws-sdk-go/service/route53/route53iface"
	"github.com/aws/aws-sdk-go/service/route53domains/route53domainsiface"
	"github.com/aws/aws-sdk-go/service/route53recoverycluster/route53recoveryclusteriface"
	"github.com/aws/aws-sdk-go/service/route53recoverycontrolconfig/route53recoverycontrolconfigiface"
	"github.com/aws/aws-sdk-go/service/route53recoveryreadiness/route53recoveryreadinessiface"
	"github.com/aws/aws-sdk-go/service/route53resolver/route53resolveriface"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/aws/aws-sdk-go/service/s3control/s3controliface"
	"github.com/aws/aws-sdk-go/service/s3outposts/s3outpostsiface"
	"github.com/aws/aws-sdk-go/service/sagemaker/sagemakeriface"
	"github.com/aws/aws-sdk-go/service/sagemakeredgemanager/sagemakeredgemanageriface"
	"github.com/aws/aws-sdk-go/service/sagemakerfeaturestoreruntime/sagemakerfeaturestoreruntimeiface"
	"github.com/aws/aws-sdk-go/service/sagemakergeospatial/sagemakergeospatialiface"
	"github.com/aws/aws-sdk-go/service/sagemakermetrics/sagemakermetricsiface"
	"github.com/aws/aws-sdk-go/service/sagemakerruntime/sagemakerruntimeiface"
	"github.com/aws/aws-sdk-go/service/savingsplans/savingsplansiface"
	"github.com/aws/aws-sdk-go/service/scheduler/scheduleriface"
	"github.com/aws/aws-sdk-go/service/schemas/schemasiface"
	"github.com/aws/aws-sdk-go/service/secretsmanager/secretsmanageriface"
	"github.com/aws/aws-sdk-go/service/securityhub/securityhubiface"
	"github.com/aws/aws-sdk-go/service/securitylake/securitylakeiface"
	"github.com/aws/aws-sdk-go/service/serverlessapplicationrepository/serverlessapplicationrepositoryiface"
	"github.com/aws/aws-sdk-go/service/servicecatalog/servicecatalogiface"
	"github.com/aws/aws-sdk-go/service/servicediscovery/servicediscoveryiface"
	"github.com/aws/aws-sdk-go/service/servicequotas/servicequotasiface"
	"github.com/aws/aws-sdk-go/service/ses/sesiface"
	"github.com/aws/aws-sdk-go/service/sesv2/sesv2iface"
	"github.com/aws/aws-sdk-go/service/sfn/sfniface"
	"github.com/aws/aws-sdk-go/service/shield/shieldiface"
	"github.com/aws/aws-sdk-go/service/signer/signeriface"
	"github.com/aws/aws-sdk-go/service/simpledb/simpledbiface"
	"github.com/aws/aws-sdk-go/service/simspaceweaver/simspaceweaveriface"
	"github.com/aws/aws-sdk-go/service/sms/smsiface"
	"github.com/aws/aws-sdk-go/service/snowball/snowballiface"
	"github.com/aws/aws-sdk-go/service/snowdevicemanagement/snowdevicemanagementiface"
	"github.com/aws/aws-sdk-go/service/sns/snsiface"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
	"github.com/aws/aws-sdk-go/service/ssm/ssmiface"
	"github.com/aws/aws-sdk-go/service/ssmcontacts/ssmcontactsiface"
	"github.com/aws/aws-sdk-go/service/ssmincidents/ssmincidentsiface"
	"github.com/aws/aws-sdk-go/service/ssmsap/ssmsapiface"
	"github.com/aws/aws-sdk-go/service/sso/ssoiface"
	"github.com/aws/aws-sdk-go/service/ssoadmin/ssoadminiface"
	"github.com/aws/aws-sdk-go/service/ssooidc/ssooidciface"
	"github.com/aws/aws-sdk-go/service/storagegateway/storagegatewayiface"
	"github.com/aws/aws-sdk-go/service/sts/stsiface"
	"github.com/aws/aws-sdk-go/service/support/supportiface"
	"github.com/aws/aws-sdk-go/service/supportapp/supportappiface"
	"github.com/aws/aws-sdk-go/service/swf/swfiface"
	"github.com/aws/aws-sdk-go/service/synthetics/syntheticsiface"
	"github.com/aws/aws-sdk-go/service/textract/textractiface"
	"github.com/aws/aws-sdk-go/service/timestreamquery/timestreamqueryiface"
	"github.com/aws/aws-sdk-go/service/timestreamwrite/timestreamwriteiface"
	"github.com/aws/aws-sdk-go/service/transcribeservice/transcribeserviceiface"
	"github.com/aws/aws-sdk-go/service/transcribestreamingservice/transcribestreamingserviceiface"
	"github.com/aws/aws-sdk-go/service/transfer/transferiface"
	"github.com/aws/aws-sdk-go/service/translate/translateiface"
	"github.com/aws/aws-sdk-go/service/voiceid/voiceidiface"
	"github.com/aws/aws-sdk-go/service/waf/wafiface"
	"github.com/aws/aws-sdk-go/service/wafregional/wafregionaliface"
	"github.com/aws/aws-sdk-go/service/wafv2/wafv2iface"
	"github.com/aws/aws-sdk-go/service/wellarchitected/wellarchitectediface"
	"github.com/aws/aws-sdk-go/service/workdocs/workdocsiface"
	"github.com/aws/aws-sdk-go/service/worklink/worklinkiface"
	"github.com/aws/aws-sdk-go/service/workmail/workmailiface"
	"github.com/aws/aws-sdk-go/service/workmailmessageflow/workmailmessageflowiface"
	"github.com/aws/aws-sdk-go/service/workspaces/workspacesiface"
	"github.com/aws/aws-sdk-go/service/workspacesweb/workspaceswebiface"
	"github.com/aws/aws-sdk-go/service/xray/xrayiface"
)

func main() {

	outputDir := flag.String("out-dir", "pkg", "directory to output generated caching clients")
	flag.Parse()

	var serviceOutDir string
	var err error
	var out string
	out, err = GenSDK[accessanalyzeriface.AccessAnalyzerAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "accessanalyzerapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/accessanalyzercacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/accessanalyzercacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[accountiface.AccountAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "accountapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/accountcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/accountcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[acmiface.ACMAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "acmapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/acmcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/acmcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[acmpcaiface.ACMPCAAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "acmpcaapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/acmpcacacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/acmpcacacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[alexaforbusinessiface.AlexaForBusinessAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "alexaforbusinessapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/alexaforbusinesscacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/alexaforbusinesscacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[amplifyiface.AmplifyAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "amplifyapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/amplifycacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/amplifycacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[amplifybackendiface.AmplifyBackendAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "amplifybackendapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/amplifybackendcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/amplifybackendcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[amplifyuibuilderiface.AmplifyUIBuilderAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "amplifyuibuilderapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/amplifyuibuildercacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/amplifyuibuildercacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[apigatewayiface.APIGatewayAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "apigatewayapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/apigatewaycacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/apigatewaycacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[apigatewaymanagementapiiface.ApiGatewayManagementApiAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "apigatewaymanagementapiapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/apigatewaymanagementapicacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/apigatewaymanagementapicacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[apigatewayv2iface.ApiGatewayV2API]()
	if err != nil {
		log.Printf("%s: %v\n", "apigatewayv2api.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/apigatewayv2cacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/apigatewayv2cacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[appconfigiface.AppConfigAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "appconfigapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/appconfigcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/appconfigcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[appconfigdataiface.AppConfigDataAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "appconfigdataapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/appconfigdatacacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/appconfigdatacacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[appflowiface.AppflowAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "appflowapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/appflowcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/appflowcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[appintegrationsserviceiface.AppIntegrationsServiceAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "appintegrationsserviceapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/appintegrationsservicecacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/appintegrationsservicecacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[applicationautoscalingiface.ApplicationAutoScalingAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "applicationautoscalingapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/applicationautoscalingcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/applicationautoscalingcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[applicationcostprofileriface.ApplicationCostProfilerAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "applicationcostprofilerapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/applicationcostprofilercacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/applicationcostprofilercacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[applicationdiscoveryserviceiface.ApplicationDiscoveryServiceAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "applicationdiscoveryserviceapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/applicationdiscoveryservicecacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/applicationdiscoveryservicecacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[applicationinsightsiface.ApplicationInsightsAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "applicationinsightsapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/applicationinsightscacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/applicationinsightscacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[appmeshiface.AppMeshAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "appmeshapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/appmeshcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/appmeshcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[appregistryiface.AppRegistryAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "appregistryapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/appregistrycacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/appregistrycacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[apprunneriface.AppRunnerAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "apprunnerapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/apprunnercacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/apprunnercacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[appstreamiface.AppStreamAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "appstreamapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/appstreamcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/appstreamcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[appsynciface.AppSyncAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "appsyncapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/appsynccacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/appsynccacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[arczonalshiftiface.ARCZonalShiftAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "arczonalshiftapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/arczonalshiftcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/arczonalshiftcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[athenaiface.AthenaAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "athenaapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/athenacacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/athenacacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[auditmanageriface.AuditManagerAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "auditmanagerapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/auditmanagercacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/auditmanagercacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[augmentedairuntimeiface.AugmentedAIRuntimeAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "augmentedairuntimeapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/augmentedairuntimecacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/augmentedairuntimecacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[autoscalingiface.AutoScalingAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "autoscalingapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/autoscalingcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/autoscalingcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[autoscalingplansiface.AutoScalingPlansAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "autoscalingplansapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/autoscalingplanscacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/autoscalingplanscacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[backupiface.BackupAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "backupapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/backupcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/backupcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[backupgatewayiface.BackupGatewayAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "backupgatewayapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/backupgatewaycacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/backupgatewaycacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[backupstorageiface.BackupStorageAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "backupstorageapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/backupstoragecacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/backupstoragecacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[batchiface.BatchAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "batchapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/batchcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/batchcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[billingconductoriface.BillingConductorAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "billingconductorapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/billingconductorcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/billingconductorcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[braketiface.BraketAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "braketapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/braketcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/braketcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[budgetsiface.BudgetsAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "budgetsapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/budgetscacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/budgetscacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[chimeiface.ChimeAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "chimeapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/chimecacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/chimecacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[chimesdkidentityiface.ChimeSDKIdentityAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "chimesdkidentityapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/chimesdkidentitycacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/chimesdkidentitycacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[chimesdkmediapipelinesiface.ChimeSDKMediaPipelinesAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "chimesdkmediapipelinesapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/chimesdkmediapipelinescacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/chimesdkmediapipelinescacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[chimesdkmeetingsiface.ChimeSDKMeetingsAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "chimesdkmeetingsapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/chimesdkmeetingscacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/chimesdkmeetingscacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[chimesdkmessagingiface.ChimeSDKMessagingAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "chimesdkmessagingapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/chimesdkmessagingcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/chimesdkmessagingcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[chimesdkvoiceiface.ChimeSDKVoiceAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "chimesdkvoiceapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/chimesdkvoicecacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/chimesdkvoicecacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[cleanroomsiface.CleanRoomsAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "cleanroomsapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/cleanroomscacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/cleanroomscacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[cloud9iface.Cloud9API]()
	if err != nil {
		log.Printf("%s: %v\n", "cloud9api.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/cloud9cacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/cloud9cacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[cloudcontrolapiiface.CloudControlApiAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "cloudcontrolapiapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/cloudcontrolapicacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/cloudcontrolapicacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[clouddirectoryiface.CloudDirectoryAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "clouddirectoryapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/clouddirectorycacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/clouddirectorycacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[cloudformationiface.CloudFormationAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "cloudformationapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/cloudformationcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/cloudformationcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[cloudfrontiface.CloudFrontAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "cloudfrontapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/cloudfrontcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/cloudfrontcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[cloudhsmiface.CloudHSMAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "cloudhsmapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/cloudhsmcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/cloudhsmcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[cloudhsmv2iface.CloudHSMV2API]()
	if err != nil {
		log.Printf("%s: %v\n", "cloudhsmv2api.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/cloudhsmv2cacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/cloudhsmv2cacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[cloudsearchiface.CloudSearchAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "cloudsearchapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/cloudsearchcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/cloudsearchcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[cloudsearchdomainiface.CloudSearchDomainAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "cloudsearchdomainapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/cloudsearchdomaincacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/cloudsearchdomaincacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[cloudtrailiface.CloudTrailAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "cloudtrailapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/cloudtrailcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/cloudtrailcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[cloudwatchiface.CloudWatchAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "cloudwatchapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/cloudwatchcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/cloudwatchcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[cloudwatcheventsiface.CloudWatchEventsAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "cloudwatcheventsapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/cloudwatcheventscacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/cloudwatcheventscacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[cloudwatchevidentlyiface.CloudWatchEvidentlyAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "cloudwatchevidentlyapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/cloudwatchevidentlycacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/cloudwatchevidentlycacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[cloudwatchlogsiface.CloudWatchLogsAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "cloudwatchlogsapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/cloudwatchlogscacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/cloudwatchlogscacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[cloudwatchrumiface.CloudWatchRUMAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "cloudwatchrumapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/cloudwatchrumcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/cloudwatchrumcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[codeartifactiface.CodeArtifactAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "codeartifactapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/codeartifactcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/codeartifactcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[codebuildiface.CodeBuildAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "codebuildapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/codebuildcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/codebuildcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[codecommitiface.CodeCommitAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "codecommitapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/codecommitcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/codecommitcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[codedeployiface.CodeDeployAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "codedeployapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/codedeploycacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/codedeploycacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[codeguruprofileriface.CodeGuruProfilerAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "codeguruprofilerapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/codeguruprofilercacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/codeguruprofilercacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[codegururevieweriface.CodeGuruReviewerAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "codegurureviewerapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/codegurureviewercacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/codegurureviewercacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[codepipelineiface.CodePipelineAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "codepipelineapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/codepipelinecacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/codepipelinecacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[codestariface.CodeStarAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "codestarapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/codestarcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/codestarcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[codestarconnectionsiface.CodeStarConnectionsAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "codestarconnectionsapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/codestarconnectionscacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/codestarconnectionscacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[codestarnotificationsiface.CodeStarNotificationsAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "codestarnotificationsapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/codestarnotificationscacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/codestarnotificationscacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[cognitoidentityiface.CognitoIdentityAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "cognitoidentityapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/cognitoidentitycacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/cognitoidentitycacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[cognitoidentityprovideriface.CognitoIdentityProviderAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "cognitoidentityproviderapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/cognitoidentityprovidercacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/cognitoidentityprovidercacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[cognitosynciface.CognitoSyncAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "cognitosyncapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/cognitosynccacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/cognitosynccacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[comprehendiface.ComprehendAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "comprehendapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/comprehendcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/comprehendcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[comprehendmedicaliface.ComprehendMedicalAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "comprehendmedicalapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/comprehendmedicalcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/comprehendmedicalcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[computeoptimizeriface.ComputeOptimizerAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "computeoptimizerapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/computeoptimizercacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/computeoptimizercacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[configserviceiface.ConfigServiceAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "configserviceapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/configservicecacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/configservicecacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[connectiface.ConnectAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "connectapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/connectcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/connectcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[connectcampaignsiface.ConnectCampaignsAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "connectcampaignsapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/connectcampaignscacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/connectcampaignscacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[connectcasesiface.ConnectCasesAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "connectcasesapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/connectcasescacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/connectcasescacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[connectcontactlensiface.ConnectContactLensAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "connectcontactlensapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/connectcontactlenscacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/connectcontactlenscacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[connectparticipantiface.ConnectParticipantAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "connectparticipantapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/connectparticipantcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/connectparticipantcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[connectwisdomserviceiface.ConnectWisdomServiceAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "connectwisdomserviceapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/connectwisdomservicecacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/connectwisdomservicecacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[controltoweriface.ControlTowerAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "controltowerapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/controltowercacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/controltowercacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[costandusagereportserviceiface.CostandUsageReportServiceAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "costandusagereportserviceapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/costandusagereportservicecacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/costandusagereportservicecacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[costexploreriface.CostExplorerAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "costexplorerapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/costexplorercacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/costexplorercacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[customerprofilesiface.CustomerProfilesAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "customerprofilesapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/customerprofilescacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/customerprofilescacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[databasemigrationserviceiface.DatabaseMigrationServiceAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "databasemigrationserviceapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/databasemigrationservicecacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/databasemigrationservicecacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[dataexchangeiface.DataExchangeAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "dataexchangeapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/dataexchangecacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/dataexchangecacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[datapipelineiface.DataPipelineAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "datapipelineapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/datapipelinecacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/datapipelinecacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[datasynciface.DataSyncAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "datasyncapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/datasynccacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/datasynccacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[daxiface.DAXAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "daxapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/daxcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/daxcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[detectiveiface.DetectiveAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "detectiveapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/detectivecacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/detectivecacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[devicefarmiface.DeviceFarmAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "devicefarmapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/devicefarmcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/devicefarmcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[devopsguruiface.DevOpsGuruAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "devopsguruapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/devopsgurucacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/devopsgurucacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[directconnectiface.DirectConnectAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "directconnectapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/directconnectcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/directconnectcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[directoryserviceiface.DirectoryServiceAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "directoryserviceapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/directoryservicecacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/directoryservicecacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[dlmiface.DLMAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "dlmapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/dlmcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/dlmcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[docdbiface.DocDBAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "docdbapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/docdbcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/docdbcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[docdbelasticiface.DocDBElasticAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "docdbelasticapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/docdbelasticcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/docdbelasticcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[drsiface.DrsAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "drsapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/drscacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/drscacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[dynamodbiface.DynamoDBAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "dynamodbapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/dynamodbcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/dynamodbcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[dynamodbstreamsiface.DynamoDBStreamsAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "dynamodbstreamsapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/dynamodbstreamscacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/dynamodbstreamscacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[ebsiface.EBSAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "ebsapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/ebscacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/ebscacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[ec2iface.EC2API]()
	if err != nil {
		log.Printf("%s: %v\n", "ec2api.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/ec2cacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/ec2cacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[ec2instanceconnectiface.EC2InstanceConnectAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "ec2instanceconnectapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/ec2instanceconnectcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/ec2instanceconnectcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[ecriface.ECRAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "ecrapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/ecrcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/ecrcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[ecrpubliciface.ECRPublicAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "ecrpublicapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/ecrpubliccacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/ecrpubliccacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[ecsiface.ECSAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "ecsapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/ecscacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/ecscacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[efsiface.EFSAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "efsapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/efscacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/efscacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[eksiface.EKSAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "eksapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/ekscacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/ekscacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[elasticacheiface.ElastiCacheAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "elasticacheapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/elasticachecacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/elasticachecacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[elasticbeanstalkiface.ElasticBeanstalkAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "elasticbeanstalkapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/elasticbeanstalkcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/elasticbeanstalkcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[elasticinferenceiface.ElasticInferenceAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "elasticinferenceapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/elasticinferencecacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/elasticinferencecacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[elasticsearchserviceiface.ElasticsearchServiceAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "elasticsearchserviceapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/elasticsearchservicecacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/elasticsearchservicecacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[elastictranscoderiface.ElasticTranscoderAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "elastictranscoderapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/elastictranscodercacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/elastictranscodercacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[elbiface.ELBAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "elbapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/elbcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/elbcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[elbv2iface.ELBV2API]()
	if err != nil {
		log.Printf("%s: %v\n", "elbv2api.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/elbv2cacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/elbv2cacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[emriface.EMRAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "emrapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/emrcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/emrcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[emrcontainersiface.EMRContainersAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "emrcontainersapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/emrcontainerscacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/emrcontainerscacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[emrserverlessiface.EMRServerlessAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "emrserverlessapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/emrserverlesscacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/emrserverlesscacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[eventbridgeiface.EventBridgeAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "eventbridgeapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/eventbridgecacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/eventbridgecacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[finspaceiface.FinspaceAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "finspaceapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/finspacecacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/finspacecacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[finspacedataiface.FinSpaceDataAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "finspacedataapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/finspacedatacacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/finspacedatacacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[firehoseiface.FirehoseAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "firehoseapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/firehosecacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/firehosecacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[fisiface.FISAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "fisapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/fiscacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/fiscacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[fmsiface.FMSAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "fmsapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/fmscacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/fmscacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[forecastqueryserviceiface.ForecastQueryServiceAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "forecastqueryserviceapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/forecastqueryservicecacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/forecastqueryservicecacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[forecastserviceiface.ForecastServiceAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "forecastserviceapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/forecastservicecacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/forecastservicecacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[frauddetectoriface.FraudDetectorAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "frauddetectorapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/frauddetectorcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/frauddetectorcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[fsxiface.FSxAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "fsxapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/fsxcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/fsxcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[gameliftiface.GameLiftAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "gameliftapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/gameliftcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/gameliftcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[gamesparksiface.GameSparksAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "gamesparksapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/gamesparkscacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/gamesparkscacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[glacieriface.GlacierAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "glacierapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/glaciercacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/glaciercacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[globalacceleratoriface.GlobalAcceleratorAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "globalacceleratorapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/globalacceleratorcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/globalacceleratorcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[glueiface.GlueAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "glueapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/gluecacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/gluecacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[gluedatabrewiface.GlueDataBrewAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "gluedatabrewapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/gluedatabrewcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/gluedatabrewcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[greengrassiface.GreengrassAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "greengrassapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/greengrasscacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/greengrasscacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[greengrassv2iface.GreengrassV2API]()
	if err != nil {
		log.Printf("%s: %v\n", "greengrassv2api.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/greengrassv2cacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/greengrassv2cacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[groundstationiface.GroundStationAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "groundstationapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/groundstationcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/groundstationcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[guarddutyiface.GuardDutyAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "guarddutyapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/guarddutycacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/guarddutycacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[healthiface.HealthAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "healthapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/healthcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/healthcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[healthlakeiface.HealthLakeAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "healthlakeapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/healthlakecacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/healthlakecacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[honeycodeiface.HoneycodeAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "honeycodeapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/honeycodecacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/honeycodecacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[iamiface.IAMAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "iamapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/iamcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/iamcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[identitystoreiface.IdentityStoreAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "identitystoreapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/identitystorecacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/identitystorecacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[imagebuilderiface.ImagebuilderAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "imagebuilderapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/imagebuildercacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/imagebuildercacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[inspectoriface.InspectorAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "inspectorapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/inspectorcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/inspectorcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[inspector2iface.Inspector2API]()
	if err != nil {
		log.Printf("%s: %v\n", "inspector2api.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/inspector2cacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/inspector2cacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[iotiface.IoTAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "iotapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/iotcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/iotcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[iot1clickdevicesserviceiface.IoT1ClickDevicesServiceAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "iot1clickdevicesserviceapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/iot1clickdevicesservicecacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/iot1clickdevicesservicecacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[iot1clickprojectsiface.IoT1ClickProjectsAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "iot1clickprojectsapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/iot1clickprojectscacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/iot1clickprojectscacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[iotanalyticsiface.IoTAnalyticsAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "iotanalyticsapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/iotanalyticscacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/iotanalyticscacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[iotdataplaneiface.IoTDataPlaneAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "iotdataplaneapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/iotdataplanecacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/iotdataplanecacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[iotdeviceadvisoriface.IoTDeviceAdvisorAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "iotdeviceadvisorapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/iotdeviceadvisorcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/iotdeviceadvisorcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[ioteventsiface.IoTEventsAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "ioteventsapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/ioteventscacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/ioteventscacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[ioteventsdataiface.IoTEventsDataAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "ioteventsdataapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/ioteventsdatacacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/ioteventsdatacacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[iotfleethubiface.IoTFleetHubAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "iotfleethubapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/iotfleethubcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/iotfleethubcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[iotfleetwiseiface.IoTFleetWiseAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "iotfleetwiseapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/iotfleetwisecacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/iotfleetwisecacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[iotjobsdataplaneiface.IoTJobsDataPlaneAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "iotjobsdataplaneapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/iotjobsdataplanecacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/iotjobsdataplanecacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[iotroborunneriface.IoTRoboRunnerAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "iotroborunnerapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/iotroborunnercacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/iotroborunnercacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[iotsecuretunnelingiface.IoTSecureTunnelingAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "iotsecuretunnelingapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/iotsecuretunnelingcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/iotsecuretunnelingcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[iotsitewiseiface.IoTSiteWiseAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "iotsitewiseapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/iotsitewisecacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/iotsitewisecacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[iotthingsgraphiface.IoTThingsGraphAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "iotthingsgraphapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/iotthingsgraphcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/iotthingsgraphcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[iottwinmakeriface.IoTTwinMakerAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "iottwinmakerapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/iottwinmakercacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/iottwinmakercacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[iotwirelessiface.IoTWirelessAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "iotwirelessapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/iotwirelesscacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/iotwirelesscacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[ivsiface.IVSAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "ivsapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/ivscacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/ivscacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[ivschatiface.IvschatAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "ivschatapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/ivschatcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/ivschatcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[kafkaiface.KafkaAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "kafkaapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/kafkacacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/kafkacacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[kafkaconnectiface.KafkaConnectAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "kafkaconnectapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/kafkaconnectcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/kafkaconnectcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[kendraiface.KendraAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "kendraapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/kendracacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/kendracacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[kendrarankingiface.KendraRankingAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "kendrarankingapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/kendrarankingcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/kendrarankingcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[keyspacesiface.KeyspacesAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "keyspacesapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/keyspacescacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/keyspacescacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[kinesisiface.KinesisAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "kinesisapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/kinesiscacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/kinesiscacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[kinesisanalyticsiface.KinesisAnalyticsAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "kinesisanalyticsapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/kinesisanalyticscacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/kinesisanalyticscacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[kinesisanalyticsv2iface.KinesisAnalyticsV2API]()
	if err != nil {
		log.Printf("%s: %v\n", "kinesisanalyticsv2api.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/kinesisanalyticsv2cacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/kinesisanalyticsv2cacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[kinesisvideoiface.KinesisVideoAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "kinesisvideoapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/kinesisvideocacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/kinesisvideocacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[kinesisvideoarchivedmediaiface.KinesisVideoArchivedMediaAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "kinesisvideoarchivedmediaapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/kinesisvideoarchivedmediacacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/kinesisvideoarchivedmediacacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[kinesisvideomediaiface.KinesisVideoMediaAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "kinesisvideomediaapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/kinesisvideomediacacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/kinesisvideomediacacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[kinesisvideosignalingchannelsiface.KinesisVideoSignalingChannelsAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "kinesisvideosignalingchannelsapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/kinesisvideosignalingchannelscacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/kinesisvideosignalingchannelscacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[kinesisvideowebrtcstorageiface.KinesisVideoWebRTCStorageAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "kinesisvideowebrtcstorageapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/kinesisvideowebrtcstoragecacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/kinesisvideowebrtcstoragecacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[kmsiface.KMSAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "kmsapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/kmscacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/kmscacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[lakeformationiface.LakeFormationAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "lakeformationapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/lakeformationcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/lakeformationcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[lambdaiface.LambdaAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "lambdaapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/lambdacacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/lambdacacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[lexmodelbuildingserviceiface.LexModelBuildingServiceAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "lexmodelbuildingserviceapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/lexmodelbuildingservicecacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/lexmodelbuildingservicecacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[lexmodelsv2iface.LexModelsV2API]()
	if err != nil {
		log.Printf("%s: %v\n", "lexmodelsv2api.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/lexmodelsv2cacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/lexmodelsv2cacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[lexruntimeserviceiface.LexRuntimeServiceAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "lexruntimeserviceapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/lexruntimeservicecacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/lexruntimeservicecacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[lexruntimev2iface.LexRuntimeV2API]()
	if err != nil {
		log.Printf("%s: %v\n", "lexruntimev2api.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/lexruntimev2cacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/lexruntimev2cacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[licensemanageriface.LicenseManagerAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "licensemanagerapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/licensemanagercacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/licensemanagercacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[licensemanagerlinuxsubscriptionsiface.LicenseManagerLinuxSubscriptionsAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "licensemanagerlinuxsubscriptionsapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/licensemanagerlinuxsubscriptionscacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/licensemanagerlinuxsubscriptionscacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[licensemanagerusersubscriptionsiface.LicenseManagerUserSubscriptionsAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "licensemanagerusersubscriptionsapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/licensemanagerusersubscriptionscacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/licensemanagerusersubscriptionscacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[lightsailiface.LightsailAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "lightsailapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/lightsailcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/lightsailcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[locationserviceiface.LocationServiceAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "locationserviceapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/locationservicecacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/locationservicecacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[lookoutequipmentiface.LookoutEquipmentAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "lookoutequipmentapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/lookoutequipmentcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/lookoutequipmentcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[lookoutforvisioniface.LookoutForVisionAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "lookoutforvisionapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/lookoutforvisioncacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/lookoutforvisioncacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[lookoutmetricsiface.LookoutMetricsAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "lookoutmetricsapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/lookoutmetricscacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/lookoutmetricscacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[m2iface.M2API]()
	if err != nil {
		log.Printf("%s: %v\n", "m2api.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/m2cacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/m2cacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[machinelearningiface.MachineLearningAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "machinelearningapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/machinelearningcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/machinelearningcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[macieiface.MacieAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "macieapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/maciecacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/maciecacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[macie2iface.Macie2API]()
	if err != nil {
		log.Printf("%s: %v\n", "macie2api.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/macie2cacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/macie2cacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[managedblockchainiface.ManagedBlockchainAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "managedblockchainapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/managedblockchaincacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/managedblockchaincacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[managedgrafanaiface.ManagedGrafanaAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "managedgrafanaapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/managedgrafanacacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/managedgrafanacacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[marketplacecatalogiface.MarketplaceCatalogAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "marketplacecatalogapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/marketplacecatalogcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/marketplacecatalogcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[marketplacecommerceanalyticsiface.MarketplaceCommerceAnalyticsAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "marketplacecommerceanalyticsapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/marketplacecommerceanalyticscacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/marketplacecommerceanalyticscacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[marketplaceentitlementserviceiface.MarketplaceEntitlementServiceAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "marketplaceentitlementserviceapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/marketplaceentitlementservicecacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/marketplaceentitlementservicecacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[marketplacemeteringiface.MarketplaceMeteringAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "marketplacemeteringapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/marketplacemeteringcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/marketplacemeteringcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[mediaconnectiface.MediaConnectAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "mediaconnectapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/mediaconnectcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/mediaconnectcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[mediaconvertiface.MediaConvertAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "mediaconvertapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/mediaconvertcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/mediaconvertcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[medialiveiface.MediaLiveAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "medialiveapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/medialivecacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/medialivecacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[mediapackageiface.MediaPackageAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "mediapackageapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/mediapackagecacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/mediapackagecacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[mediapackagevodiface.MediaPackageVodAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "mediapackagevodapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/mediapackagevodcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/mediapackagevodcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[mediastoreiface.MediaStoreAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "mediastoreapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/mediastorecacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/mediastorecacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[mediastoredataiface.MediaStoreDataAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "mediastoredataapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/mediastoredatacacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/mediastoredatacacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[mediatailoriface.MediaTailorAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "mediatailorapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/mediatailorcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/mediatailorcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[memorydbiface.MemoryDBAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "memorydbapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/memorydbcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/memorydbcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[mgniface.MgnAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "mgnapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/mgncacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/mgncacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[migrationhubiface.MigrationHubAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "migrationhubapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/migrationhubcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/migrationhubcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[migrationhubconfigiface.MigrationHubConfigAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "migrationhubconfigapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/migrationhubconfigcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/migrationhubconfigcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[migrationhuborchestratoriface.MigrationHubOrchestratorAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "migrationhuborchestratorapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/migrationhuborchestratorcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/migrationhuborchestratorcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[migrationhubrefactorspacesiface.MigrationHubRefactorSpacesAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "migrationhubrefactorspacesapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/migrationhubrefactorspacescacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/migrationhubrefactorspacescacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[migrationhubstrategyrecommendationsiface.MigrationHubStrategyRecommendationsAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "migrationhubstrategyrecommendationsapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/migrationhubstrategyrecommendationscacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/migrationhubstrategyrecommendationscacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[mobileiface.MobileAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "mobileapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/mobilecacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/mobilecacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[mobileanalyticsiface.MobileAnalyticsAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "mobileanalyticsapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/mobileanalyticscacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/mobileanalyticscacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[mqiface.MQAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "mqapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/mqcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/mqcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[mturkiface.MTurkAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "mturkapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/mturkcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/mturkcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[mwaaiface.MWAAAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "mwaaapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/mwaacacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/mwaacacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[neptuneiface.NeptuneAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "neptuneapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/neptunecacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/neptunecacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[networkfirewalliface.NetworkFirewallAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "networkfirewallapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/networkfirewallcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/networkfirewallcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[networkmanageriface.NetworkManagerAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "networkmanagerapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/networkmanagercacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/networkmanagercacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[nimblestudioiface.NimbleStudioAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "nimblestudioapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/nimblestudiocacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/nimblestudiocacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[oamiface.OAMAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "oamapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/oamcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/oamcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[omicsiface.OmicsAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "omicsapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/omicscacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/omicscacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[opensearchserverlessiface.OpenSearchServerlessAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "opensearchserverlessapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/opensearchserverlesscacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/opensearchserverlesscacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[opensearchserviceiface.OpenSearchServiceAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "opensearchserviceapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/opensearchservicecacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/opensearchservicecacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[opsworksiface.OpsWorksAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "opsworksapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/opsworkscacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/opsworkscacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[opsworkscmiface.OpsWorksCMAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "opsworkscmapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/opsworkscmcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/opsworkscmcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[organizationsiface.OrganizationsAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "organizationsapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/organizationscacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/organizationscacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[outpostsiface.OutpostsAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "outpostsapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/outpostscacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/outpostscacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[panoramaiface.PanoramaAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "panoramaapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/panoramacacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/panoramacacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[personalizeiface.PersonalizeAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "personalizeapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/personalizecacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/personalizecacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[personalizeeventsiface.PersonalizeEventsAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "personalizeeventsapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/personalizeeventscacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/personalizeeventscacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[personalizeruntimeiface.PersonalizeRuntimeAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "personalizeruntimeapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/personalizeruntimecacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/personalizeruntimecacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[piiface.PIAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "piapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/picacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/picacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[pinpointiface.PinpointAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "pinpointapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/pinpointcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/pinpointcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[pinpointemailiface.PinpointEmailAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "pinpointemailapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/pinpointemailcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/pinpointemailcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[pinpointsmsvoiceiface.PinpointSMSVoiceAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "pinpointsmsvoiceapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/pinpointsmsvoicecacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/pinpointsmsvoicecacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[pinpointsmsvoicev2iface.PinpointSMSVoiceV2API]()
	if err != nil {
		log.Printf("%s: %v\n", "pinpointsmsvoicev2api.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/pinpointsmsvoicev2cacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/pinpointsmsvoicev2cacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[pipesiface.PipesAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "pipesapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/pipescacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/pipescacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[pollyiface.PollyAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "pollyapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/pollycacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/pollycacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[pricingiface.PricingAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "pricingapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/pricingcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/pricingcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[privatenetworksiface.PrivateNetworksAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "privatenetworksapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/privatenetworkscacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/privatenetworkscacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[prometheusserviceiface.PrometheusServiceAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "prometheusserviceapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/prometheusservicecacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/prometheusservicecacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[protoniface.ProtonAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "protonapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/protoncacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/protoncacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[qldbiface.QLDBAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "qldbapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/qldbcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/qldbcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[qldbsessioniface.QLDBSessionAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "qldbsessionapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/qldbsessioncacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/qldbsessioncacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[quicksightiface.QuickSightAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "quicksightapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/quicksightcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/quicksightcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[ramiface.RAMAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "ramapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/ramcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/ramcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[rdsiface.RDSAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "rdsapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/rdscacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/rdscacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[rdsdataserviceiface.RDSDataServiceAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "rdsdataserviceapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/rdsdataservicecacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/rdsdataservicecacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[recyclebiniface.RecycleBinAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "recyclebinapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/recyclebincacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/recyclebincacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[redshiftiface.RedshiftAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "redshiftapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/redshiftcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/redshiftcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[redshiftdataapiserviceiface.RedshiftDataAPIServiceAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "redshiftdataapiserviceapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/redshiftdataapiservicecacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/redshiftdataapiservicecacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[redshiftserverlessiface.RedshiftServerlessAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "redshiftserverlessapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/redshiftserverlesscacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/redshiftserverlesscacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[rekognitioniface.RekognitionAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "rekognitionapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/rekognitioncacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/rekognitioncacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[resiliencehubiface.ResilienceHubAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "resiliencehubapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/resiliencehubcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/resiliencehubcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[resourceexplorer2iface.ResourceExplorer2API]()
	if err != nil {
		log.Printf("%s: %v\n", "resourceexplorer2api.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/resourceexplorer2cacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/resourceexplorer2cacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[resourcegroupsiface.ResourceGroupsAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "resourcegroupsapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/resourcegroupscacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/resourcegroupscacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[resourcegroupstaggingapiiface.ResourceGroupsTaggingAPIAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "resourcegroupstaggingapiapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/resourcegroupstaggingapicacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/resourcegroupstaggingapicacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[robomakeriface.RoboMakerAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "robomakerapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/robomakercacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/robomakercacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[rolesanywhereiface.RolesAnywhereAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "rolesanywhereapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/rolesanywherecacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/rolesanywherecacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[route53iface.Route53API]()
	if err != nil {
		log.Printf("%s: %v\n", "route53api.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/route53cacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/route53cacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[route53domainsiface.Route53DomainsAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "route53domainsapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/route53domainscacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/route53domainscacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[route53recoveryclusteriface.Route53RecoveryClusterAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "route53recoveryclusterapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/route53recoveryclustercacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/route53recoveryclustercacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[route53recoverycontrolconfigiface.Route53RecoveryControlConfigAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "route53recoverycontrolconfigapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/route53recoverycontrolconfigcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/route53recoverycontrolconfigcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[route53recoveryreadinessiface.Route53RecoveryReadinessAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "route53recoveryreadinessapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/route53recoveryreadinesscacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/route53recoveryreadinesscacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[route53resolveriface.Route53ResolverAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "route53resolverapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/route53resolvercacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/route53resolvercacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[s3iface.S3API]()
	if err != nil {
		log.Printf("%s: %v\n", "s3api.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/s3cacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/s3cacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[s3controliface.S3ControlAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "s3controlapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/s3controlcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/s3controlcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[s3outpostsiface.S3OutpostsAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "s3outpostsapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/s3outpostscacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/s3outpostscacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[sagemakeriface.SageMakerAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "sagemakerapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/sagemakercacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/sagemakercacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[sagemakeredgemanageriface.SagemakerEdgeManagerAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "sagemakeredgemanagerapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/sagemakeredgemanagercacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/sagemakeredgemanagercacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[sagemakerfeaturestoreruntimeiface.SageMakerFeatureStoreRuntimeAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "sagemakerfeaturestoreruntimeapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/sagemakerfeaturestoreruntimecacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/sagemakerfeaturestoreruntimecacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[sagemakergeospatialiface.SageMakerGeospatialAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "sagemakergeospatialapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/sagemakergeospatialcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/sagemakergeospatialcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[sagemakermetricsiface.SageMakerMetricsAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "sagemakermetricsapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/sagemakermetricscacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/sagemakermetricscacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[sagemakerruntimeiface.SageMakerRuntimeAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "sagemakerruntimeapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/sagemakerruntimecacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/sagemakerruntimecacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[savingsplansiface.SavingsPlansAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "savingsplansapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/savingsplanscacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/savingsplanscacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[scheduleriface.SchedulerAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "schedulerapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/schedulercacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/schedulercacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[schemasiface.SchemasAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "schemasapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/schemascacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/schemascacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[secretsmanageriface.SecretsManagerAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "secretsmanagerapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/secretsmanagercacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/secretsmanagercacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[securityhubiface.SecurityHubAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "securityhubapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/securityhubcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/securityhubcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[securitylakeiface.SecurityLakeAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "securitylakeapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/securitylakecacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/securitylakecacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[serverlessapplicationrepositoryiface.ServerlessApplicationRepositoryAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "serverlessapplicationrepositoryapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/serverlessapplicationrepositorycacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/serverlessapplicationrepositorycacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[servicecatalogiface.ServiceCatalogAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "servicecatalogapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/servicecatalogcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/servicecatalogcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[servicediscoveryiface.ServiceDiscoveryAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "servicediscoveryapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/servicediscoverycacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/servicediscoverycacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[servicequotasiface.ServiceQuotasAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "servicequotasapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/servicequotascacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/servicequotascacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[sesiface.SESAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "sesapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/sescacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/sescacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[sesv2iface.SESV2API]()
	if err != nil {
		log.Printf("%s: %v\n", "sesv2api.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/sesv2cacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/sesv2cacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[sfniface.SFNAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "sfnapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/sfncacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/sfncacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[shieldiface.ShieldAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "shieldapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/shieldcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/shieldcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[signeriface.SignerAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "signerapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/signercacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/signercacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[simpledbiface.SimpleDBAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "simpledbapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/simpledbcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/simpledbcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[simspaceweaveriface.SimSpaceWeaverAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "simspaceweaverapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/simspaceweavercacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/simspaceweavercacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[smsiface.SMSAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "smsapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/smscacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/smscacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[snowballiface.SnowballAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "snowballapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/snowballcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/snowballcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[snowdevicemanagementiface.SnowDeviceManagementAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "snowdevicemanagementapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/snowdevicemanagementcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/snowdevicemanagementcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[snsiface.SNSAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "snsapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/snscacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/snscacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[sqsiface.SQSAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "sqsapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/sqscacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/sqscacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[ssmiface.SSMAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "ssmapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/ssmcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/ssmcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[ssmcontactsiface.SSMContactsAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "ssmcontactsapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/ssmcontactscacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/ssmcontactscacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[ssmincidentsiface.SSMIncidentsAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "ssmincidentsapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/ssmincidentscacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/ssmincidentscacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[ssmsapiface.SsmSapAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "ssmsapapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/ssmsapcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/ssmsapcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[ssoiface.SSOAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "ssoapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/ssocacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/ssocacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[ssoadminiface.SSOAdminAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "ssoadminapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/ssoadmincacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/ssoadmincacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[ssooidciface.SSOOIDCAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "ssooidcapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/ssooidccacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/ssooidccacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[storagegatewayiface.StorageGatewayAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "storagegatewayapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/storagegatewaycacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/storagegatewaycacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[stsiface.STSAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "stsapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/stscacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/stscacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[supportiface.SupportAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "supportapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/supportcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/supportcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[supportappiface.SupportAppAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "supportappapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/supportappcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/supportappcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[swfiface.SWFAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "swfapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/swfcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/swfcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[syntheticsiface.SyntheticsAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "syntheticsapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/syntheticscacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/syntheticscacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[textractiface.TextractAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "textractapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/textractcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/textractcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[timestreamqueryiface.TimestreamQueryAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "timestreamqueryapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/timestreamquerycacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/timestreamquerycacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[timestreamwriteiface.TimestreamWriteAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "timestreamwriteapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/timestreamwritecacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/timestreamwritecacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[transcribeserviceiface.TranscribeServiceAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "transcribeserviceapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/transcribeservicecacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/transcribeservicecacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[transcribestreamingserviceiface.TranscribeStreamingServiceAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "transcribestreamingserviceapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/transcribestreamingservicecacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/transcribestreamingservicecacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[transferiface.TransferAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "transferapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/transfercacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/transfercacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[translateiface.TranslateAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "translateapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/translatecacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/translatecacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[voiceidiface.VoiceIDAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "voiceidapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/voiceidcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/voiceidcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[wafiface.WAFAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "wafapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/wafcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/wafcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[wafregionaliface.WAFRegionalAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "wafregionalapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/wafregionalcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/wafregionalcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[wafv2iface.WAFV2API]()
	if err != nil {
		log.Printf("%s: %v\n", "wafv2api.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/wafv2cacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/wafv2cacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[wellarchitectediface.WellArchitectedAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "wellarchitectedapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/wellarchitectedcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/wellarchitectedcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[workdocsiface.WorkDocsAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "workdocsapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/workdocscacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/workdocscacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[worklinkiface.WorkLinkAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "worklinkapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/worklinkcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/worklinkcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[workmailiface.WorkMailAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "workmailapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/workmailcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/workmailcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[workmailmessageflowiface.WorkMailMessageFlowAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "workmailmessageflowapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/workmailmessageflowcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/workmailmessageflowcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[workspacesiface.WorkSpacesAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "workspacesapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/workspacescacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/workspacescacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[workspaceswebiface.WorkSpacesWebAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "workspaceswebapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/workspaceswebcacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/workspaceswebcacher.go", serviceOutDir), []byte(out), 0644)
	}
	out, err = GenSDK[xrayiface.XRayAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "xrayapi.go", err)
	}
	serviceOutDir = fmt.Sprintf("%s/xraycacher", *outputDir)
	if out != "" {
		if _, err := os.Stat(serviceOutDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(serviceOutDir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
		os.WriteFile(fmt.Sprintf("%s/xraycacher.go", serviceOutDir), []byte(out), 0644)
	}
}
