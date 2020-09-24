package alikafka

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DeleteConsumerGroup invokes the alikafka.DeleteConsumerGroup API synchronously
// api document: https://help.aliyun.com/api/alikafka/deleteconsumergroup.html
func (client *Client) DeleteConsumerGroup(request *DeleteConsumerGroupRequest) (response *DeleteConsumerGroupResponse, err error) {
	response = CreateDeleteConsumerGroupResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteConsumerGroupWithChan invokes the alikafka.DeleteConsumerGroup API asynchronously
// api document: https://help.aliyun.com/api/alikafka/deleteconsumergroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteConsumerGroupWithChan(request *DeleteConsumerGroupRequest) (<-chan *DeleteConsumerGroupResponse, <-chan error) {
	responseChan := make(chan *DeleteConsumerGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteConsumerGroup(request)
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

// DeleteConsumerGroupWithCallback invokes the alikafka.DeleteConsumerGroup API asynchronously
// api document: https://help.aliyun.com/api/alikafka/deleteconsumergroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteConsumerGroupWithCallback(request *DeleteConsumerGroupRequest, callback func(response *DeleteConsumerGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteConsumerGroupResponse
		var err error
		defer close(result)
		response, err = client.DeleteConsumerGroup(request)
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

// DeleteConsumerGroupRequest is the request struct for api DeleteConsumerGroup
type DeleteConsumerGroupRequest struct {
	*requests.RpcRequest
	ConsumerId string `position:"Query" name:"ConsumerId"`
	InstanceId string `position:"Query" name:"InstanceId"`
}

// DeleteConsumerGroupResponse is the response struct for api DeleteConsumerGroup
type DeleteConsumerGroupResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateDeleteConsumerGroupRequest creates a request to invoke DeleteConsumerGroup API
func CreateDeleteConsumerGroupRequest() (request *DeleteConsumerGroupRequest) {
	request = &DeleteConsumerGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("alikafka", "2019-09-16", "DeleteConsumerGroup", "alikafka", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteConsumerGroupResponse creates a response to parse from DeleteConsumerGroup response
func CreateDeleteConsumerGroupResponse() (response *DeleteConsumerGroupResponse) {
	response = &DeleteConsumerGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
