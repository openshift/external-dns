package alidns

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

// DescribeDnsGtmInstanceSystemCname invokes the alidns.DescribeDnsGtmInstanceSystemCname API synchronously
func (client *Client) DescribeDnsGtmInstanceSystemCname(request *DescribeDnsGtmInstanceSystemCnameRequest) (response *DescribeDnsGtmInstanceSystemCnameResponse, err error) {
	response = CreateDescribeDnsGtmInstanceSystemCnameResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDnsGtmInstanceSystemCnameWithChan invokes the alidns.DescribeDnsGtmInstanceSystemCname API asynchronously
func (client *Client) DescribeDnsGtmInstanceSystemCnameWithChan(request *DescribeDnsGtmInstanceSystemCnameRequest) (<-chan *DescribeDnsGtmInstanceSystemCnameResponse, <-chan error) {
	responseChan := make(chan *DescribeDnsGtmInstanceSystemCnameResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDnsGtmInstanceSystemCname(request)
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

// DescribeDnsGtmInstanceSystemCnameWithCallback invokes the alidns.DescribeDnsGtmInstanceSystemCname API asynchronously
func (client *Client) DescribeDnsGtmInstanceSystemCnameWithCallback(request *DescribeDnsGtmInstanceSystemCnameRequest, callback func(response *DescribeDnsGtmInstanceSystemCnameResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDnsGtmInstanceSystemCnameResponse
		var err error
		defer close(result)
		response, err = client.DescribeDnsGtmInstanceSystemCname(request)
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

// DescribeDnsGtmInstanceSystemCnameRequest is the request struct for api DescribeDnsGtmInstanceSystemCname
type DescribeDnsGtmInstanceSystemCnameRequest struct {
	*requests.RpcRequest
	InstanceId   string `position:"Query" name:"InstanceId"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
}

// DescribeDnsGtmInstanceSystemCnameResponse is the response struct for api DescribeDnsGtmInstanceSystemCname
type DescribeDnsGtmInstanceSystemCnameResponse struct {
	*responses.BaseResponse
	SystemCname string `json:"SystemCname" xml:"SystemCname"`
	RequestId   string `json:"RequestId" xml:"RequestId"`
}

// CreateDescribeDnsGtmInstanceSystemCnameRequest creates a request to invoke DescribeDnsGtmInstanceSystemCname API
func CreateDescribeDnsGtmInstanceSystemCnameRequest() (request *DescribeDnsGtmInstanceSystemCnameRequest) {
	request = &DescribeDnsGtmInstanceSystemCnameRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "DescribeDnsGtmInstanceSystemCname", "alidns", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDnsGtmInstanceSystemCnameResponse creates a response to parse from DescribeDnsGtmInstanceSystemCname response
func CreateDescribeDnsGtmInstanceSystemCnameResponse() (response *DescribeDnsGtmInstanceSystemCnameResponse) {
	response = &DescribeDnsGtmInstanceSystemCnameResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}