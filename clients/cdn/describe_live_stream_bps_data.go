package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeLiveStreamBpsDataRequest struct {
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	StreamName    string `position:"Query" name:"StreamName"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeLiveStreamBpsDataRequest) Invoke(client goaliyun.Client) (*DescribeLiveStreamBpsDataResponse, error) {
	resp := &DescribeLiveStreamBpsDataResponse{}
	err := client.Request("cdn", "DescribeLiveStreamBpsData", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeLiveStreamBpsDataResponse struct {
	RequestId goaliyun.String
	BpsDatas  DescribeLiveStreamBpsDataDomainBpsModelList
}

type DescribeLiveStreamBpsDataDomainBpsModel struct {
	Time goaliyun.String
	Bps  goaliyun.Float
}

type DescribeLiveStreamBpsDataDomainBpsModelList []DescribeLiveStreamBpsDataDomainBpsModel

func (list *DescribeLiveStreamBpsDataDomainBpsModelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamBpsDataDomainBpsModel)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
