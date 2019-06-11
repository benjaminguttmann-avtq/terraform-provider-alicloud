package gpdb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeRdsVSwitchs invokes the gpdb.DescribeRdsVSwitchs API synchronously
// api document: https://help.aliyun.com/api/gpdb/describerdsvswitchs.html
func (client *Client) DescribeRdsVSwitchs(request *DescribeRdsVSwitchsRequest) (response *DescribeRdsVSwitchsResponse, err error) {
	response = CreateDescribeRdsVSwitchsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeRdsVSwitchsWithChan invokes the gpdb.DescribeRdsVSwitchs API asynchronously
// api document: https://help.aliyun.com/api/gpdb/describerdsvswitchs.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeRdsVSwitchsWithChan(request *DescribeRdsVSwitchsRequest) (<-chan *DescribeRdsVSwitchsResponse, <-chan error) {
	responseChan := make(chan *DescribeRdsVSwitchsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRdsVSwitchs(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DescribeRdsVSwitchsWithCallback invokes the gpdb.DescribeRdsVSwitchs API asynchronously
// api document: https://help.aliyun.com/api/gpdb/describerdsvswitchs.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeRdsVSwitchsWithCallback(request *DescribeRdsVSwitchsRequest, callback func(response *DescribeRdsVSwitchsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRdsVSwitchsResponse
		var err error
		defer close(result)
		response, err = client.DescribeRdsVSwitchs(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DescribeRdsVSwitchsRequest is the request struct for api DescribeRdsVSwitchs
type DescribeRdsVSwitchsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	VpcId                string           `position:"Query" name:"VpcId"`
	ZoneId               string           `position:"Query" name:"ZoneId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeRdsVSwitchsResponse is the response struct for api DescribeRdsVSwitchs
type DescribeRdsVSwitchsResponse struct {
	*responses.BaseResponse
	RequestId string    `json:"RequestId" xml:"RequestId"`
	VSwitches VSwitches `json:"VSwitches" xml:"VSwitches"`
}

// CreateDescribeRdsVSwitchsRequest creates a request to invoke DescribeRdsVSwitchs API
func CreateDescribeRdsVSwitchsRequest() (request *DescribeRdsVSwitchsRequest) {
	request = &DescribeRdsVSwitchsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("gpdb", "2016-05-03", "DescribeRdsVSwitchs", "gpdb", "openAPI")
	return
}

// CreateDescribeRdsVSwitchsResponse creates a response to parse from DescribeRdsVSwitchs response
func CreateDescribeRdsVSwitchsResponse() (response *DescribeRdsVSwitchsResponse) {
	response = &DescribeRdsVSwitchsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}