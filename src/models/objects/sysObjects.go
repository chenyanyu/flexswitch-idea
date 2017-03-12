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

type ComponentLogging struct {
	baseObj
	Module string `SNAPROUTE: "KEY", CATEGORY:"System", ACCESS:"w",  MULTIPLICITY:"*", DESCRIPTION: "Module name to set logging level"`
	Level  string `DESCRIPTION: "Logging level", SELECTION: "crit/err/warn/alert/emerg/notice/info/debug/trace/off", DEFAULT: "info"`
}

type IpTableAcl struct {
	baseObj
	Name         string `SNAPROUTE: "KEY", CATEGORY:"System", ACCESS:"w", MULTIPLICITY:"*", DESCRIPTION: "Ip Table ACL rule name"`
	PhysicalPort string `DESCRIPTION: "IfIndex where the acl rule is to be applied", DEFAULT: "all"`
	Action       string `DESCRIPTION: "ACCEPT or DROP", SELECTION:"ACCEPT/DROP"`
	IpAddr       string `DESCRIPTION: "ip address of subnet or host, e.g: 192.168.1.0/24, 192.168.1.1"`
	Protocol     string `DESCRITION: "protocol for which rule is to be applied, e.g TCP, UDP", SELECTION:"tcp/udp/icmp"`
	Port         string `DESCRITION: "port for protocol, e.g for dhcprelay port is 68", DEFAULT: "all"`
}

/*
type IpTableAclState struct {
	baseObj
	Name         string `SNAPROUTE: "KEY", CATEGORY:"System", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "Ip Table ACL rule name"`
	PhysicalPort string `DESCRIPTION: "IfIndex where the acl rule is to be applied", DEFAULT: "all"`
	Action       string `DESCRIPTION: "ACCEPT or DROP"`
	IpAddr       string `DESCRIPTION: "ip address of subnet or host, e.g: 192.168.1.0/24, 192.168.1.1"`
	Protocol     string `DESCRITION: "protocol for which rule is to be applied, e.g TCP, UDP"`
	Port         string `DESCRITION: "port for protocol, e.g for dhcprelay port is 68", DEFAULT: "all"`
}
*/

type DaemonState struct {
	baseObj
	Name          string `SNAPROUTE: "KEY", CATEGORY:"System", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "Daemon name", USESTATEDB:"true"`
	Enable        bool   `DESCRIPTION: "If the daemon configured to be enabled"`
	State         string `DESCRIPTION: "State of the daemon, up, starting, restarting or stopped"`
	Reason        string `DESCRIPTION: "Reason for current state of the daemon"`
	StartTime     string `DESCRIPTION: "Daemon start time"`
	KeepAlive     string `DESCRIPTION: "KeepAlive state of the daemon"`
	RestartCount  int32  `DESCRIPTION: "Number of times this daemon has been restarted"`
	RestartTime   string `DESCRIPTION: "Last restart time"`
	RestartReason string `DESCRIPTION: "Last restart reason"`
}

type SystemParam struct {
	baseObj
	Vrf         string `SNAPROUTE: "KEY", CATEGORY:"System", ACCESS:"w", MULTIPLICITY:"*", AUTOCREATE: "true", DESCRIPTION: "System Vrf", DEFAULT:"default"`
	MgmtIp      string `DESCRIPTION: "Management Ip of System"`
	Hostname    string `DESCRIPTION: "System Host Name"`
	SwitchMac   string `DESCRIPTION: "Switch Mac Address`
	SwVersion   string `DESCRIPTION: "FlexSwitch Version Information"`
	Description string `DESCRIPTION: "System Description"`
}

type SystemParamState struct {
	baseObj
	Vrf         string `SNAPROUTE: "KEY", CATEGORY:"System", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "System Vrf""`
	MgmtIp      string `DESCRIPTION: "Management Ip of System"`
	Hostname    string `DESCRIPTION: "System Host Name"`
	SwitchMac   string `DESCRIPTION: "Switch Mac Address`
	SwVersion   string `DESCRIPTION: "FlexSwitch Version Information"`
	Description string `DESCRIPTION: "System Description"`
	Distro      string `DESCRIPTION: "Linux distro running on this system"`
	Kernel      string `DESCRIPTION: "Kernel version running on this system"`
}

type TacacsConfig struct {
	baseObj
	ServerIp       string `SNAPROUTE: "KEY", CATEGORY:"System", ACCESS:"w", MULTIPLICITY:"*", DESCRIPTION: "TACACS Server IP"`
	SourceIntf     string `DESCRIPTION: "Interface on which the source IP is configured", DEFAULT: "ma1"`
	AuthService    string `DESCRIPTION: "Service requesting auth to grant user different privileges", DEFAULT:"ppp", SELECTION:"login/enable/ppp/arap/pt/rcmd/x25/nasi/fwproxy"`
	Secret         string `DESCRIPTION: "TACACS Server Secret"`
	Port           int16 `DESCRIPTION: "TCP port of TACACS+ Server", DEFAULT:49`
	PrivilegeLevel int32  `DESCRIPTION: "Privilege level for the cli user", MIN:1, MAX:15, DEFAULT:15`
	Debug          int32  `DESCRIPTION: "Debug level for the cli user", SELECTION:0/1, DEFAULT:0`
}

type TacacsState struct {
	baseObj
	ServerIp       string `SNAPROUTE: "KEY", CATEGORY:"System", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "TACACS Server IP"`
	SourceIntf     string `DESCRIPTION: "Interface on which the source IP is configured"`
	AuthService    string `DESCRIPTION: "Service requesting auth to grant user different privileges"`
	Secret         string `DESCRIPTION: "TACACS Server Secret"`
	Port           int16 `DESCRIPTION: "TCP port of TACACS+ Server"`
	PrivilegeLevel int32  `DESCRIPTION: "Privilege level for the cli user"`
	Debug          int32  `DESCRIPTION: "Debug level for the cli user, used to display TACAC's debug information"`
	ConnFailReason string `DESCRIPTION: "Connection to server failure reason`
}

type TacacsGlobalConfig struct {
	baseObj
	ProfileName string `SNAPROUTE: "KEY", CATEGORY:"System", ACCESS:"w", MULTIPLICITY:"1", AUTOCREATE: "true", DESCRIPTION: "System Profile Name", DEFAULT:"default"`
	Enable      string `DESCRIPTION: "Enable TACACS Servers", DEFAULT: "false", SELECTION: "true/false"`
	Timeout     int32  `DESCRIPTION: TCP timeout value used to determine if a connection to server is dead. DEFAULT:10`
}

type TacacsGlobalState struct {
	baseObj
	ProfileName       string `SNAPROUTE: "KEY", CATEGORY:"System", ACCESS:"r", MULTIPLICITY:"1", DESCRIPTION: "System Profile Name"`
	OperStatus        string `DESCRIPTION: "Operational Status of TACACS Servers"`
	NumActiveSessions int32  `DESCRIPTION: Number of active sessions`
}
