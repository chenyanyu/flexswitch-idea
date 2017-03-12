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
	"github.com/garyburd/redigo/redis"
)

func (obj EventStats) GetKey() string {
	key := "EventStats#" + fmt.Sprintf("%d", obj.EventId)
	return key
}

func (obj EventStats) StoreObjectInDb(dbHdl redis.Conn) error {
	_, err := dbHdl.Do("HMSET", redis.Args{}.Add(obj.GetKey()).AddFlat(obj)...)
	if err != nil {
		fmt.Println("Failed to store object in DB", obj)
		return err
	}
	return nil
}

func (obj EventStats) GetObjectFromDb(objKey string, dbHdl redis.Conn) (EventObj, error) {
	var object EventStats
	val, err := redis.Values(dbHdl.Do("HGETALL", objKey))
	if err != nil || len(val) == 0 {
		fmt.Println("Failed to get obj from DB", obj)
		return object, errors.New("Failed to get obj from DB")
	}
	_ = redis.ScanStruct(val, &object)

	return object, nil
}

func (obj EventStats) GetAllObjFromDb(dbHdl redis.Conn) (objList []EventObj, err error) {
	keyStr := "EventStats#*"
	keys, err := redis.Strings(dbHdl.Do("KEYS", keyStr))
	if err != nil {
		fmt.Println("Failed to get all object keys from db", obj)
		return nil, err
	}
	for idx := 0; idx < len(keys); idx++ {
		object, err := obj.GetObjectFromDb(keys[idx], dbHdl)
		if err != nil {
			fmt.Println("Failed to get object from db", obj)
			return nil, err
		}
		objList = append(objList, object)
	}
	return objList, nil
}

func (obj EventStats) UnmarshalObject(body []byte) (EventObj, error) {
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("EventStats, unmarshal failed", obj, err)
		}
	}
	return obj, err
}
