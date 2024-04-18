package aws_client

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/openshift-online/ocm-common/pkg/aws/api"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/sts"

	"github.com/openshift-online/ocm-common/pkg/log"

	elb "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing"
	"github.com/aws/aws-sdk-go-v2/service/route53"
)

type AwsClient interface {
	ListVPCByName(vpcName string) ([]types.Vpc, error)
	CreateVpc(cidr string, name ...string) (*ec2.CreateVpcOutput, error)
}

type awsClient struct {
	config    aws.Config
	accountID string
	ec2Client api.Ec2ApiClient
	stsClient api.StsApiClient
	iamClient api.IamApiClient

	route53Client        *route53.Client
	stackFormationClient *cloudformation.Client
	elbClient            *elb.Client
	kmsClient            *kms.Client
}

type AwsClientSpec struct {
	Config    aws.Config
	AccountID string
	Ec2Client api.Ec2ApiClient
	StsClient api.StsApiClient
	IamClient api.IamApiClient

	Route53Client        *route53.Client
	StackFormationClient *cloudformation.Client
	ElbClient            *elb.Client
	KmsClient            *kms.Client
}

func NewAwsClient(spec AwsClientSpec) AwsClient {
	return &awsClient{
		config:    spec.Config,
		ec2Client: spec.Ec2Client,
		stsClient: spec.StsClient,
		iamClient: spec.IamClient,

		route53Client:        spec.Route53Client,
		stackFormationClient: spec.StackFormationClient,
		elbClient:            spec.ElbClient,
		kmsClient:            spec.KmsClient,
	}
}

func CreateAWSClient(profileName string, region string) (AwsClient, error) {
	var cfg aws.Config
	var err error

	if envCredential() {
		log.LogInfo("Got AWS_ACCESS_KEY_ID env settings, going to build the config with the env")
		cfg, err = config.LoadDefaultConfig(context.TODO(),
			config.WithRegion(region),
			config.WithCredentialsProvider(
				credentials.NewStaticCredentialsProvider(
					os.Getenv("AWS_ACCESS_KEY_ID"),
					os.Getenv("AWS_SECRET_ACCESS_KEY"),
					"")),
		)
	} else {
		if envAwsProfile() {
			file := os.Getenv("AWS_SHARED_CREDENTIALS_FILE")
			log.LogInfo("Got file path: %s from env variable AWS_SHARED_CREDENTIALS_FILE\n", file)
			cfg, err = config.LoadDefaultConfig(context.TODO(),
				config.WithRegion(region),
				config.WithSharedCredentialsFiles([]string{file}),
			)
		} else {
			cfg, err = config.LoadDefaultConfig(context.TODO(),
				config.WithRegion(region),
				config.WithSharedConfigProfile(profileName),
			)
		}

	}

	if err != nil {
		return nil, err
	}

	return NewAwsClient(AwsClientSpec{
		Ec2Client: ec2.NewFromConfig(cfg),
		StsClient: sts.NewFromConfig(cfg),
		IamClient: iam.NewFromConfig(cfg),
	}), nil
}
