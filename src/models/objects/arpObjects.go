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

/*
 * This DS will be used while Created/Deleting Arp Config
 */
type ArpGlobal struct {
	baseObj
	// placeholder to create a key
	Vrf     string `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"w", MULTIPLICITY:"1", AUTOCREATE: "true", DESCRIPTION: "System Vrf", DEFAULT:"default"`
	Timeout int32  `DESCRIPTION: "Global Arp entry timeout value. Default value: 600 seconds, Minimum Possible Value: 300 seconds, Unit: second", MIN:300, MAX:1500, DEFAULT: "600"`
}

type ArpEntryState struct {
	baseObj
	IpAddr         string `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "Neighbor's IP Address"`
	MacAddr        string `DESCRIPTION: "MAC address of the neighbor machine with corresponding IP Address"`
	Vlan           string `DESCRIPTION: "Vlan ID of the Router Interface to which neighbor is attached to"`
	Intf           string `DESCRIPTION: "Router Interface to which neighbor is attached to"`
	ExpiryTimeLeft string `DESCRIPTION: "Time left before entry expires in case neighbor departs"`
}

type ArpLinuxEntryState struct {
	baseObj
	IpAddr  string `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "Neighbor's IP Address"`
	HWType  string `DESCRIPTION: "Hardware Type"`
	MacAddr string `DESCRIPTION: "MAC address of neighbor"`
	IfName  string `DESCRIPTION: "Interface name"`
}
