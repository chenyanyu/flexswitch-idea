//
//Copyright [2016] [SnapRoute Inc]
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.
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

type DWDMModuleKey struct {
	ModuleId uint8
}

//Module wide event id's
const (
	ModuleTempHiAlarm EventId = iota + 0
	ModuleTempHiAlarmClear
	ModuleTempHiWarn
	ModuleTempHiWarnClear
	ModuleTempLoAlarm
	ModuleTempLoAlarmClear
	ModuleTempLoWarn
	ModuleTempLoWarnClear
	ModuleVoltageHiAlarm
	ModuleVoltageHiAlarmClear
	ModuleVoltageHiWarn
	ModuleVoltageHiWarnClear
	ModuleVoltageLoAlarm
	ModuleVoltageLoAlarmClear
	ModuleVoltageLoWarn
	ModuleVoltageLoWarnClear
)

type DWDMModuleNwIntfKey struct {
	ModuleId uint8
	NwIntfId uint8
}

//Module NW Intf event id's
const (
	RXLOS EventId = iota + 128
	RXLOSClear
	TxPwrHiAlarm
	TxPwrHiAlarmClear
	TxPwrHiWarn
	TxPwrHiWarnClear
	TxPwrLoWarn
	TxPwrLoWarnClear
	TxPwrLoAlarm
	TxPwrLoAlarmClear
	RxPwrHiAlarm
	RxPwrHiAlarmClear
	RxPwrHiWarn
	RxPwrHiWarnClear
	RxPwrLoWarn
	RxPwrLoWarnClear
	RxPwrLoAlarm
	RxPwrLoAlarmClear
)

var OpticdEventKeyMap KeyMap = KeyMap{
	"DWDMModule":       DWDMModuleKey{},
	"DWDMModuleNwIntf": DWDMModuleNwIntfKey{},
}

func (obj DWDMModuleKey) GetObjDBKey(bytes []byte) (string, string, error) {
	var err error
	if len(bytes) == 0 {
		return "", "", errors.New("Empty byte stream")
	}
	if err = json.Unmarshal(bytes, &obj); err != nil {
		return "", "", err
	}
	return fmt.Sprintf("ModuleId:%d", obj.ModuleId), fmt.Sprintf("DWDMModule#%d", obj.ModuleId), nil
}

func (obj DWDMModuleNwIntfKey) GetObjDBKey(bytes []byte) (string, string, error) {
	var err error
	if len(bytes) == 0 {
		return "", "", errors.New("Empty byte stream")
	}
	if err = json.Unmarshal(bytes, &obj); err != nil {
		return "", "", err
	}
	return fmt.Sprintf("ModuleId:%d NwIntfId:%d", obj.ModuleId, obj.NwIntfId), fmt.Sprintf("DWDMModuleNwIntf#%d#%d", obj.ModuleId, obj.NwIntfId), nil
}
