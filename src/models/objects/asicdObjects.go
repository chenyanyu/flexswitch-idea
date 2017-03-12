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

type AsicGlobalState struct {
	baseObj
	ModuleId   uint8   `SNAPROUTE: "KEY", CATEGORY:"System", ACCESS:"r", MULTIPLICITY: "1", DESCRIPTION:"Module identifier"`
	VendorId   string  `DESCRIPTION: "Vendor identification value"`
	PartNumber string  `DESCRIPTION: "Part number of underlying switching asic"`
	RevisionId string  `DESCRIPTION: "Revision ID of underlying switching asic"`
	ModuleTemp float64 `DESCRIPTION: "Current module temperature", UNIT: degC`
}

type PMData struct {
	TimeStamp string  `DESCRIPTION: "Timestamp at which data is collected"`
	Value     float64 `DESCRIPTION: "PM Data Value"`
}

type AsicGlobalPM struct {
	baseObj
	ModuleId           uint8   `SNAPROUTE: "KEY", CATEGORY:"Performance", ACCESS:"rw", MULTIPLICITY: "1", AUTODISCOVER:"true", DESCRIPTION:"Module identifier, DEFAULT: 0"`
	Resource           string  `SNAPROUTE: "KEY", CATEGORY:"Performance", DESCRIPTION: "Resource identifier", SELECTION: "Temperature"`
	PMClassAEnable     bool    `DESCRIPTION: "Enable/Disable control for CLASS-A PM", DEFAULT:true`
	PMClassBEnable     bool    `DESCRIPTION: "Enable/Disable control for CLASS-B PM", DEFAULT:true`
	PMClassCEnable     bool    `DESCRIPTION: "Enable/Disable control for CLASS-C PM", DEFAULT:true`
	HighAlarmThreshold float64 `DESCRIPTION: "High alarm threshold value for this PM", DEFAULT: 100000`
	HighWarnThreshold  float64 `DESCRIPTION: "High warning threshold value for this PM", DEFAULT: 100000`
	LowAlarmThreshold  float64 `DESCRIPTION: "Low alarm threshold value for this PM", DEFAULT: -100000`
	LowWarnThreshold   float64 `DESCRIPTION: "Low warning threshold value for this PM", DEFAULT: -100000`
}

type AsicGlobalPMState struct {
	baseObj
	ModuleId     uint8    `SNAPROUTE: "KEY", CATEGORY:"Performance", ACCESS:"r", MULTIPLICITY: "1", DESCRIPTION:"Module identifier"`
	Resource     string   `SNAPROUTE: "KEY", CATEGORY:"Performance", DESCRIPTION: "Resource identifier"`
	ClassAPMData []PMData `DESCRIPTION: "PM Data corresponding to PM Class A"`
	ClassBPMData []PMData `DESCRIPTION: "PM Data corresponding to PM Class B"`
	ClassCPMData []PMData `DESCRIPTION: "PM Data corresponding to PM Class C"`
}

type EthernetPM struct {
	baseObj
	IntfRef            string  `SNAPROUTE: "KEY", CATEGORY:"Performance", ACCESS:"rw", MULTIPLICITY: "*", AUTODISCOVER:"true", DESCRIPTION: "Interface name of port"`
	Resource           string  `SNAPROUTE: "KEY", CATEGORY:"Performance", DESCRIPTION: "Resource identifier", SELECTION:"StatUnderSizePkts/StatOverSizePkts/StatFragments/StatCRCAlignErrors/StatJabber/StatEtherPkts/StatMCPkts/StatBCPkts/Stat64OctOrLess/Stat65OctTo126Oct/Stat128OctTo255Oct/Stat128OctTo255Oct/Stat256OctTo511Oct/Stat512OctTo1023Oct/Statc1024OctTo1518Oct"`
	PMClassAEnable     bool    `DESCRIPTION: "Enable/Disable control for CLASS-A PM", DEFAULT:true`
	PMClassBEnable     bool    `DESCRIPTION: "Enable/Disable control for CLASS-B PM", DEFAULT:true`
	PMClassCEnable     bool    `DESCRIPTION: "Enable/Disable control for CLASS-C PM", DEFAULT:true`
	HighAlarmThreshold float64 `DESCRIPTION: "High alarm threshold value for this PM", DEFAULT: 100000`
	HighWarnThreshold  float64 `DESCRIPTION: "High warning threshold value for this PM", DEFAULT: 100000`
	LowAlarmThreshold  float64 `DESCRIPTION: "Low alarm threshold value for this PM", DEFAULT: -100000`
	LowWarnThreshold   float64 `DESCRIPTION: "Low warning threshold value for this PM", DEFAULT: -100000`
}

