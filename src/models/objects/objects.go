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

import (
	"github.com/garyburd/redigo/redis"
)

type ConfigObj interface {
	UnmarshalObject(data []byte) (ConfigObj, error)
	UnmarshalObjectData(queryMap map[string][]string) (ConfigObj, error)
	StoreObjectInDb(dbHdl redis.Conn) error
	StoreObjectDefaultInDb(dbHdl redis.Conn) error
	DeleteObjectFromDb(dbHdl redis.Conn) error
	GetKey() string
	GetObjectFromDbByKey(objKey string, dbHdl redis.Conn) (ConfigObj, error)
	GetObjectFromDb(objKey string, dbHdl redis.Conn) (ConfigObj, error)
	CompareObjectsAndDiff(updateKeys map[string]bool, dbObj ConfigObj) ([]bool, error)
	CompareObjectDefaultAndDiff(dbObj ConfigObj) ([]bool, error)
	MergeDbAndConfigObjForPatchUpdate(dbObj ConfigObj, patchOpInfoSlice []PatchOpInfo) (ConfigObj, []bool, error)
	MergeDbAndConfigObj(dbObj ConfigObj, attrSet []bool) (ConfigObj, error)
	UpdateObjectInDb(dbV4Route ConfigObj, attrSet []bool, dbHdl redis.Conn) error
	GetAllObjFromDb(dbHdl redis.Conn) ([]ConfigObj, error)
	GetBulkObjFromDb(startIndex int64, count int64, dbHdl redis.Conn) (error, int64, int64, bool, []ConfigObj)
	MergeDbObjKeys(dbObj ConfigObj) (ConfigObj, error)
	SortObjList(objList []ConfigObj) []ConfigObj
}
type PatchOpInfo struct {
	Op    string
	Path  string
	Value string
}
type baseObj struct {
}

func (obj baseObj) UnmarshalObject(data []byte) (ConfigObj, error) {
	return nil, nil
}
func (obj baseObj) UnmarshalObjectData(data map[string][]string) (ConfigObj, error) {
	return nil, nil
}
func (obj baseObj) StoreObjectInDb(dbHdl redis.Conn) error {
	return nil
}
func (obj baseObj) StoreObjectDefaultInDb(dbHdl redis.Conn) error {
	return nil
}
func (obj baseObj) DeleteObjectFromDb(dbHdl redis.Conn) error {
	return nil
}
func (obj baseObj) GetKey() string {
	return ""
}
func (obj baseObj) GetObjectFromDbByKey(objKey string, dbHdl redis.Conn) (ConfigObj, error) {
	return nil, nil
}
func (obj baseObj) GetObjectFromDb(objKey string, dbHdl redis.Conn) (ConfigObj, error) {
	return nil, nil
}
func (obj baseObj) CompareObjectsAndDiff(updateKeys map[string]bool, dbObj ConfigObj) ([]bool, error) {
	return nil, nil
}
func (obj baseObj) CompareObjectDefaultAndDiff(dbObj ConfigObj) ([]bool, error) {
	return nil, nil
}
func (obj baseObj) MergeDbAndConfigObj(dbObj ConfigObj, attrSet []bool) (ConfigObj, error) {
	return nil, nil
}
func (obj baseObj) MergeDbAndConfigObjForPatchUpdate(dbObj ConfigObj, patchOpInfoSlice []PatchOpInfo) (ConfigObj, []bool, error) {
	return nil, nil, nil
}
func (obj baseObj) UpdateObjectInDb(dbV4Route ConfigObj, attrSet []bool, dbHdl redis.Conn) error {
	return nil
}
func (obj baseObj) GetAllObjFromDb(dbHdl redis.Conn) ([]ConfigObj, error) {
	return nil, nil
}
func (obj baseObj) GetBulkObjFromDb(startIndex int64, count int64, dbHdl redis.Conn) (error, int64, int64, bool, []ConfigObj) {
	return nil, 0, 0, false, nil
}
func (obj baseObj) MergeDbObjKeys(dbObj ConfigObj) (ConfigObj, error) {
	return nil, nil
}
func (obj baseObj) SortObjList(objList []ConfigObj) []ConfigObj {
	return nil
}
