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
//   This is a auto-generated file, please do not edit!
// _______   __       __________   ___      _______.____    __    ____  __  .___________.  ______  __    __
// |   ____||  |     |   ____\  \ /  /     /       |\   \  /  \  /   / |  | |           | /      ||  |  |  |
// |  |__   |  |     |  |__   \  V  /     |   (----  \   \/    \/   /  |  |  ---|  |---- |  ,---- |  |__|  |
// |   __|  |  |     |   __|   >   <       \   \      \            /   |  |     |  |     |  |     |   __   |
// |  |     |  `----.|  |____ /  .  \  .----)   |      \    /\    /    |  |     |  |     |  `----.|  |  |  |
// |__|     |_______||_______/__/ \__\ |_______/        \__/  \__/     |__|     |__|      \______||__|  |__|
//

package clients

import (
	"models/actions"
	"models/objects"
	"utils/clntUtils/clntIntfs"
	"utils/clntUtils/clntIntfs/asicdClntIntfs"
	"utils/dbutils"
	"utils/logging"
)

//Dummy import
var _ actions.ActionObj

type ASICDClient struct {
	ClientBase
	ClntInitParams *clntIntfs.BaseClntInitParams
	Logger         logging.LoggerIntf
	ClntPlugin     asicdClntIntfs.AsicdClntIntf
	ParamsDir      string
}

func (clnt *ASICDClient) Initialize(name string, address string, logger logging.LoggerIntf, paramsDir string) {
	var err error
	clnt.Logger = logger
	clnt.Name = name
	clnt.Enabled = true
	clnt.ParamsDir = paramsDir + "/"
	clnt.ClntInitParams, err = clntIntfs.NewBaseClntInitParams(name, logger, nil, paramsDir+"/")
	if err != nil {
		logger.Err("Error Initializing base clnt for", name)
		panic(err)
	}
	return
}

func (clnt *ASICDClient) ConnectToServer() bool {
	var err error
	clnt.ClntPlugin, err = asicdClntIntfs.NewAsicdClntInit(clnt.ClntInitParams)
	if err != nil {
		clnt.Logger.Err("Error Initializing new clnt for", err)
		clnt.IsConnected = false
		return false
	}
	if clnt.ClntPlugin != nil {
		clnt.IsConnected = true
	} else {
		clnt.IsConnected = false
	}
	return true
}

func (clnt *ASICDClient) DisconnectFromServer() bool {
	if clnt.IsConnectedToServer() {
		asicdClntIntfs.AsicdClntDeinit(clnt.ClntPlugin)
		clnt.ClntPlugin = nil
		clnt.IsConnected = false
	}
	return true
}

func (clnt *ASICDClient) IsConnectedToServer() bool {
	return clnt.IsConnected
}

func (clnt *ASICDClient) DisableServer() bool {
	clnt.Enabled = false
	return true
}

func (clnt *ASICDClient) IsServerEnabled() bool {
	return clnt.Enabled
}

func (clnt *ASICDClient) GetServerName() string {
	return clnt.Name
}

func (clnt *ASICDClient) LockApiHandler() {
}

func (clnt *ASICDClient) UnlockApiHandler() {
}

