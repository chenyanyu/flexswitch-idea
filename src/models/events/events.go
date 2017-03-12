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
	"github.com/garyburd/redigo/redis"
)

type OwnerId uint8
type EventId uint32

type KeyMap map[string]EventObjKeyIntf

var EventKeyMap map[string]KeyMap = map[string]KeyMap{
	"ASICD":     AsicdEventKeyMap,
	"ARPD":      ArpdEventKeyMap,
	"OPTICD":    OpticdEventKeyMap,
	"BGPD":      BGPdEventKeyMap,
	"LLDP":      LLDPEventKeyMap,
	"PLATFORMD": PlatformdEventKeyMap,
	"LACPD":     LacpdEventKeyMap,
}

type Event struct {
	OwnerId        int32
	EventId        int32
	OwnerName      string
	EventName      string
	TimeStamp      string
	SrcObjName     string
	SrcObjKey      interface{}
	Description    string
	AdditionalData interface{}
}

type EventStats struct {
	EventId       EventId
	EventName     string
	NumEvents     uint32
	LastEventTime string
}

type EventObj interface {
	UnmarshalObject([]byte) (EventObj, error)
	GetKey() string
	StoreObjectInDb(redis.Conn) error
	GetObjectFromDb(string, redis.Conn) (EventObj, error)
	GetAllObjFromDb(redis.Conn) ([]EventObj, error)
}

var EventObjectMap = map[string]EventObj{
	"events":     Event{},
	"eventstats": EventStats{},
}

type EventObjKeyIntf interface {
	GetObjDBKey([]byte) (string, string, error)
}
