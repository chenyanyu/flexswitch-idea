package clients

import (
	"models/objects"
	"utils/clntUtils/clntIntfs/asicdClntIntfs"
	"utils/dbutils"
)

const (
	BREAKOUT_MODE_ATTR_OFFSET = 11
)

func (clnt *ASICDClient) PostUpdateProcessing(dbObj, obj objects.ConfigObj, attrSet []bool, dbHdl *dbutils.DBUtil) error {
	switch obj.(type) {
	case objects.Port:
		if attrSet[BREAKOUT_MODE_ATTR_OFFSET] {
			objList, err := clnt.ClntPlugin.GetAllPortsWithDirtyCache()
			if err != nil {
				clnt.Logger.Err("Asicd post processing failed")
				return err
			}
			for _, obj := range objList {
				var dbObj objects.Port
				asicdClntIntfs.ConvertFromClntDefsToObjectPort(obj, &dbObj)
				dbHdl.StoreObjectInDb(dbObj)
			}
		}
	default:
		clnt.Logger.Info("Post update processing not implemented for this config object")
	}
	return nil
}