func (clnt *ASICDClient) CreateObject(obj objects.ConfigObj, dbHdl *dbutils.DBUtil) (error, bool) {
	var err error
	var ok bool
	ok = true
	switch obj.(type) {

	case objects.Vlan:
		data := obj.(objects.Vlan)
		ok, err = clnt.ClntPlugin.CreateVlan(&data)

		if err == nil && ok == true {
			err = dbHdl.StoreObjectInDb(data)
			if err != nil {
				clnt.Logger.Err("Store object in DB failed:", err)
				return err, false
			}
		} else {
			clnt.Logger.Err("Create failed:", err)
			return err, false
		}
		break

	case objects.AclMacFilter:
		data := obj.(objects.AclMacFilter)
		ok, err = clnt.ClntPlugin.CreateAclMacFilter(&data)

		if err == nil && ok == true {
			err = dbHdl.StoreObjectInDb(data)
			if err != nil {
				clnt.Logger.Err("Store object in DB failed:", err)
				return err, false
			}
		} else {
			clnt.Logger.Err("Create failed:", err)
			return err, false
		}
		break

	case objects.IPv4Intf:
		data := obj.(objects.IPv4Intf)
		ok, err = clnt.ClntPlugin.CreateIPv4Intf(&data)

		if err == nil && ok == true {
			err = dbHdl.StoreObjectInDb(data)
			if err != nil {
				clnt.Logger.Err("Store object in DB failed:", err)
				return err, false
			}
		} else {
			clnt.Logger.Err("Create failed:", err)
			return err, false
		}
		break

	case objects.EthernetPM:
		data := obj.(objects.EthernetPM)
		ok, err = clnt.ClntPlugin.CreateEthernetPM(&data)

		if err == nil && ok == true {
			err = dbHdl.StoreObjectInDb(data)
			if err != nil {
				clnt.Logger.Err("Store object in DB failed:", err)
				return err, false
			}
		} else {
			clnt.Logger.Err("Create failed:", err)
			return err, false
		}
		break

	case objects.AclIpv4Filter:
		data := obj.(objects.AclIpv4Filter)
		ok, err = clnt.ClntPlugin.CreateAclIpv4Filter(&data)

		if err == nil && ok == true {
			err = dbHdl.StoreObjectInDb(data)
			if err != nil {
				clnt.Logger.Err("Store object in DB failed:", err)
				return err, false
			}
		} else {
			clnt.Logger.Err("Create failed:", err)
			return err, false
		}
		break

	case objects.LogicalIntf:
		data := obj.(objects.LogicalIntf)
		ok, err = clnt.ClntPlugin.CreateLogicalIntf(&data)

		if err == nil && ok == true {
			err = dbHdl.StoreObjectInDb(data)
			if err != nil {
				clnt.Logger.Err("Store object in DB failed:", err)
				return err, false
			}
		} else {
			clnt.Logger.Err("Create failed:", err)
			return err, false
		}
		break

	case objects.AclIpv6Filter:
		data := obj.(objects.AclIpv6Filter)
		ok, err = clnt.ClntPlugin.CreateAclIpv6Filter(&data)

		if err == nil && ok == true {
			err = dbHdl.StoreObjectInDb(data)
			if err != nil {
				clnt.Logger.Err("Store object in DB failed:", err)
				return err, false
			}
		} else {
			clnt.Logger.Err("Create failed:", err)
			return err, false
		}
		break

	case objects.Acl:
		data := obj.(objects.Acl)
		ok, err = clnt.ClntPlugin.CreateAcl(&data)

		if err == nil && ok == true {
			err = dbHdl.StoreObjectInDb(data)
			if err != nil {
				clnt.Logger.Err("Store object in DB failed:", err)
				return err, false
			}
		} else {
			clnt.Logger.Err("Create failed:", err)
			return err, false
		}
		break

	case objects.AclGlobal:
		data := obj.(objects.AclGlobal)
		ok, err = clnt.ClntPlugin.CreateAclGlobal(&data)

		if err == nil && ok == true {
			err = dbHdl.StoreObjectInDb(data)
			if err != nil {
				clnt.Logger.Err("Store object in DB failed:", err)
				return err, false
			}
		} else {
			clnt.Logger.Err("Create failed:", err)
			return err, false
		}
		break

	case objects.SubIPv6Intf:
		data := obj.(objects.SubIPv6Intf)
		ok, err = clnt.ClntPlugin.CreateSubIPv6Intf(&data)

		if err == nil && ok == true {
			err = dbHdl.StoreObjectInDb(data)
			if err != nil {
				clnt.Logger.Err("Store object in DB failed:", err)
				return err, false
			}
		} else {
			clnt.Logger.Err("Create failed:", err)
			return err, false
		}
		break

	case objects.IPv6Intf:
		data := obj.(objects.IPv6Intf)
		ok, err = clnt.ClntPlugin.CreateIPv6Intf(&data)

		if err == nil && ok == true {
			err = dbHdl.StoreObjectInDb(data)
			if err != nil {
				clnt.Logger.Err("Store object in DB failed:", err)
				return err, false
			}
		} else {
			clnt.Logger.Err("Create failed:", err)
			return err, false
		}
		break

	case objects.AsicGlobalPM:
		data := obj.(objects.AsicGlobalPM)
		ok, err = clnt.ClntPlugin.CreateAsicGlobalPM(&data)

		if err == nil && ok == true {
			err = dbHdl.StoreObjectInDb(data)
			if err != nil {
				clnt.Logger.Err("Store object in DB failed:", err)
				return err, false
			}
		} else {
			clnt.Logger.Err("Create failed:", err)
			return err, false
		}
		break

	case objects.SubIPv4Intf:
		data := obj.(objects.SubIPv4Intf)
		ok, err = clnt.ClntPlugin.CreateSubIPv4Intf(&data)

		if err == nil && ok == true {
			err = dbHdl.StoreObjectInDb(data)
			if err != nil {
				clnt.Logger.Err("Store object in DB failed:", err)
				return err, false
			}
		} else {
			clnt.Logger.Err("Create failed:", err)
			return err, false
		}
		break

	case objects.Port:
		data := obj.(objects.Port)
		ok, err = clnt.ClntPlugin.CreatePort(&data)

		if err == nil && ok == true {
			err = dbHdl.StoreObjectInDb(data)
			if err != nil {
				clnt.Logger.Err("Store object in DB failed:", err)
				return err, false
			}
		} else {
			clnt.Logger.Err("Create failed:", err)
			return err, false
		}
		break
	default:
		break
	}

	return err, ok
}

