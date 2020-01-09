package live

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

// AddLiveSnapshotDetectPornConfig invokes the live.AddLiveSnapshotDetectPornConfig API synchronously
// api document: https://help.aliyun.com/api/live/addlivesnapshotdetectpornconfig.html
func (client *Client) AddLiveSnapshotDetectPornConfig(request *AddLiveSnapshotDetectPornConfigRequest) (response *AddLiveSnapshotDetectPornConfigResponse, err error) {
	response = CreateAddLiveSnapshotDetectPornConfigResponse()
	err = client.DoAction(request, response)
	return
}

// AddLiveSnapshotDetectPornConfigWithChan invokes the live.AddLiveSnapshotDetectPornConfig API asynchronously
// api document: https://help.aliyun.com/api/live/addlivesnapshotdetectpornconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddLiveSnapshotDetectPornConfigWithChan(request *AddLiveSnapshotDetectPornConfigRequest) (<-chan *AddLiveSnapshotDetectPornConfigResponse, <-chan error) {
	responseChan := make(chan *AddLiveSnapshotDetectPornConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddLiveSnapshotDetectPornConfig(request)
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

// AddLiveSnapshotDetectPornConfigWithCallback invokes the live.AddLiveSnapshotDetectPornConfig API asynchronously
// api document: https://help.aliyun.com/api/live/addlivesnapshotdetectpornconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddLiveSnapshotDetectPornConfigWithCallback(request *AddLiveSnapshotDetectPornConfigRequest, callback func(response *AddLiveSnapshotDetectPornConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddLiveSnapshotDetectPornConfigResponse
		var err error
		defer close(result)
		response, err = client.AddLiveSnapshotDetectPornConfig(request)
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

// AddLiveSnapshotDetectPornConfigRequest is the request struct for api AddLiveSnapshotDetectPornConfig
type AddLiveSnapshotDetectPornConfigRequest struct {
	*requests.RpcRequest
	OssEndpoint   string           `position:"Query" name:"OssEndpoint"`
	OssObject     string           `position:"Query" name:"OssObject"`
	Scene         *[]string        `position:"Query" name:"Scene"  type:"Repeated"`
	AppName       string           `position:"Query" name:"AppName"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	OssBucket     string           `position:"Query" name:"OssBucket"`
	DomainName    string           `position:"Query" name:"DomainName"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	Interval      requests.Integer `position:"Query" name:"Interval"`
}

// AddLiveSnapshotDetectPornConfigResponse is the response struct for api AddLiveSnapshotDetectPornConfig
type AddLiveSnapshotDetectPornConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAddLiveSnapshotDetectPornConfigRequest creates a request to invoke AddLiveSnapshotDetectPornConfig API
func CreateAddLiveSnapshotDetectPornConfigRequest() (request *AddLiveSnapshotDetectPornConfigRequest) {
	request = &AddLiveSnapshotDetectPornConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "AddLiveSnapshotDetectPornConfig", "live", "openAPI")
	return
}

// CreateAddLiveSnapshotDetectPornConfigResponse creates a response to parse from AddLiveSnapshotDetectPornConfig response
func CreateAddLiveSnapshotDetectPornConfigResponse() (response *AddLiveSnapshotDetectPornConfigResponse) {
	response = &AddLiveSnapshotDetectPornConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}