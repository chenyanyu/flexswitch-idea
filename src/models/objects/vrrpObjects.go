//
//Copyright [2016] [SnapRoute Inc]
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//	 Unless required by applicable law or agreed to in writing, software
//	 distributed under the License is distributed on an "AS IS" BASIS,
//	 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//	 See the License for the specific language governing permissions and
//	 limitations under the License.
//
// _______  __       __________   ___      _______.____    __    ____  __  .___________.  ______  __    __
// |   ____||  |     |   ____\  \ /  /     /       |\   \  /  \  /   / |  | |           | /      ||  |  |  |
// |  |__   |  |     |  |__   \  V  /     |   (----` \   \/    \/   /  |  | `---|  |----`|  ,----'|  |__|  |
// |   __|  |  |     |   __|   >   <       \   \      \            /   |  |     |  |     |  |     |   __   |
// |  |     |  `----.|  |____ /  .  \  .----)   |      \    /\    /    |  |     |  |     |  `----.|  |  |  |
// |__|     |_______||_______/__/ \__\ |_______/        \__/  \__/     |__|     |__|      \______||__|  |__|
//

package objects

type VrrpGlobal struct {
	baseObj
	Vrf    string `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"w", MULTIPLICITY:"1", AUTOCREATE: "true", DESCRIPTION: "System Vrf", DEFAULT:"default"`
	Enable bool   `DESCRIPTION: "Enable/Disable VRRP Globally", DEFAULT:false`
}

type VrrpGlobalState struct {
	baseObj
	Vrf           string `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"r", MULTIPLICITY:"1", DESCRIPTION: "System Vrf"`
	Status        string `DESCRIPTION: "Enable/Disable VRRP Globally"`
	V4Intfs       int32  `DESCRIPTION: "vrrp v4 interfaces configured"`
	V6Intfs       int32  `DESCRIPTION: "vrrp v6 interfaces configured"`
	TotalRxFrames int32  `DESCRIPTION: "total vrrp advertisement received`
	TotalTxFrames int32  `DESCRIPTION: "total vrrp advertisement send out"`
}

type VrrpV4Intf struct {
	baseObj
	IntfRef               string `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"w", MULTIPLICITY:"*", DESCRIPTION: "Interface (name) for which VRRP Version 2 aka VRRP with ipv4 Config needs to be done"`
	VRID                  int32  `SNAPROUTE: "KEY", CATEGORY:"L3", DESCRIPTION: "Virtual Router's Unique Identifier", MIN:1, MAX:10`
	Version               string `DESCRIPTION: "vrrp should be running in which version, SELECTION:"version2/version3", DEFAULT:"version3"`
	Priority              int32  `DESCRIPTION: "Sending VRRP router's priority for the virtual router", DEFAULT:100, MIN:1, MAX:255`
	Address               string `DESCRIPTION: "Virtual Router IPv4 address", STRLEN:"17"`
	AdvertisementInterval int32  `DESCRIPTION: "Time interval between ADVERTISEMENTS", DEFAULT:1, MIN:1, MAX:4095`
	PreemptMode           bool   `DESCRIPTION: "Controls whether a (starting or restarting) higher-priority Backup router preempts a lower-priority Master router", DEFAULT: true`
	AcceptMode            bool   `DESCRIPTION: "Controls whether a virtual router in Master state will accept packets addressed to the address owner's IPv4 address as its own if it is not the IPv4 address owner.", DEFAULT:false`
	AdminState            string `DESCRIPTION:"Vrrp State up or down", DEFAULT:"DOWN", SELECTION:"UP/DOWN"`
}