func (clnt *ASICDClient) DeleteObject(obj objects.ConfigObj, objKey string, dbHdl *dbutils.DBUtil) (error, bool) {
	var err error
	var ok bool
	ok = true
	switch obj.(type) {

	case objects.Vlan:
		data := obj.(objects.Vlan)
		ok, err = clnt.ClntPlugin.DeleteVlan(&data)

		if err == nil && ok == true {
			err = dbHdl.DeleteObjectFromDb(data)
			if err != nil {
				clnt.Logger.Err("Delete object from DB failed:", err)
				return err, false
			}
		} else {
			clnt.Logger.Err("Delete failed:", err)
			return err, false
		}
		break

	case objects.AclMacFilter:
		data := obj.(objects.AclMacFilter)
		ok, err = clnt.ClntPlugin.DeleteAclMacFilter(&data)

		if err == nil && ok == true {
			err = dbHdl.DeleteObjectFromDb(data)
			if err != nil {
				clnt.Logger.Err("Delete object from DB failed:", err)
				return err, false
			}
		} else {
			clnt.Logger.Err("Delete failed:", err)
			return err, false
		}
		break

	case objects.IPv4Intf:
		data := obj.(objects.IPv4Intf)
		ok, err = clnt.ClntPlugin.DeleteIPv4Intf(&data)

		if err == nil && ok == true {
			err = dbHdl.DeleteObjectFromDb(data)
			if err != nil {
				clnt.Logger.Err("Delete object from DB failed:", err)
				return err, false
			}
		} else {
			clnt.Logger.Err("Delete failed:", err)
			return err, false
		}
		break

	case objects.EthernetPM:
		data := obj.(objects.EthernetPM)
		ok, err = clnt.ClntPlugin.DeleteEthernetPM(&data)

		if err == nil && ok == true {
			err = dbHdl.DeleteObjectFromDb(data)
			if err != nil {
				clnt.Logger.Err("Delete object from DB failed:", err)
				return err, false
			}
		} else {
			clnt.Logger.Err("Delete failed:", err)
			return err, false
		}
		break

	case objects.AclIpv4Filter:
		data := obj.(objects.AclIpv4Filter)
		ok, err = clnt.ClntPlugin.DeleteAclIpv4Filter(&data)

		if err == nil && ok == true {
			err = dbHdl.DeleteObjectFromDb(data)
			if err != nil {
				clnt.Logger.Err("Delete object from DB failed:", err)
				return err, false
			}
		} else {
			clnt.Logger.Err("Delete failed:", err)
			return err, false
		}
		break

	case objects.LogicalIntf:
		data := obj.(objects.LogicalIntf)
		ok, err = clnt.ClntPlugin.DeleteLogicalIntf(&data)

		if err == nil && ok == true {
			err = dbHdl.DeleteObjectFromDb(data)
			if err != nil {
				clnt.Logger.Err("Delete object from DB failed:", err)
				return err, false
			}
		} else {
			clnt.Logger.Err("Delete failed:", err)
			return err, false
		}
		break

	case objects.AclIpv6Filter:
		data := obj.(objects.AclIpv6Filter)
		ok, err = clnt.ClntPlugin.DeleteAclIpv6Filter(&data)

		if err == nil && ok == true {
			err = dbHdl.DeleteObjectFromDb(data)
			if err != nil {
				clnt.Logger.Err("Delete object from DB failed:", err)
				return err, false
			}
		} else {
			clnt.Logger.Err("Delete failed:", err)
			return err, false
		}
		break

	case objects.Acl:
		data := obj.(objects.Acl)
		ok, err = clnt.ClntPlugin.DeleteAcl(&data)

		if err == nil && ok == true {
			err = dbHdl.DeleteObjectFromDb(data)
			if err != nil {
				clnt.Logger.Err("Delete object from DB failed:", err)
				return err, false
			}
		} else {
			clnt.Logger.Err("Delete failed:", err)
			return err, false
		}
		break

	case objects.AclGlobal:
		data := obj.(objects.AclGlobal)
		ok, err = clnt.ClntPlugin.DeleteAclGlobal(&data)

		if err == nil && ok == true {
			err = dbHdl.DeleteObjectFromDb(data)
			if err != nil {
				clnt.Logger.Err("Delete object from DB failed:", err)
				return err, false
			}
		} else {
			clnt.Logger.Err("Delete failed:", err)
			return err, false
		}
		break

	case objects.SubIPv6Intf:
		data := obj.(objects.SubIPv6Intf)
		ok, err = clnt.ClntPlugin.DeleteSubIPv6Intf(&data)

		if err == nil && ok == true {
			err = dbHdl.DeleteObjectFromDb(data)
			if err != nil {
				clnt.Logger.Err("Delete object from DB failed:", err)
				return err, false
			}
		} else {
			clnt.Logger.Err("Delete failed:", err)
			return err, false
		}
		break

	case objects.IPv6Intf:
		data := obj.(objects.IPv6Intf)
		ok, err = clnt.ClntPlugin.DeleteIPv6Intf(&data)

		if err == nil && ok == true {
			err = dbHdl.DeleteObjectFromDb(data)
			if err != nil {
				clnt.Logger.Err("Delete object from DB failed:", err)
				return err, false
			}
		} else {
			clnt.Logger.Err("Delete failed:", err)
			return err, false
		}
		break

	case objects.AsicGlobalPM:
		data := obj.(objects.AsicGlobalPM)
		ok, err = clnt.ClntPlugin.DeleteAsicGlobalPM(&data)

		if err == nil && ok == true {
			err = dbHdl.DeleteObjectFromDb(data)
			if err != nil {
				clnt.Logger.Err("Delete object from DB failed:", err)
				return err, false
			}
		} else {
			clnt.Logger.Err("Delete failed:", err)
			return err, false
		}
		break

	case objects.SubIPv4Intf:
		data := obj.(objects.SubIPv4Intf)
		ok, err = clnt.ClntPlugin.DeleteSubIPv4Intf(&data)

		if err == nil && ok == true {
			err = dbHdl.DeleteObjectFromDb(data)
			if err != nil {
				clnt.Logger.Err("Delete object from DB failed:", err)
				return err, false
			}
		} else {
			clnt.Logger.Err("Delete failed:", err)
			return err, false
		}
		break

	case objects.Port:
		data := obj.(objects.Port)
		ok, err = clnt.ClntPlugin.DeletePort(&data)

		if err == nil && ok == true {
			err = dbHdl.DeleteObjectFromDb(data)
			if err != nil {
				clnt.Logger.Err("Delete object from DB failed:", err)
				return err, false
			}
		} else {
			clnt.Logger.Err("Delete failed:", err)
			return err, false
		}
		break
	default:
		break
	}

	return err, ok
}

