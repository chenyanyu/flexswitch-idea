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
 * This DS will be used while Created/Deleting DHCP Config
 */
type DhcpGlobalConfig struct {
	baseObj
	// placeholder to create a key
	DhcpConfigKey    string `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"w", MULTIPLICITY:"1", DESCRIPTION: "DHCP global config", DEFAULT:"default"`
	Enable           bool   `DESCRIPTION: "DHCP Server enable/disable control DEFAULT: false"`
	DefaultLeaseTime uint32 `DESCRIPTION: "Default Lease Time in seconds DEFAULT: 600"`
	MaxLeaseTime     uint32 `DESCRIPTION: "Max Lease Time in seconds DEFAULT: 7200"`
}

type DhcpIntfConfig struct {
	baseObj
	// placeholder to create a key
	IntfRef       string `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"w", MULTIPLICITY:"*", DESCRIPTION: "Interface name or ifindex of L3 interface object on which Dhcp Server need to be configured"`
	Subnet        string `DESCRIPTION: "Subnet"`
	SubnetMask    string `DESCRIPTION: "Subnet Mask"`
	IPAddrRange   string `DESCRIPTION: "Range of IP Addresses DEFAULT: All the IP Addresses in the given subnet"`
	BroadcastAddr string `DESCRIPTION: "Broadcast Address DEFAULT: Last IP Addresses in the given subnet"`
	RouterAddr    string `DESCRIPTION: "Router Address DEFAULT: First IP Addresses in the given subnet"`
	DNSServerAddr string `DESCRIPTION: "Comma seperated List of DNS Server Address DEFAULT: none"`
	DomainName    string `DESCRIPTION: "Domain Name Address DEFAULT: none"`
	Enable        bool   `DESCRIPTION: "Enable and Disable Control DEFAULT: false"`
}
