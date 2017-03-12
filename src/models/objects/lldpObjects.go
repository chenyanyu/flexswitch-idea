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

type LLDPGlobal struct {
	baseObj
	Vrf             string `SNAPROUTE: "KEY", CATEGORY:"L2", ACCESS:"w", MULTIPLICITY:"1", DESCRIPTION: "LLDP Global Config For Default VRF", DEFAULT:"default", AUTOCREATE:"true"`
	Enable          bool   `DESCRIPTION: "Enable/Disable LLDP Globally", DEFAULT:false`
	TranmitInterval int32  `DESCRIPTION: "LLDP Re-Transmit Interval in seconds", DEFAULT:30`
	TxRxMode        string `DESCRIPTION: "Transmit/Receive mode configruration for the LLDP agent", SELECTION:"TxOnly/RxOnly/TxRx", DEFAULT:"TxRx"`
	SnoopAndDrop    bool   `DESCRIPTION: "Operational mode to determine whether LLDP traffic is bi-directionally forwarded. This configuration is only available on select platforms", DEFAULT:false`
}

type LLDPGlobalState struct {
	baseObj
	Vrf             string `SNAPROUTE: "KEY", CATEGORY:"L2", ACCESS:"r", MULTIPLICITY:"1", DESCRIPTION:"Vrf where LLDP Global Config is running"`
	Enable          bool   `DESCRIPTION: "Enable/Disable LLDP Globally"`
	TranmitInterval int32  `DESCRIPTION: "LLDP Re-Transmit Interval in seconds"`
	Neighbors       int32  `DESCRIPTION: "Total lldp Neighbors learned on the system"`
	TotalTxFrames   int32  `DESCRIPTION: "Total no.of lldp frames send out by the system"`
	TotalRxFrames   int32  `DESCRIPTION: "Total no.of lldp frames received by the system"`
}

type LLDPIntf struct {
	baseObj
	IntfRef  string `SNAPROUTE: "KEY", CATEGORY:"L2", ACCESS:"rw", MULTIPLICITY:"*", DESCRIPTION: "IfIndex where lldp needs is enabled/disabled", DEFAULT: "None", AUTODISCOVER:"true"`
	Enable   bool   `DESCRIPTION: "Enable/Disable lldp config Per Port", DEFAULT:true`
	TxRxMode string `DESCRIPTION: "Transmit/Receive mode configruration for the LLDP agent specific to an interface", SELECTION:"TxOnly/RxOnly/TxRx", DEFAULT:"TxRx"`
}

type LLDPIntfState struct {
	baseObj
	IntfRef             string `SNAPROUTE: "KEY", CATEGORY:"L2", ACCESS:"r", MULTIPLICITY:"*",DESCRIPTION: "IntfRef where lldp is configured"`
	IfIndex             int32  `DESCRIPTION: "IfIndex where lldp needs to be configured"`
	SendFrames          int32  `DESCRIPTION: "Total Frames send to the neighbor"`
	ReceivedFrames      int32  `DESCRIPTION: "Total Frames received from neighbor"`
	Enable              bool   `DESCRIPTION: "Enable/Disable lldp config"`
	LocalPort           string `DESCRIPTION: "Local interface"`
	PeerMac             string `DESCRIPTION: "Mac address of direct connection"`
	PeerPort            string `DESCRIPTION: "Name of directtly connected pors"`
	PeerHostName        string `DESCRIPTION: "Name of the peer host`
	PortDescription     string `DESCRIPTION: "peer port description"`
	VlanId              int32  `DESCRIPTION: "Port Vlan Identifier"`
	HoldTime            string `DESCRIPTION: "Validity of the peer information"`
	SystemDescription   string `DESCRIPTION: "System Description of the peer port"`
	SystemCapabilities  string `DESCRIPTION: "System Capabilities of the peer port"`
	EnabledCapabilities string `DESCRIPTION: "Enabled Capabilities of the peer port"`
}
