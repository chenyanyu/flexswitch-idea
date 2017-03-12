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

package events

import (
	"encoding/json"
	"errors"
	"fmt"
)

type PortKey struct {
	IntfRef string
}

type VlanKey struct {
	VlanId int32
}

type IPv4IntfKey struct {
	IntfRef string
}

type IPv6IntfKey struct {
	IntfRef string
}

type AsicGlobalPMTCAKey struct {
	ModuleId uint8
	Resource string
}

type EthernetPMTCAKey struct {
	IntfRef  string
	Resource string
}

const (
	_ EventId = iota
	PortOperStateUp
	PortOperStateDown
	VlanOperStateUp
	VlanOperStateDown
	IPv4IntfOperStateUp
	IPv4IntfOperStateDown
	IPv6IntfOperStateUp
	IPv6IntfOperStateDown
	TemperatureHighAlarm
	TemperatureHighAlarmClear
	TemperatureHighWarn
	TemperatureHighWarnClear
	TemperatureLowAlarm
	TemperatureLowAlarmClear
	TemperatureLowWarn
	TemperatureLowWarnClear
	EthIfUnderSizePktsHighAlarm
	EthIfUnderSizePktsHighAlarmClear
	EthIfUnderSizePktsHighWarn
	EthIfUnderSizePktsHighWarnClear
	EthIfUnderSizePktsLowAlarm
	EthIfUnderSizePktsLowAlarmClear
	EthIfUnderSizePktsLowWarn
	EthIfUnderSizePktsLowWarnClear
	EthIfOverSizePktsHighAlarm
	EthIfOverSizePktsHighAlarmClear
	EthIfOverSizePktsHighWarn
	EthIfOverSizePktsHighWarnClear
	EthIfOverSizePktsLowAlarm
	EthIfOverSizePktsLowAlarmClear
	EthIfOverSizePktsLowWarn
	EthIfOverSizePktsLowWarnClear
	EthIfFragmentedPktsHighAlarm
	EthIfFragmentedPktsHighAlarmClear
	EthIfFragmentedPktsHighWarn
	EthIfFragmentedPktsHighWarnClear
	EthIfFragmentedPktsLowAlarm
	EthIfFragmentedPktsLowAlarmClear
	EthIfFragmentedPktsLowWarn
	EthIfFragmentedPktsLowWarnClear
	EthIfCRCAlignErrHighAlarm
	EthIfCRCAlignErrHighAlarmClear
	EthIfCRCAlignErrHighWarn
	EthIfCRCAlignErrHighWarnClear
	EthIfCRCAlignErrLowAlarm
	EthIfCRCAlignErrLowAlarmClear
	EthIfCRCAlignErrLowWarn
	EthIfCRCAlignErrLowWarnClear
	EthIfJabberFramesHighAlarm
	EthIfJabberFramesHighAlarmClear
	EthIfJabberFramesHighWarn
	EthIfJabberFramesHighWarnClear
	EthIfJabberFramesLowAlarm
	EthIfJabberFramesLowAlarmClear
	EthIfJabberFramesLowWarn
	EthIfJabberFramesLowWarnClear
	EthIfEthernetPktsHighAlarm
	EthIfEthernetPktsHighAlarmClear
	EthIfEthernetPktsHighWarn
	EthIfEthernetPktsHighWarnClear
	EthIfEthernetPktsLowAlarm
	EthIfEthernetPktsLowAlarmClear
	EthIfEthernetPktsLowWarn
	EthIfEthernetPktsLowWarnClear
	EthIfMCPktsHighAlarm
	EthIfMCPktsHighAlarmClear
	EthIfMCPktsHighWarn
	EthIfMCPktsHighWarnClear
	EthIfMCPktsLowAlarm
	EthIfMCPktsLowAlarmClear
	EthIfMCPktsLowWarn
	EthIfMCPktsLowWarnClear
	EthIfBCPktsHighAlarm
	EthIfBCPktsHighAlarmClear
	EthIfBCPktsHighWarn
	EthIfBCPktsHighWarnClear
	EthIfBCPktsLowAlarm
	EthIfBCPktsLowAlarmClear
	EthIfBCPktsLowWarn
	EthIfBCPktsLowWarnClear
	EthIf64BOrLessPktsHighAlarm
	EthIf64BOrLessPktsHighAlarmClear
	EthIf64BOrLessPktsHighWarn
	EthIf64BOrLessPktsHighWarnClear
	EthIf64BOrLessPktsLowAlarm
	EthIf64BOrLessPktsLowAlarmClear
	EthIf64BOrLessPktsLowWarn
	EthIf64BOrLessPktsLowWarnClear
	EthIf65BTo127BPktsHighAlarm
	EthIf65BTo127BPktsHighAlarmClear
	EthIf65BTo127BPktsHighWarn
	EthIf65BTo127BPktsHighWarnClear
	EthIf65BTo127BPktsLowAlarm
	EthIf65BTo127BPktsLowAlarmClear
	EthIf65BTo127BPktsLowWarn
	EthIf65BTo127BPktsLowWarnClear
	EthIf128BTo255BPktsHighAlarm
	EthIf128BTo255BPktsHighAlarmClear
	EthIf128BTo255BPktsHighWarn
	EthIf128BTo255BPktsHighWarnClear
	EthIf128BTo255BPktsLowAlarm
	EthIf128BTo255BPktsLowAlarmClear
	EthIf128BTo255BPktsLowWarn
	EthIf128BTo255BPktsLowWarnClear
	EthIf256BTo511BPktsHighAlarm
	EthIf256BTo511BPktsHighAlarmClear
	EthIf256BTo511BPktsHighWarn
	EthIf256BTo511BPktsHighWarnClear
	EthIf256BTo511BPktsLowAlarm
	EthIf256BTo511BPktsLowAlarmClear
	EthIf256BTo511BPktsLowWarn
	EthIf256BTo511BPktsLowWarnClear
	EthIf512BTo1023BPktsHighAlarm
	EthIf512BTo1023BPktsHighAlarmClear
	EthIf512BTo1023BPktsHighWarn
	EthIf512BTo1023BPktsHighWarnClear
	EthIf512BTo1023BPktsLowAlarm
	EthIf512BTo1023BPktsLowAlarmClear
	EthIf512BTo1023BPktsLowWarn
	EthIf512BTo1023BPktsLowWarnClear
	EthIf1024BTo1518BPktsHighAlarm
	EthIf1024BTo1518BPktsHighAlarmClear
	EthIf1024BTo1518BPktsHighWarn
	EthIf1024BTo1518BPktsHighWarnClear
	EthIf1024BTo1518BPktsLowAlarm
	EthIf1024BTo1518BPktsLowAlarmClear
	EthIf1024BTo1518BPktsLowWarn
	EthIf1024BTo1518BPktsLowWarnClear
)

