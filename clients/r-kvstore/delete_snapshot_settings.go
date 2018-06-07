package r_kvstore

import (
	"github.com/morlay/goaliyun"
)

type DeleteSnapshotSettingsRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DeleteSnapshotSettingsRequest) Invoke(client goaliyun.Client) (*DeleteSnapshotSettingsResponse, error) {
	resp := &DeleteSnapshotSettingsResponse{}
	err := client.Request("r-kvstore", "DeleteSnapshotSettings", "2015-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteSnapshotSettingsResponse struct {
	RequestId goaliyun.String
}
