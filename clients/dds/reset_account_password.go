package dds

import (
	"github.com/morlay/goaliyun"
)

type ResetAccountPasswordRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	AccountPassword      string `position:"Query" name:"AccountPassword"`
	AccountName          string `position:"Query" name:"AccountName"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ResetAccountPasswordRequest) Invoke(client goaliyun.Client) (*ResetAccountPasswordResponse, error) {
	resp := &ResetAccountPasswordResponse{}
	err := client.Request("dds", "ResetAccountPassword", "2015-12-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ResetAccountPasswordResponse struct {
	RequestId goaliyun.String
}
