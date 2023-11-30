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

// Strategy is a nested struct in alidns response
type Strategy struct {
	StrategyId                    string                             `json:"StrategyId" xml:"StrategyId"`
	StrategyName                  string                             `json:"StrategyName" xml:"StrategyName"`
	CreateTime                    string                             `json:"CreateTime" xml:"CreateTime"`
	CreateTimestamp               int64                              `json:"CreateTimestamp" xml:"CreateTimestamp"`
	DefaultAddrPoolId             string                             `json:"DefaultAddrPoolId" xml:"DefaultAddrPoolId"`
	DefaultAddrPoolName           string                             `json:"DefaultAddrPoolName" xml:"DefaultAddrPoolName"`
	FailoverAddrPoolId            string                             `json:"FailoverAddrPoolId" xml:"FailoverAddrPoolId"`
	FailoverAddrPoolName          string                             `json:"FailoverAddrPoolName" xml:"FailoverAddrPoolName"`
	AccessMode                    string                             `json:"AccessMode" xml:"AccessMode"`
	AccessStatus                  string                             `json:"AccessStatus" xml:"AccessStatus"`
	StrategyMode                  string                             `json:"StrategyMode" xml:"StrategyMode"`
	InstanceId                    string                             `json:"InstanceId" xml:"InstanceId"`
	DefaultAddrPoolStatus         string                             `json:"DefaultAddrPoolStatus" xml:"DefaultAddrPoolStatus"`
	FailoverAddrPoolStatus        string                             `json:"FailoverAddrPoolStatus" xml:"FailoverAddrPoolStatus"`
	DefaultAddrPoolMonitorStatus  string                             `json:"DefaultAddrPoolMonitorStatus" xml:"DefaultAddrPoolMonitorStatus"`
	FailoverAddrPoolMonitorStatus string                             `json:"FailoverAddrPoolMonitorStatus" xml:"FailoverAddrPoolMonitorStatus"`
	Lines                         LinesInDescribeGtmAccessStrategies `json:"Lines" xml:"Lines"`
}