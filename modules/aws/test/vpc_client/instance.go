package vpc_client

import (
	"github.com/openshift-online/ocm-common/module/aws/consts"
	"github.com/openshift-online/ocm-common/module/aws/log"
	"strings"
)

func (vpc *VPC) TerminateVPCInstances(nonClusterOnly bool) error {
	filters := []map[string][]string{
		{
			"vpc-id": []string{
				vpc.VpcID,
			},
		},
	}
	if nonClusterOnly {
		filters = append(filters, map[string][]string{
			"tag:Name": {
				consts.ProxyName,
				consts.BastionName,
			},
		})
	}
	insts, err := vpc.AWSClient.ListInstances([]string{}, filters...)

	if err != nil {
		log.LogError("Error happened when list instances for vpc %s: %s", vpc.VpcID, err)
		return err
	}
	needTermination := []string{}
	for _, inst := range insts {
		needTermination = append(needTermination, *inst.InstanceId)
	}
	err = vpc.AWSClient.TerminateInstances(needTermination, true, 20)
	if err != nil {
		log.LogError("Terminating instances %s meet error: %s", strings.Join(needTermination, ","), err)
	} else {
		log.LogInfo("Terminating instances %s successfully", strings.Join(needTermination, ","))
	}
	return err

}
