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

type FanSensorKey struct {
	Name string
}

type TemperatureSensorKey struct {
	Name string
}

type VoltageSensorKey struct {
	Name string
}

type PowerConverterSensorKey struct {
	Name string
}

type QsfpKey struct {
	QsfpId int32
}

type QsfpChannelKey struct {
	QsfpId     int32
	ChannelNum int32
}

const (
	FanHigherTCAAlarmClear             EventId = 1
	FanHigherTCAAlarm                  EventId = 2
	FanHigherTCAWarnClear              EventId = 3
	FanHigherTCAWarn                   EventId = 4
	FanLowerTCAAlarmClear              EventId = 5
	FanLowerTCAAlarm                   EventId = 6
	FanLowerTCAWarnClear               EventId = 7
	FanLowerTCAWarn                    EventId = 8
	TemperatureHigherTCAAlarmClear     EventId = 9
	TemperatureHigherTCAAlarm          EventId = 10
	TemperatureHigherTCAWarnClear      EventId = 11
	TemperatureHigherTCAWarn           EventId = 12
	TemperatureLowerTCAAlarmClear      EventId = 13
	TemperatureLowerTCAAlarm           EventId = 14
	TemperatureLowerTCAWarnClear       EventId = 15
	TemperatureLowerTCAWarn            EventId = 16
	VoltageHigherTCAAlarmClear         EventId = 17
	VoltageHigherTCAAlarm              EventId = 18
	VoltageHigherTCAWarnClear          EventId = 19
	VoltageHigherTCAWarn               EventId = 20
	VoltageLowerTCAAlarmClear          EventId = 21
	VoltageLowerTCAAlarm               EventId = 22
	VoltageLowerTCAWarnClear           EventId = 23
	VoltageLowerTCAWarn                EventId = 24
	PowerConverterHigherTCAAlarmClear  EventId = 25
	PowerConverterHigherTCAAlarm       EventId = 26
	PowerConverterHigherTCAWarnClear   EventId = 27
	PowerConverterHigherTCAWarn        EventId = 28
	PowerConverterLowerTCAAlarmClear   EventId = 29
	PowerConverterLowerTCAAlarm        EventId = 30
	PowerConverterLowerTCAWarnClear    EventId = 31
	PowerConverterLowerTCAWarn         EventId = 32
	QsfpTemperatureHigherTCAAlarmClear EventId = 33
	QsfpTemperatureHigherTCAAlarm      EventId = 34
	QsfpTemperatureHigherTCAWarnClear  EventId = 35
	QsfpTemperatureHigherTCAWarn       EventId = 36
	QsfpTemperatureLowerTCAAlarmClear  EventId = 37
	QsfpTemperatureLowerTCAAlarm       EventId = 38
	QsfpTemperatureLowerTCAWarnClear   EventId = 39
	QsfpTemperatureLowerTCAWarn        EventId = 40
	QsfpVoltageHigherTCAAlarmClear     EventId = 41
	QsfpVoltageHigherTCAAlarm          EventId = 42
	QsfpVoltageHigherTCAWarnClear      EventId = 43
	QsfpVoltageHigherTCAWarn           EventId = 44
	QsfpVoltageLowerTCAAlarmClear      EventId = 45
	QsfpVoltageLowerTCAAlarm           EventId = 46
	QsfpVoltageLowerTCAWarnClear       EventId = 47
	QsfpVoltageLowerTCAWarn            EventId = 48
	QsfpRXPowerHigherTCAAlarmClear     EventId = 49
	QsfpRXPowerHigherTCAAlarm          EventId = 50
	QsfpRXPowerHigherTCAWarnClear      EventId = 51
	QsfpRXPowerHigherTCAWarn           EventId = 52
	QsfpRXPowerLowerTCAAlarmClear      EventId = 53
	QsfpRXPowerLowerTCAAlarm           EventId = 54
	QsfpRXPowerLowerTCAWarnClear       EventId = 55
	QsfpRXPowerLowerTCAWarn            EventId = 56
	QsfpTXPowerHigherTCAAlarmClear     EventId = 57
	QsfpTXPowerHigherTCAAlarm          EventId = 58
	QsfpTXPowerHigherTCAWarnClear      EventId = 59
	QsfpTXPowerHigherTCAWarn           EventId = 60
	QsfpTXPowerLowerTCAAlarmClear      EventId = 61
	QsfpTXPowerLowerTCAAlarm           EventId = 62
	QsfpTXPowerLowerTCAWarnClear       EventId = 63
	QsfpTXPowerLowerTCAWarn            EventId = 64
	QsfpTXBiasHigherTCAAlarmClear      EventId = 65
	QsfpTXBiasHigherTCAAlarm           EventId = 66
	QsfpTXBiasHigherTCAWarnClear       EventId = 67
	QsfpTXBiasHigherTCAWarn            EventId = 68
	QsfpTXBiasLowerTCAAlarmClear       EventId = 69
	QsfpTXBiasLowerTCAAlarm            EventId = 70
	QsfpTXBiasLowerTCAWarnClear        EventId = 71
	QsfpTXBiasLowerTCAWarn             EventId = 72
)