func (clnt *ASICDClient) UpdateObject(dbObj objects.ConfigObj, obj objects.ConfigObj, attrSet []bool, patchOpInfo []objects.PatchOpInfo, objKey string, dbHdl *dbutils.DBUtil) (error, bool) {
	var ok bool
	var err error
	ok = true
	err = nil

	var patchOp []*objects.PatchOpInfo = make([]*objects.PatchOpInfo, 0)
	for opIdx := 0; opIdx < len(patchOpInfo); opIdx++ {
		patchOp = append(patchOp, &patchOpInfo[opIdx])
	}
	switch obj.(type) {

	case objects.Vlan:
		// cast original object
		origdata := dbObj.(objects.Vlan)
		updatedata := obj.(objects.Vlan)
		ok, err = clnt.ClntPlugin.UpdateVlan(&origdata, &updatedata, attrSet, patchOp)
		if err == nil && ok == true {

			err = dbHdl.UpdateObjectInDb(updatedata, dbObj, attrSet)
			if err != nil {
				clnt.Logger.Err("Update object in DB failed:", err)
				return err, false
			}
		} else {
			return err, false
		}
		break

	case objects.AclMacFilter:
		// cast original object
		origdata := dbObj.(objects.AclMacFilter)
		updatedata := obj.(objects.AclMacFilter)
		ok, err = clnt.ClntPlugin.UpdateAclMacFilter(&origdata, &updatedata, attrSet, patchOp)
		if err == nil && ok == true {

			err = dbHdl.UpdateObjectInDb(updatedata, dbObj, attrSet)
			if err != nil {
				clnt.Logger.Err("Update object in DB failed:", err)
				return err, false
			}
		} else {
			return err, false
		}
		break

	case objects.IPv4Intf:
		// cast original object
		origdata := dbObj.(objects.IPv4Intf)
		updatedata := obj.(objects.IPv4Intf)
		ok, err = clnt.ClntPlugin.UpdateIPv4Intf(&origdata, &updatedata, attrSet, patchOp)
		if err == nil && ok == true {

			err = dbHdl.UpdateObjectInDb(updatedata, dbObj, attrSet)
			if err != nil {
				clnt.Logger.Err("Update object in DB failed:", err)
				return err, false
			}
		} else {
			return err, false
		}
		break

	case objects.EthernetPM:
		// cast original object
		origdata := dbObj.(objects.EthernetPM)
		updatedata := obj.(objects.EthernetPM)
		ok, err = clnt.ClntPlugin.UpdateEthernetPM(&origdata, &updatedata, attrSet, patchOp)
		if err == nil && ok == true {

			err = dbHdl.UpdateObjectInDb(updatedata, dbObj, attrSet)
			if err != nil {
				clnt.Logger.Err("Update object in DB failed:", err)
				return err, false
			}
		} else {
			return err, false
		}
		break

	case objects.AclIpv4Filter:
		// cast original object
		origdata := dbObj.(objects.AclIpv4Filter)
		updatedata := obj.(objects.AclIpv4Filter)
		ok, err = clnt.ClntPlugin.UpdateAclIpv4Filter(&origdata, &updatedata, attrSet, patchOp)
		if err == nil && ok == true {

			err = dbHdl.UpdateObjectInDb(updatedata, dbObj, attrSet)
			if err != nil {
				clnt.Logger.Err("Update object in DB failed:", err)
				return err, false
			}
		} else {
			return err, false
		}
		break

	case objects.LogicalIntf:
		// cast original object
		origdata := dbObj.(objects.LogicalIntf)
		updatedata := obj.(objects.LogicalIntf)
		ok, err = clnt.ClntPlugin.UpdateLogicalIntf(&origdata, &updatedata, attrSet, patchOp)
		if err == nil && ok == true {

			err = dbHdl.UpdateObjectInDb(updatedata, dbObj, attrSet)
			if err != nil {
				clnt.Logger.Err("Update object in DB failed:", err)
				return err, false
			}
		} else {
			return err, false
		}
		break

	case objects.AclIpv6Filter:
		// cast original object
		origdata := dbObj.(objects.AclIpv6Filter)
		updatedata := obj.(objects.AclIpv6Filter)
		ok, err = clnt.ClntPlugin.UpdateAclIpv6Filter(&origdata, &updatedata, attrSet, patchOp)
		if err == nil && ok == true {

			err = dbHdl.UpdateObjectInDb(updatedata, dbObj, attrSet)
			if err != nil {
				clnt.Logger.Err("Update object in DB failed:", err)
				return err, false
			}
		} else {
			return err, false
		}
		break

	case objects.Acl:
		// cast original object
		origdata := dbObj.(objects.Acl)
		updatedata := obj.(objects.Acl)
		ok, err = clnt.ClntPlugin.UpdateAcl(&origdata, &updatedata, attrSet, patchOp)
		if err == nil && ok == true {

			err = dbHdl.UpdateObjectInDb(updatedata, dbObj, attrSet)
			if err != nil {
				clnt.Logger.Err("Update object in DB failed:", err)
				return err, false
			}
		} else {
			return err, false
		}
		break

	case objects.AclGlobal:
		// cast original object
		origdata := dbObj.(objects.AclGlobal)
		updatedata := obj.(objects.AclGlobal)
		ok, err = clnt.ClntPlugin.UpdateAclGlobal(&origdata, &updatedata, attrSet, patchOp)
		if err == nil && ok == true {

			err = dbHdl.UpdateObjectInDb(updatedata, dbObj, attrSet)
			if err != nil {
				clnt.Logger.Err("Update object in DB failed:", err)
				return err, false
			}
		} else {
			return err, false
		}
		break

	case objects.SubIPv6Intf:
		// cast original object
		origdata := dbObj.(objects.SubIPv6Intf)
		updatedata := obj.(objects.SubIPv6Intf)
		ok, err = clnt.ClntPlugin.UpdateSubIPv6Intf(&origdata, &updatedata, attrSet, patchOp)
		if err == nil && ok == true {

			err = dbHdl.UpdateObjectInDb(updatedata, dbObj, attrSet)
			if err != nil {
				clnt.Logger.Err("Update object in DB failed:", err)
				return err, false
			}
		} else {
			return err, false
		}
		break

	case objects.IPv6Intf:
		// cast original object
		origdata := dbObj.(objects.IPv6Intf)
		updatedata := obj.(objects.IPv6Intf)
		ok, err = clnt.ClntPlugin.UpdateIPv6Intf(&origdata, &updatedata, attrSet, patchOp)
		if err == nil && ok == true {

			err = dbHdl.UpdateObjectInDb(updatedata, dbObj, attrSet)
			if err != nil {
				clnt.Logger.Err("Update object in DB failed:", err)
				return err, false
			}
		} else {
			return err, false
		}
		break

	case objects.AsicGlobalPM:
		// cast original object
		origdata := dbObj.(objects.AsicGlobalPM)
		updatedata := obj.(objects.AsicGlobalPM)
		ok, err = clnt.ClntPlugin.UpdateAsicGlobalPM(&origdata, &updatedata, attrSet, patchOp)
		if err == nil && ok == true {

			err = dbHdl.UpdateObjectInDb(updatedata, dbObj, attrSet)
			if err != nil {
				clnt.Logger.Err("Update object in DB failed:", err)
				return err, false
			}
		} else {
			return err, false
		}
		break

	case objects.SubIPv4Intf:
		// cast original object
		origdata := dbObj.(objects.SubIPv4Intf)
		updatedata := obj.(objects.SubIPv4Intf)
		ok, err = clnt.ClntPlugin.UpdateSubIPv4Intf(&origdata, &updatedata, attrSet, patchOp)
		if err == nil && ok == true {

			err = dbHdl.UpdateObjectInDb(updatedata, dbObj, attrSet)
			if err != nil {
				clnt.Logger.Err("Update object in DB failed:", err)
				return err, false
			}
		} else {
			return err, false
		}
		break

	case objects.Port:
		// cast original object
		origdata := dbObj.(objects.Port)
		updatedata := obj.(objects.Port)
		ok, err = clnt.ClntPlugin.UpdatePort(&origdata, &updatedata, attrSet, patchOp)
		if err == nil && ok == true {

			err = dbHdl.UpdateObjectInDb(updatedata, dbObj, attrSet)
			if err != nil {
				clnt.Logger.Err("Update object in DB failed:", err)
				return err, false
			}
		} else {
			return err, false
		}
		break

	default:
		break
	}
	return err, ok

}

