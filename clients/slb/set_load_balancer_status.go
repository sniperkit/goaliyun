package slb

import (
	"github.com/morlay/goaliyun"
)

type SetLoadBalancerStatusRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	LoadBalancerId       string `position:"Query" name:"LoadBalancerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	LoadBalancerStatus   string `position:"Query" name:"LoadBalancerStatus"`
	Tags                 string `position:"Query" name:"Tags"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *SetLoadBalancerStatusRequest) Invoke(client goaliyun.Client) (*SetLoadBalancerStatusResponse, error) {
	resp := &SetLoadBalancerStatusResponse{}
	err := client.Request("slb", "SetLoadBalancerStatus", "2014-05-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetLoadBalancerStatusResponse struct {
	RequestId goaliyun.String
}
