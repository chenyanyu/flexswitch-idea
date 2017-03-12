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

type SystemStatusState struct {
	baseObj
	Name           string        `SNAPROUTE: "KEY", CATEGORY:"System", ACCESS:"r",  MULTIPLICITY:"1", DESCRIPTION: "Name of the system"`
	Ready          bool          `DESCRIPTION: "System is ready to accept api calls"`
	Reason         string        `DESCRIPTION: "Reason if system not ready"`
	UpTime         string        `DESCRIPTION: "Uptime of this system"`
	NumCreateCalls string        `DESCRIPTION: "Number of create api calls made"`
	NumDeleteCalls string        `DESCRIPTION: "Number of delete api calls made"`
	NumUpdateCalls string        `DESCRIPTION: "Number of update api calls made"`
	NumGetCalls    string        `DESCRIPTION: "Number of get api calls made"`
	NumActionCalls string        `DESCRIPTION: "Number of action api calls made"`
	FlexDaemons    []DaemonState `DESCRIPTION: "Daemon states"`
}

type RepoInfo struct {
	Name   string `DESCRIPTION: "Name of the git repo"`
	Sha1   string `DESCRIPTION: "Git commit Sha1"`
	Branch string `DESCRIPTION: "Branch name"`
	Time   string `DESCRIPTION: "Build time"`
}

type SystemSwVersionState struct {
	baseObj
	FlexswitchVersion string     `SNAPROUTE: "KEY", CATEGORY:"System", ACCESS:"r",  MULTIPLICITY:"1", DESCRIPTION: "Flexswitch version"`
	Repos             []RepoInfo `DESCRIPTION: "Git repo details"`
}

type ConfigLogState struct {
	baseObj
	SeqNum    uint32 `SNAPROUTE: "KEY", CATEGORY:"System", DESCRIPTION: "Sequence number of the API call"`
	Time      string `SNAPROUTE: "KEY", CATEGORY:"System", ACCESS:"r", MULTIPLICITY: "*", DESCRIPTION: "When the API was called", USESTATEDB:"true"`
	API       string `SNAPROUTE: "KEY", CATEGORY:"System", DESCRIPTION: "Name of the API called"`
	Operation string `DESCRIPTION: "Oprtation executed on this API"`
	Data      string `DESCRIPTION: "User provided data"`
	Result    string `DESCRIPTION: "Result of the API call"`
	UserAddr  string `DESCRIPTION: "Host address from where the call was made"`
	UserName  string `DESCRIPTION: "User who made the call"`
}

//Voyager specific config object
type XponderGlobal struct {
	baseObj
	XponderId          uint8  `SNAPROUTE: "KEY", CATEGORY:"System", ACCESS:"w", MULTIPLICITY: "1", DESCRIPTION: "Xponder module identifier", DEFAULT:0, AUTOCREATE:"true"`
	XponderMode        string `DESCRIPTION: "Global operational mode of Xponder module", SELECTION:"InServiceWire"/"InServiceRegen"/"InServiceOverSub"/"InServicePacketOptical"/"OutOfService", DEFAULT:"OutOfService"`
	XponderDescription string `DESCRIPTION: "User configurable description string for the xponder module", DEFAULT:"This is a Voyager platform"`
}

type XponderGlobalState struct {
	baseObj
	XponderId          uint8  `SNAPROUTE: "KEY", CATEGORY:"System", ACCESS:"r", MULTIPLICITY: "1", DESCRIPTION: "Xponder module identifier"`
	XponderMode        string `DESCRIPTION: "Global operational mode of Xponder module"`
	XponderDescription string `DESCRIPTION: "User configurable description string for the xponder module"`
}

type ApiInfoState struct {
	baseObj
	Url  string   `SNAPROUTE: "KEY", CATEGORY:"System", ACCESS:"r", MULTIPLICITY: "1", DESCRIPTION: "URL"`
	Info []string `DESCRIPTION: "APIs available under this URL or details of the specific API"`
}
