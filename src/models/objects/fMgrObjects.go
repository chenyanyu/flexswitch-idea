//
//Copyright [2016] [SnapRoute Inc]
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//	 Unless required by applicable law or agreed to in writing, software
//	 distributed under the License is distributed on an "AS IS" BASIS,
//	 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//	 See the License for the specific language governing permissions and
//	 limitations under the License.
//
// _______  __       __________   ___      _______.____    __    ____  __  .___________.  ______  __    __
// |   ____||  |     |   ____\  \ /  /     /       |\   \  /  \  /   / |  | |           | /      ||  |  |  |
// |  |__   |  |     |  |__   \  V  /     |   (----` \   \/    \/   /  |  | `---|  |----`|  ,----'|  |__|  |
// |   __|  |  |     |   __|   >   <       \   \      \            /   |  |     |  |     |  |     |   __   |
// |  |     |  `----.|  |____ /  .  \  .----)   |      \    /\    /    |  |     |  |     |  `----.|  |  |  |
// |__|     |_______||_______/__/ \__\ |_______/        \__/  \__/     |__|     |__|      \______||__|  |__|
//

package objects

/*
 * This DS will be used while Created/Deleting Fault Manager Config
 */
type FMgrGlobal struct {
	baseObj
	// placeholder to create a key
	Vrf    string `SNAPROUTE: "KEY", CATEGORY:"Faults and Alarms", ACCESS:"w", MULTIPLICITY:"1", DESCRIPTION: "System Vrf", DEFAULT:"default"`
	Enable bool   `DESCRIPTION: "Enable Fault Manager"`
}

type FaultState struct {
	baseObj
	OwnerId          int32  `SNAPROUTE: "KEY", CATEGORY:"Faults and Alarms", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "Fault owner daemon Id picked up from events.json"`
	EventId          int32  `SNAPROUTE: "KEY", CATEGORY:"Faults and Alarms", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "Fault event id picked up from events.json"`
	OwnerName        string `SNAPROUTE: "KEY", CATEGORY:"Faults and Alarms", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "Fault owner daemon name picked up from events.json"`
	EventName        string `SNAPROUTE: "KEY", CATEGORY:"Faults and Alarms", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "Fault event name picked up from events.json"`
	SrcObjName       string `SNAPROUTE: "KEY", CATEGORY:"Faults and Alarms", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "Fault event name picked up from events.json"`
	Description      string `DESCRIPTION: "Description explaining the fault"`
	OccuranceTime    string `DESCRIPTION: "Timestamp at which fault occured"`
	SrcObjKey        string `DESCRIPTION: "Fault Object Key"`
	SrcObjUUID       string `DESCRIPTION: "Fault Object UUID"`
	ResolutionTime   string `DESCRIPTION: "Resolution Time stamp"`
	ResolutionReason string `DESCRIPTION: "Cleared/Disabled"`
}

type AlarmState struct {
	baseObj
	OwnerId          int32  `SNAPROUTE: "KEY", CATEGORY:"Faults and Alarms", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "Alarm owner daemon Id picked up from events.json"`
	EventId          int32  `SNAPROUTE: "KEY", CATEGORY:"Faults and Alarms", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "Alarm event id picked up from events.json"`
	OwnerName        string `SNAPROUTE: "KEY", CATEGORY:"Faults and Alarms", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "Alarm owner daemon name picked up from events.json"`
	EventName        string `SNAPROUTE: "KEY", CATEGORY:"Faults and Alarms", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "Alarm event name picked up from events.json"`
	SrcObjName       string `SNAPROUTE: "KEY", CATEGORY:"Faults and Alarms", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "Alarm event name picked up from events.json"`
	Severity         string `DESCRIPTION: "Alarm Severity"`
	Description      string `DESCRIPTION: "Description explaining the fault"`
	OccuranceTime    string `DESCRIPTION: "Timestamp at which fault occured"`
	SrcObjKey        string `DESCRIPTION: "Fault Object Key"`
	SrcObjUUID       string `DESCRIPTION: "Fault Object UUID"`
	ResolutionTime   string `DESCRIPTION: "Resolution Time stamp"`
	ResolutionReason string `DESCRIPTION: "Cleared/Disabled"`
}
