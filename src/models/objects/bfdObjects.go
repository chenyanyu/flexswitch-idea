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

import ()

/*
 * Global Config for BFD
 */
type BfdGlobal struct {
	baseObj
	Vrf    string `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"w",  MULTIPLICITY:"1", AUTOCREATE: "true", DEFAULT: "default", DESCRIPTION: "VRF id where BFD is globally enabled or disabled"`
	Enable bool   `DESCRIPTION: "Global BFD state in this VRF", DEFAULT: "true"`
}

/*
 * Global State
 */
type BfdGlobalState struct {
	baseObj
	Vrf                  string `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"r",  MULTIPLICITY:"1", DESCRIPTION: "VRF id for which global BFD state is requested"`
	Enable               bool   `DESCRIPTION: "Global BFD state in this VRF"`
	NumTotalSessions     uint32 `DESCRIPTION: "Total number of BFD sessions"`
	NumUpSessions        uint32 `DESCRIPTION: "Number of BFD sessions in up state"`
	NumDownSessions      uint32 `DESCRIPTION: "Number of BFD sessions in down state"`
	NumAdminDownSessions uint32 `DESCRIPTION: "Number of BFD sessions in admin down state"`
}

/*
 * BFD Session config
 */
type BfdSession struct {
	baseObj
	IpAddr    string `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"w",  MULTIPLICITY:"*", DESCRIPTION: "BFD neighbor IP address"`
	ParamName string `DESCRIPTION: "Name of the session parameters object to be applied on this session", DEFAULT: "default"`
	Interface string `DESCRIPTION: "Name of the interface this session has to be established on", DEFAULT: "None"`
	PerLink   bool   `DESCRIPTION: "Run BFD sessions on individual link of a LAG if the neighbor is reachable through LAG", DEFAULT: "false"`
	Owner     string `DESCRIPTION: "Module requesting BFD session configuration", DEFAULT: "user"`
}

/*
 * BFD Session state
 */
type BfdSessionState struct {
	baseObj
	IpAddr                    string `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"r",  MULTIPLICITY:"*",DESCRIPTION: "Neighbor IP address"`
	SessionId                 int32  `DESCRIPTION: "Session index"`
	ParamName                 string `DESCRIPTION: "Session parameters config"`
	IntfRef                   string `DESCRIPTION: "Interface on which this session is running"`
	InterfaceSpecific         bool   `DESCRIPTION: "This session is tied to an interface"`
	PerLinkSession            bool   `DESCRIPTION: "This is a perlink session on LAG"`
	LocalMacAddr              string `DESCRIPTION: "My MAC address"`
	RemoteMacAddr             string `DESCRIPTION: "Neighbor MAC address"`
	RegisteredProtocols       string `DESCRIPTION: "Registered owners"`
	SessionState              string `DESCRIPTION: "My state"`
	RemoteSessionState        string `DESCRIPTION: "Neighbor state"`
	LocalDiscriminator        uint32 `DESCRIPTION: "My discriminator"`
	RemoteDiscriminator       uint32 `DESCRIPTION: "Neighbor discriminator"`
	LocalDiagType             string `DESCRIPTION: "My diagnostic"`
	DesiredMinTxInterval      string `DESCRIPTION: "My desired minimum tx interval"`
	RequiredMinRxInterval     string `DESCRIPTION: "My required minimum rx interval"`
	RemoteMinRxInterval       string `DESCRIPTION: "Neighbor minimum rx interval"`
	DetectionMultiplier       uint32 `DESCRIPTION: "My detection multiplier"`
	RemoteDetectionMultiplier uint32 `DESCRIPTION: "Neighbor detection multiplier"`
	DemandMode                bool   `DESCRIPTION: "My demand mode"`
	RemoteDemandMode          bool   `DESCRIPTION: "Neighbor demand mode"`
	AuthSeqKnown              bool   `DESCRIPTION: "Authentication sequence known"`
	AuthType                  string `DESCRIPTION: "My Authentication type"`
	ReceivedAuthSeq           uint32 `DESCRIPTION: "Received authentication sequence number"`
	SentAuthSeq               uint32 `DESCRIPTION: "Sent authentication sequence number"`
	NumTxPackets              uint32 `DESCRIPTION: "Number of control packets sent"`
	NumRxPackets              uint32 `DESCRIPTION: "Number of control packets received"`
	ToDownCount               uint32 `DESCRIPTION: "Number of times this session have moved to down state"`
	ToUpCount                 uint32 `DESCRIPTION: "Number of times this session have moved to up state"`
	UpDuration                string `DESCRIPTION: "Duration of this session in up state"`
}

/*
 * BFD Session param config
 */
type BfdSessionParam struct {
	baseObj
	Name                      string `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"w",  MULTIPLICITY:"*", DESCRIPTION: "Session parameters"`
	LocalMultiplier           uint32 `DESCRIPTION: "Detection multiplier", DEFAULT: "3"`
	DesiredMinTxInterval      uint32 `DESCRIPTION: "Desired minimum tx interval in ms", DEFAULT: "1000"`
	RequiredMinRxInterval     uint32 `DESCRIPTION: "Required minimum rx interval in ms", DEFAULT: "1000"`
	RequiredMinEchoRxInterval uint32 `DESCRIPTION: "Required minimum echo rx interval in ms", DEFAULT: "0"`
	DemandEnabled             bool   `DESCRIPTION: "Enable or disable demand mode", DEFAULT: "false"`
	AuthenticationEnabled     bool   `DESCRIPTION: "Enable or disable authentication", DEFAULT: "false"`
	AuthType                  string `DESCRIPTION: "Authentication type", SELECTION: "metmd5/keyedmd5/metsha1/keyedsha1/simple", DEFAULT: "simple"`
	AuthKeyId                 uint32 `DESCRIPTION: "Authentication key id", DEFAULT: "1"`
	AuthData                  string `DESCRIPTION: "Authentication password", DEFAULT: "snaproute"`
}

/*
 * BFD Session param state
 */
type BfdSessionParamState struct {
	baseObj
	Name                      string `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"r",  MULTIPLICITY:"*", DESCRIPTION: "Session parameters"`
	NumSessions               int32  `DESCRIPTION: "Number of sessions using these params"`
	LocalMultiplier           int32  `DESCRIPTION: "Detection multiplier"`
	DesiredMinTxInterval      string `DESCRIPTION: "Desired minimum tx interval"`
	RequiredMinRxInterval     string `DESCRIPTION: "Required minimum rx interval"`
	RequiredMinEchoRxInterval string `DESCRIPTION: "Required minimum echo rx interval"`
	DemandEnabled             bool   `DESCRIPTION: "Demand mode enabled"`
	AuthenticationEnabled     bool   `DESCRIPTION: "Authentication enabled"`
	AuthenticationType        string `DESCRIPTION: "Authentication type"`
	AuthenticationKeyId       int32  `DESCRIPTION: "Authentication key id"`
	AuthenticationData        string `DESCRIPTION: "Authentication password"`
}