func (clnt *ASICDClient) GetBulkObject(obj objects.ConfigObj, dbHdl *dbutils.DBUtil, currMarker int64, count int64) (err error,
	objCount int64,
	nextMarker int64,
	more bool,
	objs []objects.ConfigObj) {
	switch obj.(type) {

	case objects.LinkScopeIpState:

		err, objCount, nextMarker, more, objs = dbHdl.GetBulkObjFromDb(obj, currMarker, count)
		if err != nil {
			return nil, objCount, nextMarker, more, objs
		}
		break
	case objects.BufferGlobalStatState:
		bulkInfo, err := clnt.ClntPlugin.GetBulkBufferGlobalStatState(int(currMarker), int(count))
		if err == nil && bulkInfo != nil {
			objCount = int64(bulkInfo.Count)
			nextMarker = int64(bulkInfo.EndIdx)
			more = bulkInfo.More
			for i := 0; i < len(bulkInfo.BufferGlobalStatStateList); i++ {
				objs = append(objs, bulkInfo.BufferGlobalStatStateList[i])
			}
		}
		break

	case objects.IPv6IntfState:
		bulkInfo, err := clnt.ClntPlugin.GetBulkIPv6IntfState(int(currMarker), int(count))
		if err == nil && bulkInfo != nil {
			objCount = int64(bulkInfo.Count)
			nextMarker = int64(bulkInfo.EndIdx)
			more = bulkInfo.More
			for i := 0; i < len(bulkInfo.IPv6IntfStateList); i++ {
				objs = append(objs, bulkInfo.IPv6IntfStateList[i])
			}
		}
		break

	case objects.AsicGlobalState:
		bulkInfo, err := clnt.ClntPlugin.GetBulkAsicGlobalState(int(currMarker), int(count))
		if err == nil && bulkInfo != nil {
			objCount = int64(bulkInfo.Count)
			nextMarker = int64(bulkInfo.EndIdx)
			more = bulkInfo.More
			for i := 0; i < len(bulkInfo.AsicGlobalStateList); i++ {
				objs = append(objs, bulkInfo.AsicGlobalStateList[i])
			}
		}
		break

	case objects.SubIPv6IntfState:
		bulkInfo, err := clnt.ClntPlugin.GetBulkSubIPv6IntfState(int(currMarker), int(count))
		if err == nil && bulkInfo != nil {
			objCount = int64(bulkInfo.Count)
			nextMarker = int64(bulkInfo.EndIdx)
			more = bulkInfo.More
			for i := 0; i < len(bulkInfo.SubIPv6IntfStateList); i++ {
				objs = append(objs, bulkInfo.SubIPv6IntfStateList[i])
			}
		}
		break

	case objects.NdpEntryHwState:
		bulkInfo, err := clnt.ClntPlugin.GetBulkNdpEntryHwState(int(currMarker), int(count))
		if err == nil && bulkInfo != nil {
			objCount = int64(bulkInfo.Count)
			nextMarker = int64(bulkInfo.EndIdx)
			more = bulkInfo.More
			for i := 0; i < len(bulkInfo.NdpEntryHwStateList); i++ {
				objs = append(objs, bulkInfo.NdpEntryHwStateList[i])
			}
		}
		break

	case objects.AclState:
		bulkInfo, err := clnt.ClntPlugin.GetBulkAclState(int(currMarker), int(count))
		if err == nil && bulkInfo != nil {
			objCount = int64(bulkInfo.Count)
			nextMarker = int64(bulkInfo.EndIdx)
			more = bulkInfo.More
			for i := 0; i < len(bulkInfo.AclStateList); i++ {
				objs = append(objs, bulkInfo.AclStateList[i])
			}
		}
		break

	case objects.IPv4IntfState:
		bulkInfo, err := clnt.ClntPlugin.GetBulkIPv4IntfState(int(currMarker), int(count))
		if err == nil && bulkInfo != nil {
			objCount = int64(bulkInfo.Count)
			nextMarker = int64(bulkInfo.EndIdx)
			more = bulkInfo.More
			for i := 0; i < len(bulkInfo.IPv4IntfStateList); i++ {
				objs = append(objs, bulkInfo.IPv4IntfStateList[i])
			}
		}
		break

	case objects.PortState:
		bulkInfo, err := clnt.ClntPlugin.GetBulkPortState(int(currMarker), int(count))
		if err == nil && bulkInfo != nil {
			objCount = int64(bulkInfo.Count)
			nextMarker = int64(bulkInfo.EndIdx)
			more = bulkInfo.More
			for i := 0; i < len(bulkInfo.PortStateList); i++ {
				objs = append(objs, bulkInfo.PortStateList[i])
			}
		}
		break

	case objects.EthernetPM:
		bulkInfo, err := clnt.ClntPlugin.GetBulkEthernetPM(int(currMarker), int(count))
		if err == nil && bulkInfo != nil {
			objCount = int64(bulkInfo.Count)
			nextMarker = int64(bulkInfo.EndIdx)
			more = bulkInfo.More
			for i := 0; i < len(bulkInfo.EthernetPMList); i++ {
				objs = append(objs, bulkInfo.EthernetPMList[i])
			}
		}
		break

	case objects.SubIPv4IntfState:
		bulkInfo, err := clnt.ClntPlugin.GetBulkSubIPv4IntfState(int(currMarker), int(count))
		if err == nil && bulkInfo != nil {
			objCount = int64(bulkInfo.Count)
			nextMarker = int64(bulkInfo.EndIdx)
			more = bulkInfo.More
			for i := 0; i < len(bulkInfo.SubIPv4IntfStateList); i++ {
				objs = append(objs, bulkInfo.SubIPv4IntfStateList[i])
			}
		}
		break

	case objects.MacTableEntryState:

		err, objCount, nextMarker, more, objs = dbHdl.GetBulkObjFromDb(obj, currMarker, count)
		if err != nil {
			return nil, objCount, nextMarker, more, objs
		}
		break
	case objects.BufferPortStatState:
		bulkInfo, err := clnt.ClntPlugin.GetBulkBufferPortStatState(int(currMarker), int(count))
		if err == nil && bulkInfo != nil {
			objCount = int64(bulkInfo.Count)
			nextMarker = int64(bulkInfo.EndIdx)
			more = bulkInfo.More
			for i := 0; i < len(bulkInfo.BufferPortStatStateList); i++ {
				objs = append(objs, bulkInfo.BufferPortStatStateList[i])
			}
		}
		break

	case objects.AsicSummaryState:
		bulkInfo, err := clnt.ClntPlugin.GetBulkAsicSummaryState(int(currMarker), int(count))
		if err == nil && bulkInfo != nil {
			objCount = int64(bulkInfo.Count)
			nextMarker = int64(bulkInfo.EndIdx)
			more = bulkInfo.More
			for i := 0; i < len(bulkInfo.AsicSummaryStateList); i++ {
				objs = append(objs, bulkInfo.AsicSummaryStateList[i])
			}
		}
		break

	case objects.AsicGlobalPM:
		bulkInfo, err := clnt.ClntPlugin.GetBulkAsicGlobalPM(int(currMarker), int(count))
		if err == nil && bulkInfo != nil {
			objCount = int64(bulkInfo.Count)
			nextMarker = int64(bulkInfo.EndIdx)
			more = bulkInfo.More
			for i := 0; i < len(bulkInfo.AsicGlobalPMList); i++ {
				objs = append(objs, bulkInfo.AsicGlobalPMList[i])
			}
		}
		break

	case objects.ArpEntryHwState:
		bulkInfo, err := clnt.ClntPlugin.GetBulkArpEntryHwState(int(currMarker), int(count))
		if err == nil && bulkInfo != nil {
			objCount = int64(bulkInfo.Count)
			nextMarker = int64(bulkInfo.EndIdx)
			more = bulkInfo.More
			for i := 0; i < len(bulkInfo.ArpEntryHwStateList); i++ {
				objs = append(objs, bulkInfo.ArpEntryHwStateList[i])
			}
		}
		break

	case objects.CoppStatState:
		bulkInfo, err := clnt.ClntPlugin.GetBulkCoppStatState(int(currMarker), int(count))
		if err == nil && bulkInfo != nil {
			objCount = int64(bulkInfo.Count)
			nextMarker = int64(bulkInfo.EndIdx)
			more = bulkInfo.More
			for i := 0; i < len(bulkInfo.CoppStatStateList); i++ {
				objs = append(objs, bulkInfo.CoppStatStateList[i])
			}
		}
		break

	case objects.EthernetPMState:
		bulkInfo, err := clnt.ClntPlugin.GetBulkEthernetPMState(int(currMarker), int(count))
		if err == nil && bulkInfo != nil {
			objCount = int64(bulkInfo.Count)
			nextMarker = int64(bulkInfo.EndIdx)
			more = bulkInfo.More
			for i := 0; i < len(bulkInfo.EthernetPMStateList); i++ {
				objs = append(objs, bulkInfo.EthernetPMStateList[i])
			}
		}
		break

	case objects.IPv6RouteHwState:
		bulkInfo, err := clnt.ClntPlugin.GetBulkIPv6RouteHwState(int(currMarker), int(count))
		if err == nil && bulkInfo != nil {
			objCount = int64(bulkInfo.Count)
			nextMarker = int64(bulkInfo.EndIdx)
			more = bulkInfo.More
			for i := 0; i < len(bulkInfo.IPv6RouteHwStateList); i++ {
				objs = append(objs, bulkInfo.IPv6RouteHwStateList[i])
			}
		}
		break

	case objects.VlanState:
		bulkInfo, err := clnt.ClntPlugin.GetBulkVlanState(int(currMarker), int(count))
		if err == nil && bulkInfo != nil {
			objCount = int64(bulkInfo.Count)
			nextMarker = int64(bulkInfo.EndIdx)
			more = bulkInfo.More
			for i := 0; i < len(bulkInfo.VlanStateList); i++ {
				objs = append(objs, bulkInfo.VlanStateList[i])
			}
		}
		break

	case objects.AsicGlobalPMState:
		bulkInfo, err := clnt.ClntPlugin.GetBulkAsicGlobalPMState(int(currMarker), int(count))
		if err == nil && bulkInfo != nil {
			objCount = int64(bulkInfo.Count)
			nextMarker = int64(bulkInfo.EndIdx)
			more = bulkInfo.More
			for i := 0; i < len(bulkInfo.AsicGlobalPMStateList); i++ {
				objs = append(objs, bulkInfo.AsicGlobalPMStateList[i])
			}
		}
		break

	case objects.LogicalIntfState:
		bulkInfo, err := clnt.ClntPlugin.GetBulkLogicalIntfState(int(currMarker), int(count))
		if err == nil && bulkInfo != nil {
			objCount = int64(bulkInfo.Count)
			nextMarker = int64(bulkInfo.EndIdx)
			more = bulkInfo.More
			for i := 0; i < len(bulkInfo.LogicalIntfStateList); i++ {
				objs = append(objs, bulkInfo.LogicalIntfStateList[i])
			}
		}
		break

	case objects.IPv4RouteHwState:
		bulkInfo, err := clnt.ClntPlugin.GetBulkIPv4RouteHwState(int(currMarker), int(count))
		if err == nil && bulkInfo != nil {
			objCount = int64(bulkInfo.Count)
			nextMarker = int64(bulkInfo.EndIdx)
			more = bulkInfo.More
			for i := 0; i < len(bulkInfo.IPv4RouteHwStateList); i++ {
				objs = append(objs, bulkInfo.IPv4RouteHwStateList[i])
			}
		}
		break

	case objects.Port:
		bulkInfo, err := clnt.ClntPlugin.GetBulkPort(int(currMarker), int(count))
		if err == nil && bulkInfo != nil {
			objCount = int64(bulkInfo.Count)
			nextMarker = int64(bulkInfo.EndIdx)
			more = bulkInfo.More
			for i := 0; i < len(bulkInfo.PortList); i++ {
				objs = append(objs, bulkInfo.PortList[i])
			}
		}
		break

	default:
		break
	}
	return nil, objCount, nextMarker, more, objs
}

