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
	"net"
	"utils/clntUtils/clntDefs/asicdClntDefs"
	"utils/netUtils"
)

/*
#include <stdint.h>
#include <stdlib.h>
#include <cps.h>
#cgo CFLAGS: -I. -I/usr/include/
#cgo LDFLAGS: -L/usr/lib/x86_64-linux-gnu/ -lopx_common -lopx_cps_api_common
*/
import "C"

func (asicdClientMgr *CPSAsicdClntMgr) CreateIPv4Neighbor(ipAddr, macAddr string, vlanId, ifIndex int32) (val int32, err error) {
	cpsAsicdMutex.Lock()
	defer cpsAsicdMutex.Unlock()

	var intfName string
	parsedMacAddr, err := net.ParseMAC(macAddr)
	if err != nil {
		Logger.Info("Error CreateIPv4Neighbor, ParseMAC", macAddr)
		return -1, errors.New("Error CreateIPv4Neighbor, ParseMAC")
	}
	if vlanId != -1 {
		intfName = fmt.Sprintf("br%d", vlanId)
	} else {
		intf, err := net.InterfaceByIndex(int(ifIndex))
		if err != nil {
			Logger.Info("Error CreateIPv4Neighbor, Invalid ifIndex", ifIndex)
			return -1, errors.New("Error CreateIPv4Neighbor, Invalid ifIndex")
		}
		intfName = intf.Name
	}
	rv := int(C.CPSCreateIPv4Neighbor(C.CString(ipAddr), C.CString(intfName), (*C.uint8_t)(&parsedMacAddr[0])))
	if rv != 0 {
		Logger.Info("Error Creating IPv4 Neighbor", ipAddr, macAddr, vlanId, ifIndex)
		return -1, errors.New("Error Creating IPv4 Neighbor")
	}
	return val, err
}

func (asicdClientMgr *CPSAsicdClntMgr) UpdateIPv4Neighbor(ipAddr, macAddr string, vlanId, ifIndex int32) (val int32, err error) {
	return val, err
}

func (asicdClientMgr *CPSAsicdClntMgr) DeleteIPv4Neighbor(ipAddr, macAddr string, vlanId, ifIndex int32) (val int32, err error) {
	cpsAsicdMutex.Lock()
	defer cpsAsicdMutex.Unlock()

	var intfName string
	parsedMacAddr, err := net.ParseMAC(macAddr)
	if err != nil {
		Logger.Info("Error CreateIPv4Neighbor, ParseMAC", macAddr)
		return -1, errors.New("Error CreateIPv4Neighbor, ParseMAC")
	}
	if vlanId != -1 {
		intfName = fmt.Sprintf("br%d", vlanId)
	} else {
		intf, err := net.InterfaceByIndex(int(ifIndex))
		if err != nil {
			Logger.Info("Error CreateIPv4Neighbor, Invalid ifIndex", ifIndex)
			return -1, errors.New("Error CreateIPv4Neighbor, Invalid ifIndex")
		}
		intfName = intf.Name
	}
	rv := int(C.CPSDeleteIPv4Neighbor(C.CString(ipAddr), C.CString(intfName), (*C.uint8_t)(&parsedMacAddr[0])))
	if rv != 0 {
		Logger.Info("Error Creating IPv4 Neighbor", ipAddr, macAddr, vlanId, ifIndex)
		return -1, errors.New("Error Creating IPv4 Neighbor")
	}
	return val, err
}

func (asicdClientMgr *CPSAsicdClntMgr) CreateIPv6Neighbor(ipAddr, macAddr string, vlanId, ifIndex int32) (val int32, err error) {
	return val, err
}

func (asicdClientMgr *CPSAsicdClntMgr) UpdateIPv6Neighbor(ipAddr, macAddr string, vlanId, ifIndex int32) (val int32, err error) {
	return val, err
}

func (asicdClientMgr *CPSAsicdClntMgr) DeleteIPv6Neighbor(ipAddr, macAddr string, vlanId, ifIndex int32) (val int32, err error) {
	return val, err
}

