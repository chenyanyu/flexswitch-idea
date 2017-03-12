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

type LLDPIntfKey struct {
	IfIndex int32
}

const (
	NeighborLearned EventId = 1
	NeighborUpdated EventId = 2
	NeighborRemoved EventId = 3
)

var LLDPEventKeyMap KeyMap = KeyMap{
	"LLDPIntf": LLDPIntfKey{},
}

func (obj LLDPIntfKey) GetObjDBKey(bytes []byte) (string, string, error) {
	var err error
	if len(bytes) == 0 {
		return "", "", errors.New("Empty byte stream")
	}
	if err = json.Unmarshal(bytes, &obj); err != nil {
		return "", "", err
	}
	return fmt.Sprintf("IfIndex:%d", obj.IfIndex), fmt.Sprintf("LLDPIntf#%d", obj.IfIndex), nil
}
