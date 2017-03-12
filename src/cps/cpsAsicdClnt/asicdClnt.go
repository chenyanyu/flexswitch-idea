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

package cpsAsicdClnt

import (
	"utils/clntUtils/clntIntfs"
)

//var asicdMutex *sync.Mutex = &sync.Mutex{}

func (asicdClientMgr *CPSAsicdClntMgr) GetAsicdThriftClientHdl(clntInitParams *clntIntfs.BaseClntInitParams) error {
	/*
		asicdClientMgr.Name = "asicd"
		port, err := cfgParser.GetDmnPortFromClientJson(asicdClientMgr.Name, clntInitParams.ParamsFile)
		if err != nil {
			Logger.Err("Error opening client connection for asicd", err)
			return err
		}
		Logger.Debug("found asicd at port", port)
		asicdClientMgr.Address = "localhost:" + strconv.Itoa(port)
		asicdClientMgr.TTransport, asicdClientMgr.PtrProtocolFactory, err = ipcutils.CreateIPCHandles(asicdClientMgr.Address)
		if err != nil {
			Logger.Err("Failed to connect to Asicd, retrying until connection is successful")
			count := 0
			ticker := time.NewTicker(time.Duration(1000) * time.Millisecond)
			for _ = range ticker.C {
				asicdClientMgr.TTransport, asicdClientMgr.PtrProtocolFactory, err = ipcutils.CreateIPCHandles(asicdClientMgr.Address)
				if err == nil {
					ticker.Stop()
					break
				}
				count++
				if (count % 10) == 0 {
					Logger.Err("Still can't connect to Asicd, retrying..")
				}
			}

		}
		Logger.Info("Connected to Asicd")
		asicdClientMgr.ClientHdl = asicdServices.NewASICDServicesClientFactory(asicdClientMgr.TTransport, asicdClientMgr.PtrProtocolFactory)
	*/
	return nil
}

func (asicdClientMgr *CPSAsicdClntMgr) DeinitAsicdThriftClientHdl() {
	//asicdClientMgr.CloseIPCHandles()
}