type EthernetPMState struct {
	baseObj
	IntfRef      string   `SNAPROUTE: "KEY", CATEGORY:"Performance", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION:"Interface name of port"`
	Resource     string   `SNAPROUTE: "KEY", CATEGORY:"Performance", DESCRIPTION: "Resource identifier"`
	ClassAPMData []PMData `DESCRIPTION: "PM Data corresponding to PM Class A"`
	ClassBPMData []PMData `DESCRIPTION: "PM Data corresponding to PM Class B"`
	ClassCPMData []PMData `DESCRIPTION: "PM Data corresponding to PM Class C"`
}

type AsicSummaryState struct {
	baseObj
	ModuleId      uint8 `SNAPROUTE: "KEY", CATEGORY:"System", ACCESS:"r", MULTIPLICITY: "1", DESCRIPTION:"Module identifier"`
	NumPortsUp    int32 `DESCRIPTION: Summary stating number of ports that have operstate UP`
	NumPortsDown  int32 `DESCRIPTION: Summary stating number of ports that have operstate DOWN`
	NumVlans      int32 `DESCRIPTION: Summary stating number of vlans configured in the asic`
	NumV4Intfs    int32 `DESCRIPTION: Summary stating number of IPv4 interfaces configured in the asic`
	NumV6Intfs    int32 `DESCRIPTION: Summary stating number of IPv6 interfaces configured in the asic`
	NumV4Adjs     int32 `DESCRIPTION: Summary stating number of IPv4 adjacencies configured in the asic`
	NumV6Adjs     int32 `DESCRIPTION: Summary stating number of IPv6 adjacencies configured in the asic`
	NumV4Routes   int32 `DESCRIPTION: Summary stating number of IPv4 routes configured in the asic`
	NumV6Routes   int32 `DESCRIPTION: Summary stating number of IPv6 routes configured in the asic`
	NumECMPRoutes int32 `DESCRIPTION: Summary stating number of ECMP routes configured in the asic`
}

type Vlan struct {
	baseObj
	VlanId        int32    `SNAPROUTE: "KEY", CATEGORY:"L2", ACCESS:"w", MULTIPLICITY: "*", MIN:"1", MAX: "4094", DESCRIPTION: "802.1Q tag/Vlan ID for vlan being provisioned"`
	AdminState    string   `DESCRIPTION: "Administrative state of this vlan interface", SELECTION:"UP/DOWN", DEFAULT:"UP"`
	Description   string   `DESCRIPTION: "Description about the vlan interface", DEFAULT:"none"`
	AutoState     string   `DESCRIPTION: Auto State of this vlan interface", SELECTION:"UP/DOWN", DEFAULT:"UP"`
	IntfList      []string `DESCRIPTION: "List of interface names or ifindex values to  be added as tagged members of the vlan"`
	UntagIntfList []string `DESCRIPTION: "List of interface names or ifindex values to  be added as untagged members of the vlan"`
}

type VlanState struct {
	baseObj
	VlanId                 int32  `SNAPROUTE: "KEY", CATEGORY:"L2", ACCESS:"r", MULTIPLICITY: "*", DESCRIPTION: "802.1Q tag/Vlan ID for vlan being provisioned"`
	VlanName               string `DESCRIPTION: "System assigned vlan name"`
	OperState              string `DESCRIPTION: "Operational state of vlan interface"`
	IfIndex                int32  `DESCRIPTION: "System assigned interface id for this vlan interface"`
	SysInternalDescription string `DESCRIPTION: "This is a system generated string that explains the operstate value"`
}

type IPv4Intf struct {
	baseObj
	IntfRef    string `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"w", DESCRIPTION: "Interface name or ifindex of port/lag or vlan on which this IPv4 object is configured", RELTN:"DEP:[Vlan, Port]`
	IpAddr     string `DESCRIPTION: "Interface IP/Net mask in CIDR format to provision on switch interface", STRLEN:"18"`
	AdminState string `DESCRIPTION: "Administrative state of this IP interface", SELECTION:"UP/DOWN", DEFAULT:"UP"`
}