func (asicdClientMgr *CPSAsicdClntMgr) GetBulkVlan(curMark, count int) (*asicdClntDefs.VlanGetInfo, error) {
	var retObj asicdClntDefs.VlanGetInfo
	return &retObj, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) CreateLag(ifName string, hashType int32, ifIndexList string) (val int32, err error) {
	return val, err
}

func (asicdClientMgr *CPSAsicdClntMgr) DeleteLag(ifIndex int32) (val int32, err error) {
	return val, err
}

func (asicdClientMgr *CPSAsicdClntMgr) UpdateLag(ifIndex, hashType int32, ifIndexList string) (val int32, err error) {
	return val, err
}

func (asicdClientMgr *CPSAsicdClntMgr) GetBulkLag(fromIndex, count int) (*asicdClntDefs.LagGetInfo, error) {
	var retObj asicdClntDefs.LagGetInfo
	return &retObj, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) CreateLagCfgIntfList(ifName string, ifIndexList []int32) (rv bool, err error) {
	return rv, err
}

func (asicdClientMgr *CPSAsicdClntMgr) UpdateLagCfgIntfList(ifName string, ifIndexList []int32) (rv bool, err error) {
	return rv, err
}

func (asicdClientMgr *CPSAsicdClntMgr) DeleteLagCfgIntfList(ifName string, ifIndexList []int32) (rv bool, err error) {
	return rv, err
}

func (asicdClientMgr *CPSAsicdClntMgr) IsLinuxOnlyPlugin() (bool, error) {
	return true, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) GetAllPortsWithDirtyCache() ([]*asicdClntDefs.Port, error) {
	return nil, nil
}

func (asicdClientMgr *CPSAsicdClntMgr) OnewayCreateIPv4Route(ipv4RouteList []*asicdClntDefs.IPv4Route) {
	for idx := 0; idx < len(ipv4RouteList); idx++ {
		destNw := ipv4RouteList[idx].DestinationNw
		prefixLen, err := netUtils.GetPrefixLen(net.ParseIP(ipv4RouteList[idx].NetworkMask))
		if err != nil {
			Logger.Err("Error getting the prefixlen for:", ipv4RouteList[idx])
			continue
		}
		for idx1 := 0; idx1 < len(ipv4RouteList[idx].NextHopList); idx1++ {
			nhIP := ipv4RouteList[idx].NextHopList[idx1].NextHopIp
			rv := int(C.CPSCreateIPv4Route(C.CString(destNw), C.uint32_t(prefixLen), C.CString(nhIP)))
			if rv != 0 {
				Logger.Info("Error Creating IPv4Route", destNw, prefixLen, nhIP)
			}
		}
	}
}
func (asicdClientMgr *CPSAsicdClntMgr) OnewayDeleteIPv4Route(ipv4RouteList []*asicdClntDefs.IPv4Route) {
	for idx := 0; idx < len(ipv4RouteList); idx++ {
		destNw := ipv4RouteList[idx].DestinationNw
		prefixLen, err := netUtils.GetPrefixLen(net.ParseIP(ipv4RouteList[idx].NetworkMask))
		if err != nil {
			Logger.Err("Error getting the prefixlen for:", ipv4RouteList[idx])
			continue
		}
		for idx1 := 0; idx1 < len(ipv4RouteList[idx].NextHopList); idx1++ {
			nhIP := ipv4RouteList[idx].NextHopList[idx1].NextHopIp
			rv := int(C.CPSDeleteIPv4Route(C.CString(destNw), C.uint32_t(prefixLen), C.CString(nhIP)))
			if rv != 0 {
				Logger.Info("Error Creating IPv4Route", destNw, prefixLen, nhIP)
			}
		}
	}
}
func (asicdClientMgr *CPSAsicdClntMgr) OnewayCreateIPv6Route(ipv6RouteList []*asicdClntDefs.IPv6Route) {
}
func (asicdClientMgr *CPSAsicdClntMgr) OnewayDeleteIPv6Route(ipv6RouteList []*asicdClntDefs.IPv6Route) {
}
