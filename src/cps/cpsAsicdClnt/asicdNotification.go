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
	"github.com/vishvananda/netlink"
	"net"
	"strings"
	"strconv"
	"utils/clntUtils/clntDefs/asicdClntDefs"
	"utils/clntUtils/clntIntfs"
)

//#include "cps.h"
import "C"

var nHdl clntIntfs.NotificationHdl

func (asicdClientMgr *CPSAsicdClntMgr) InitFSAsicdSubscriber(clntInitParams *clntIntfs.BaseClntInitParams) error {
	//Init notification hdl
	nHdl = clntInitParams.NHdl
	//Register with CPS to receive notifications
	ret := int(C.CPSRegisterNotificationHandler())
	if ret < 0 {
		Logger.Err("InitFSAsicdSubscriber failed")
	}
	asicdClientMgr.AddrSubCh = make(chan netlink.AddrUpdate)
	asicdClientMgr.AddrSubDone = make(chan struct{})
	err := netlink.AddrSubscribe(asicdClientMgr.AddrSubCh, asicdClientMgr.AddrSubDone)
	if err != nil {
		Logger.Err("Error opening AddrSubscriber()")
		return err
	}
	go asicdClientMgr.IPAddrNetLinkSubscriber()
	asicdClientMgr.LinkSubCh = make(chan netlink.LinkUpdate)
	asicdClientMgr.LinkSubDone = make(chan struct{})
	err = netlink.LinkSubscribe(asicdClientMgr.LinkSubCh, asicdClientMgr.LinkSubDone)
	if err != nil {
		Logger.Err("Error opening LinkSubscriber()")
		return err
	}
	go asicdClientMgr.LinkNetLinkSubscriber()
	return nil
}

func (asicdClientMgr *CPSAsicdClntMgr) DeinitFSAsicdSubscriber() {
	//Deregister notification handler
	C.CPSUnregisterNotificationHandler()
	close(asicdClientMgr.AddrSubDone)
	close(asicdClientMgr.LinkSubDone)
}

//export HandleLinkNotifications
func HandleLinkNotifications(ifIndex, operState C.int) {
	var ifState uint8

	switch operState {
	case C.IF_INTERFACES_STATE_INTERFACE_ADMIN_STATUS_UP:
		ifState = asicdClntDefs.INTF_STATE_UP

	case C.IF_INTERFACES_STATE_INTERFACE_ADMIN_STATUS_DOWN:
		ifState = asicdClntDefs.INTF_STATE_DOWN

	default:
		Logger.Err("Unrecognized operstate value received")
		return
	}
	msg := asicdClntDefs.L2IntfStateNotifyMsg{
		MsgType: asicdClntDefs.NOTIFY_L2INTF_STATE_CHANGE,
		IfIndex: int32(ifIndex),
		IfState: ifState,
	}
	Logger.Debug("Sending port state change notification for : ", ifIndex, ifState)
	cpsAsicdMutex.Lock()
	if nHdl != nil {
		nHdl.ProcessNotification(msg)
	}
	cpsAsicdMutex.Unlock()
}

func (asicdClientMgr *CPSAsicdClntMgr) SendIPv6IntfNotifyMsg(msgType int, ipAddr string, ifIdx int32, intfRef string) {
	var msg asicdClntDefs.IPv6IntfNotifyMsg
	msg.MsgType = uint8(msgType)
	msg.IpAddr = ipAddr
	msg.IfIndex = ifIdx
	msg.IntfRef = intfRef
	asicdClientMgr.SendNotification(msg)
}

func (asicdClientMgr *CPSAsicdClntMgr) SendIPv6IntfStateNotifyMsg(msgType int, ipAddr string, ifIdx int32, state uint8) {
	var msg asicdClntDefs.IPv6L3IntfStateNotifyMsg
	msg.MsgType = uint8(msgType)
	msg.IpAddr = ipAddr
	msg.IfIndex = ifIdx
	msg.IfState = state
	asicdClientMgr.SendNotification(msg)
}

func (asicdClientMgr *CPSAsicdClntMgr) SendIPv4IntfNotifyMsg(msgType int, ipAddr string, ifIdx int32, intfRef string) {
	var msg asicdClntDefs.IPv4IntfNotifyMsg
	msg.MsgType = uint8(msgType)
	msg.IpAddr = ipAddr
	msg.IfIndex = ifIdx
	msg.IntfRef = intfRef
	asicdClientMgr.SendNotification(msg)
}

func (asicdClientMgr *CPSAsicdClntMgr) SendIPv4IntfStateNotifyMsg(msgType int, ipAddr string, ifIdx int32, state uint8) {
	var msg asicdClntDefs.IPv4L3IntfStateNotifyMsg
	msg.MsgType = uint8(msgType)
	msg.IpAddr = ipAddr
	msg.IfIndex = ifIdx
	msg.IfState = state
	asicdClientMgr.SendNotification(msg)
}