type IPv4IntfState struct {
	baseObj
	IntfRef           string `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"r", DESCRIPTION: "System assigned interface id of L2 interface (port/lag/vlan) to which this IPv4 object is linked"`
	IfIndex           int32  `DESCRIPTION: "System assigned interface id for this IPv4 interface"`
	IpAddr            string `DESCRIPTION: "Interface IP/Net mask in CIDR format to provision on switch interface"`
	OperState         string `DESCRIPTION: "Operational state of this IP interface"`
	NumUpEvents       int32  `DESCRIPTION: "Number of times the operational state transitioned from DOWN to UP"`
	LastUpEventTime   string `DESCRIPTION: "Timestamp corresponding to the last DOWN to UP operational state change event"`
	NumDownEvents     int32  `DESCRIPTION: "Number of times the operational state transitioned from UP to DOWN"`
	LastDownEventTime string `DESCRIPTION: "Timestamp corresponding to the last UP to DOWN operational state change event"`
	L2IntfType        string `DESCRIPTION: "Type of L2 interface on which IP has been configured (Port/Lag/Vlan)"`
	L2IntfId          int32  `DESCRIPTION: "Id of the L2 interface. Port number/lag id/vlan id."`
}

type Port struct {
	baseObj
	IntfRef        string `SNAPROUTE: "KEY", CATEGORY:"Physical", ACCESS:"rw", MULTIPLICITY:"*", AUTODISCOVER:"true", DESCRIPTION: "Front panel port name or system assigned interface id"`
	IfIndex        int32  `DESCRIPTION: "System assigned interface id for this port. Read only attribute"`
	Description    string `DESCRIPTION: "User provided string description", DEFAULT:"FP Port", STRLEN:"64"`
	PhyIntfType    string `DESCRIPTION: "Type of internal phy interface", STRLEN:"16" SELECTION:"GMII/SGMII/QSMII/SFI/XFI/XAUI/XLAUI/RXAUI/CR/CR2/CR4/KR/KR2/KR4/SR/SR2/SR4/SR10/LR/LR4"`
	AdminState     string `DESCRIPTION: "Administrative state of this port", STRLEN:"4" SELECTION:"UP/DOWN", DEFAULT:"DOWN"`
	MacAddr        string `DESCRIPTION: "Mac address associated with this port", STRLEN:"17"`
	Speed          int32  `DESCRIPTION: "Port speed in Mbps", MIN: 10, MAX: "100000"`
	Duplex         string `DESCRIPTION: "Duplex setting for this port", STRLEN:"16" SELECTION:"Half_Duplex/Full_Duplex", DEFAULT:"Full_Duplex"`
	Autoneg        string `DESCRIPTION: "Autonegotiation setting for this port", STRLEN:"4" SELECTION:"ON/OFF", DEFAULT:"OFF"`
	MediaType      string `DESCRIPTION: "Type of media inserted into this port", STRLEN:"16"`
	Mtu            int32  `DESCRIPTION: "Maximum transmission unit size for this port"`
	BreakOutMode   string `DESCRIPTION: "Break out mode for the port. Only applicable on ports that support breakout.", STRLEN:"6" SELECTION:"1x100/1x40/2x50/4x25/4x10"`
	LoopbackMode   string `DESCRIPTION: "Desired loopback setting for this port", SELECTION:"NONE/MAC/PHY/RMT", DEFAULT:"NONE"`
	EnableFEC      bool   `DESCRIPTION: "Enable/Disable 802.3bj FEC on this interface", DEFAULT: false`
	PRBSTxEnable   bool   `DESCRIPTION: "Enable/Disable generation of PRBS on this port", DEFAULT: false`
	PRBSRxEnable   bool   `DESCRIPTION: "Enable/Disable PRBS checker on this port", DEFAULT: false`
	PRBSPolynomial string `DESCRIPTION: "PRBS polynomial to use for generation/checking", DEFAULT:2^7, SELECTION:"2^7/2^23/2^31"`
}