var PlatformdEventKeyMap KeyMap = KeyMap{
	"FanSensor":            FanSensorKey{},
	"TemperatureSensor":    TemperatureSensorKey{},
	"VoltageSensor":        VoltageSensorKey{},
	"PowerConverterSensor": PowerConverterSensorKey{},
	"Qsfp":                 QsfpKey{},
	"QsfpChannel":          QsfpChannelKey{},
}

func (obj FanSensorKey) GetObjDBKey(bytes []byte) (string, string, error) {
	var err error
	if len(bytes) == 0 {
		return "", "", errors.New("Empty byte stream")
	}
	if err = json.Unmarshal(bytes, &obj); err != nil {
		return "", "", err
	}
	return fmt.Sprintf("Name:%s", obj.Name), fmt.Sprintf("FanSensor#%s", obj.Name), nil
}

func (obj TemperatureSensorKey) GetObjDBKey(bytes []byte) (string, string, error) {
	var err error
	if len(bytes) == 0 {
		return "", "", errors.New("Empty byte stream")
	}
	if err = json.Unmarshal(bytes, &obj); err != nil {
		return "", "", err
	}
	return fmt.Sprintf("Name:%s", obj.Name), fmt.Sprintf("TemperatureSensor#%s", obj.Name), nil
}

func (obj VoltageSensorKey) GetObjDBKey(bytes []byte) (string, string, error) {
	var err error
	if len(bytes) == 0 {
		return "", "", errors.New("Empty byte stream")
	}
	if err = json.Unmarshal(bytes, &obj); err != nil {
		return "", "", err
	}
	return fmt.Sprintf("Name:%s", obj.Name), fmt.Sprintf("VoltageSensor#%s", obj.Name), nil
}

func (obj PowerConverterSensorKey) GetObjDBKey(bytes []byte) (string, string, error) {
	var err error
	if len(bytes) == 0 {
		return "", "", errors.New("Empty byte stream")
	}
	if err = json.Unmarshal(bytes, &obj); err != nil {
		return "", "", err
	}
	return fmt.Sprintf("Name:%s", obj.Name), fmt.Sprintf("PowerConverterSensor#%s", obj.Name), nil
}

func (obj QsfpKey) GetObjDBKey(bytes []byte) (string, string, error) {
	var err error
	if len(bytes) == 0 {
		return "", "", errors.New("empty byte stream")
	}
	if err = json.Unmarshal(bytes, &obj); err != nil {
		return "", "", err
	}
	return fmt.Sprintf("QsfpId:%d", obj.QsfpId), fmt.Sprintf("Qsfp#%d", obj.QsfpId), nil
}

func (obj QsfpChannelKey) GetObjDBKey(bytes []byte) (string, string, error) {
	var err error
	if len(bytes) == 0 {
		return "", "", errors.New("empty byte stream")
	}
	if err = json.Unmarshal(bytes, &obj); err != nil {
		return "", "", err
	}
	return fmt.Sprintf("QsfpId:%d ChannelNum: %d", obj.QsfpId, obj.ChannelNum), fmt.Sprintf("QsfpChannel#%d#%d", obj.QsfpId, obj.ChannelNum), nil
}
