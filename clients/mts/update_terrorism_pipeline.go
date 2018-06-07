package mts

import (
	"github.com/morlay/goaliyun"
)

type UpdateTerrorismPipelineRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	State                string `position:"Query" name:"State"`
	NotifyConfig         string `position:"Query" name:"NotifyConfig"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Priority             int64  `position:"Query" name:"Priority"`
	PipelineId           string `position:"Query" name:"PipelineId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *UpdateTerrorismPipelineRequest) Invoke(client goaliyun.Client) (*UpdateTerrorismPipelineResponse, error) {
	resp := &UpdateTerrorismPipelineResponse{}
	err := client.Request("mts", "UpdateTerrorismPipeline", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UpdateTerrorismPipelineResponse struct {
	RequestId goaliyun.String
	Pipeline  UpdateTerrorismPipelinePipeline
}

type UpdateTerrorismPipelinePipeline struct {
	Id           goaliyun.String
	Name         goaliyun.String
	State        goaliyun.String
	Priority     goaliyun.Integer
	NotifyConfig UpdateTerrorismPipelineNotifyConfig
}

type UpdateTerrorismPipelineNotifyConfig struct {
	Topic goaliyun.String
	Queue goaliyun.String
}