type PortState struct {
	baseObj
	IntfRef                     string `SNAPROUTE: "KEY", CATEGORY:"Physical", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "Front panel port name or system assigned interface id"`
	IfIndex                     int32  `DESCRIPTION: "System assigned interface id for this port"`
	Name                        string `DESCRIPTION: "System assigned vlan name"`
	OperState                   string `DESCRIPTION: "Operational state of front panel port"`
	NumUpEvents                 int32  `DESCRIPTION: "Number of times the operational state transitioned from DOWN to UP"`
	LastUpEventTime             string `DESCRIPTION: "Timestamp corresponding to the last DOWN to UP operational state change event"`
	NumDownEvents               int32  `DESCRIPTION: "Number of times the operational state transitioned from UP to DOWN"`
	LastDownEventTime           string `DESCRIPTION: "Timestamp corresponding to the last UP to DOWN operational state change event"`
	Pvid                        int32  `DESCRIPTION: "The vlanid assigned to untagged traffic ingressing this port"`
	IfInOctets                  int64  `DESCRIPTION: "RFC2233 Total number of octets received on this port"`
	IfInUcastPkts               int64  `DESCRIPTION: "RFC2233 Total number of unicast packets received on this port"`
	IfInDiscards                int64  `DESCRIPTION: "RFC2233 Total number of inbound packets that were discarded"`
	IfInErrors                  int64  `DESCRIPTION: "RFC2233 Total number of inbound packets that contained an error"`
	IfInUnknownProtos           int64  `DESCRIPTION: "RFC2233 Total number of inbound packets discarded due to unknown protocol"`
	IfOutOctets                 int64  `DESCRIPTION: "RFC2233 Total number of octets transmitted on this port"`
	IfOutUcastPkts              int64  `DESCRIPTION: "RFC2233 Total number of unicast packets transmitted on this port"`
	IfOutDiscards               int64  `DESCRIPTION: "RFC2233 Total number of error free packets discarded and not transmitted"`
	IfOutErrors                 int64  `DESCRIPTION: "RFC2233 Total number of packets discarded and not transmitted due to packet errors"`
	IfEtherUnderSizePktCnt      int64  `DESCRIPTION: "RFC 1757 Total numbe of undersized packets received and transmitted"`
	IfEtherOverSizePktCnt       int64  `DESCRIPTION: "RFC 1757 Total number of oversized packets received and transmitted"`
	IfEtherFragments            int64  `DESCRIPTION: "RFC1757 Total number of ethernet fragments received and transmitted"`
	IfEtherCRCAlignError        int64  `DESCRIPTION: "RFC 1757 Total number of CRC alignment errors"`
	IfEtherJabber               int64  `DESCRIPTION: "RFC 1757 Total number of jabber frames received and transmitted"`
	IfEtherPkts                 int64  `DESCRIPTION: "RFC 1757 Total number of ethernet packets received and transmitted"`
	IfEtherMCPkts               int64  `DESCRIPTION: "RFC 1757 Total number of multicast packets received and transmitted"`
	IfEtherBcastPkts            int64  `DESCRIPTION: "RFC 1757 Total number of ethernet broadcast packets received and transmitted"`
	IfEtherPkts64OrLessOctets   int64  `DESCRIPTION: "RFC1757 Total number of ethernet packets sized 64 bytes or lesser"`
	IfEtherPkts65To127Octets    int64  `DESCRIPTION: "RFC 1757 Total number of ethernet packets sized between 65 and 127 bytes"`
	IfEtherPkts128To255Octets   int64  `DESCRIPTION: "RFC 1757 Total number of ethernet packets sized between 128 and 255 bytes"`
	IfEtherPkts256To511Octets   int64  `DESCRIPTION: "RFC 1757 Total number of ethernet packets sized between 256 and 511 bytes"`
	IfEtherPkts512To1023Octets  int64  `DESCRIPTION: "RFC 1757 Total number of ethernet packets sized between 512 and 1023 bytes"`
	IfEtherPkts1024To1518Octets int64  `DESCRIPTION: "RFC 1757 Total number of ethernet packets sized between 1024 and 1518 bytes"`
	ErrDisableReason            string `DESCRIPTION: "Reason explaining why port has been disabled by protocol code"`
	PresentInHW                 string `DESCRIPTION: "Indication of whether this port object maps to a physical port. Set to 'No' for ports that are not broken out."`
	ConfigMode                  string `DESCRIPTION: "The current mode of configuration on this port (L2/L3/Internal)"`
	PRBSRxErrCnt                int64  `DESCRIPTION: "Receive error count reported by PRBS checker"`
}

type MacTableEntryState struct {
	baseObj
	MacAddr string `SNAPROUTE: "KEY", CATEGORY:"L2", ACCESS:"r", DESCRIPTION: "MAC Address", USESTATEDB:"true"`
	VlanId  int32  `DESCRIPTION: "Vlan id corresponding to which mac was learned", DEFAULT:0`
	Port    int32  `DESCRIPTION: "Port number on which mac was learned", DEFAULT:0`
}

type IPv4RouteHwState struct {
	baseObj
	DestinationNw    string `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "IP address of the route in CIDR format"`
	NextHopIps       string `DESCRIPTION: "next hop ip list for the route"`
	RouteCreatedTime string `DESCRIPTION :"Time when the route was added"`
	RouteUpdatedTime string `DESCRIPTION :"Time when the route was last updated"`
}

