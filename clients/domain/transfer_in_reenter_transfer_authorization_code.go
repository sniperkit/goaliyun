package domain

import (
	"github.com/morlay/goaliyun"
)

type TransferInReenterTransferAuthorizationCodeRequest struct {
	TransferAuthorizationCode string `position:"Query" name:"TransferAuthorizationCode"`
	DomainName                string `position:"Query" name:"DomainName"`
	UserClientIp              string `position:"Query" name:"UserClientIp"`
	Lang                      string `position:"Query" name:"Lang"`
	RegionId                  string `position:"Query" name:"RegionId"`
}

func (req *TransferInReenterTransferAuthorizationCodeRequest) Invoke(client goaliyun.Client) (*TransferInReenterTransferAuthorizationCodeResponse, error) {
	resp := &TransferInReenterTransferAuthorizationCodeResponse{}
	err := client.Request("domain", "TransferInReenterTransferAuthorizationCode", "2018-01-29", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type TransferInReenterTransferAuthorizationCodeResponse struct {
	RequestId goaliyun.String
}
