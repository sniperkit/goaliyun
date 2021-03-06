package slb

import (
	"github.com/morlay/goaliyun"
)

type SetLoadBalancerTCPListenerAttributeRequest struct {
	Access_key_id             string `position:"Query" name:"Access_key_id"`
	HealthCheckConnectTimeout int64  `position:"Query" name:"HealthCheckConnectTimeout"`
	ResourceOwnerId           int64  `position:"Query" name:"ResourceOwnerId"`
	HealthCheckURI            string `position:"Query" name:"HealthCheckURI"`
	UnhealthyThreshold        int64  `position:"Query" name:"UnhealthyThreshold"`
	HealthyThreshold          int64  `position:"Query" name:"HealthyThreshold"`
	AclStatus                 string `position:"Query" name:"AclStatus"`
	Scheduler                 string `position:"Query" name:"Scheduler"`
	AclType                   string `position:"Query" name:"AclType"`
	MasterSlaveServerGroup    string `position:"Query" name:"MasterSlaveServerGroup"`
	EstablishedTimeout        int64  `position:"Query" name:"EstablishedTimeout"`
	MaxConnection             int64  `position:"Query" name:"MaxConnection"`
	PersistenceTimeout        int64  `position:"Query" name:"PersistenceTimeout"`
	VpcIds                    string `position:"Query" name:"VpcIds"`
	VServerGroupId            string `position:"Query" name:"VServerGroupId"`
	AclId                     string `position:"Query" name:"AclId"`
	ListenerPort              int64  `position:"Query" name:"ListenerPort"`
	HealthCheckType           string `position:"Query" name:"HealthCheckType"`
	ResourceOwnerAccount      string `position:"Query" name:"ResourceOwnerAccount"`
	Bandwidth                 int64  `position:"Query" name:"Bandwidth"`
	HealthCheckDomain         string `position:"Query" name:"HealthCheckDomain"`
	OwnerAccount              string `position:"Query" name:"OwnerAccount"`
	SynProxy                  string `position:"Query" name:"SynProxy"`
	OwnerId                   int64  `position:"Query" name:"OwnerId"`
	Tags                      string `position:"Query" name:"Tags"`
	LoadBalancerId            string `position:"Query" name:"LoadBalancerId"`
	MasterSlaveServerGroupId  string `position:"Query" name:"MasterSlaveServerGroupId"`
	HealthCheckInterval       int64  `position:"Query" name:"HealthCheckInterval"`
	HealthCheckConnectPort    int64  `position:"Query" name:"HealthCheckConnectPort"`
	HealthCheckHttpCode       string `position:"Query" name:"HealthCheckHttpCode"`
	VServerGroup              string `position:"Query" name:"VServerGroup"`
	RegionId                  string `position:"Query" name:"RegionId"`
}

func (req *SetLoadBalancerTCPListenerAttributeRequest) Invoke(client goaliyun.Client) (*SetLoadBalancerTCPListenerAttributeResponse, error) {
	resp := &SetLoadBalancerTCPListenerAttributeResponse{}
	err := client.Request("slb", "SetLoadBalancerTCPListenerAttribute", "2014-05-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetLoadBalancerTCPListenerAttributeResponse struct {
	RequestId goaliyun.String
}
