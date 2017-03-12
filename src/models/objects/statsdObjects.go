//
//Copyright [2016] [SnapRoute Inc]
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.
//
// _______  __       __________   ___      _______.____    __    ____  __  .___________.  ______  __    __
// |   ____||  |     |   ____\  \ /  /     /       |\   \  /  \  /   / |  | |           | /      ||  |  |  |
// |  |__   |  |     |  |__   \  V  /     |   (----` \   \/    \/   /  |  | `---|  |----`|  ,----'|  |__|  |
// |   __|  |  |     |   __|   >   <       \   \      \            /   |  |     |  |     |  |     |   __   |
// |  |     |  `----.|  |____ /  .  \  .----)   |      \    /\    /    |  |     |  |     |  `----.|  |  |  |
// |__|     |_______||_______/__/ \__\ |_______/        \__/  \__/     |__|     |__|      \______||__|  |__|
//

package objects

type SflowGlobal struct {
	baseObj
	Vrf                 string `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY: "1", DESCRIPTION: "VRF that this sflow instance is associated with", DEFAULT: "default"`
	AdminState          string `DESCRIPTION: "Administrative state of all sflow collectors and interfaces setup in the system", SELECTION: "UP"/"DOWN", DEFAULT: "UP"`
	AgentIpAddr         string `DESCRIPTION: "Source ip address to use for the Sflow agent"`
	MaxSampledSize      int32  `DESCRIPTION: "Maximum number of bytes sampled per packet", MIN: 64, MAX: 256, DEFAULT: 128`
	CounterPollInterval int32  `DESCRIPTION: "Interval between successive poll cycles of sflow counters in seconds. Set to 0 to disable polling", DEFAULT:20`
	MaxDatagramSize     int32  `DESCRIPTION: "Maximum number of data bytes that can be sent in a single datagram", DEFAULT: 1400`
}

type SflowCollector struct {
	baseObj
	IpAddr     string `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY: "*", DESCRIPTION: "IP address corresponding to the sflow collector/analyzer"`
	UdpPort    int32  `DESCRIPTION: "UDP port number that the sflow collector/analyzer is listening on", DEFAULT: 6343`
	AdminState string `DESCRIPTION: "Administrative state for this collector. When set to 'DOWN', sflow data is not exported to this collector", SELECTION: "UP"/"DOWN", DEFAULT: "UP"`
}

type SflowCollectorState struct {
	baseObj
	IpAddr                  string `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY: "*", DESCRIPTION: "IP address corresponding to the sflow collector/analyzer"`
	OperState               string `DESCRIPTION: "Operational state of this sflow collector"`
	NumSflowSamplesExported int32  `DESCRIPTION: "Total number of sflow records exported to this collector"`
	NumDatagramExported     int32  `DESCRIPTION: "Number of sflow datagrams that have been exported to this collector/analyzer"`
}

type SflowIntf struct {
	baseObj
	IntfRef      string `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY: "*", DESCRIPTION: "Physical interface name that this interface configuration applies to" `
	AdminState   string `DESCRIPTION:"Adminstratively enable/disable sflow sampling for this interface", SELECTION: "UP"/"DOWN", DEFAULT: "DOWN"`
	SamplingRate int32  `DESCRIPTION: "Sampling rate to use for this interface. Set to 'n' to sample 1/n th of the packets", DEFAULT: 100`
}

type SflowIntfState struct {
	baseObj
	IntfRef                 string `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY: "*", DESCRIPTION: "Physical interface name that this interface configuration applies to" `
	OperState               string `DESCRIPTION: "Operational state of this sflow collector"`
	NumSflowSamplesExported int32  `DESCRIPTION: "Number of sflow records that have been exported for this interface"`
}
