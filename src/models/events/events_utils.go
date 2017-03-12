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
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func (obj Event) GetKey() string {
	return ""
}

func (obj Event) StoreObjectInDb(dbHdl redis.Conn) error {
	return nil
}

func (obj Event) GetObjectFromDb(objKey string, dbHdl redis.Conn) (EventObj, error) {
	return obj, nil
}

func (obj Event) GetAllObjFromDb(dbHdl redis.Conn) (objList []EventObj, err error) {
	return nil, nil
}

func (obj Event) UnmarshalObject(body []byte) (EventObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("Event unmarshal failed", obj, err)
		}
	}
	return obj, err
}
