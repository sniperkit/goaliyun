package cms

import (
	"github.com/morlay/goaliyun"
)

type QueryCustomEventHistogramRequest struct {
	QueryJson string `position:"Query" name:"QueryJson"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *QueryCustomEventHistogramRequest) Invoke(client goaliyun.Client) (*QueryCustomEventHistogramResponse, error) {
	resp := &QueryCustomEventHistogramResponse{}
	err := client.Request("cms", "QueryCustomEventHistogram", "2018-03-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryCustomEventHistogramResponse struct {
	Code      goaliyun.String
	Message   goaliyun.String
	Data      goaliyun.String
	RequestId goaliyun.String
	Success   goaliyun.String
}
