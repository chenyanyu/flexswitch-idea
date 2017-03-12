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
	"models/objects"
	"net"
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

type PortCfg C.struct_portCfg_s

func (asicdClientMgr *CPSAsicdClntMgr) GetSysRsvdVlan() int {
	return -1
}

func (asicdClientMgr *CPSAsicdClntMgr) GetIntfIdFromIfIndex(ifIndex int32) int {
	id, exist := asicdClientMgr.IfIdxToIfIdMap[ifIndex]
	if !exist {
		return -1
	}
	return int(id)
}

func (asicdClientMgr *CPSAsicdClntMgr) GetIntfTypeFromIfIndex(ifIndex int32) int {
	ifType, exist := asicdClientMgr.IfIdxToIfTypeMap[ifIndex]
	if !exist {
		return -1
	}
	return int(ifType)
}

func (asicdClientMgr *CPSAsicdClntMgr) GetIfIndexFromIntfIdAndIntfType(ifId int, ifType int) int32 {
	switch ifType {
	case asicdClntDefs.IfTypePort:
		for ifIdx, id := range asicdClientMgr.IfIdxToIfIdMap {
			if id == int32(ifId) {
				if asicdClientMgr.IfIdxToIfTypeMap[ifIdx] == asicdClntDefs.IfTypePort {
					return ifIdx
				}
			}
		}
	case asicdClntDefs.IfTypeVlan:
		for ifIdx, id := range asicdClientMgr.IfIdxToIfIdMap {
			if id == int32(ifId) {
				if asicdClientMgr.IfIdxToIfTypeMap[ifIdx] == asicdClntDefs.IfTypeVlan {
					return ifIdx
				}
			}
		}
	}
	return int32(-1)
}

func (asicdClientMgr *CPSAsicdClntMgr) GetMinSysPort() int {
	return 0
}

func (asicdClientMgr *CPSAsicdClntMgr) GetAllSubIPv4IntfState() ([]*objects.SubIPv4IntfState, error) {
	retObj := make([]*objects.SubIPv4IntfState, 0)
	return retObj, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetAllPortConfig() error {
	portCfg := make([]PortCfg, 256)
	count := uint8(0)
	ret := int(C.CPSGetAllPortCfg((*C.struct_portCfg_s)(&portCfg[0]), (*C.uint8_t)(&count)))
	if ret != 0 {
		Logger.Err("Error GetAllPortCfg()", ret)
		return errors.New("Error GetAllPortCfg()")
	}
	Logger.Info("Port Config:", portCfg)
	portMap := make(map[string]bool)
	for idx := 0; idx < int(count); idx++ {
		portName := C.GoString(&(portCfg[idx].PortName[0]))
		_, exist := portMap[portName]
		if !exist {
			var port Port
			Logger.Info("Port IfName:", portName, idx)
			port.IntfRef = portName
			intf, err := net.InterfaceByName(portName)
			if err != nil {
				Logger.Err("Error getting InterfaceByName", err)
				continue
			}
			port.MacAddr = intf.HardwareAddr.String()
			port.IfIndex = int32(intf.Index)
			if (intf.Flags & net.FlagUp) != 0 {
				port.OperState = "UP"
				port.AdminState = "UP"
			} else {
				port.OperState = "DOWN"
				port.AdminState = "DOWN"
			}
			asicdClientMgr.PortDB = append(asicdClientMgr.PortDB, port)
			portMap[portName] = true
			asicdClientMgr.IfIdxToIfIdMap[port.IfIndex] = port.IfIndex
			asicdClientMgr.IfIdxToIfTypeMap[port.IfIndex] = asicdClntDefs.IfTypePort
		}
	}
	return nil
}