type IPv6RouteHwState struct {
	baseObj
	DestinationNw    string `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "IP address of the route in CIDR format"`
	NextHopIps       string `DESCRIPTION: "next hop ip list for the route"`
	RouteCreatedTime string `DESCRIPTION :"Time when the route was added"`
	RouteUpdatedTime string `DESCRIPTION :"Time when the route was last updated"`
}

type ArpEntryHwState struct {
	baseObj
	IpAddr  string `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"r", MULTIPLICITY:"*", QPARAM: "optional" ,DESCRIPTION: "Neighbor's IP Address"`
	MacAddr string `DESCRIPTION: "MAC address of the neighbor machine with corresponding IP Address", QPARAM: "optional" `
	Vlan    string `DESCRIPTION: "Vlan ID of the Router Interface to which neighbor is attached to", QPARAM: "optional" `
	Port    string `DESCRIPTION: "Router Interface to which neighbor is attached to", QPARAM: "optional" `
}

type NdpEntryHwState struct {
	baseObj
	IpAddr  string `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"r", MULTIPLICITY:"*", QPARAM: "optional" ,DESCRIPTION: "Neighbor's IP Address"`
	MacAddr string `DESCRIPTION: "MAC address of the neighbor machine with corresponding IP Address", QPARAM: "optional" `
	Vlan    string `DESCRIPTION: "Vlan ID of the Router Interface to which neighbor is attached to", QPARAM: "optional" `
	Port    string `DESCRIPTION: "Router Interface to which neighbor is attached to", QPARAM: "optional" `
}

type LogicalIntf struct {
	baseObj
	Name string `SNAPROUTE: "KEY", CATEGORY:"Physical", ACCESS:"w", DESCRIPTION: "Name of logical interface"`
	Type string `DESCRIPTION: "Type of logical interface (e.x. loopback)", SELECTION:"Loopback", DEFAULT:"Loopback", STRLEN:"16"`
}

type LogicalIntfState struct {
	baseObj
	Name              string `SNAPROUTE: "KEY", CATEGORY:"Physical", ACCESS:"r", DESCRIPTION: "Name of logical interface"`
	IfIndex           int32  `DESCRIPTION: "System assigned interface id for this logical interface"`
	SrcMac            string `DESCRIPTION: "Source Mac assigned to the interface"`
	OperState         string `DESCRIPTION: "Operational state of logical interface"`
	IfInOctets        int64  `DESCRIPTION: "RFC2233 Total number of octets received on this port"`
	IfInUcastPkts     int64  `DESCRIPTION: "RFC2233 Total number of unicast packets received on this port"`
	IfInDiscards      int64  `DESCRIPTION: "RFC2233 Total number of inbound packets that were discarded"`
	IfInErrors        int64  `DESCRIPTION: "RFC2233 Total number of inbound packets that contained an error"`
	IfInUnknownProtos int64  `DESCRIPTION: "RFC2233 Total number of inbound packets discarded due to unknown protocol"`
	IfOutOctets       int64  `DESCRIPTION: "RFC2233 Total number of octets transmitted on this port"`
	IfOutUcastPkts    int64  `DESCRIPTION: "RFC2233 Total number of unicast packets transmitted on this port"`
	IfOutDiscards     int64  `DESCRIPTION: "RFC2233 Total number of error free packets discarded and not transmitted"`
	IfOutErrors       int64  `DESCRIPTION: "RFC2233 Total number of packets discarded and not transmitted due to packet errors"`
}

type SubIPv4Intf struct {
	baseObj
	IntfRef string `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"w", DESCRIPTION:"Intf name for which ipv4Intf sub interface is to be configured"`
	Type    string `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"w", DESCRIPTION:"Type of interface, e.g. Secondary", STRLEN:"16", SELECTION: "Secondary`
	IpAddr  string `DESCRIPTION:"Ip Address for sub interface", STRLEN:"18"`
	MacAddr string `DESCRIPTION:"Mac address to be used for the sub interface. If none specified IPv4Intf mac address will be used", STRLEN:"17", DEFAULT:""`
	Enable  bool   `DESCRIPTION:"Enable or disable this interface", DEFAULT:true`
}

