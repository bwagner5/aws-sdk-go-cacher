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

	outputDir := flag.String("out-dir", "pkg/cacher", "directory to output generated caching clients")
	flag.Parse()

	var out string
	var err error
	out, err = GenSDK[accessanalyzeriface.AccessAnalyzerAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "accessanalyzerapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/accessanalyzerapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[accountiface.AccountAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "accountapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/accountapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[acmiface.ACMAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "acmapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/acmapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[acmpcaiface.ACMPCAAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "acmpcaapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/acmpcaapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[alexaforbusinessiface.AlexaForBusinessAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "alexaforbusinessapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/alexaforbusinessapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[amplifyiface.AmplifyAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "amplifyapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/amplifyapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[amplifybackendiface.AmplifyBackendAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "amplifybackendapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/amplifybackendapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[amplifyuibuilderiface.AmplifyUIBuilderAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "amplifyuibuilderapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/amplifyuibuilderapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[apigatewayiface.APIGatewayAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "apigatewayapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/apigatewayapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[apigatewaymanagementapiiface.ApiGatewayManagementApiAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "apigatewaymanagementapiapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/apigatewaymanagementapiapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[apigatewayv2iface.ApiGatewayV2API]()
	if err != nil {
		log.Printf("%s: %v\n", "apigatewayv2api.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/apigatewayv2api.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[appconfigiface.AppConfigAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "appconfigapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/appconfigapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[appconfigdataiface.AppConfigDataAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "appconfigdataapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/appconfigdataapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[appflowiface.AppflowAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "appflowapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/appflowapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[appintegrationsserviceiface.AppIntegrationsServiceAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "appintegrationsserviceapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/appintegrationsserviceapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[applicationautoscalingiface.ApplicationAutoScalingAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "applicationautoscalingapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/applicationautoscalingapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[applicationcostprofileriface.ApplicationCostProfilerAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "applicationcostprofilerapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/applicationcostprofilerapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[applicationdiscoveryserviceiface.ApplicationDiscoveryServiceAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "applicationdiscoveryserviceapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/applicationdiscoveryserviceapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[applicationinsightsiface.ApplicationInsightsAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "applicationinsightsapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/applicationinsightsapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[appmeshiface.AppMeshAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "appmeshapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/appmeshapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[appregistryiface.AppRegistryAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "appregistryapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/appregistryapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[apprunneriface.AppRunnerAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "apprunnerapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/apprunnerapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[appstreamiface.AppStreamAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "appstreamapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/appstreamapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[appsynciface.AppSyncAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "appsyncapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/appsyncapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[arczonalshiftiface.ARCZonalShiftAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "arczonalshiftapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/arczonalshiftapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[athenaiface.AthenaAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "athenaapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/athenaapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[auditmanageriface.AuditManagerAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "auditmanagerapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/auditmanagerapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[augmentedairuntimeiface.AugmentedAIRuntimeAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "augmentedairuntimeapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/augmentedairuntimeapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[autoscalingiface.AutoScalingAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "autoscalingapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/autoscalingapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[autoscalingplansiface.AutoScalingPlansAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "autoscalingplansapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/autoscalingplansapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[backupiface.BackupAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "backupapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/backupapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[backupgatewayiface.BackupGatewayAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "backupgatewayapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/backupgatewayapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[backupstorageiface.BackupStorageAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "backupstorageapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/backupstorageapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[batchiface.BatchAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "batchapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/batchapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[billingconductoriface.BillingConductorAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "billingconductorapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/billingconductorapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[braketiface.BraketAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "braketapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/braketapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[budgetsiface.BudgetsAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "budgetsapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/budgetsapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[chimeiface.ChimeAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "chimeapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/chimeapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[chimesdkidentityiface.ChimeSDKIdentityAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "chimesdkidentityapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/chimesdkidentityapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[chimesdkmediapipelinesiface.ChimeSDKMediaPipelinesAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "chimesdkmediapipelinesapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/chimesdkmediapipelinesapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[chimesdkmeetingsiface.ChimeSDKMeetingsAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "chimesdkmeetingsapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/chimesdkmeetingsapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[chimesdkmessagingiface.ChimeSDKMessagingAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "chimesdkmessagingapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/chimesdkmessagingapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[chimesdkvoiceiface.ChimeSDKVoiceAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "chimesdkvoiceapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/chimesdkvoiceapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[cleanroomsiface.CleanRoomsAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "cleanroomsapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/cleanroomsapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[cloud9iface.Cloud9API]()
	if err != nil {
		log.Printf("%s: %v\n", "cloud9api.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/cloud9api.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[cloudcontrolapiiface.CloudControlApiAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "cloudcontrolapiapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/cloudcontrolapiapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[clouddirectoryiface.CloudDirectoryAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "clouddirectoryapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/clouddirectoryapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[cloudformationiface.CloudFormationAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "cloudformationapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/cloudformationapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[cloudfrontiface.CloudFrontAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "cloudfrontapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/cloudfrontapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[cloudhsmiface.CloudHSMAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "cloudhsmapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/cloudhsmapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[cloudhsmv2iface.CloudHSMV2API]()
	if err != nil {
		log.Printf("%s: %v\n", "cloudhsmv2api.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/cloudhsmv2api.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[cloudsearchiface.CloudSearchAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "cloudsearchapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/cloudsearchapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[cloudsearchdomainiface.CloudSearchDomainAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "cloudsearchdomainapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/cloudsearchdomainapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[cloudtrailiface.CloudTrailAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "cloudtrailapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/cloudtrailapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[cloudwatchiface.CloudWatchAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "cloudwatchapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/cloudwatchapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[cloudwatcheventsiface.CloudWatchEventsAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "cloudwatcheventsapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/cloudwatcheventsapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[cloudwatchevidentlyiface.CloudWatchEvidentlyAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "cloudwatchevidentlyapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/cloudwatchevidentlyapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[cloudwatchlogsiface.CloudWatchLogsAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "cloudwatchlogsapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/cloudwatchlogsapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[cloudwatchrumiface.CloudWatchRUMAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "cloudwatchrumapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/cloudwatchrumapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[codeartifactiface.CodeArtifactAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "codeartifactapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/codeartifactapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[codebuildiface.CodeBuildAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "codebuildapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/codebuildapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[codecommitiface.CodeCommitAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "codecommitapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/codecommitapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[codedeployiface.CodeDeployAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "codedeployapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/codedeployapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[codeguruprofileriface.CodeGuruProfilerAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "codeguruprofilerapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/codeguruprofilerapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[codegururevieweriface.CodeGuruReviewerAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "codegurureviewerapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/codegurureviewerapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[codepipelineiface.CodePipelineAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "codepipelineapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/codepipelineapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[codestariface.CodeStarAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "codestarapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/codestarapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[codestarconnectionsiface.CodeStarConnectionsAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "codestarconnectionsapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/codestarconnectionsapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[codestarnotificationsiface.CodeStarNotificationsAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "codestarnotificationsapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/codestarnotificationsapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[cognitoidentityiface.CognitoIdentityAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "cognitoidentityapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/cognitoidentityapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[cognitoidentityprovideriface.CognitoIdentityProviderAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "cognitoidentityproviderapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/cognitoidentityproviderapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[cognitosynciface.CognitoSyncAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "cognitosyncapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/cognitosyncapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[comprehendiface.ComprehendAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "comprehendapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/comprehendapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[comprehendmedicaliface.ComprehendMedicalAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "comprehendmedicalapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/comprehendmedicalapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[computeoptimizeriface.ComputeOptimizerAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "computeoptimizerapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/computeoptimizerapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[configserviceiface.ConfigServiceAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "configserviceapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/configserviceapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[connectiface.ConnectAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "connectapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/connectapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[connectcampaignsiface.ConnectCampaignsAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "connectcampaignsapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/connectcampaignsapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[connectcasesiface.ConnectCasesAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "connectcasesapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/connectcasesapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[connectcontactlensiface.ConnectContactLensAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "connectcontactlensapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/connectcontactlensapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[connectparticipantiface.ConnectParticipantAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "connectparticipantapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/connectparticipantapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[connectwisdomserviceiface.ConnectWisdomServiceAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "connectwisdomserviceapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/connectwisdomserviceapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[controltoweriface.ControlTowerAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "controltowerapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/controltowerapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[costandusagereportserviceiface.CostandUsageReportServiceAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "costandusagereportserviceapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/costandusagereportserviceapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[costexploreriface.CostExplorerAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "costexplorerapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/costexplorerapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[customerprofilesiface.CustomerProfilesAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "customerprofilesapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/customerprofilesapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[databasemigrationserviceiface.DatabaseMigrationServiceAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "databasemigrationserviceapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/databasemigrationserviceapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[dataexchangeiface.DataExchangeAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "dataexchangeapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/dataexchangeapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[datapipelineiface.DataPipelineAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "datapipelineapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/datapipelineapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[datasynciface.DataSyncAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "datasyncapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/datasyncapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[daxiface.DAXAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "daxapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/daxapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[detectiveiface.DetectiveAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "detectiveapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/detectiveapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[devicefarmiface.DeviceFarmAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "devicefarmapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/devicefarmapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[devopsguruiface.DevOpsGuruAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "devopsguruapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/devopsguruapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[directconnectiface.DirectConnectAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "directconnectapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/directconnectapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[directoryserviceiface.DirectoryServiceAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "directoryserviceapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/directoryserviceapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[dlmiface.DLMAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "dlmapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/dlmapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[docdbiface.DocDBAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "docdbapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/docdbapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[docdbelasticiface.DocDBElasticAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "docdbelasticapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/docdbelasticapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[drsiface.DrsAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "drsapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/drsapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[dynamodbiface.DynamoDBAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "dynamodbapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/dynamodbapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[dynamodbstreamsiface.DynamoDBStreamsAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "dynamodbstreamsapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/dynamodbstreamsapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[ebsiface.EBSAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "ebsapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/ebsapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[ec2iface.EC2API]()
	if err != nil {
		log.Printf("%s: %v\n", "ec2api.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/ec2api.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[ec2instanceconnectiface.EC2InstanceConnectAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "ec2instanceconnectapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/ec2instanceconnectapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[ecriface.ECRAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "ecrapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/ecrapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[ecrpubliciface.ECRPublicAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "ecrpublicapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/ecrpublicapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[ecsiface.ECSAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "ecsapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/ecsapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[efsiface.EFSAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "efsapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/efsapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[eksiface.EKSAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "eksapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/eksapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[elasticacheiface.ElastiCacheAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "elasticacheapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/elasticacheapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[elasticbeanstalkiface.ElasticBeanstalkAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "elasticbeanstalkapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/elasticbeanstalkapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[elasticinferenceiface.ElasticInferenceAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "elasticinferenceapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/elasticinferenceapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[elasticsearchserviceiface.ElasticsearchServiceAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "elasticsearchserviceapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/elasticsearchserviceapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[elastictranscoderiface.ElasticTranscoderAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "elastictranscoderapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/elastictranscoderapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[elbiface.ELBAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "elbapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/elbapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[elbv2iface.ELBV2API]()
	if err != nil {
		log.Printf("%s: %v\n", "elbv2api.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/elbv2api.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[emriface.EMRAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "emrapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/emrapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[emrcontainersiface.EMRContainersAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "emrcontainersapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/emrcontainersapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[emrserverlessiface.EMRServerlessAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "emrserverlessapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/emrserverlessapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[eventbridgeiface.EventBridgeAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "eventbridgeapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/eventbridgeapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[finspaceiface.FinspaceAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "finspaceapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/finspaceapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[finspacedataiface.FinSpaceDataAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "finspacedataapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/finspacedataapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[firehoseiface.FirehoseAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "firehoseapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/firehoseapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[fisiface.FISAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "fisapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/fisapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[fmsiface.FMSAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "fmsapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/fmsapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[forecastqueryserviceiface.ForecastQueryServiceAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "forecastqueryserviceapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/forecastqueryserviceapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[forecastserviceiface.ForecastServiceAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "forecastserviceapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/forecastserviceapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[frauddetectoriface.FraudDetectorAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "frauddetectorapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/frauddetectorapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[fsxiface.FSxAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "fsxapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/fsxapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[gameliftiface.GameLiftAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "gameliftapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/gameliftapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[gamesparksiface.GameSparksAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "gamesparksapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/gamesparksapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[glacieriface.GlacierAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "glacierapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/glacierapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[globalacceleratoriface.GlobalAcceleratorAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "globalacceleratorapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/globalacceleratorapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[glueiface.GlueAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "glueapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/glueapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[gluedatabrewiface.GlueDataBrewAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "gluedatabrewapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/gluedatabrewapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[greengrassiface.GreengrassAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "greengrassapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/greengrassapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[greengrassv2iface.GreengrassV2API]()
	if err != nil {
		log.Printf("%s: %v\n", "greengrassv2api.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/greengrassv2api.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[groundstationiface.GroundStationAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "groundstationapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/groundstationapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[guarddutyiface.GuardDutyAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "guarddutyapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/guarddutyapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[healthiface.HealthAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "healthapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/healthapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[healthlakeiface.HealthLakeAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "healthlakeapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/healthlakeapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[honeycodeiface.HoneycodeAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "honeycodeapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/honeycodeapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[iamiface.IAMAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "iamapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/iamapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[identitystoreiface.IdentityStoreAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "identitystoreapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/identitystoreapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[imagebuilderiface.ImagebuilderAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "imagebuilderapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/imagebuilderapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[inspectoriface.InspectorAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "inspectorapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/inspectorapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[inspector2iface.Inspector2API]()
	if err != nil {
		log.Printf("%s: %v\n", "inspector2api.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/inspector2api.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[iotiface.IoTAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "iotapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/iotapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[iot1clickdevicesserviceiface.IoT1ClickDevicesServiceAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "iot1clickdevicesserviceapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/iot1clickdevicesserviceapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[iot1clickprojectsiface.IoT1ClickProjectsAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "iot1clickprojectsapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/iot1clickprojectsapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[iotanalyticsiface.IoTAnalyticsAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "iotanalyticsapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/iotanalyticsapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[iotdataplaneiface.IoTDataPlaneAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "iotdataplaneapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/iotdataplaneapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[iotdeviceadvisoriface.IoTDeviceAdvisorAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "iotdeviceadvisorapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/iotdeviceadvisorapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[ioteventsiface.IoTEventsAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "ioteventsapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/ioteventsapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[ioteventsdataiface.IoTEventsDataAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "ioteventsdataapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/ioteventsdataapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[iotfleethubiface.IoTFleetHubAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "iotfleethubapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/iotfleethubapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[iotfleetwiseiface.IoTFleetWiseAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "iotfleetwiseapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/iotfleetwiseapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[iotjobsdataplaneiface.IoTJobsDataPlaneAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "iotjobsdataplaneapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/iotjobsdataplaneapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[iotroborunneriface.IoTRoboRunnerAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "iotroborunnerapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/iotroborunnerapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[iotsecuretunnelingiface.IoTSecureTunnelingAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "iotsecuretunnelingapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/iotsecuretunnelingapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[iotsitewiseiface.IoTSiteWiseAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "iotsitewiseapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/iotsitewiseapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[iotthingsgraphiface.IoTThingsGraphAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "iotthingsgraphapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/iotthingsgraphapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[iottwinmakeriface.IoTTwinMakerAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "iottwinmakerapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/iottwinmakerapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[iotwirelessiface.IoTWirelessAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "iotwirelessapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/iotwirelessapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[ivsiface.IVSAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "ivsapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/ivsapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[ivschatiface.IvschatAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "ivschatapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/ivschatapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[kafkaiface.KafkaAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "kafkaapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/kafkaapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[kafkaconnectiface.KafkaConnectAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "kafkaconnectapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/kafkaconnectapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[kendraiface.KendraAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "kendraapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/kendraapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[kendrarankingiface.KendraRankingAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "kendrarankingapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/kendrarankingapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[keyspacesiface.KeyspacesAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "keyspacesapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/keyspacesapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[kinesisiface.KinesisAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "kinesisapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/kinesisapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[kinesisanalyticsiface.KinesisAnalyticsAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "kinesisanalyticsapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/kinesisanalyticsapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[kinesisanalyticsv2iface.KinesisAnalyticsV2API]()
	if err != nil {
		log.Printf("%s: %v\n", "kinesisanalyticsv2api.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/kinesisanalyticsv2api.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[kinesisvideoiface.KinesisVideoAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "kinesisvideoapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/kinesisvideoapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[kinesisvideoarchivedmediaiface.KinesisVideoArchivedMediaAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "kinesisvideoarchivedmediaapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/kinesisvideoarchivedmediaapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[kinesisvideomediaiface.KinesisVideoMediaAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "kinesisvideomediaapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/kinesisvideomediaapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[kinesisvideosignalingchannelsiface.KinesisVideoSignalingChannelsAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "kinesisvideosignalingchannelsapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/kinesisvideosignalingchannelsapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[kinesisvideowebrtcstorageiface.KinesisVideoWebRTCStorageAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "kinesisvideowebrtcstorageapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/kinesisvideowebrtcstorageapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[kmsiface.KMSAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "kmsapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/kmsapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[lakeformationiface.LakeFormationAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "lakeformationapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/lakeformationapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[lambdaiface.LambdaAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "lambdaapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/lambdaapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[lexmodelbuildingserviceiface.LexModelBuildingServiceAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "lexmodelbuildingserviceapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/lexmodelbuildingserviceapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[lexmodelsv2iface.LexModelsV2API]()
	if err != nil {
		log.Printf("%s: %v\n", "lexmodelsv2api.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/lexmodelsv2api.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[lexruntimeserviceiface.LexRuntimeServiceAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "lexruntimeserviceapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/lexruntimeserviceapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[lexruntimev2iface.LexRuntimeV2API]()
	if err != nil {
		log.Printf("%s: %v\n", "lexruntimev2api.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/lexruntimev2api.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[licensemanageriface.LicenseManagerAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "licensemanagerapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/licensemanagerapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[licensemanagerlinuxsubscriptionsiface.LicenseManagerLinuxSubscriptionsAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "licensemanagerlinuxsubscriptionsapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/licensemanagerlinuxsubscriptionsapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[licensemanagerusersubscriptionsiface.LicenseManagerUserSubscriptionsAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "licensemanagerusersubscriptionsapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/licensemanagerusersubscriptionsapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[lightsailiface.LightsailAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "lightsailapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/lightsailapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[locationserviceiface.LocationServiceAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "locationserviceapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/locationserviceapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[lookoutequipmentiface.LookoutEquipmentAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "lookoutequipmentapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/lookoutequipmentapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[lookoutforvisioniface.LookoutForVisionAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "lookoutforvisionapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/lookoutforvisionapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[lookoutmetricsiface.LookoutMetricsAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "lookoutmetricsapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/lookoutmetricsapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[m2iface.M2API]()
	if err != nil {
		log.Printf("%s: %v\n", "m2api.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/m2api.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[machinelearningiface.MachineLearningAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "machinelearningapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/machinelearningapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[macieiface.MacieAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "macieapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/macieapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[macie2iface.Macie2API]()
	if err != nil {
		log.Printf("%s: %v\n", "macie2api.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/macie2api.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[managedblockchainiface.ManagedBlockchainAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "managedblockchainapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/managedblockchainapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[managedgrafanaiface.ManagedGrafanaAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "managedgrafanaapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/managedgrafanaapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[marketplacecatalogiface.MarketplaceCatalogAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "marketplacecatalogapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/marketplacecatalogapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[marketplacecommerceanalyticsiface.MarketplaceCommerceAnalyticsAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "marketplacecommerceanalyticsapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/marketplacecommerceanalyticsapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[marketplaceentitlementserviceiface.MarketplaceEntitlementServiceAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "marketplaceentitlementserviceapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/marketplaceentitlementserviceapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[marketplacemeteringiface.MarketplaceMeteringAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "marketplacemeteringapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/marketplacemeteringapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[mediaconnectiface.MediaConnectAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "mediaconnectapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/mediaconnectapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[mediaconvertiface.MediaConvertAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "mediaconvertapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/mediaconvertapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[medialiveiface.MediaLiveAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "medialiveapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/medialiveapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[mediapackageiface.MediaPackageAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "mediapackageapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/mediapackageapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[mediapackagevodiface.MediaPackageVodAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "mediapackagevodapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/mediapackagevodapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[mediastoreiface.MediaStoreAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "mediastoreapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/mediastoreapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[mediastoredataiface.MediaStoreDataAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "mediastoredataapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/mediastoredataapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[mediatailoriface.MediaTailorAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "mediatailorapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/mediatailorapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[memorydbiface.MemoryDBAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "memorydbapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/memorydbapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[mgniface.MgnAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "mgnapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/mgnapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[migrationhubiface.MigrationHubAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "migrationhubapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/migrationhubapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[migrationhubconfigiface.MigrationHubConfigAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "migrationhubconfigapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/migrationhubconfigapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[migrationhuborchestratoriface.MigrationHubOrchestratorAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "migrationhuborchestratorapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/migrationhuborchestratorapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[migrationhubrefactorspacesiface.MigrationHubRefactorSpacesAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "migrationhubrefactorspacesapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/migrationhubrefactorspacesapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[migrationhubstrategyrecommendationsiface.MigrationHubStrategyRecommendationsAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "migrationhubstrategyrecommendationsapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/migrationhubstrategyrecommendationsapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[mobileiface.MobileAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "mobileapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/mobileapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[mobileanalyticsiface.MobileAnalyticsAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "mobileanalyticsapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/mobileanalyticsapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[mqiface.MQAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "mqapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/mqapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[mturkiface.MTurkAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "mturkapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/mturkapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[mwaaiface.MWAAAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "mwaaapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/mwaaapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[neptuneiface.NeptuneAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "neptuneapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/neptuneapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[networkfirewalliface.NetworkFirewallAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "networkfirewallapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/networkfirewallapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[networkmanageriface.NetworkManagerAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "networkmanagerapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/networkmanagerapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[nimblestudioiface.NimbleStudioAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "nimblestudioapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/nimblestudioapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[oamiface.OAMAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "oamapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/oamapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[omicsiface.OmicsAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "omicsapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/omicsapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[opensearchserverlessiface.OpenSearchServerlessAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "opensearchserverlessapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/opensearchserverlessapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[opensearchserviceiface.OpenSearchServiceAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "opensearchserviceapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/opensearchserviceapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[opsworksiface.OpsWorksAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "opsworksapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/opsworksapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[opsworkscmiface.OpsWorksCMAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "opsworkscmapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/opsworkscmapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[organizationsiface.OrganizationsAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "organizationsapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/organizationsapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[outpostsiface.OutpostsAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "outpostsapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/outpostsapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[panoramaiface.PanoramaAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "panoramaapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/panoramaapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[personalizeiface.PersonalizeAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "personalizeapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/personalizeapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[personalizeeventsiface.PersonalizeEventsAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "personalizeeventsapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/personalizeeventsapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[personalizeruntimeiface.PersonalizeRuntimeAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "personalizeruntimeapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/personalizeruntimeapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[piiface.PIAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "piapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/piapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[pinpointiface.PinpointAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "pinpointapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/pinpointapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[pinpointemailiface.PinpointEmailAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "pinpointemailapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/pinpointemailapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[pinpointsmsvoiceiface.PinpointSMSVoiceAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "pinpointsmsvoiceapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/pinpointsmsvoiceapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[pinpointsmsvoicev2iface.PinpointSMSVoiceV2API]()
	if err != nil {
		log.Printf("%s: %v\n", "pinpointsmsvoicev2api.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/pinpointsmsvoicev2api.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[pipesiface.PipesAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "pipesapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/pipesapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[pollyiface.PollyAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "pollyapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/pollyapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[pricingiface.PricingAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "pricingapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/pricingapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[privatenetworksiface.PrivateNetworksAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "privatenetworksapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/privatenetworksapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[prometheusserviceiface.PrometheusServiceAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "prometheusserviceapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/prometheusserviceapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[protoniface.ProtonAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "protonapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/protonapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[qldbiface.QLDBAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "qldbapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/qldbapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[qldbsessioniface.QLDBSessionAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "qldbsessionapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/qldbsessionapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[quicksightiface.QuickSightAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "quicksightapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/quicksightapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[ramiface.RAMAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "ramapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/ramapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[rdsiface.RDSAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "rdsapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/rdsapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[rdsdataserviceiface.RDSDataServiceAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "rdsdataserviceapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/rdsdataserviceapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[recyclebiniface.RecycleBinAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "recyclebinapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/recyclebinapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[redshiftiface.RedshiftAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "redshiftapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/redshiftapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[redshiftdataapiserviceiface.RedshiftDataAPIServiceAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "redshiftdataapiserviceapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/redshiftdataapiserviceapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[redshiftserverlessiface.RedshiftServerlessAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "redshiftserverlessapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/redshiftserverlessapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[rekognitioniface.RekognitionAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "rekognitionapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/rekognitionapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[resiliencehubiface.ResilienceHubAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "resiliencehubapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/resiliencehubapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[resourceexplorer2iface.ResourceExplorer2API]()
	if err != nil {
		log.Printf("%s: %v\n", "resourceexplorer2api.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/resourceexplorer2api.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[resourcegroupsiface.ResourceGroupsAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "resourcegroupsapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/resourcegroupsapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[resourcegroupstaggingapiiface.ResourceGroupsTaggingAPIAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "resourcegroupstaggingapiapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/resourcegroupstaggingapiapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[robomakeriface.RoboMakerAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "robomakerapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/robomakerapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[rolesanywhereiface.RolesAnywhereAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "rolesanywhereapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/rolesanywhereapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[route53iface.Route53API]()
	if err != nil {
		log.Printf("%s: %v\n", "route53api.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/route53api.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[route53domainsiface.Route53DomainsAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "route53domainsapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/route53domainsapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[route53recoveryclusteriface.Route53RecoveryClusterAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "route53recoveryclusterapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/route53recoveryclusterapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[route53recoverycontrolconfigiface.Route53RecoveryControlConfigAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "route53recoverycontrolconfigapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/route53recoverycontrolconfigapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[route53recoveryreadinessiface.Route53RecoveryReadinessAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "route53recoveryreadinessapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/route53recoveryreadinessapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[route53resolveriface.Route53ResolverAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "route53resolverapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/route53resolverapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[s3iface.S3API]()
	if err != nil {
		log.Printf("%s: %v\n", "s3api.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/s3api.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[s3controliface.S3ControlAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "s3controlapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/s3controlapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[s3outpostsiface.S3OutpostsAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "s3outpostsapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/s3outpostsapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[sagemakeriface.SageMakerAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "sagemakerapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/sagemakerapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[sagemakeredgemanageriface.SagemakerEdgeManagerAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "sagemakeredgemanagerapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/sagemakeredgemanagerapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[sagemakerfeaturestoreruntimeiface.SageMakerFeatureStoreRuntimeAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "sagemakerfeaturestoreruntimeapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/sagemakerfeaturestoreruntimeapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[sagemakergeospatialiface.SageMakerGeospatialAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "sagemakergeospatialapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/sagemakergeospatialapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[sagemakermetricsiface.SageMakerMetricsAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "sagemakermetricsapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/sagemakermetricsapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[sagemakerruntimeiface.SageMakerRuntimeAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "sagemakerruntimeapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/sagemakerruntimeapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[savingsplansiface.SavingsPlansAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "savingsplansapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/savingsplansapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[scheduleriface.SchedulerAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "schedulerapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/schedulerapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[schemasiface.SchemasAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "schemasapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/schemasapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[secretsmanageriface.SecretsManagerAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "secretsmanagerapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/secretsmanagerapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[securityhubiface.SecurityHubAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "securityhubapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/securityhubapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[securitylakeiface.SecurityLakeAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "securitylakeapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/securitylakeapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[serverlessapplicationrepositoryiface.ServerlessApplicationRepositoryAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "serverlessapplicationrepositoryapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/serverlessapplicationrepositoryapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[servicecatalogiface.ServiceCatalogAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "servicecatalogapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/servicecatalogapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[servicediscoveryiface.ServiceDiscoveryAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "servicediscoveryapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/servicediscoveryapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[servicequotasiface.ServiceQuotasAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "servicequotasapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/servicequotasapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[sesiface.SESAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "sesapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/sesapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[sesv2iface.SESV2API]()
	if err != nil {
		log.Printf("%s: %v\n", "sesv2api.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/sesv2api.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[sfniface.SFNAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "sfnapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/sfnapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[shieldiface.ShieldAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "shieldapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/shieldapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[signeriface.SignerAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "signerapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/signerapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[simpledbiface.SimpleDBAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "simpledbapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/simpledbapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[simspaceweaveriface.SimSpaceWeaverAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "simspaceweaverapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/simspaceweaverapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[smsiface.SMSAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "smsapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/smsapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[snowballiface.SnowballAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "snowballapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/snowballapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[snowdevicemanagementiface.SnowDeviceManagementAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "snowdevicemanagementapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/snowdevicemanagementapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[snsiface.SNSAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "snsapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/snsapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[sqsiface.SQSAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "sqsapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/sqsapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[ssmiface.SSMAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "ssmapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/ssmapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[ssmcontactsiface.SSMContactsAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "ssmcontactsapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/ssmcontactsapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[ssmincidentsiface.SSMIncidentsAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "ssmincidentsapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/ssmincidentsapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[ssmsapiface.SsmSapAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "ssmsapapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/ssmsapapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[ssoiface.SSOAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "ssoapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/ssoapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[ssoadminiface.SSOAdminAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "ssoadminapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/ssoadminapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[ssooidciface.SSOOIDCAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "ssooidcapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/ssooidcapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[storagegatewayiface.StorageGatewayAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "storagegatewayapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/storagegatewayapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[stsiface.STSAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "stsapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/stsapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[supportiface.SupportAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "supportapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/supportapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[supportappiface.SupportAppAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "supportappapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/supportappapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[swfiface.SWFAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "swfapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/swfapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[syntheticsiface.SyntheticsAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "syntheticsapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/syntheticsapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[textractiface.TextractAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "textractapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/textractapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[timestreamqueryiface.TimestreamQueryAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "timestreamqueryapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/timestreamqueryapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[timestreamwriteiface.TimestreamWriteAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "timestreamwriteapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/timestreamwriteapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[transcribeserviceiface.TranscribeServiceAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "transcribeserviceapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/transcribeserviceapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[transcribestreamingserviceiface.TranscribeStreamingServiceAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "transcribestreamingserviceapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/transcribestreamingserviceapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[transferiface.TransferAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "transferapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/transferapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[translateiface.TranslateAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "translateapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/translateapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[voiceidiface.VoiceIDAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "voiceidapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/voiceidapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[wafiface.WAFAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "wafapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/wafapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[wafregionaliface.WAFRegionalAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "wafregionalapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/wafregionalapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[wafv2iface.WAFV2API]()
	if err != nil {
		log.Printf("%s: %v\n", "wafv2api.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/wafv2api.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[wellarchitectediface.WellArchitectedAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "wellarchitectedapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/wellarchitectedapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[workdocsiface.WorkDocsAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "workdocsapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/workdocsapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[worklinkiface.WorkLinkAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "worklinkapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/worklinkapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[workmailiface.WorkMailAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "workmailapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/workmailapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[workmailmessageflowiface.WorkMailMessageFlowAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "workmailmessageflowapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/workmailmessageflowapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[workspacesiface.WorkSpacesAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "workspacesapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/workspacesapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[workspaceswebiface.WorkSpacesWebAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "workspaceswebapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/workspaceswebapi.go", *outputDir), []byte(out), 0644)
	}
	out, err = GenSDK[xrayiface.XRayAPI]()
	if err != nil {
		log.Printf("%s: %v\n", "xrayapi.go", err)
	}
	if out != "" {
		os.WriteFile(fmt.Sprintf("%s/xrayapi.go", *outputDir), []byte(out), 0644)
	}
}
