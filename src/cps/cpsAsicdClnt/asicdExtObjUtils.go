//
//Copyright [2016] [SnapRoute Inc]
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//       Unless required by applicable law or agreed to in writing, software
//       distributed under the License is distributed on an "AS IS" BASIS,
//       WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//       See the License for the specific language governing permissions and
//       limitations under the License.
//
// _______  __       __________   ___      _______.____    __    ____  __  .___________.  ______  __    __
// |   ____||  |     |   ____\  \ /  /     /       |\   \  /  \  /   / |  | |           | /      ||  |  |  |
// |  |__   |  |     |  |__   \  V  /     |   (----` \   \/    \/   /  |  | `---|  |----`|  ,----'|  |__|  |
// |   __|  |  |     |   __|   >   <       \   \      \            /   |  |     |  |     |  |     |   __   |
// |  |     |  `----.|  |____ /  .  \  .----)   |      \    /\    /    |  |     |  |     |  `----.|  |  |  |
// |__|     |_______||_______/__/ \__\ |_______/        \__/  \__/     |__|     |__|      \______||__|  |__|
//

package cpsAsicdClnt

/*
import (
	"asicdServices"
	"models/objects"
)

func convertFromThriftToAsicGlobalState(obj *asicdServices.AsicGlobalState) *objects.AsicGlobalState {
	return &objects.AsicGlobalState{
		ModuleId:   uint8(obj.ModuleId),
		VendorId:   string(obj.VendorId),
		PartNumber: string(obj.PartNumber),
		RevisionId: string(obj.RevisionId),
		ModuleTemp: float64(obj.ModuleTemp),
	}
}

func convertToThriftFromAsicGlobalPM(obj *objects.AsicGlobalPM) *asicdServices.AsicGlobalPM {
	return &asicdServices.AsicGlobalPM{
		ModuleId:           int8(obj.ModuleId),
		Resource:           string(obj.Resource),
		PMClassAEnable:     bool(obj.PMClassAEnable),
		PMClassBEnable:     bool(obj.PMClassBEnable),
		PMClassCEnable:     bool(obj.PMClassCEnable),
		HighAlarmThreshold: float64(obj.HighAlarmThreshold),
		HighWarnThreshold:  float64(obj.HighWarnThreshold),
		LowAlarmThreshold:  float64(obj.LowAlarmThreshold),
		LowWarnThreshold:   float64(obj.LowWarnThreshold),
	}
}

func convertFromThriftToAsicGlobalPM(obj *asicdServices.AsicGlobalPM) *objects.AsicGlobalPM {
	return &objects.AsicGlobalPM{
		ModuleId:           uint8(obj.ModuleId),
		Resource:           string(obj.Resource),
		PMClassAEnable:     bool(obj.PMClassAEnable),
		PMClassBEnable:     bool(obj.PMClassBEnable),
		PMClassCEnable:     bool(obj.PMClassCEnable),
		HighAlarmThreshold: float64(obj.HighAlarmThreshold),
		HighWarnThreshold:  float64(obj.HighWarnThreshold),
		LowAlarmThreshold:  float64(obj.LowAlarmThreshold),
		LowWarnThreshold:   float64(obj.LowWarnThreshold),
	}
}

func convertFromThriftToPMData(obj []*asicdServices.PMData) []objects.PMData {
	var retObj []objects.PMData
	for _, pmData := range obj {
		data := objects.PMData{
			TimeStamp: string(pmData.TimeStamp),
			Value:     float64(pmData.Value),
		}
		retObj = append(retObj, data)
	}
	return retObj
}

func convertFromThriftToAsicGlobalPMState(obj *asicdServices.AsicGlobalPMState) *objects.AsicGlobalPMState {
	classAPMData := convertFromThriftToPMData(obj.ClassAPMData)
	classBPMData := convertFromThriftToPMData(obj.ClassBPMData)
	classCPMData := convertFromThriftToPMData(obj.ClassCPMData)
	return &objects.AsicGlobalPMState{
		ModuleId:     uint8(obj.ModuleId),
		Resource:     string(obj.Resource),
		ClassAPMData: classAPMData,
		ClassBPMData: classBPMData,
		ClassCPMData: classCPMData,
	}
}

func convertFromThriftToEthernetPM(obj *asicdServices.EthernetPM) *objects.EthernetPM {
	return &objects.EthernetPM{
		IntfRef:            string(obj.IntfRef),
		Resource:           string(obj.Resource),
		PMClassAEnable:     bool(obj.PMClassAEnable),
		PMClassBEnable:     bool(obj.PMClassBEnable),
		PMClassCEnable:     bool(obj.PMClassCEnable),
		HighAlarmThreshold: float64(obj.HighAlarmThreshold),
		HighWarnThreshold:  float64(obj.HighWarnThreshold),
		LowAlarmThreshold:  float64(obj.LowAlarmThreshold),
		LowWarnThreshold:   float64(obj.LowWarnThreshold),
	}
}

func convertToThriftFromEthernetPM(obj *objects.EthernetPM) *asicdServices.EthernetPM {
	return &asicdServices.EthernetPM{
		IntfRef:            string(obj.IntfRef),
		Resource:           string(obj.Resource),
		PMClassAEnable:     bool(obj.PMClassAEnable),
		PMClassBEnable:     bool(obj.PMClassBEnable),
		PMClassCEnable:     bool(obj.PMClassCEnable),
		HighAlarmThreshold: float64(obj.HighAlarmThreshold),
		HighWarnThreshold:  float64(obj.HighWarnThreshold),
		LowAlarmThreshold:  float64(obj.LowAlarmThreshold),
		LowWarnThreshold:   float64(obj.LowWarnThreshold),
	}
}

func convertFromThriftToEthernetPMState(obj *asicdServices.EthernetPMState) *objects.EthernetPMState {
	classAPMData := convertFromThriftToPMData(obj.ClassAPMData)
	classBPMData := convertFromThriftToPMData(obj.ClassBPMData)
	classCPMData := convertFromThriftToPMData(obj.ClassCPMData)
	return &objects.EthernetPMState{
		IntfRef:      string(obj.IntfRef),
		Resource:     string(obj.Resource),
		ClassAPMData: classAPMData,
		ClassBPMData: classBPMData,
		ClassCPMData: classCPMData,
	}
}

func convertFromThriftToAsicSummaryState(obj *asicdServices.AsicSummaryState) *objects.AsicSummaryState {
	return &objects.AsicSummaryState{
		ModuleId:      uint8(obj.ModuleId),
		NumPortsUp:    int32(obj.NumPortsUp),
		NumPortsDown:  int32(obj.NumPortsDown),
		NumVlans:      int32(obj.NumVlans),
		NumV4Intfs:    int32(obj.NumV4Intfs),
		NumV6Intfs:    int32(obj.NumV6Intfs),
		NumV4Adjs:     int32(obj.NumV4Adjs),
		NumV6Adjs:     int32(obj.NumV6Adjs),
		NumV4Routes:   int32(obj.NumV4Routes),
		NumV6Routes:   int32(obj.NumV6Routes),
		NumECMPRoutes: int32(obj.NumECMPRoutes),
	}
}

func convertFromThriftToVlan(obj *asicdServices.Vlan) *objects.Vlan {
	var intfList []string
	for _, intf := range obj.IntfList {
		intfList = append(intfList, string(intf))
	}
	var untagIntfList []string
	for _, intf := range obj.UntagIntfList {
		untagIntfList = append(untagIntfList, string(intf))
	}
	return &objects.Vlan{
		VlanId:        int32(obj.VlanId),
		AdminState:    string(obj.AdminState),
		Description:   string(obj.Description),
		AutoState:     string(obj.AutoState),
		IntfList:      intfList,
		UntagIntfList: untagIntfList,
	}
}

func convertToThriftFromVlan(obj *objects.Vlan) *asicdServices.Vlan {
	var intfList []string
	for _, intf := range obj.IntfList {
		intfList = append(intfList, string(intf))
	}
	var untagIntfList []string
	for _, intf := range obj.UntagIntfList {
		untagIntfList = append(untagIntfList, string(intf))
	}
	return &asicdServices.Vlan{
		VlanId:        int32(obj.VlanId),
		AdminState:    string(obj.AdminState),
		Description:   string(obj.Description),
		AutoState:     string(obj.AutoState),
		IntfList:      intfList,
		UntagIntfList: untagIntfList,
	}
}

func convertFromThriftToVlanState(obj *asicdServices.VlanState) *objects.VlanState {
	return &objects.VlanState{
		VlanId:                 int32(obj.VlanId),
		VlanName:               string(obj.VlanName),
		OperState:              string(obj.OperState),
		IfIndex:                int32(obj.IfIndex),
		SysInternalDescription: string(obj.SysInternalDescription),
	}
}

func convertFromThriftToIPv4Intf(obj *asicdServices.IPv4Intf) *objects.IPv4Intf {
	return &objects.IPv4Intf{
		IntfRef:    string(obj.IntfRef),
		IpAddr:     string(obj.IpAddr),
		AdminState: string(obj.AdminState),
	}
}

func convertToThriftFromIPv4Intf(obj *objects.IPv4Intf) *asicdServices.IPv4Intf {
	return &asicdServices.IPv4Intf{
		IntfRef:    string(obj.IntfRef),
		IpAddr:     string(obj.IpAddr),
		AdminState: string(obj.AdminState),
	}
}

func convertFromThriftToIPv4IntfState(obj *asicdServices.IPv4IntfState) *objects.IPv4IntfState {
	return &objects.IPv4IntfState{
		IntfRef:           string(obj.IntfRef),
		IfIndex:           int32(obj.IfIndex),
		IpAddr:            string(obj.IpAddr),
		OperState:         string(obj.OperState),
		NumUpEvents:       int32(obj.NumUpEvents),
		LastUpEventTime:   string(obj.LastUpEventTime),
		NumDownEvents:     int32(obj.NumDownEvents),
		LastDownEventTime: string(obj.LastDownEventTime),
		L2IntfType:        string(obj.L2IntfType),
		L2IntfId:          int32(obj.L2IntfId),
	}
}

func convertToThriftFromPort(obj *objects.Port) *asicdServices.Port {
	return &asicdServices.Port{
		IntfRef:        string(obj.IntfRef),
		IfIndex:        int32(obj.IfIndex),
		Description:    string(obj.Description),
		PhyIntfType:    string(obj.PhyIntfType),
		AdminState:     string(obj.AdminState),
		MacAddr:        string(obj.MacAddr),
		Speed:          int32(obj.Speed),
		Duplex:         string(obj.Duplex),
		Autoneg:        string(obj.Autoneg),
		MediaType:      string(obj.MediaType),
		Mtu:            int32(obj.Mtu),
		BreakOutMode:   string(obj.BreakOutMode),
		LoopbackMode:   string(obj.LoopbackMode),
		EnableFEC:      bool(obj.EnableFEC),
		PRBSTxEnable:   bool(obj.PRBSTxEnable),
		PRBSRxEnable:   bool(obj.PRBSRxEnable),
		PRBSPolynomial: string(obj.PRBSPolynomial),
	}
}

func convertFromThriftToPort(obj *asicdServices.Port) *objects.Port {
	return &objects.Port{
		IntfRef:        string(obj.IntfRef),
		IfIndex:        int32(obj.IfIndex),
		Description:    string(obj.Description),
		PhyIntfType:    string(obj.PhyIntfType),
		AdminState:     string(obj.AdminState),
		MacAddr:        string(obj.MacAddr),
		Speed:          int32(obj.Speed),
		Duplex:         string(obj.Duplex),
		Autoneg:        string(obj.Autoneg),
		MediaType:      string(obj.MediaType),
		Mtu:            int32(obj.Mtu),
		BreakOutMode:   string(obj.BreakOutMode),
		LoopbackMode:   string(obj.LoopbackMode),
		EnableFEC:      bool(obj.EnableFEC),
		PRBSTxEnable:   bool(obj.PRBSTxEnable),
		PRBSRxEnable:   bool(obj.PRBSRxEnable),
		PRBSPolynomial: string(obj.PRBSPolynomial),
	}
}

func convertFromThriftToPortState(obj *asicdServices.PortState) *objects.PortState {
	return &objects.PortState{
		IntfRef:                     string(obj.IntfRef),
		IfIndex:                     int32(obj.IfIndex),
		Name:                        string(obj.Name),
		OperState:                   string(obj.OperState),
		NumUpEvents:                 int32(obj.NumUpEvents),
		LastUpEventTime:             string(obj.LastUpEventTime),
		NumDownEvents:               int32(obj.NumDownEvents),
		LastDownEventTime:           string(obj.LastDownEventTime),
		Pvid:                        int32(obj.Pvid),
		IfInOctets:                  int64(obj.IfInOctets),
		IfInUcastPkts:               int64(obj.IfInUcastPkts),
		IfInDiscards:                int64(obj.IfInDiscards),
		IfInErrors:                  int64(obj.IfInErrors),
		IfInUnknownProtos:           int64(obj.IfInUnknownProtos),
		IfOutOctets:                 int64(obj.IfOutOctets),
		IfOutUcastPkts:              int64(obj.IfOutUcastPkts),
		IfOutDiscards:               int64(obj.IfOutDiscards),
		IfEtherUnderSizePktCnt:      int64(obj.IfEtherUnderSizePktCnt),
		IfEtherOverSizePktCnt:       int64(obj.IfEtherOverSizePktCnt),
		IfEtherFragments:            int64(obj.IfEtherFragments),
		IfEtherCRCAlignError:        int64(obj.IfEtherCRCAlignError),
		IfEtherJabber:               int64(obj.IfEtherJabber),
		IfEtherPkts:                 int64(obj.IfEtherPkts),
		IfEtherMCPkts:               int64(obj.IfEtherMCPkts),
		IfEtherBcastPkts:            int64(obj.IfEtherBcastPkts),
		IfEtherPkts64OrLessOctets:   int64(obj.IfEtherPkts64OrLessOctets),
		IfEtherPkts65To127Octets:    int64(obj.IfEtherPkts65To127Octets),
		IfEtherPkts128To255Octets:   int64(obj.IfEtherPkts128To255Octets),
		IfEtherPkts256To511Octets:   int64(obj.IfEtherPkts256To511Octets),
		IfEtherPkts512To1023Octets:  int64(obj.IfEtherPkts512To1023Octets),
		IfEtherPkts1024To1518Octets: int64(obj.IfEtherPkts1024To1518Octets),
		ErrDisableReason:            string(obj.ErrDisableReason),
		PresentInHW:                 string(obj.PresentInHW),
		ConfigMode:                  string(obj.ConfigMode),
		PRBSRxErrCnt:                int64(obj.PRBSRxErrCnt),
	}
}

func convertFromThriftToMacTableEntryState(obj *asicdServices.MacTableEntryState) *objects.MacTableEntryState {
	return &objects.MacTableEntryState{
		MacAddr: string(obj.MacAddr),
		VlanId:  int32(obj.VlanId),
		Port:    int32(obj.Port),
	}
}

func convertFromThriftToIPv4RouteHwState(obj *asicdServices.IPv4RouteHwState) *objects.IPv4RouteHwState {
	return &objects.IPv4RouteHwState{
		DestinationNw:    string(obj.DestinationNw),
		NextHopIps:       string(obj.NextHopIps),
		RouteCreatedTime: string(obj.RouteCreatedTime),
		RouteUpdatedTime: string(obj.RouteUpdatedTime),
	}
}

func convertFromThriftToIPv6RouteHwState(obj *asicdServices.IPv6RouteHwState) *objects.IPv6RouteHwState {
	return &objects.IPv6RouteHwState{
		DestinationNw:    string(obj.DestinationNw),
		NextHopIps:       string(obj.NextHopIps),
		RouteCreatedTime: string(obj.RouteCreatedTime),
		RouteUpdatedTime: string(obj.RouteUpdatedTime),
	}
}

func convertFromThriftToArpEntryHwState(obj *asicdServices.ArpEntryHwState) *objects.ArpEntryHwState {
	return &objects.ArpEntryHwState{
		IpAddr:  string(obj.IpAddr),
		MacAddr: string(obj.MacAddr),
		Vlan:    string(obj.Vlan),
		Port:    string(obj.Port),
	}
}

func convertFromThriftToNdpEntryHwState(obj *asicdServices.NdpEntryHwState) *objects.NdpEntryHwState {
	return &objects.NdpEntryHwState{
		IpAddr:  string(obj.IpAddr),
		MacAddr: string(obj.MacAddr),
		Vlan:    string(obj.Vlan),
		Port:    string(obj.Port),
	}
}

func convertFromThriftToLogicalIntf(obj *asicdServices.LogicalIntf) *objects.LogicalIntf {
	return &objects.LogicalIntf{
		Name: string(obj.Name),
		Type: string(obj.Type),
	}
}

func convertToThriftFromLogicalIntf(obj *objects.LogicalIntf) *asicdServices.LogicalIntf {
	return &asicdServices.LogicalIntf{
		Name: string(obj.Name),
		Type: string(obj.Type),
	}
}

func convertFromThriftToLogicalIntfState(obj *asicdServices.LogicalIntfState) *objects.LogicalIntfState {
	return &objects.LogicalIntfState{
		Name:              string(obj.Name),
		IfIndex:           int32(obj.IfIndex),
		SrcMac:            string(obj.SrcMac),
		OperState:         string(obj.OperState),
		IfInOctets:        int64(obj.IfInOctets),
		IfInUcastPkts:     int64(obj.IfInUcastPkts),
		IfInDiscards:      int64(obj.IfInDiscards),
		IfInErrors:        int64(obj.IfInErrors),
		IfInUnknownProtos: int64(obj.IfInUnknownProtos),
		IfOutOctets:       int64(obj.IfOutOctets),
		IfOutUcastPkts:    int64(obj.IfOutUcastPkts),
		IfOutDiscards:     int64(obj.IfOutDiscards),
		IfOutErrors:       int64(obj.IfOutErrors),
	}
}

func convertFromThriftToSubIPv4Intf(obj *asicdServices.SubIPv4Intf) *objects.SubIPv4Intf {
	return &objects.SubIPv4Intf{
		IntfRef: string(obj.IntfRef),
		Type:    string(obj.Type),
		IpAddr:  string(obj.IpAddr),
		MacAddr: string(obj.MacAddr),
		Enable:  bool(obj.Enable),
	}
}

func convertToThriftFromSubIPv4Intf(obj *objects.SubIPv4Intf) *asicdServices.SubIPv4Intf {
	return &asicdServices.SubIPv4Intf{
		IntfRef: string(obj.IntfRef),
		Type:    string(obj.Type),
		IpAddr:  string(obj.IpAddr),
		MacAddr: string(obj.MacAddr),
		Enable:  bool(obj.Enable),
	}
}

func convertFromThriftToSubIPv4IntfState(obj *asicdServices.SubIPv4IntfState) *objects.SubIPv4IntfState {
	return &objects.SubIPv4IntfState{
		IntfRef:       obj.IntfRef,
		Type:          obj.Type,
		IfIndex:       int32(obj.IfIndex),
		IfName:        obj.IfName,
		ParentIfIndex: int32(obj.ParentIfIndex),
		IpAddr:        obj.IpAddr,
		MacAddr:       obj.MacAddr,
		OperState:     obj.OperState,
	}
}

func convertFromThriftToIPv6Intf(obj *asicdServices.IPv6Intf) *objects.IPv6Intf {
	return &objects.IPv6Intf{
		IntfRef:    string(obj.IntfRef),
		IpAddr:     string(obj.IpAddr),
		LinkIp:     bool(obj.LinkIp),
		AdminState: string(obj.AdminState),
	}
}

func convertToThriftFromIPv6Intf(obj *objects.IPv6Intf) *asicdServices.IPv6Intf {
	return &asicdServices.IPv6Intf{
		IntfRef:    string(obj.IntfRef),
		IpAddr:     string(obj.IpAddr),
		LinkIp:     bool(obj.LinkIp),
		AdminState: string(obj.AdminState),
	}
}

func convertFromThriftToIPv6IntfState(obj *asicdServices.IPv6IntfState) *objects.IPv6IntfState {
	return &objects.IPv6IntfState{
		IntfRef:           string(obj.IntfRef),
		IfIndex:           int32(obj.IfIndex),
		IpAddr:            string(obj.IpAddr),
		OperState:         string(obj.OperState),
		NumUpEvents:       int32(obj.NumUpEvents),
		LastUpEventTime:   string(obj.LastUpEventTime),
		NumDownEvents:     int32(obj.NumDownEvents),
		LastDownEventTime: string(obj.LastDownEventTime),
		L2IntfType:        string(obj.L2IntfType),
		L2IntfId:          int32(obj.L2IntfId),
	}
}

func convertToThriftFromSubIPv6Intf(obj *objects.SubIPv6Intf) *asicdServices.SubIPv6Intf {
	return &asicdServices.SubIPv6Intf{
		IntfRef: string(obj.IntfRef),
		Type:    string(obj.Type),
		IpAddr:  string(obj.IpAddr),
		MacAddr: string(obj.MacAddr),
		LinkIp:  bool(obj.LinkIp),
		Enable:  bool(obj.Enable),
	}
}

func convertFromThriftToSubIPv6Intf(obj *asicdServices.SubIPv6Intf) *objects.SubIPv6Intf {
	return &objects.SubIPv6Intf{
		IntfRef: string(obj.IntfRef),
		Type:    string(obj.Type),
		IpAddr:  string(obj.IpAddr),
		MacAddr: string(obj.MacAddr),
		LinkIp:  bool(obj.LinkIp),
		Enable:  bool(obj.Enable),
	}
}

func convertFromThriftToSubIPv6IntfState(obj *asicdServices.SubIPv6IntfState) *objects.SubIPv6IntfState {
	return &objects.SubIPv6IntfState{
		IntfRef:       string(obj.IntfRef),
		Type:          string(obj.Type),
		IfIndex:       int32(obj.IfIndex),
		IfName:        string(obj.IfName),
		ParentIfIndex: int32(obj.ParentIfIndex),
		IpAddr:        string(obj.IpAddr),
		MacAddr:       string(obj.MacAddr),
		OperState:     string(obj.OperState),
	}
}

func convertFromThriftToBufferPortStatState(obj *asicdServices.BufferPortStatState) *objects.BufferPortStatState {
	return &objects.BufferPortStatState{
		IntfRef:        string(obj.IntfRef),
		IfIndex:        int32(obj.IfIndex),
		EgressPort:     uint64(obj.EgressPort),
		IngressPort:    uint64(obj.IngressPort),
		PortBufferStat: uint64(obj.PortBufferStat),
	}
}

func convertFromThriftToBufferGlobalStatState(obj *asicdServices.BufferGlobalStatState) *objects.BufferGlobalStatState {
	return &objects.BufferGlobalStatState{
		DeviceId:          uint32(obj.DeviceId),
		BufferStat:        uint64(obj.BufferStat),
		EgressBufferStat:  uint64(obj.EgressBufferStat),
		IngressBufferStat: uint64(obj.IngressBufferStat),
	}
}

func convertToThriftFromAclGlobal(obj *objects.AclGlobal) *asicdServices.AclGlobal {
	return &asicdServices.AclGlobal{
		AclGlobal:        string(obj.AclGlobal),
		GlobalDropEnable: string(obj.GlobalDropEnable),
	}
}

func convertFromThriftToAclGlobal(obj *asicdServices.AclGlobal) *objects.AclGlobal {
	return &objects.AclGlobal{
		AclGlobal:        string(obj.AclGlobal),
		GlobalDropEnable: string(obj.GlobalDropEnable),
	}
}

func convertToThriftFromAcl(obj *objects.Acl) *asicdServices.Acl {
	var intfList []string
	for _, intf := range obj.IntfList {
		intfList = append(intfList, string(intf))
	}
	return &asicdServices.Acl{
		AclName:    string(obj.AclName),
		IntfList:   intfList,
		Stage:      string(obj.Stage),
		Priority:   int32(obj.Priority),
		AclType:    string(obj.AclType),
		Action:     string(obj.Action),
		FilterName: string(obj.FilterName),
	}
}

func convertFromThriftToAcl(obj *asicdServices.Acl) *objects.Acl {
	var intfList []string
	for _, intf := range obj.IntfList {
		intfList = append(intfList, string(intf))
	}
	return &objects.Acl{
		AclName:    string(obj.AclName),
		IntfList:   intfList,
		Stage:      string(obj.Stage),
		Priority:   int32(obj.Priority),
		AclType:    string(obj.AclType),
		Action:     string(obj.Action),
		FilterName: string(obj.FilterName),
	}
}

func convertToThriftFromAclIpv4Filter(obj *objects.AclIpv4Filter) *asicdServices.AclIpv4Filter {
	return &asicdServices.AclIpv4Filter{
		FilterName:  string(obj.FilterName),
		SourceIp:    string(obj.SourceIp),
		DestIp:      string(obj.DestIp),
		SourceMask:  string(obj.SourceMask),
		DestMask:    string(obj.DestMask),
		Proto:       string(obj.Proto),
		SrcIntf:     string(obj.SrcIntf),
		DstIntf:     string(obj.DstIntf),
		L4SrcPort:   int32(obj.L4SrcPort),
		L4DstPort:   int32(obj.L4DstPort),
		L4PortMatch: string(obj.L4PortMatch),
		L4MinPort:   int32(obj.L4MinPort),
		L4MaxPort:   int32(obj.L4MaxPort),
	}
}

func convertFromThriftToAclIpv4Filter(obj *asicdServices.AclIpv4Filter) *objects.AclIpv4Filter {
	return &objects.AclIpv4Filter{
		FilterName:  string(obj.FilterName),
		SourceIp:    string(obj.SourceIp),
		DestIp:      string(obj.DestIp),
		SourceMask:  string(obj.SourceMask),
		DestMask:    string(obj.DestMask),
		Proto:       string(obj.Proto),
		SrcIntf:     string(obj.SrcIntf),
		DstIntf:     string(obj.DstIntf),
		L4SrcPort:   int32(obj.L4SrcPort),
		L4DstPort:   int32(obj.L4DstPort),
		L4PortMatch: string(obj.L4PortMatch),
		L4MinPort:   int32(obj.L4MinPort),
		L4MaxPort:   int32(obj.L4MaxPort),
	}
}

func convertToThriftFromAclMacFilter(obj *objects.AclMacFilter) *asicdServices.AclMacFilter {
	return &asicdServices.AclMacFilter{
		FilterName: string(obj.FilterName),
		SourceMac:  string(obj.SourceMac),
		DestMac:    string(obj.DestMac),
		SourceMask: string(obj.SourceMask),
		DestMask:   string(obj.DestMask),
	}
}

func convertFromThriftToAclMacFilter(obj *asicdServices.AclMacFilter) *objects.AclMacFilter {
	return &objects.AclMacFilter{
		FilterName: string(obj.FilterName),
		SourceMac:  string(obj.SourceMac),
		DestMac:    string(obj.DestMac),
		SourceMask: string(obj.SourceMask),
		DestMask:   string(obj.DestMask),
	}
}

func convertToThriftFromAclIpv6Filter(obj *objects.AclIpv6Filter) *asicdServices.AclIpv6Filter {
	return &asicdServices.AclIpv6Filter{
		FilterName:   string(obj.FilterName),
		SourceIpv6:   string(obj.SourceIpv6),
		DestIpv6:     string(obj.DestIpv6),
		SourceMaskv6: string(obj.SourceMaskv6),
		DestMaskv6:   string(obj.DestMaskv6),
		Proto:        string(obj.Proto),
		SrcIntf:      string(obj.SrcIntf),
		DstIntf:      string(obj.DstIntf),
		L4SrcPort:    int32(obj.L4SrcPort),
		L4DstPort:    int32(obj.L4DstPort),
		L4PortMatch:  string(obj.L4PortMatch),
		L4MinPort:    int32(obj.L4MinPort),
		L4MaxPort:    int32(obj.L4MaxPort),
	}
}

func convertFromThriftToAclIpv6Filter(obj *asicdServices.AclIpv6Filter) *objects.AclIpv6Filter {
	return &objects.AclIpv6Filter{
		FilterName:   string(obj.FilterName),
		SourceIpv6:   string(obj.SourceIpv6),
		DestIpv6:     string(obj.DestIpv6),
		SourceMaskv6: string(obj.SourceMaskv6),
		DestMaskv6:   string(obj.DestMaskv6),
		Proto:        string(obj.Proto),
		SrcIntf:      string(obj.SrcIntf),
		DstIntf:      string(obj.DstIntf),
		L4SrcPort:    int32(obj.L4SrcPort),
		L4DstPort:    int32(obj.L4DstPort),
		L4PortMatch:  string(obj.L4PortMatch),
		L4MinPort:    int32(obj.L4MinPort),
		L4MaxPort:    int32(obj.L4MaxPort),
	}
}

func convertFromThriftToAclState(obj *asicdServices.AclState) *objects.AclState {
	var intfList []string
	for _, intf := range obj.IntfList {
		intfList = append(intfList, string(intf))
	}
	return &objects.AclState{
		AclName:    string(obj.AclName),
		Priority:   int32(obj.Priority),
		AclType:    string(obj.AclType),
		IntfList:   intfList,
		HwPresence: string(obj.HwPresence),
		HitCount:   uint64(obj.HitCount),
	}
}

func convertFromThriftToLinkScopeIpState(obj *asicdServices.LinkScopeIpState) *objects.LinkScopeIpState {
	return &objects.LinkScopeIpState{
		LinkScopeIp: string(obj.LinkScopeIp),
		IntfRef:     string(obj.IntfRef),
		IfIndex:     int32(obj.IfIndex),
		Used:        bool(obj.Used),
	}
}

func convertFromThriftToCoppStatState(obj *asicdServices.CoppStatState) *objects.CoppStatState {
	return &objects.CoppStatState{
		Protocol:     string(obj.Protocol),
		PeakRate:     int32(obj.PeakRate),
		BurstRate:    int32(obj.BurstRate),
		GreenPackets: int64(obj.GreenPackets),
		RedPackets:   int64(obj.RedPackets),
	}
}
*/