func (clnt *ASICDClient) GetObject(obj objects.ConfigObj, dbHdl *dbutils.DBUtil) (error, objects.ConfigObj) {
	switch obj.(type) {

	case objects.LinkScopeIpState:

		retObj, err := dbHdl.GetObjectFromDb(obj, obj.GetKey())
		if err != nil {
			return err, nil
		} else {
			return nil, retObj
		}
		break
	case objects.BufferGlobalStatState:
		data := obj.(objects.BufferGlobalStatState)
		retObj, err := clnt.ClntPlugin.GetBufferGlobalStatState(data.DeviceId)
		if err == nil {
			return err, retObj
		} else {
			return err, nil
		}
		break

	case objects.IPv6IntfState:
		data := obj.(objects.IPv6IntfState)
		retObj, err := clnt.ClntPlugin.GetIPv6IntfState(data.IntfRef)
		if err == nil {
			return err, retObj
		} else {
			return err, nil
		}
		break

	case objects.AsicGlobalState:
		data := obj.(objects.AsicGlobalState)
		retObj, err := clnt.ClntPlugin.GetAsicGlobalState(data.ModuleId)
		if err == nil {
			return err, retObj
		} else {
			return err, nil
		}
		break

	case objects.SubIPv6IntfState:
		data := obj.(objects.SubIPv6IntfState)
		retObj, err := clnt.ClntPlugin.GetSubIPv6IntfState(data.IntfRef, data.Type)
		if err == nil {
			return err, retObj
		} else {
			return err, nil
		}
		break

	case objects.NdpEntryHwState:
		data := obj.(objects.NdpEntryHwState)
		retObj, err := clnt.ClntPlugin.GetNdpEntryHwState(data.IpAddr)
		if err == nil {
			return err, retObj
		} else {
			return err, nil
		}
		break

	case objects.AclState:
		data := obj.(objects.AclState)
		retObj, err := clnt.ClntPlugin.GetAclState(data.AclName)
		if err == nil {
			return err, retObj
		} else {
			return err, nil
		}
		break

	case objects.IPv4IntfState:
		data := obj.(objects.IPv4IntfState)
		retObj, err := clnt.ClntPlugin.GetIPv4IntfState(data.IntfRef)
		if err == nil {
			return err, retObj
		} else {
			return err, nil
		}
		break

	case objects.PortState:
		data := obj.(objects.PortState)
		retObj, err := clnt.ClntPlugin.GetPortState(data.IntfRef)
		if err == nil {
			return err, retObj
		} else {
			return err, nil
		}
		break

	case objects.EthernetPM:
		data := obj.(objects.EthernetPM)
		retObj, err := clnt.ClntPlugin.GetEthernetPM(data.IntfRef, data.Resource)
		if err == nil {
			return err, retObj
		} else {
			return err, nil
		}
		break

	case objects.SubIPv4IntfState:
		data := obj.(objects.SubIPv4IntfState)
		retObj, err := clnt.ClntPlugin.GetSubIPv4IntfState(data.IntfRef, data.Type)
		if err == nil {
			return err, retObj
		} else {
			return err, nil
		}
		break

	case objects.MacTableEntryState:

		retObj, err := dbHdl.GetObjectFromDb(obj, obj.GetKey())
		if err != nil {
			return err, nil
		} else {
			return nil, retObj
		}
		break
	case objects.BufferPortStatState:
		data := obj.(objects.BufferPortStatState)
		retObj, err := clnt.ClntPlugin.GetBufferPortStatState(data.IntfRef)
		if err == nil {
			return err, retObj
		} else {
			return err, nil
		}
		break

	case objects.AsicSummaryState:
		data := obj.(objects.AsicSummaryState)
		retObj, err := clnt.ClntPlugin.GetAsicSummaryState(data.ModuleId)
		if err == nil {
			return err, retObj
		} else {
			return err, nil
		}
		break

	case objects.AsicGlobalPM:
		data := obj.(objects.AsicGlobalPM)
		retObj, err := clnt.ClntPlugin.GetAsicGlobalPM(data.ModuleId, data.Resource)
		if err == nil {
			return err, retObj
		} else {
			return err, nil
		}
		break

	case objects.ArpEntryHwState:
		data := obj.(objects.ArpEntryHwState)
		retObj, err := clnt.ClntPlugin.GetArpEntryHwState(data.IpAddr)
		if err == nil {
			return err, retObj
		} else {
			return err, nil
		}
		break

	case objects.CoppStatState:
		data := obj.(objects.CoppStatState)
		retObj, err := clnt.ClntPlugin.GetCoppStatState(data.Protocol)
		if err == nil {
			return err, retObj
		} else {
			return err, nil
		}
		break

	case objects.EthernetPMState:
		data := obj.(objects.EthernetPMState)
		retObj, err := clnt.ClntPlugin.GetEthernetPMState(data.IntfRef, data.Resource)
		if err == nil {
			return err, retObj
		} else {
			return err, nil
		}
		break

	case objects.IPv6RouteHwState:
		data := obj.(objects.IPv6RouteHwState)
		retObj, err := clnt.ClntPlugin.GetIPv6RouteHwState(data.DestinationNw)
		if err == nil {
			return err, retObj
		} else {
			return err, nil
		}
		break

	case objects.VlanState:
		data := obj.(objects.VlanState)
		retObj, err := clnt.ClntPlugin.GetVlanState(data.VlanId)
		if err == nil {
			return err, retObj
		} else {
			return err, nil
		}
		break

	case objects.AsicGlobalPMState:
		data := obj.(objects.AsicGlobalPMState)
		retObj, err := clnt.ClntPlugin.GetAsicGlobalPMState(data.ModuleId, data.Resource)
		if err == nil {
			return err, retObj
		} else {
			return err, nil
		}
		break

	case objects.LogicalIntfState:
		data := obj.(objects.LogicalIntfState)
		retObj, err := clnt.ClntPlugin.GetLogicalIntfState(data.Name)
		if err == nil {
			return err, retObj
		} else {
			return err, nil
		}
		break

	case objects.IPv4RouteHwState:
		data := obj.(objects.IPv4RouteHwState)
		retObj, err := clnt.ClntPlugin.GetIPv4RouteHwState(data.DestinationNw)
		if err == nil {
			return err, retObj
		} else {
			return err, nil
		}
		break

	case objects.Port:
		data := obj.(objects.Port)
		retObj, err := clnt.ClntPlugin.GetPort(data.IntfRef)
		if err == nil {
			return err, retObj
		} else {
			return err, nil
		}
		break

	default:
		break
	}
	return nil, nil
}

func (clnt *ASICDClient) ExecuteAction(obj actions.ActionObj) error {
	switch obj.(type) {

	case actions.AsicdClearCounters:
		data := obj.(actions.AsicdClearCounters)
		_, err := clnt.ClntPlugin.ExecuteActionAsicdClearCounters(&data)
		return err
	case actions.FlushMacTableEntry:
		data := obj.(actions.FlushMacTableEntry)
		_, err := clnt.ClntPlugin.ExecuteActionFlushMacTableEntry(&data)
		return err
	default:
		break
	}
	return nil
}
