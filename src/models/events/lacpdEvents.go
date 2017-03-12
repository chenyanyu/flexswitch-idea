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

type LacpEntryKey struct {
	IntfRef string // Port_channel/lag group
}

type LacpPortEntryKey struct {
	IntfRef string // Port_channel/lag group
}

type DrcpEntryKey struct {
	DrcpName string // Distributed Relay Group
}

type DrcpIppEntryKey struct {
	IntfRef string // Distributed Relay Ipp Link
}

const (
	LacpdEventNone EventId = iota
	LacpdEventPortOperStateUp
	LacpdEventPortOperStateDown
	LacpdEventPortPartnerInfoSync
	LacpdEventPortPartnerInfoMismatch
	LacpdEventGroupOperStateUp
	LacpdEventGroupOperStateDown
	LacpdEventDrcpNeighborInfoSync
	LacpdEventDrcpNeighborInfoMismatch
	LacpdEventDrcpIppLinkUp
	LacpdEventDrcpIppLinkDown
)

var LacpdEventKeyMap KeyMap = KeyMap{
	"LaPortChannel":                 LacpEntryKey{},
	"LaPortChannelIntfRefListState": LacpPortEntryKey{},
	"DistributedRelay":              DrcpEntryKey{},
	"IppLinkState":                  DrcpIppEntryKey{},
}

func (obj LacpEntryKey) GetObjDBKey(bytes []byte) (string, string, error) {
	var err error
	if len(bytes) == 0 {
		return "", "", errors.New("Empty byte stream")
	}
	if err = json.Unmarshal(bytes, &obj); err != nil {
		return "", "", err
	}
	return fmt.Sprintf("IntfRef:%s", obj.IntfRef), fmt.Sprintf("LaPortChannel#%s", obj.IntfRef), nil
}

func (obj LacpPortEntryKey) GetObjDBKey(bytes []byte) (string, string, error) {
	var err error
	if len(bytes) == 0 {
		return "", "", errors.New("Empty byte stream")
	}
	if err = json.Unmarshal(bytes, &obj); err != nil {
		return "", "", err
	}
	return fmt.Sprintf("IntfRef:%s", obj.IntfRef), fmt.Sprintf("LaPortChannelIntfRefListState#%s", obj.IntfRef), nil
}

func (obj DrcpEntryKey) GetObjDBKey(bytes []byte) (string, string, error) {
	var err error
	if len(bytes) == 0 {
		return "", "", errors.New("Empty byte stream")
	}
	if err = json.Unmarshal(bytes, &obj); err != nil {
		return "", "", err
	}
	return fmt.Sprintf("IntfRef:%s", obj.DrcpName), fmt.Sprintf("DistributedRelay#%s", obj.DrcpName), nil
}

func (obj DrcpIppEntryKey) GetObjDBKey(bytes []byte) (string, string, error) {
	var err error
	if len(bytes) == 0 {
		return "", "", errors.New("Empty byte stream")
	}
	if err = json.Unmarshal(bytes, &obj); err != nil {
		return "", "", err
	}
	return fmt.Sprintf("IntfRef:%s", obj.IntfRef), fmt.Sprintf("IppLinkState#%s", obj.IntfRef), nil
}
