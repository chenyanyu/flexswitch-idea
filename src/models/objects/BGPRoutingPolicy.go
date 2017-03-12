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

// BGPRoutingPolicy.go
package objects

import (
	"encoding/json"
	"fmt"
)

/*type BGPPolicyPrefix struct {
	IpPrefix string
	MaskLengthRange string
}
*/
type BGPPolicyPrefixSet struct {
	ConfigObj
	PrefixSetName string `SNAPROUTE: "KEY"`
	IpPrefixList  []PolicyPrefix
}

func (obj BGPPolicyPrefixSet) UnmarshalObject(body []byte) (ConfigObj, error) {
	var policyPrefixSet BGPPolicyPrefixSet
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &policyPrefixSet); err != nil {
			fmt.Println("### Trouble in unmarshaling BGPPolicyPrefixSet from Json", body)
		}
	}
	return policyPrefixSet, err
}

/*
type BGPPolicyDstIpMatchPrefixSetCondition struct {
	//yang_name: prefix-set class: leaf
	PrefixSet string
	//yang_name: match-set-options class: leaf
    Prefix PolicyPrefix
}
*/

type BGPPolicyAggregateAction struct {
	GenerateASSet   bool
	SendSummaryOnly bool
}
