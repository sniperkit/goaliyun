package emr

import (
	"github.com/morlay/goaliyun"
)

type ModifyFlowForWebRequest struct {
	CronExpr        string `position:"Query" name:"CronExpr"`
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Periodic        string `position:"Query" name:"Periodic"`
	StartSchedule   int64  `position:"Query" name:"StartSchedule"`
	Description     string `position:"Query" name:"Description"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	Graph           string `position:"Query" name:"Graph"`
	CreateCluster   string `position:"Query" name:"CreateCluster"`
	Name            string `position:"Query" name:"Name"`
	EndSchedule     int64  `position:"Query" name:"EndSchedule"`
	Id              string `position:"Query" name:"Id"`
	ProjectId       string `position:"Query" name:"ProjectId"`
	Status          string `position:"Query" name:"Status"`
	ParentCategory  string `position:"Query" name:"ParentCategory"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ModifyFlowForWebRequest) Invoke(client goaliyun.Client) (*ModifyFlowForWebResponse, error) {
	resp := &ModifyFlowForWebResponse{}
	err := client.Request("emr", "ModifyFlowForWeb", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyFlowForWebResponse struct {
	RequestId goaliyun.String
	Data      bool
}