type SubIPv4IntfState struct {
	baseObj
	IntfRef       string `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"r", DESCRIPTION:"Intf name for which ipv4Intf sub interface is to be configured"`
	Type          string `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"r", DESCRIPTION:"Type of interface, e.g. Secondary or Virtual"`
	IfIndex       int32  `DESCRIPTION:"System assigned interface id for this sub IPv4 interface"`
	IfName        string `DESCRIPTION:"System generated sub interface name"`
	ParentIfIndex int32  `DESCRIPTION:"System assigned interface id for interface parent interface"`
	IpAddr        string `DESCRIPTION:"Ip Address for sub interface"`
	MacAddr       string `DESCRIPTION:"Mac address to be used for the sub interface. If none specified IPv4Intf mac address will be used"`
	OperState     string `DESCRIPTION:"Operational state of this SubIPv4 interface"`
}

type IPv6Intf struct {
	baseObj
	IntfRef    string `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"w", DESCRIPTION: "Interface name or ifindex of port/lag or vlan on which this IPv4 object is configured", RELTN:"DEP:[Vlan, Port]`
	IpAddr     string `DESCRIPTION: "Interface Global Scope IP Address/Prefix-Length to provision on switch interface", STRLEN:"43", DEFAULT:""`
	LinkIp     bool   `DESCRIPTION: "Interface Link Scope IP Address auto-configured", DEFAULT:true`
	AdminState string `DESCRIPTION: "Administrative state of this IP interface", SELECTION:"UP/DOWN", DEFAULT:"UP"`
}

type IPv6IntfState struct {
	baseObj
	IntfRef           string `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"r", DESCRIPTION: "System assigned interface id of L2 interface (port/lag/vlan) to which this IPv4 object is linked"`
	IfIndex           int32  `DESCRIPTION: "System assigned interface id for this IPv4 interface"`
	IpAddr            string `DESCRIPTION: "Interface IP Address/Prefix-Lenght to provisioned on switch interface", STRLEN:"43"`
	OperState         string `DESCRIPTION: "Operational state of this IP interface"`
	NumUpEvents       int32  `DESCRIPTION: "Number of times the operational state transitioned from DOWN to UP"`
	LastUpEventTime   string `DESCRIPTION: "Timestamp corresponding to the last DOWN to UP operational state change event"`
	NumDownEvents     int32  `DESCRIPTION: "Number of times the operational state transitioned from UP to DOWN"`
	LastDownEventTime string `DESCRIPTION: "Timestamp corresponding to the last UP to DOWN operational state change event"`
	L2IntfType        string `DESCRIPTION: "Type of L2 interface on which IP has been configured (Port/Lag/Vlan)"`
	L2IntfId          int32  `DESCRIPTION: "Id of the L2 interface. Port number/lag id/vlan id."`
}

type SubIPv6Intf struct {
	baseObj
	IntfRef string `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"w", DESCRIPTION:"Intf name for which ipv6Intf sub interface is to be configured"`
	Type    string `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"w", DESCRIPTION:"Type of interface, e.g. Secondary", STRLEN:"16", SELECTION: "Secondary"`
	IpAddr  string `DESCRIPTION:"Ip Address for sub interface", STRLEN:"43"`
	MacAddr string `DESCRIPTION:"Mac address to be used for the sub interface. If none specified IPv6Intf mac address will be used", STRLEN:"17", DEFAULT:""`
	LinkIp  bool   `DESCRIPTION: "Interface Link Scope IP Address auto-configured", DEFAULT:true`
	Enable  bool   `DESCRIPTION:"Enable or disable this interface", DEFAULT:true`
}

type SubIPv6IntfState struct {
	baseObj
	IntfRef       string `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"r", DESCRIPTION:"Intf name for which ipv6Intf sub interface is to be configured"`
	Type          string `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"r", DESCRIPTION:"Type of interface, e.g. Secondary or Virtual"`
	IfIndex       int32  `DESCRIPTION:"System assigned interface id for this sub IPv6 interface"`
	IfName        string `DESCRIPTION:"System generated sub interface name"`
	ParentIfIndex int32  `DESCRIPTION:"System assigned interface id for interface parent interface"`
	IpAddr        string `DESCRIPTION:"Ip Address for sub interface"`
	MacAddr       string `DESCRIPTION:"Mac address to be used for the sub interface. If none specified IPv6Intf mac address will be used"`
	OperState     string `DESCRIPTION:"Operational state of this SubIPv6 interface"`
}

type BufferPortStatState struct {
	baseObj
	IntfRef        string `SNAPROUTE: "KEY", CATEGORY:"System", ACCESS:"r", DESCRIPTION: "Front panel port name interface id"`
	IfIndex        int32  `DESCRIPTION: "System assigned interface id for this port. Read only attribute"`
	EgressPort     uint64 `DESCRIPTION: "Egress port buffer stats "`
	IngressPort    uint64 `DESCRIPTION: "Ingress port buffer stats "`
	PortBufferStat uint64 `DESCRIPTION: "Per port buffer stats"`
}