var AsicdEventKeyMap KeyMap = KeyMap{
	"Port":         PortKey{},
	"Vlan":         VlanKey{},
	"IPv4Intf":     IPv4IntfKey{},
	"IPv6Intf":     IPv6IntfKey{},
	"AsicGlobalPM": AsicGlobalPMTCAKey{},
	"EthernetPM":   EthernetPMTCAKey{},
}

func (obj PortKey) GetObjDBKey(bytes []byte) (string, string, error) {
	var err error
	if len(bytes) == 0 {
		return "", "", errors.New("Empty byte stream")
	}
	if err = json.Unmarshal(bytes, &obj); err != nil {
		return "", "", err
	}
	return fmt.Sprintf("IntfRef:%s", obj.IntfRef), fmt.Sprintf("Port#%s", obj.IntfRef), nil
}

func (obj VlanKey) GetObjDBKey(bytes []byte) (string, string, error) {
	var err error
	if len(bytes) == 0 {
		return "", "", errors.New("Empty byte stream")
	}
	if err = json.Unmarshal(bytes, &obj); err != nil {
		return "", "", err
	}
	return fmt.Sprintf("VlanId:%d", obj.VlanId), fmt.Sprintf("Vlan#%d", obj.VlanId), nil
}

func (obj IPv4IntfKey) GetObjDBKey(bytes []byte) (string, string, error) {
	var err error
	if len(bytes) == 0 {
		return "", "", errors.New("Empty byte stream")
	}
	if err = json.Unmarshal(bytes, &obj); err != nil {
		return "", "", err
	}
	return fmt.Sprintf("IntfRef:%s", obj.IntfRef), fmt.Sprintf("IPv4Intf#%s", obj.IntfRef), nil
}

func (obj IPv6IntfKey) GetObjDBKey(bytes []byte) (string, string, error) {
	var err error
	if len(bytes) == 0 {
		return "", "", errors.New("Empty byte stream")
	}
	if err = json.Unmarshal(bytes, &obj); err != nil {
		return "", "", err
	}
	return fmt.Sprintf("IntfRef:%s", obj.IntfRef), fmt.Sprintf("IPv6Intf#%s", obj.IntfRef), nil
}

func (obj AsicGlobalPMTCAKey) GetObjDBKey(bytes []byte) (string, string, error) {
	var err error
	if len(bytes) == 0 {
		return "", "", errors.New("Empty byte stream")
	}
	if err = json.Unmarshal(bytes, &obj); err != nil {
		return "", "", err
	}
	return fmt.Sprintf("ModuleId:%d Resource:%s", obj.ModuleId, obj.Resource), fmt.Sprintf("AsicGlobalPM#%d#", obj.ModuleId) + obj.Resource, nil
}

func (obj EthernetPMTCAKey) GetObjDBKey(bytes []byte) (string, string, error) {
	var err error
	if len(bytes) == 0 {
		return "", "", errors.New("Empty byte stream")
	}
	if err = json.Unmarshal(bytes, &obj); err != nil {
		return "", "", err
	}
	return fmt.Sprintf("IntfRef:%s Resource:%s", obj.IntfRef, obj.Resource), fmt.Sprintf("EthernetPM#%s#", obj.IntfRef) + obj.Resource, nil
}
