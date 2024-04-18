package api

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

// Ec2ApiClient is an interface that defines the methods that we want to use
// from the Client type in the AWS SDK (github.com/aws/aws-sdk-go-v2/service/ec2)
// The AIM is to only contain methods that are defined in the AWS SDK's EC2
// Client.
// For the cases where logic is desired to be implemened combining EC2 calls and
// other logic use the pkg/aws.Client type.
// If you need to use a method provided by the AWS SDK's EC2 Client but it
// is not defined in this interface then it has to be added and all
// the types implementing this interface have to implement the new method.
// The reason this interface has been defined is so we can perform unit testing
// on methods that make use of the AWS EC2 service.
//
//go:generate mockgen -source=ec2_api.go -package=api -destination=ec2_api_mock.go
type Ec2ApiClient interface {
	AttachInternetGateway(ctx context.Context, params *ec2.AttachInternetGatewayInput, optFns ...func(*ec2.Options),
	) (*ec2.AttachInternetGatewayOutput, error)
	AuthorizeSecurityGroupIngress(ctx context.Context, params *ec2.AuthorizeSecurityGroupIngressInput, optFns ...func(*ec2.Options),
	) (*ec2.AuthorizeSecurityGroupIngressOutput, error)
	AssociateRouteTable(ctx context.Context, params *ec2.AssociateRouteTableInput, optFns ...func(*ec2.Options),
	) (*ec2.AssociateRouteTableOutput, error)

	AllocateAddress(ctx context.Context, params *ec2.AllocateAddressInput, optFns ...func(*ec2.Options),
	) (*ec2.AllocateAddressOutput, error)
	AssociateAddress(ctx context.Context, params *ec2.AssociateAddressInput, optFns ...func(*ec2.Options),
	) (*ec2.AssociateAddressOutput, error)

	CopyImage(ctx context.Context, params *ec2.CopyImageInput, optFns ...func(*ec2.Options),
	) (*ec2.CopyImageOutput, error)
	CreateKeyPair(ctx context.Context, params *ec2.CreateKeyPairInput, optFns ...func(*ec2.Options),
	) (*ec2.CreateKeyPairOutput, error)
	CreateInternetGateway(ctx context.Context, params *ec2.CreateInternetGatewayInput, optFns ...func(*ec2.Options),
	) (*ec2.CreateInternetGatewayOutput, error)
	DescribeInstanceStatus(context.Context, *ec2.DescribeInstanceStatusInput, ...func(*ec2.Options),
	) (*ec2.DescribeInstanceStatusOutput, error)
	DescribeInstances(ctx context.Context, params *ec2.DescribeInstancesInput, optFns ...func(*ec2.Options),
	) (*ec2.DescribeInstancesOutput, error)
	DescribeImages(ctx context.Context, params *ec2.DescribeImagesInput, optFns ...func(*ec2.Options),
	) (*ec2.DescribeImagesOutput, error)
	DisassociateAddress(ctx context.Context, params *ec2.DisassociateAddressInput, optFns ...func(*ec2.Options),
	) (*ec2.DisassociateAddressOutput, error)
	DescribeSecurityGroups(ctx context.Context, params *ec2.DescribeSecurityGroupsInput, optFns ...func(*ec2.Options),
	) (*ec2.DescribeSecurityGroupsOutput, error)
	DescribeNatGateways(ctx context.Context, params *ec2.DescribeNatGatewaysInput, optFns ...func(*ec2.Options),
	) (*ec2.DescribeNatGatewaysOutput, error)
	DescribeAddresses(ctx context.Context, params *ec2.DescribeAddressesInput, optFns ...func(*ec2.Options),
	) (*ec2.DescribeAddressesOutput, error)
	DescribeSecurityGroupRules(ctx context.Context,
		params *ec2.DescribeSecurityGroupRulesInput, optFns ...func(*ec2.Options),
	) (*ec2.DescribeSecurityGroupRulesOutput, error)
	DescribeVpcAttribute(ctx context.Context, params *ec2.DescribeVpcAttributeInput, optFns ...func(*ec2.Options),
	) (*ec2.DescribeVpcAttributeOutput, error)
	DescribeAvailabilityZones(ctx context.Context,
		params *ec2.DescribeAvailabilityZonesInput, optFns ...func(*ec2.Options),
	) (*ec2.DescribeAvailabilityZonesOutput, error)
	DescribeRegions(ctx context.Context, params *ec2.DescribeRegionsInput, optFns ...func(*ec2.Options),
	) (*ec2.DescribeRegionsOutput, error)
	DescribeReservedInstancesOfferings(ctx context.Context,
		params *ec2.DescribeReservedInstancesOfferingsInput, optFns ...func(*ec2.Options),
	) (*ec2.DescribeReservedInstancesOfferingsOutput, error)
	DescribeRouteTables(ctx context.Context, params *ec2.DescribeRouteTablesInput, optFns ...func(*ec2.Options),
	) (*ec2.DescribeRouteTablesOutput, error)
	DescribeSubnets(ctx context.Context, params *ec2.DescribeSubnetsInput, optFns ...func(*ec2.Options),
	) (*ec2.DescribeSubnetsOutput, error)
	DescribeVpcs(ctx context.Context, params *ec2.DescribeVpcsInput, optFns ...func(*ec2.Options),
	) (*ec2.DescribeVpcsOutput, error)
	DescribeNetworkInterfaces(ctx context.Context,
		params *ec2.DescribeNetworkInterfacesInput, optFns ...func(*ec2.Options),
	) (*ec2.DescribeNetworkInterfacesOutput, error)
	DescribeInternetGateways(ctx context.Context, params *ec2.DescribeInternetGatewaysInput, optFns ...func(*ec2.Options),
	) (*ec2.DescribeInternetGatewaysOutput, error)
	DescribeInstanceTypeOfferings(ctx context.Context,
		params *ec2.DescribeInstanceTypeOfferingsInput, optFns ...func(*ec2.Options),
	) (*ec2.DescribeInstanceTypeOfferingsOutput, error)
	DeleteKeyPair(ctx context.Context, params *ec2.DeleteKeyPairInput, optFns ...func(*ec2.Options),
	) (*ec2.DeleteKeyPairOutput, error)
	DescribeNetworkAcls(ctx context.Context, params *ec2.DescribeNetworkAclsInput, optFns ...func(*ec2.Options),
	) (*ec2.DescribeNetworkAclsOutput, error)
	DeleteNetworkAclEntry(ctx context.Context, params *ec2.DeleteNetworkAclEntryInput, optFns ...func(*ec2.Options),
	) (*ec2.DeleteNetworkAclEntryOutput, error)
	DescribeVolumes(ctx context.Context, params *ec2.DescribeVolumesInput, optFns ...func(*ec2.Options),
	) (*ec2.DescribeVolumesOutput, error)
	DeleteNetworkInterface(ctx context.Context, params *ec2.DeleteNetworkInterfaceInput, optFns ...func(*ec2.Options),
	) (*ec2.DeleteNetworkInterfaceOutput, error)

	CreateNetworkAclEntry(ctx context.Context, params *ec2.CreateNetworkAclEntryInput, optFns ...func(*ec2.Options),
	) (*ec2.CreateNetworkAclEntryOutput, error)

	CreateSecurityGroup(ctx context.Context, params *ec2.CreateSecurityGroupInput, optFns ...func(*ec2.Options),
	) (*ec2.CreateSecurityGroupOutput, error)
	CreateSubnet(ctx context.Context, params *ec2.CreateSubnetInput, optFns ...func(*ec2.Options),
	) (*ec2.CreateSubnetOutput, error)
	CreateTags(ctx context.Context, params *ec2.CreateTagsInput, optFns ...func(*ec2.Options),
	) (*ec2.CreateTagsOutput, error)
	CreateRoute(ctx context.Context, params *ec2.CreateRouteInput, optFns ...func(*ec2.Options),
	) (*ec2.CreateRouteOutput, error)
	CreateVolume(ctx context.Context, params *ec2.CreateVolumeInput, optFns ...func(*ec2.Options),
	) (*ec2.CreateVolumeOutput, error)
	CreateVpc(ctx context.Context, params *ec2.CreateVpcInput, optFns ...func(*ec2.Options),
	) (*ec2.CreateVpcOutput, error)
	CreateNatGateway(ctx context.Context, params *ec2.CreateNatGatewayInput, optFns ...func(*ec2.Options),
	) (*ec2.CreateNatGatewayOutput, error)
	CreateRouteTable(ctx context.Context, params *ec2.CreateRouteTableInput, optFns ...func(*ec2.Options),
	) (*ec2.CreateRouteTableOutput, error)

	DeleteSecurityGroup(ctx context.Context, params *ec2.DeleteSecurityGroupInput, optFns ...func(*ec2.Options),
	) (*ec2.DeleteSecurityGroupOutput, error)
	DeleteSubnet(ctx context.Context, params *ec2.DeleteSubnetInput, optFns ...func(*ec2.Options),
	) (*ec2.DeleteSubnetOutput, error)
	DeleteTags(ctx context.Context, params *ec2.DeleteTagsInput, optFns ...func(*ec2.Options),
	) (*ec2.DeleteTagsOutput, error)
	DeleteVolume(ctx context.Context, params *ec2.DeleteVolumeInput, optFns ...func(*ec2.Options),
	) (*ec2.DeleteVolumeOutput, error)
	DeleteVpc(ctx context.Context, params *ec2.DeleteVpcInput, optFns ...func(*ec2.Options),
	) (*ec2.DeleteVpcOutput, error)
	DeleteVpcEndpoints(ctx context.Context, params *ec2.DeleteVpcEndpointsInput, optFns ...func(*ec2.Options),
	) (*ec2.DeleteVpcEndpointsOutput, error)
	DescribeVpcEndpoints(context.Context, *ec2.DescribeVpcEndpointsInput, ...func(*ec2.Options),
	) (*ec2.DescribeVpcEndpointsOutput, error)
	DeleteNatGateway(ctx context.Context, params *ec2.DeleteNatGatewayInput, optFns ...func(*ec2.Options),
	) (*ec2.DeleteNatGatewayOutput, error)
	DeleteInternetGateway(ctx context.Context, params *ec2.DeleteInternetGatewayInput, optFns ...func(*ec2.Options),
	) (*ec2.DeleteInternetGatewayOutput, error)
	DetachInternetGateway(ctx context.Context, params *ec2.DetachInternetGatewayInput, optFns ...func(*ec2.Options),
	) (*ec2.DetachInternetGatewayOutput, error)
	DeleteRouteTable(ctx context.Context, params *ec2.DeleteRouteTableInput, optFns ...func(*ec2.Options),
	) (*ec2.DeleteRouteTableOutput, error)
	DisassociateRouteTable(ctx context.Context, params *ec2.DisassociateRouteTableInput, optFns ...func(*ec2.Options),
	) (*ec2.DisassociateRouteTableOutput, error)

	ModifyVpcAttribute(ctx context.Context, params *ec2.ModifyVpcAttributeInput, optFns ...func(*ec2.Options),
	) (*ec2.ModifyVpcAttributeOutput, error)

	RunInstances(ctx context.Context, params *ec2.RunInstancesInput, optFns ...func(*ec2.Options),
	) (*ec2.RunInstancesOutput, error)
	ReleaseAddress(ctx context.Context, params *ec2.ReleaseAddressInput, optFns ...func(*ec2.Options),
	) (*ec2.ReleaseAddressOutput, error)
	RevokeSecurityGroupIngress(ctx context.Context, params *ec2.RevokeSecurityGroupIngressInput, optFns ...func(*ec2.Options),
	) (*ec2.RevokeSecurityGroupIngressOutput, error)
	RevokeSecurityGroupEgress(ctx context.Context, params *ec2.RevokeSecurityGroupEgressInput, optFns ...func(*ec2.Options),
	) (*ec2.RevokeSecurityGroupEgressOutput, error)

	TerminateInstances(ctx context.Context, params *ec2.TerminateInstancesInput, optFns ...func(*ec2.Options),
	) (*ec2.TerminateInstancesOutput, error)
}

// interface guard to ensure that all methods defined in the Ec2ApiClient
// interface are implemented by the real AWS EC2 client. This interface
// guard should always compile
var _ Ec2ApiClient = (*ec2.Client)(nil)
