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

import (
	"errors"
	"fmt"
	"models/objects"
	"net"
	"sync"
	"utils/clntUtils/clntDefs/asicdClntDefs"
)

/*
#include <stdint.h>
#include <stdlib.h>
#include <cps.h>
#cgo CFLAGS: -I. -I/usr/include/
#cgo LDFLAGS: -L/usr/lib/x86_64-linux-gnu/ -lopx_common -lopx_cps_api_common
*/
import "C"

var cpsAsicdMutex *sync.Mutex = &sync.Mutex{}

func (asicdClientMgr *CPSAsicdClntMgr) GetBulkAsicGlobalState(fromIndex int, count int) (*asicdClntDefs.AsicGlobalStateGetInfo, error) {
	var retObj asicdClntDefs.AsicGlobalStateGetInfo
	return &retObj, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetAsicGlobalState(ModuleId uint8) (*objects.AsicGlobalState, error) {
	return nil, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) CreateAsicGlobalPM(cfg *objects.AsicGlobalPM) (bool, error) {
	return false, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) UpdateAsicGlobalPM(origCfg, newCfg *objects.AsicGlobalPM, attrset []bool, op []*objects.PatchOpInfo) (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) DeleteAsicGlobalPM(cfg *objects.AsicGlobalPM) (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetBulkAsicGlobalPM(fromIndex int, count int) (*asicdClntDefs.AsicGlobalPMGetInfo, error) {
	var retObj asicdClntDefs.AsicGlobalPMGetInfo
	return &retObj, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetAsicGlobalPM(ModuleId uint8, Resource string) (*objects.AsicGlobalPM, error) {
	return nil, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetBulkAsicGlobalPMState(fromIndex int, count int) (*asicdClntDefs.AsicGlobalPMStateGetInfo, error) {
	var retObj asicdClntDefs.AsicGlobalPMStateGetInfo
	return &retObj, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetAsicGlobalPMState(ModuleId uint8, Resource string) (*objects.AsicGlobalPMState, error) {
	return nil, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) CreateEthernetPM(cfg *objects.EthernetPM) (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) UpdateEthernetPM(origCfg, newCfg *objects.EthernetPM, attrset []bool, op []*objects.PatchOpInfo) (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) DeleteEthernetPM(cfg *objects.EthernetPM) (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetBulkEthernetPM(fromIndex int, count int) (*asicdClntDefs.EthernetPMGetInfo, error) {
	var retObj asicdClntDefs.EthernetPMGetInfo
	return &retObj, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetEthernetPM(IntfRef string, Resource string) (*objects.EthernetPM, error) {
	return nil, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetBulkEthernetPMState(fromIndex int, count int) (*asicdClntDefs.EthernetPMStateGetInfo, error) {
	var retObj asicdClntDefs.EthernetPMStateGetInfo
	return &retObj, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetEthernetPMState(IntfRef string, Resource string) (*objects.EthernetPMState, error) {
	return nil, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetBulkAsicSummaryState(fromIndex int, count int) (*asicdClntDefs.AsicSummaryStateGetInfo, error) {
	var retObj asicdClntDefs.AsicSummaryStateGetInfo
	return &retObj, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetAsicSummaryState(ModuleId uint8) (*objects.AsicSummaryState, error) {
	return nil, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) CreateVlan(cfg *objects.Vlan) (bool, error) {
	var tagPortIfIdx, untagPortIfIdx []int32
	cpsAsicdMutex.Lock()
	defer cpsAsicdMutex.Unlock()
	Logger.Info("Calling CPS CreateVlan:", cfg)
	tagPorts := C.MakeCharArray(C.int(len(cfg.IntfList)))
	defer C.FreeCharArray(tagPorts, C.int(len(cfg.IntfList)))
	for idx, intf := range cfg.IntfList {
		C.SetArrayString(tagPorts, C.CString(intf), C.int(idx))
	}
	untagPorts := C.MakeCharArray(C.int(len(cfg.UntagIntfList)))
	defer C.FreeCharArray(untagPorts, C.int(len(cfg.UntagIntfList)))
	for idx, intf := range cfg.UntagIntfList {
		C.SetArrayString(untagPorts, C.CString(intf), C.int(idx))
	}
	rv := int(C.CPSCreateVlan(C.uint32_t(cfg.VlanId), C.uint32_t(len(cfg.IntfList)), tagPorts, C.uint32_t(len(cfg.UntagIntfList)), untagPorts))
	if rv != 0 {
		return false, errors.New("Error Creating Vlan")
	}
	vlanName := fmt.Sprintf("br%d", cfg.VlanId)
	intf, err := net.InterfaceByName(vlanName)
	if err != nil {
		Logger.Err("1 Error Create Vlan, Unable to get vlanName", vlanName)
		return false, errors.New("Error CreateVlan, Unable to get interface vlanName")
	}
	ifIdx := int32(intf.Index)
	vlanEnt, _ := asicdClientMgr.VlanDB[cfg.VlanId]
	vlanEnt.VlanName = vlanName
	vlanEnt.OperState = "UP"
	vlanEnt.AdminState = "UP"
	vlanEnt.IfIndex = ifIdx
	vlanEnt.IntfList = make(map[int32]bool)
	for _, tagIntf := range cfg.IntfList {
		intf, err := net.InterfaceByName(tagIntf)
		if err != nil {
			Logger.Err("Error Create Vlan, Unable to get vlanName", tagIntf)
			continue
		}
		vlanEnt.IntfList[int32(intf.Index)] = true
		tagPortIfIdx = append(tagPortIfIdx, int32(intf.Index))
	}
	vlanEnt.UntagIntfList = make(map[int32]bool)
	for _, untagIntf := range cfg.UntagIntfList {
		intf, err := net.InterfaceByName(untagIntf)
		if err != nil {
			Logger.Info("Error Create Vlan, Unable to get vlanName", untagIntf)
			continue
		}
		vlanEnt.UntagIntfList[int32(intf.Index)] = true
		untagPortIfIdx = append(untagPortIfIdx, int32(intf.Index))
	}
	asicdClientMgr.VlanDB[cfg.VlanId] = vlanEnt
	asicdClientMgr.VlanList = append(asicdClientMgr.VlanList, cfg.VlanId)
	asicdClientMgr.IfIdxToIfIdMap[ifIdx] = cfg.VlanId
	asicdClientMgr.IfIdxToIfTypeMap[ifIdx] = asicdClntDefs.IfTypeVlan
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) DeleteVlan(cfg *objects.Vlan) (bool, error) {
	cpsAsicdMutex.Lock()
	defer cpsAsicdMutex.Unlock()
	Logger.Info("Calling CPS DeleteVlan:", cfg)
	vlanName := fmt.Sprintf("br%d", cfg.VlanId)
	intf, err := net.InterfaceByName(vlanName)
	if err != nil {
		Logger.Info("Error Create Vlan, Unable to get vlanName", vlanName)
		return false, errors.New("Error CreateVlan, Unable to get interface vlanName")
	}
	ifIdx := int32(intf.Index)
	rv := int(C.CPSDeleteVlan(C.uint32_t(cfg.VlanId)))
	if rv != 0 {
		return false, errors.New("Error Delete Vlan")
	}
	delete(asicdClientMgr.VlanDB, cfg.VlanId)
	var vlanList []int32
	for idx := 0; idx < len(asicdClientMgr.VlanList); idx++ {
		if cfg.VlanId != asicdClientMgr.VlanList[idx] {
			vlanList = append(vlanList, cfg.VlanId)
		}
	}
	asicdClientMgr.VlanList = asicdClientMgr.VlanList[:0]
	asicdClientMgr.VlanList = nil
	asicdClientMgr.VlanList = append(asicdClientMgr.VlanList, vlanList...)
	delete(asicdClientMgr.IfIdxToIfIdMap, ifIdx)
	delete(asicdClientMgr.IfIdxToIfTypeMap, ifIdx)

	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) UpdateVlan(origCfg, newCfg *objects.Vlan, attrset []bool, op []*objects.PatchOpInfo) (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetBulkVlanState(fromIdx, count int) (*asicdClntDefs.VlanStateGetInfo, error) {
	var retObj asicdClntDefs.VlanStateGetInfo
	cpsAsicdMutex.Lock()
	defer cpsAsicdMutex.Unlock()
	var idx, numEntries int
	if (fromIdx > len(asicdClientMgr.VlanList)) || (fromIdx < 0) {
		Logger.Err("Invalid fromIdx vlan argument in get bulk vlan state")
		return nil, errors.New("Invalid fromIdx vlan argument in get bulk vlan state")
	}
	if count < 0 {
		Logger.Err("Invalid count in get bulk port config")
		return nil, errors.New("Invalid count int get bulk port config")
	}
	for idx = fromIdx; idx < len(asicdClientMgr.VlanList); idx++ {
		if numEntries == count {
			retObj.More = true
			break
		}
		var vlanState objects.VlanState
		vlanState.VlanId = asicdClientMgr.VlanList[idx]
		vlanState.VlanName = asicdClientMgr.VlanDB[vlanState.VlanId].VlanName
		vlanState.IfIndex = asicdClientMgr.VlanDB[vlanState.VlanId].IfIndex
		vlanState.OperState = asicdClientMgr.VlanDB[vlanState.VlanId].OperState
		retObj.VlanStateList = append(retObj.VlanStateList, &vlanState)
		numEntries++
	}
	retObj.EndIdx = int32(idx)
	retObj.Count = int32(numEntries)
	if idx == len(asicdClientMgr.VlanDB) {
		retObj.More = false
	}
	return &retObj, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetVlanState(VlanId int32) (*objects.VlanState, error) {
	cpsAsicdMutex.Lock()
	defer cpsAsicdMutex.Unlock()
	vlanEnt, exist := asicdClientMgr.VlanDB[VlanId]
	if !exist {
		return nil, errors.New("Invalid VlanId")
	}
	var retObj objects.VlanState
	retObj.VlanId = VlanId
	retObj.VlanName = vlanEnt.VlanName
	retObj.IfIndex = vlanEnt.IfIndex
	retObj.OperState = vlanEnt.OperState
	return &retObj, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) CreateIPv4Intf(cfg *objects.IPv4Intf) (bool, error) {
	cpsAsicdMutex.Lock()
	defer cpsAsicdMutex.Unlock()
	Logger.Info("Calling CPS CreateIPv4Intf:", cfg)
	ipAddr, ipNet, err := net.ParseCIDR(cfg.IpAddr)
	if err != nil {
		return false, errors.New(fmt.Sprintln("Invalid IP Address", err))
	}
	prefixLen, _ := ipNet.Mask.Size()
	if err != nil {
		return false, errors.New(fmt.Sprintln("Invalid IP Address", err))
	}

	rv := int(C.CPSCreateIPv4Intf(C.CString(cfg.IntfRef), C.CString(ipAddr.String()), C.uint32_t(prefixLen)))
	if rv != 0 {
		return false, errors.New("Error Creating IPv4 Intf")
	}
	intf, err := net.InterfaceByName(cfg.IntfRef)
	if err != nil {
		Logger.Info("Error Create IPv4Intf, Unable to get interface details", cfg.IntfRef)
		return false, errors.New(fmt.Sprintln("Error Create IPv4Intf, Unable to get interface details", cfg.IntfRef))
	}
	ipv4IntfEnt, _ := asicdClientMgr.IPv4IntfDB[cfg.IntfRef]
	ipv4IntfEnt.IpAddr = cfg.IpAddr
	ipv4IntfEnt.AdminState = "UP"
	ipv4IntfEnt.IfIdx = int32(intf.Index)
	ipv4IntfEnt.OperState = "UP"
	asicdClientMgr.IPv4IntfDB[cfg.IntfRef] = ipv4IntfEnt
	asicdClientMgr.IPv4IntfList = append(asicdClientMgr.IPv4IntfList, cfg.IntfRef)
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) DeleteIPv4Intf(cfg *objects.IPv4Intf) (bool, error) {
	cpsAsicdMutex.Lock()
	defer cpsAsicdMutex.Unlock()
	Logger.Info("Calling CPS DeleteIPv4Intf:", cfg)
	ipAddr, ipNet, err := net.ParseCIDR(cfg.IpAddr)
	if err != nil {
		return false, errors.New(fmt.Sprintln("Invalid IP Address", err))
	}
	prefixLen, _ := ipNet.Mask.Size()
	if err != nil {
		return false, errors.New(fmt.Sprintln("Invalid IP Address", err))
	}

	rv := int(C.CPSDeleteIPv4Intf(C.CString(cfg.IntfRef), C.CString(ipAddr.String()), C.uint32_t(prefixLen)))
	if rv != 0 {
		return false, errors.New("Error Creating IPv4 Intf")
	}
	delete(asicdClientMgr.IPv4IntfDB, cfg.IntfRef)
	var ipv4IntfList []string
	for idx := 0; idx < len(asicdClientMgr.IPv4IntfList); idx++ {
		if asicdClientMgr.IPv4IntfList[idx] != cfg.IntfRef {
			ipv4IntfList = append(ipv4IntfList, asicdClientMgr.IPv4IntfList[idx])
		}
	}
	asicdClientMgr.IPv4IntfList = asicdClientMgr.IPv4IntfList[:0]
	asicdClientMgr.IPv4IntfList = nil
	asicdClientMgr.IPv4IntfList = append(asicdClientMgr.IPv4IntfList, ipv4IntfList...)
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) UpdateIPv4Intf(origCfg, newCfg *objects.IPv4Intf, attrset []bool, op []*objects.PatchOpInfo) (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetBulkIPv4IntfState(fromIdx, count int) (*asicdClntDefs.IPv4IntfStateGetInfo, error) {
	var retObj asicdClntDefs.IPv4IntfStateGetInfo
	cpsAsicdMutex.Lock()
	defer cpsAsicdMutex.Unlock()
	var idx, numEntries int
	if (fromIdx > len(asicdClientMgr.IPv4IntfList)) || (fromIdx < 0) {
		Logger.Err("Invalid fromIdx ipv4Intf argument in get bulk ipv4Intf state")
		return nil, errors.New("Invalid fromIdx ipv4Intf argument in get bulk ipv4Intf state")
	}
	if count < 0 {
		Logger.Err("Invalid count in get bulk IPv4IntfList")
		return nil, errors.New("Invalid count int get bulk IPv4IntList")
	}
	for idx = fromIdx; idx < len(asicdClientMgr.IPv4IntfList); idx++ {
		if numEntries == count {
			retObj.More = true
			break
		}
		var ipv4IntfState objects.IPv4IntfState
		intfRef := asicdClientMgr.IPv4IntfList[idx]
		ipv4IntfState.IntfRef = intfRef
		ipv4IntfState.IfIndex = asicdClientMgr.IPv4IntfDB[intfRef].IfIdx
		ipv4IntfState.IpAddr = asicdClientMgr.IPv4IntfDB[intfRef].IpAddr
		ipv4IntfState.OperState = asicdClientMgr.IPv4IntfDB[intfRef].OperState
		ipv4IntfState.L2IntfId = int32(asicdClientMgr.GetIntfIdFromIfIndex(ipv4IntfState.IfIndex))
		ifType := asicdClientMgr.GetIntfTypeFromIfIndex(ipv4IntfState.IfIndex)
		switch ifType {
		case asicdClntDefs.IfTypePort:
			ipv4IntfState.L2IntfType = "Port"
		case asicdClntDefs.IfTypeVlan:
			ipv4IntfState.L2IntfType = "Vlan"
		}
		retObj.IPv4IntfStateList = append(retObj.IPv4IntfStateList, &ipv4IntfState)
		numEntries++
	}
	retObj.EndIdx = int32(idx)
	retObj.Count = int32(numEntries)
	if idx == len(asicdClientMgr.IPv4IntfList) {
		retObj.More = false
	}
	return &retObj, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetIPv4IntfState(IntfRef string) (*objects.IPv4IntfState, error) {
	cpsAsicdMutex.Lock()
	defer cpsAsicdMutex.Unlock()
	ipv4Ent, exist := asicdClientMgr.IPv4IntfDB[IntfRef]
	if !exist {
		return nil, errors.New("Invalid IntfRef")
	}
	var retObj objects.IPv4IntfState
	retObj.IntfRef = IntfRef
	retObj.IfIndex = ipv4Ent.IfIdx
	retObj.IpAddr = ipv4Ent.IpAddr
	retObj.OperState = ipv4Ent.OperState
	retObj.L2IntfId = int32(asicdClientMgr.GetIntfIdFromIfIndex(retObj.IfIndex))
	ifType := asicdClientMgr.GetIntfTypeFromIfIndex(retObj.IfIndex)
	switch ifType {
	case asicdClntDefs.IfTypePort:
		retObj.L2IntfType = "Port"
	case asicdClntDefs.IfTypeVlan:
		retObj.L2IntfType = "Vlan"
	}
	return &retObj, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) CreatePort(cfg *objects.Port) (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) UpdatePort(origCfg, newCfg *objects.Port, attrset []bool, op []*objects.PatchOpInfo) (bool, error) {
	cpsAsicdMutex.Lock()
	defer cpsAsicdMutex.Unlock()
	var mask uint32
	var idx int

	for idx, val := range attrset {
		if val {
			switch idx {
			case 0:
				//Object Key IntfRef
			case 1:
				mask |= asicdClntDefs.PORT_UPDATE_ATTR_IF_INDEX
			case 2:
				mask |= asicdClntDefs.PORT_UPDATE_ATTR_DESCRIPTION
			case 3:
				mask |= asicdClntDefs.PORT_UPDATE_ATTR_PHY_INTF_TYPE
			case 4:
				mask |= asicdClntDefs.PORT_UPDATE_ATTR_ADMIN_STATE
			case 5:
				mask |= asicdClntDefs.PORT_UPDATE_ATTR_MAC_ADDR
			case 6:
				mask |= asicdClntDefs.PORT_UPDATE_ATTR_SPEED
			case 7:
				mask |= asicdClntDefs.PORT_UPDATE_ATTR_DUPLEX
			case 8:
				mask |= asicdClntDefs.PORT_UPDATE_ATTR_AUTONEG
			case 9:
				mask |= asicdClntDefs.PORT_UPDATE_ATTR_MEDIA_TYPE
			case 10:
				mask |= asicdClntDefs.PORT_UPDATE_ATTR_MTU
			case 11:
				mask |= asicdClntDefs.PORT_UPDATE_ATTR_BREAK_OUT_MODE
			case 12:
				mask |= asicdClntDefs.PORT_UPDATE_ATTR_LOOPBACK_MODE
			case 13:
				mask |= asicdClntDefs.PORT_UPDATE_ATTR_ENABLE_FEC
			case 14:
				mask |= asicdClntDefs.PORT_UPDATE_ATTR_PRBS_TX_ENABLE
			case 15:
				mask |= asicdClntDefs.PORT_UPDATE_ATTR_PRBS_RX_ENABLE
			case 16:
				mask |= asicdClntDefs.PORT_UPDATE_ATTR_PRBS_POLYNOMIAL
			}
		}
	}

	for idx = 0; idx < len(asicdClientMgr.PortDB); idx++ {
		if asicdClientMgr.PortDB[idx].IntfRef == newCfg.IntfRef {
			break
		}
	}
	if idx == len(asicdClientMgr.PortDB) {
		return false, errors.New(fmt.Sprintln("Error Invalid IntfRef", newCfg.IntfRef))
	}

	if (mask & asicdClntDefs.PORT_UPDATE_ATTR_ADMIN_STATE) == asicdClntDefs.PORT_UPDATE_ATTR_ADMIN_STATE {
		var val uint8
		if newCfg.AdminState == "UP" {
			val = 1
		}
		rv := int(C.CPSSetPortAdminState(C.CString(newCfg.IntfRef), C.uint8_t(val)))
		if rv != 0 {
			return false, errors.New("Error Setting Port AdminState")
		}
		asicdClientMgr.PortDB[idx].AdminState = newCfg.AdminState
		asicdClientMgr.PortDB[idx].OperState = newCfg.AdminState
	}

	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) DeletePort(cfg *objects.Port) (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetBulkPort(fromIdx, count int) (*asicdClntDefs.PortGetInfo, error) {
	cpsAsicdMutex.Lock()
	defer cpsAsicdMutex.Unlock()
	var retObj asicdClntDefs.PortGetInfo
	var idx, numEntries int
	if (fromIdx > len(asicdClientMgr.PortDB)) || (fromIdx < 0) {
		Logger.Err("Invalid fromIdx port argument in get bulk port config")
		return nil, errors.New("Invalid fromIdx port argument in get bulk port config")
	}
	if count < 0 {
		Logger.Err("Invalid count in get bulk port config")
		return nil, errors.New("Invalid count int get bulk port config")
	}
	for idx = fromIdx; idx < len(asicdClientMgr.PortDB); idx++ {
		if numEntries == count {
			retObj.More = true
			break
		}
		var portCfg objects.Port
		portCfg.IntfRef = asicdClientMgr.PortDB[idx].IntfRef
		portCfg.IfIndex = asicdClientMgr.PortDB[idx].IfIndex
		portCfg.MacAddr = asicdClientMgr.PortDB[idx].MacAddr
		portCfg.AdminState = asicdClientMgr.PortDB[idx].AdminState
		retObj.PortList = append(retObj.PortList, &portCfg)
		numEntries++
	}
	retObj.EndIdx = int32(idx)
	if idx == len(asicdClientMgr.PortDB) {
		retObj.More = false
	}
	return &retObj, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetPort(IntfRef string) (*objects.Port, error) {
	cpsAsicdMutex.Lock()
	defer cpsAsicdMutex.Unlock()
	var flag bool
	var idx int
	for idx = 0; idx < len(asicdClientMgr.PortDB); idx++ {
		if asicdClientMgr.PortDB[idx].IntfRef == IntfRef {
			flag = true
			break
		}
	}
	if flag == false {
		return nil, errors.New("Invalid IntfRef")
	}
	var retObj objects.Port
	retObj.IntfRef = asicdClientMgr.PortDB[idx].IntfRef
	retObj.IfIndex = asicdClientMgr.PortDB[idx].IfIndex
	retObj.MacAddr = asicdClientMgr.PortDB[idx].MacAddr
	retObj.AdminState = asicdClientMgr.PortDB[idx].AdminState
	return &retObj, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetBulkPortState(fromIdx, count int) (*asicdClntDefs.PortStateGetInfo, error) {
	cpsAsicdMutex.Lock()
	defer cpsAsicdMutex.Unlock()
	var retObj asicdClntDefs.PortStateGetInfo
	var idx, numEntries int
	if (fromIdx > len(asicdClientMgr.PortDB)) || (fromIdx < 0) {
		Logger.Err("Invalid fromIdx port argument in get bulk port config")
		return nil, errors.New("Invalid fromIdx port argument in get bulk port config")
	}
	if count < 0 {
		Logger.Err("Invalid count in get bulk port config")
		return nil, errors.New("Invalid count int get bulk port config")
	}
	for idx = fromIdx; idx < len(asicdClientMgr.PortDB); idx++ {
		if numEntries == count {
			retObj.More = true
			break
		}
		var portState objects.PortState
		portState.IntfRef = asicdClientMgr.PortDB[idx].IntfRef
		portState.IfIndex = asicdClientMgr.PortDB[idx].IfIndex
		portState.Name = asicdClientMgr.PortDB[idx].IntfRef
		portState.OperState = asicdClientMgr.PortDB[idx].OperState
		retObj.PortStateList = append(retObj.PortStateList, &portState)
		numEntries++
	}
	retObj.EndIdx = int32(idx)
	if idx == len(asicdClientMgr.PortDB) {
		retObj.More = false
	}
	retObj.Count = int32(numEntries)
	return &retObj, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetPortState(IntfRef string) (*objects.PortState, error) {
	cpsAsicdMutex.Lock()
	defer cpsAsicdMutex.Unlock()
	var flag bool
	var idx int
	for idx = 0; idx < len(asicdClientMgr.PortDB); idx++ {
		if asicdClientMgr.PortDB[idx].IntfRef == IntfRef {
			flag = true
			break
		}
	}
	if flag == false {
		return nil, errors.New("Invalid IntfRef")
	}
	var retObj objects.PortState
	retObj.IntfRef = asicdClientMgr.PortDB[idx].IntfRef
	retObj.IfIndex = asicdClientMgr.PortDB[idx].IfIndex
	retObj.Name = asicdClientMgr.PortDB[idx].IntfRef
	retObj.OperState = asicdClientMgr.PortDB[idx].OperState
	return &retObj, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetBulkMacTableEntryState(fromIndex int, count int) (*asicdClntDefs.MacTableEntryStateGetInfo, error) {
	var retObj asicdClntDefs.MacTableEntryStateGetInfo
	return &retObj, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetMacTableEntryState(MacAddr string) (*objects.MacTableEntryState, error) {
	return nil, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetBulkIPv4RouteHwState(fromIndex int, count int) (*asicdClntDefs.IPv4RouteHwStateGetInfo, error) {
	var retObj asicdClntDefs.IPv4RouteHwStateGetInfo
	return &retObj, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetIPv4RouteHwState(DestinationNw string) (*objects.IPv4RouteHwState, error) {
	return nil, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetBulkIPv6RouteHwState(fromIndex int, count int) (*asicdClntDefs.IPv6RouteHwStateGetInfo, error) {
	var retObj asicdClntDefs.IPv6RouteHwStateGetInfo
	return &retObj, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetIPv6RouteHwState(DestinationNw string) (*objects.IPv6RouteHwState, error) {
	return nil, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetBulkArpEntryHwState(fromIndex int, count int) (*asicdClntDefs.ArpEntryHwStateGetInfo, error) {
	var retObj asicdClntDefs.ArpEntryHwStateGetInfo
	return &retObj, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetArpEntryHwState(IpAddr string) (*objects.ArpEntryHwState, error) {
	return nil, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetBulkNdpEntryHwState(fromIndex int, count int) (*asicdClntDefs.NdpEntryHwStateGetInfo, error) {
	var retObj asicdClntDefs.NdpEntryHwStateGetInfo
	return &retObj, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetNdpEntryHwState(IpAddr string) (*objects.NdpEntryHwState, error) {
	return nil, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) CreateLogicalIntf(cfg *objects.LogicalIntf) (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) UpdateLogicalIntf(origCfg, newCfg *objects.LogicalIntf, attrset []bool, op []*objects.PatchOpInfo) (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) DeleteLogicalIntf(cfg *objects.LogicalIntf) (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetBulkLogicalIntfState(fromIndex int, count int) (*asicdClntDefs.LogicalIntfStateGetInfo, error) {
	var retObj asicdClntDefs.LogicalIntfStateGetInfo
	return &retObj, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetLogicalIntfState(Name string) (*objects.LogicalIntfState, error) {
	return nil, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) CreateSubIPv4Intf(cfg *objects.SubIPv4Intf) (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) UpdateSubIPv4Intf(origCfg, newCfg *objects.SubIPv4Intf, attrset []bool, op []*objects.PatchOpInfo) (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) DeleteSubIPv4Intf(cfg *objects.SubIPv4Intf) (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetBulkSubIPv4IntfState(fromIndex int, count int) (*asicdClntDefs.SubIPv4IntfStateGetInfo, error) {
	var retObj asicdClntDefs.SubIPv4IntfStateGetInfo
	return &retObj, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetSubIPv4IntfState(IntfRef string, Type string) (*objects.SubIPv4IntfState, error) {
	return nil, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) CreateIPv6Intf(cfg *objects.IPv6Intf) (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) UpdateIPv6Intf(origCfg, newCfg *objects.IPv6Intf, attrset []bool, op []*objects.PatchOpInfo) (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) DeleteIPv6Intf(cfg *objects.IPv6Intf) (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetBulkIPv6IntfState(fromIndex int, count int) (*asicdClntDefs.IPv6IntfStateGetInfo, error) {
	var retObj asicdClntDefs.IPv6IntfStateGetInfo
	return &retObj, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetIPv6IntfState(IntfRef string) (*objects.IPv6IntfState, error) {
	return nil, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) CreateSubIPv6Intf(cfg *objects.SubIPv6Intf) (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) UpdateSubIPv6Intf(origCfg, newCfg *objects.SubIPv6Intf, attrset []bool, op []*objects.PatchOpInfo) (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) DeleteSubIPv6Intf(cfg *objects.SubIPv6Intf) (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetBulkSubIPv6IntfState(fromIndex int, count int) (*asicdClntDefs.SubIPv6IntfStateGetInfo, error) {
	var retObj asicdClntDefs.SubIPv6IntfStateGetInfo
	return &retObj, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetSubIPv6IntfState(IntfRef string, Type string) (*objects.SubIPv6IntfState, error) {
	return nil, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetBulkBufferPortStatState(fromIndex int, count int) (*asicdClntDefs.BufferPortStatStateGetInfo, error) {
	var retObj asicdClntDefs.BufferPortStatStateGetInfo
	return &retObj, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetBufferPortStatState(IntfRef string) (*objects.BufferPortStatState, error) {
	return nil, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetBulkBufferGlobalStatState(fromIndex int, count int) (*asicdClntDefs.BufferGlobalStatStateGetInfo, error) {
	var retObj asicdClntDefs.BufferGlobalStatStateGetInfo
	return &retObj, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetBufferGlobalStatState(DeviceId uint32) (*objects.BufferGlobalStatState, error) {
	return nil, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) CreateAclGlobal(cfg *objects.AclGlobal) (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) UpdateAclGlobal(origCfg, newCfg *objects.AclGlobal, attrset []bool, op []*objects.PatchOpInfo) (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) DeleteAclGlobal(cfg *objects.AclGlobal) (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) CreateAcl(cfg *objects.Acl) (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) UpdateAcl(origCfg, newCfg *objects.Acl, attrset []bool, op []*objects.PatchOpInfo) (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) DeleteAcl(cfg *objects.Acl) (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) CreateAclIpv4Filter(cfg *objects.AclIpv4Filter) (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) UpdateAclIpv4Filter(origCfg, newCfg *objects.AclIpv4Filter, attrset []bool, op []*objects.PatchOpInfo) (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) DeleteAclIpv4Filter(cfg *objects.AclIpv4Filter) (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) CreateAclMacFilter(cfg *objects.AclMacFilter) (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) UpdateAclMacFilter(origCfg, newCfg *objects.AclMacFilter, attrset []bool, op []*objects.PatchOpInfo) (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) DeleteAclMacFilter(cfg *objects.AclMacFilter) (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) CreateAclIpv6Filter(cfg *objects.AclIpv6Filter) (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) UpdateAclIpv6Filter(origCfg, newCfg *objects.AclIpv6Filter, attrset []bool, op []*objects.PatchOpInfo) (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) DeleteAclIpv6Filter(cfg *objects.AclIpv6Filter) (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetBulkAclState(fromIndex int, count int) (*asicdClntDefs.AclStateGetInfo, error) {
	var retObj asicdClntDefs.AclStateGetInfo
	return &retObj, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetAclState(AclName string) (*objects.AclState, error) {
	return nil, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetBulkLinkScopeIpState(fromIndex int, count int) (*asicdClntDefs.LinkScopeIpStateGetInfo, error) {
	var retObj asicdClntDefs.LinkScopeIpStateGetInfo
	return &retObj, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetLinkScopeIpState(LinkScopeIp string) (*objects.LinkScopeIpState, error) {
	return nil, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetBulkCoppStatState(fromIndex int, count int) (*asicdClntDefs.CoppStatStateGetInfo, error) {
	var retObj asicdClntDefs.CoppStatStateGetInfo
	return &retObj, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetCoppStatState(Protocol string) (*objects.CoppStatState, error) {
	return nil, nil
}