func (asicdClientMgr *CPSAsicdClntMgr) SendNotification(msg clntIntfs.NotifyMsg) {
	cpsAsicdMutex.Lock()
	Logger.Info("Send Notification", msg)
	asicdClientMgr.BaseClntInitParams.NHdl.ProcessNotification(msg)
	cpsAsicdMutex.Unlock()
}

func (asicdClientMgr *CPSAsicdClntMgr) IPAddrNetLinkSubscriber() {
	for {
		addrUpdate := <-asicdClientMgr.AddrSubCh
		intf, err := net.InterfaceByIndex(addrUpdate.LinkIndex)
		if err != nil {
			Logger.Err("Error getting the interface using ifIndex")
			continue
		}
		if !strings.HasPrefix(intf.Name, "br") {
			flag := false
			for idx := 0; idx < len(asicdClientMgr.PortDB); idx++ {
				if intf.Name == asicdClientMgr.PortDB[idx].IntfRef {
					flag = true
					break
				}
			}
			if flag == false {
				continue
			}
		}
		if addrUpdate.NewAddr {
			if len(addrUpdate.LinkAddress.IP) == net.IPv4len {
				asicdClientMgr.SendIPv4IntfNotifyMsg(asicdClntDefs.NOTIFY_IPV4INTF_CREATE,
					addrUpdate.LinkAddress.String(), int32(intf.Index), intf.Name)
				asicdClientMgr.SendIPv4IntfStateNotifyMsg(asicdClntDefs.NOTIFY_IPV4_L3INTF_STATE_CHANGE,
					addrUpdate.LinkAddress.String(), int32(intf.Index), uint8(asicdClntDefs.INTF_STATE_UP))
			} else if len(addrUpdate.LinkAddress.IP) == net.IPv6len {
				asicdClientMgr.SendIPv4IntfNotifyMsg(asicdClntDefs.NOTIFY_IPV4INTF_CREATE,
					addrUpdate.LinkAddress.String(), int32(intf.Index), intf.Name)
				asicdClientMgr.SendIPv4IntfStateNotifyMsg(asicdClntDefs.NOTIFY_IPV4_L3INTF_STATE_CHANGE,
					addrUpdate.LinkAddress.String(), int32(intf.Index), uint8(asicdClntDefs.INTF_STATE_UP))
			}
		} else {
			if len(addrUpdate.LinkAddress.IP) == net.IPv4len {
				asicdClientMgr.SendIPv4IntfStateNotifyMsg(asicdClntDefs.NOTIFY_IPV4_L3INTF_STATE_CHANGE,
					addrUpdate.LinkAddress.String(), int32(intf.Index), uint8(asicdClntDefs.INTF_STATE_DOWN))
				asicdClientMgr.SendIPv4IntfNotifyMsg(asicdClntDefs.NOTIFY_IPV4INTF_DELETE,
					addrUpdate.LinkAddress.String(), int32(intf.Index), intf.Name)
			} else if len(addrUpdate.LinkAddress.IP) == net.IPv6len {
				asicdClientMgr.SendIPv4IntfStateNotifyMsg(asicdClntDefs.NOTIFY_IPV4_L3INTF_STATE_CHANGE,
					addrUpdate.LinkAddress.String(), int32(intf.Index), uint8(asicdClntDefs.INTF_STATE_DOWN))
				asicdClientMgr.SendIPv4IntfNotifyMsg(asicdClntDefs.NOTIFY_IPV4INTF_DELETE,
					addrUpdate.LinkAddress.String(), int32(intf.Index), intf.Name)
			}
		}
	}
}

func (asicdClientMgr *CPSAsicdClntMgr) SendVlanNotifyMsg(msgType int, vlanId uint16, ifIdx int32, vlanName string) {
	var msg asicdClntDefs.VlanNotifyMsg
	msg.MsgType = uint8(msgType)
	msg.VlanId = vlanId
	msg.VlanIfIndex = ifIdx
	msg.VlanName = vlanName
	asicdClientMgr.SendNotification(msg)
}

func (asicdClientMgr *CPSAsicdClntMgr) LinkNetLinkSubscriber() {
	for {
		linkUpdate := <-asicdClientMgr.LinkSubCh
		linkAttr := linkUpdate.Link.Attrs()
		intf, err := net.InterfaceByIndex(linkAttr.Index)
		if err != nil {
			Logger.Err("Error getting the interface using ifIndex")
			continue
		}
		if !strings.HasPrefix(intf.Name, "br") {
			continue
		}

		vlanIdStr := strings.TrimPrefix(intf.Name, "br")
		vlanId, err := strconv.Atoi(vlanIdStr)
		if err != nil {
			Logger.Err("Error getting the vlan Id")
			continue
		}
		if linkAttr.Flags & net.FlagUp == 0 {
			asicdClientMgr.SendVlanNotifyMsg(asicdClntDefs.NOTIFY_VLAN_DELETE,
				uint16(vlanId), int32(linkAttr.Index), intf.Name)
		} else {
			asicdClientMgr.SendVlanNotifyMsg(asicdClntDefs.NOTIFY_VLAN_CREATE,
				uint16(vlanId), int32(linkAttr.Index), intf.Name)
		}
	}
}