type BufferGlobalStatState struct {
	baseObj
	DeviceId          uint32 `SNAPROUTE: "KEY", CATEGORY:"System", ACCESS:"r", DESCRIPTION: "Device id"`
	BufferStat        uint64 `DESCRIPTION: "Buffer stats for the device "`
	EgressBufferStat  uint64 `DESCRIPTION: "Egress Buffer stats "`
	IngressBufferStat uint64 `DESCRIPTION: "Ingress buffer stats "`
}

type AclGlobal struct {
	baseObj
	AclGlobal        string `SNAPROUTE:"KEY", CATEGORY:"System", AUTOCREATE:"true", ACCESS:"w",MULTIPLICITY: "1", DESCRIPTION: "Indicates aclGlobal instance.", DEFAULT:"default"`
	GlobalDropEnable string `SELECTION:"TRUE/FALSE", DEFAULT:"FALSE", DESCRIPTION:"Global traffic drop  flag"`
}

type Acl struct {
	baseObj
	AclName    string   `SNAPROUTE: "KEY", CATEGORY:"System", MULTIPLICITY: "*", ACCESS:"w", DESCRIPTION: "Acl name."`
	IntfList   []string `DESCRIPTION: "list of IntfRef can be port/lag object"`
	Stage      string   `DESCRIPTION: "Ingress or Egress where ACL to be applied", SELECTION:"IN/OUT", DEFAULT:"IN"`
	Priority   int32    `DESCRIPTION: "Acl priority. Acls with higher priority will have precedence over with lower.", DEFAULT:1`
	AclType    string   `DESCRIPTION: "Acl type IPv4/Mac/Ipv6", SELECTION:"IPv4/Mac/IPv6", DEFAULT:"IPv4", STRLEN:"16"`
	Action     string   `DESCRIPTION: "Type of action (ALLOW/DENY)",SELECTION:"ALLOW/DENY",  DEFAULT:"ALLOW", STRLEN:"16"`
	FilterName string   `DESCRIPTION: "Filter name for acl . ", DEFAULT:""`
}

type AclIpv4Filter struct {
	baseObj
	FilterName  string `SNAPROUTE: "KEY", MULTIPLICITY: "*", ACCESS:"w", DESCRIPTION: "AClIpv4 filter name ."`
	SourceIp    string `DESCRIPTION: "Source IP address", DEFAULT:""`
	DestIp      string `DESCRIPTION: "Destination IP address", DEFAULT:""`
	SourceMask  string `DESCRIPTION: "Network mask for source IP", DEFAULT:""`
	DestMask    string `DESCRIPTION: "Network mark for dest IP", DEFAULT:""`
	Proto       string `DESCRIPTION: "Protocol type TCP/UDP/ICMPv4/ICMPv6", SELECTION:"TCP/UDP/ICMPv4/ICMPv6", DEFAULT:""`
	SrcIntf     string `DESCRIPTION: "Source Intf(used for mlag)", DEFAULT:""`
	DstIntf     string `DESCRIPTION: "Dest Intf(used for mlag)", DEFAULT:""`
	L4SrcPort   int32  `DESCRIPTION: "TCP/UDP source port", DEFAULT:0`
	L4DstPort   int32  `DESCRIPTION: "TCP/UDP destionation port", DEFAULT:0`
	L4PortMatch string `DESCRIPTION: "match condition can be EQ(equal) , NEQ(not equal), RANGE(port range)",SELECTION:"EQ/NEQ/RANGE", DEFAULT:""`
	L4MinPort   int32  `DESCRIPTION: "Min port when l4 port is specified as range", DEFAULT:0`
	L4MaxPort   int32  `DESCRIPTION: "Max port when l4 port is specified as range", DEFAULT:0`
}

type AclMacFilter struct {
	baseObj
	FilterName string `SNAPROUTE:"KEY", MULTIPLICITY: "*", ACCESS:"w", DESCRIPTION: "MAC filter name ."`
	SourceMac  string `DESCRIPTION: "Source MAC address.", DEFAULT:""`
	DestMac    string `DESCRIPTION: "Destination MAC address", DEFAULT:""`
	SourceMask string `DESCRIPTION: "Destination MAC address", DEFAULT:FF:FF:FF:FF:FF:FF`
	DestMask   string `DESCRIPTION: "Source MAC address", DEFAULT:FF:FF:FF:FF:FF:FF`
}

