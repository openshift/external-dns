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

// StatisticItem is a nested struct in alidns response
type StatisticItem struct {
	UdpTotalCount    int64        `json:"UdpTotalCount" xml:"UdpTotalCount"`
	ThreatCount      int64        `json:"ThreatCount" xml:"ThreatCount"`
	DomainCount      int64        `json:"DomainCount" xml:"DomainCount"`
	HttpCount        int64        `json:"HttpCount" xml:"HttpCount"`
	TotalCount       int64        `json:"TotalCount" xml:"TotalCount"`
	Timestamp        int64        `json:"Timestamp" xml:"Timestamp"`
	DomainName       string       `json:"DomainName" xml:"DomainName"`
	V6HttpCount      int64        `json:"V6HttpCount" xml:"V6HttpCount"`
	SubDomain        string       `json:"SubDomain" xml:"SubDomain"`
	MaxThreatLevel   string       `json:"MaxThreatLevel" xml:"MaxThreatLevel"`
	V4HttpCount      int64        `json:"V4HttpCount" xml:"V4HttpCount"`
	V6HttpsCount     int64        `json:"V6HttpsCount" xml:"V6HttpsCount"`
	IpCount          int64        `json:"IpCount" xml:"IpCount"`
	HttpsCount       int64        `json:"HttpsCount" xml:"HttpsCount"`
	LatestThreatTime int64        `json:"LatestThreatTime" xml:"LatestThreatTime"`
	DohTotalCount    int64        `json:"DohTotalCount" xml:"DohTotalCount"`
	SourceIp         string       `json:"SourceIp" xml:"SourceIp"`
	V4Count          int64        `json:"V4Count" xml:"V4Count"`
	V6Count          int64        `json:"V6Count" xml:"V6Count"`
	ThreatLevel      string       `json:"ThreatLevel" xml:"ThreatLevel"`
	V4HttpsCount     int64        `json:"V4HttpsCount" xml:"V4HttpsCount"`
	ThreatType       string       `json:"ThreatType" xml:"ThreatType"`
	ThreatInfo       []ThreatItem `json:"ThreatInfo" xml:"ThreatInfo"`
}
