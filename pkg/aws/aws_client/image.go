package aws_client

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/openshift-online/ocm-common/pkg/log"
)

func (client *awsClient) CopyImage(sourceImageID string, sourceRegion string, name string) (string, error) {
	copyImageInput := &ec2.CopyImageInput{
		Name:          &name,
		SourceImageId: &sourceImageID,
		SourceRegion:  &sourceRegion,
	}
	output, err := client.ec2Client.CopyImage(context.TODO(), copyImageInput)
	if err != nil {
		log.LogError("Error happens when copy image: %s", err)
		return "", err
	}
	return *output.ImageId, nil
}

func (client *awsClient) DescribeImage(imageID string) (*ec2.DescribeImagesOutput, error) {
	describeImageInput := &ec2.DescribeImagesInput{
		ImageIds: []string{imageID},
	}
	output, err := client.ec2Client.DescribeImages(context.TODO(), describeImageInput)
	if err != nil {
		log.LogError("Describe image %s meet error: %s", imageID, err)
		return nil, err
	}

	return output, nil
}
