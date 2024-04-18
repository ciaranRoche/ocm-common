package aws_client

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/openshift-online/ocm-common/pkg/aws/api"
	"go.uber.org/mock/gomock"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("VPC Client", func() {
	Describe("ListVPCByName", func() {
		Context("when DescribeVpcs returns Vpcs", func() {
			It("should return the list of Vpcs", func() {
				ctrl := gomock.NewController(GinkgoT())
				mockEc2 := api.NewMockEc2ApiClient(ctrl)
				client := NewAwsClient(AwsClientSpec{
					Ec2Client: mockEc2,
				})

				name := "my-vpc"
				mockEc2.EXPECT().DescribeVpcs(gomock.Any(), &ec2.DescribeVpcsInput{
					Filters: []types.Filter{
						{
							Name:   aws.String("tag:Name"),
							Values: []string{name},
						},
					},
				}).Return(&ec2.DescribeVpcsOutput{
					Vpcs: []types.Vpc{
						{
							Tags: []types.Tag{
								{
									Key:   aws.String("Name"),
									Value: aws.String(name),
								},
							},
						},
					},
				}, nil).Times(1)

				vpcs, err := client.ListVPCByName(name)
				Expect(err).ToNot(HaveOccurred())
				Expect(len(vpcs)).To(Equal(1))
				Expect(len(vpcs[0].Tags)).To(Equal(1))
				Expect(aws.ToString(vpcs[0].Tags[0].Key)).To(Equal("Name"))
				Expect(aws.ToString(vpcs[0].Tags[0].Value)).To(Equal(name))

			})
		})
	})

	Describe("CreateVpc", func() {
		Context("when CreateVpc is called with CIDR and optional name", func() {
			It("should create a VPC with the provided CIDR and name", func() {
				ctrl := gomock.NewController(GinkgoT())
				mockEc2 := api.NewMockEc2ApiClient(ctrl)
				client := NewAwsClient(AwsClientSpec{
					Ec2Client: mockEc2,
				})

				cidr := "10.0.0.0/16"
				name := "my-vpc"

				mockEc2.EXPECT().DescribeVpcs(gomock.Any(), gomock.Any()).Return(&ec2.DescribeVpcsOutput{
					Vpcs: []types.Vpc{
						{
							VpcId: aws.String("vpc-12345678"),
						},
					},
				}, nil).AnyTimes()

				mockEc2.EXPECT().CreateTags(gomock.Any(), gomock.Any(), gomock.Any()).Return(
					&ec2.CreateTagsOutput{}, nil).Times(1)

				mockEc2.EXPECT().CreateVpc(gomock.Any(), &ec2.CreateVpcInput{
					CidrBlock:         aws.String(cidr),
					DryRun:            nil,
					Ipv4IpamPoolId:    nil,
					Ipv4NetmaskLength: nil,
					TagSpecifications: nil,
				}).DoAndReturn(func(ctx context.Context, input *ec2.CreateVpcInput, optFns ...func(*ec2.Options)) (*ec2.CreateVpcOutput, error) {
					Expect(input.CidrBlock).To(Equal(aws.String(cidr)))
					return &ec2.CreateVpcOutput{
						Vpc: &types.Vpc{
							VpcId: aws.String("vpc-12345678"),
						},
					}, nil
				}).Times(1)

				_, err := client.CreateVpc(cidr, name)
				Expect(err).ToNot(HaveOccurred())
			})
		})
	})
})
