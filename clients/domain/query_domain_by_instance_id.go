package domain

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryDomainByInstanceIdRequest struct {
	InstanceId   string `position:"Query" name:"InstanceId"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *QueryDomainByInstanceIdRequest) Invoke(client goaliyun.Client) (*QueryDomainByInstanceIdResponse, error) {
	resp := &QueryDomainByInstanceIdResponse{}
	err := client.Request("domain", "QueryDomainByInstanceId", "2018-01-29", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryDomainByInstanceIdResponse struct {
	UserId                       goaliyun.String
	DomainName                   goaliyun.String
	InstanceId                   goaliyun.String
	RegistrationDate             goaliyun.String
	ExpirationDate               goaliyun.String
	RegistrantOrganization       goaliyun.String
	RegistrantName               goaliyun.String
	Email                        goaliyun.String
	UpdateProhibitionLock        goaliyun.String
	TransferProhibitionLock      goaliyun.String
	DomainNameProxyService       bool
	Premium                      bool
	EmailVerificationStatus      goaliyun.Integer
	EmailVerificationClientHold  bool
	RealNameStatus               goaliyun.String
	RegistrantUpdatingStatus     goaliyun.String
	TransferOutStatus            goaliyun.String
	RegistrantType               goaliyun.String
	DomainNameVerificationStatus goaliyun.String
	RequestId                    goaliyun.String
	DnsList                      QueryDomainByInstanceIdDnsListList
}

type QueryDomainByInstanceIdDnsListList []goaliyun.String

func (list *QueryDomainByInstanceIdDnsListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]goaliyun.String)
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