type AclIpv6Filter struct {
	baseObj
	FilterName   string `SNAPROUTE:"KEY", MULTIPLICITY: "*", ACCESS:"w", DESCRIPTION: "AClIpv6 filter name ."`
	SourceIpv6   string `DESCRIPTION: "Source IPv6 address", DEFAULT:""`
	DestIpv6     string `DESCRIPTION: "Destination IPv6 address", DEFAULT:""`
	SourceMaskv6 string `DESCRIPTION: "Network mask for source IPv6", DEFAULT:""`
	DestMaskv6   string `DESCRIPTION: "Network mark for dest IPv6", DEFAULT:""`
	Proto        string `DESCRIPTION: "Protocol type TCP/UDP/ICMPv4/ICMPv6", SELECTION:"TCP/UDP/ICMPv4/ICMPv6", DEFAULT:""`
	SrcIntf      string `DESCRIPTION: "Source Intf(used for mlag)", DEFAULT:""`
	DstIntf      string `DESCRIPTION: "Dest Intf(used for mlag)", DEFAULT:""`
	L4SrcPort    int32  `DESCRIPTION: "TCP/UDP source port", DEFAULT:0`
	L4DstPort    int32  `DESCRIPTION: "TCP/UDP destionation port", DEFAULT:0`
	L4PortMatch  string `DESCRIPTION: "match condition can be EQ(equal) , NEQ(not equal), RANGE(port range)",SELECTION:"EQ/NEQ/LT/GT/RANGE", DEFAULT:""`
	L4MinPort    int32  `DESCRIPTION: "Min port when l4 port is specified as range", DEFAULT:0`
	L4MaxPort    int32  `DESCRIPTION: "Max port when l4 port is specified as range", DEFAULT:0`
}

type AclState struct {
	baseObj
	AclName    string   `SNAPROUTE: "KEY", CATEGORY:"L3", MULTIPLICITY: "*", ACCESS:"r", DESCRIPTION: "Acl rule name"`
	Priority   int32    `DESCRIPTION: "Acl priority "`
	AclType    string   `DESCRIPTION: "Type can be IPv4/MAC/IPv6"`
	IntfList   []string `DESCRIPTION: "list of IntfRef can be port/lag object"`
	HwPresence string   `DESCRIPTION: "Check if the rule is installed in hardware. Applied/Not Applied/Failed"`
	HitCount   uint64   `DESCRIPTION: "No of  packets hit the rule if applied."`
}

// NEED TO ADD SUPPORT TO MAKE THIS INTERNAL ONLY
type LinkScopeIpState struct {
	baseObj
	LinkScopeIp string `SNAPROUTE: "KEY", CATEGORY:"L3", MULTIPLICITY: "*", ACCESS:"r", DESCRIPTION:"Link scope IP Address", USESTATEDB:"true"`
	IntfRef     string `DESCRIPTION: "Interface where the link scope ip is configured"`
	IfIndex     int32  `DESCRIPTION: "System Generated Unique Interface Id"`
	Used        bool   `DESCRIPTION : "states whether the ip being used"`
}

type Copp struct {
	baseObj
	Protocol                string `SNAPROUTE: "KEY", MULTIPLICITY: "*", ACCESS:"rw", AUTODISCOVER:"true", SELECTION:"ArpUC/ArpMC/BGP/ICMPv4UC/ICMPv4BC/STP/LACP/BFD/ICMPv6/LLDP", DESCRIPTION:"Protocol for which COPP is configured"`
	Cpuqueue                int32  `DESCRIPTION: "CPU queue to which the protocol traffic is mapped."`
	PolicerPeakRatePPS      int32  `DESCRIPTION: "Policer peak rate in pps"`
	PolicerPeakRateBurstPPS int32  `DESCRIPTION: "Policer peak burst rate in pps"`
}

type CoppStatState struct {
	baseObj
	Protocol     string `SNAPROUTE: "KEY", CATEGORY:"Physical", MULTIPLICITY: "*", ACCESS:"r", DESCRIPTION:"Protocol type for which CoPP is configured."`
	PeakRate     int32  `DESCRIPTION:"Peak rate (packets) for policer."`
	BurstRate    int32  `DESCRIPTION:"Burst rate (packets) for policer."`
	GreenPackets int64  `DESCRIPTION:"Packets marked with green for tri color policer."`
	RedPackets   int64  `DESCRIPTION:"Dropped packets. Packets marked with red for tri color policer. "`
}