type VrrpV6Intf struct {
	baseObj
	IntfRef               string `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"w", MULTIPLICITY:"*", DESCRIPTION: ""Interface (name) for which VRRP Version 3 aka VRRP with ipv6 Config needs to be done"`
	VRID                  int32  `SNAPROUTE: "KEY", CATEGORY:"L3", DESCRIPTION: "Virtual Router's Unique Identifier",MIN:1, MAX:10`
	Priority              int32  `DESCRIPTION: "Sending VRRP router's priority for the virtual router", DEFAULT:100, MIN:1, MAX:255`
	Address               string `DESCRIPTION: "Virtual Router IPv6 Address", STRLEN:"43"`
	AdvertisementInterval int32  `DESCRIPTION: "Time interval between ADVERTISEMENTS", DEFAULT:1, MIN:1, MAX:4095`
	PreemptMode           bool   `DESCRIPTION: "Controls whether a (starting or restarting) higher-priority Backup router preempts a lower-priority Master router", DEFAULT: true`
	AcceptMode            bool   `DESCRIPTION: "Controls whether a virtual router in Master state will accept packets addressed to the address owner's IPv6 address as its own if it is not the IPv6 address owner.", DEFAULT:false`
	AdminState            string `DESCRIPTION:"Vrrp State up or down", DEFAULT:"DOWN", SELECTION:"UP/DOWN"`
}

type VrrpV4IntfState struct {
	baseObj
	IntfRef               string `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "Interface (name) for which VRRP Version 2 aka VRRP with ipv4 state information needs to be retreived"`
	VRID                  int32  `SNAPROUTE: "KEY", CATEGORY:"L3", DESCRIPTION: "Virtual Router's Unique Identifier"`
	OperState             string `DESCRIPTION: "Informs whether vrrp is up or down"`
	CurrentState          string `DESCRIPTION: "Current vrrp state i.e. backup or master"`
	MasterIp              string `DESCRIPTION:"Ip Address of the Master VRRP"`
	AdverRx               int32  `DESCRIPTION:"Total number of advertisement packets received"`
	AdverTx               int32  `DESCRIPTION:"Total number of advertisement packets send"`
	LastAdverRx           string `DESCRIPTION:"Time when last advertisement packet was received"`
	LastAdverTx           string `DESCRIPTION:"Time when last advertisement packet was send out"`
	IntfIpAddr            string `DESCRIPTION: "Ip Address of l3 Interface where VRRP is configured"`
	VirtualAddress        string `DESCRIPTION: "Ip Address of Virtual Router"`
	VirtualMACAddress     string `DESCRIPTION: "VRRP router's Mac Address"`
	AdvertisementInterval int32  `DESCRIPTION: "Time interval between ADVERTISEMENTS", DEFAULT:1, MIN:1, MAX:4095`
	DownTimer             int32  `DESCRIPTION: "Time interval for Backup to declare Master down"`
}

type VrrpV6IntfState struct {
	baseObj
	IntfRef               string `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "Interface (name) for which VRRP Version 3 aka VRRP with ipv4 state information needs to be retreived"`
	VRID                  int32  `SNAPROUTE: "KEY", CATEGORY:"L3", DESCRIPTION: "Virtual Router's Unique Identifier"`
	OperState             string `DESCRIPTION: "Informs whether vrrp is up or down"`
	CurrentState          string `DESCRIPTION: "Current vrrp state i.e. backup or master"`
	MasterIp              string `DESCRIPTION:"Ip Address of the Master VRRP"`
	AdverRx               int32  `DESCRIPTION:"Total number of advertisement packets received"`
	AdverTx               int32  `DESCRIPTION:"Total number of advertisement packets send"`
	LastAdverRx           string `DESCRIPTION:"Time when last advertisement packet was received"`
	LastAdverTx           string `DESCRIPTION:"Time when last advertisement packet was send out"`
	IntfIpAddr            string `DESCRIPTION: "Ipv6 Address of l3 Interface where VRRP is configured"`
	VirtualAddress        string `DESCRIPTION: "Ipv6 Address of Virtual Router"`
	VirtualMACAddress     string `DESCRIPTION: "VRRP router's Mac Address"`
	AdvertisementInterval int32  `DESCRIPTION: "Time interval between ADVERTISEMENTS", DEFAULT:1, MIN:1, MAX:4095`
	DownTimer             int32  `DESCRIPTION: "Time interval for Backup to declare Master down"`
}
