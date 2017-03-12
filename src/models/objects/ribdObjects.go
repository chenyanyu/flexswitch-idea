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

type NextHopInfo struct {
	NextHopIp     string `DESCRIPTION: "next hop ip of the route, DEFAULT:"0.0.0.0""`
	NextHopIntRef string `DESCRIPTION: "Interface name or ifindex of port/lag or vlan on which this next hop is configured", OPTIONAL`
	Weight        int32  `DESCRIPTION : "Weight of the next hop",DEFAULT:0, MIN:0, MAX:31, OPTIONAL `
}

type NextBestRouteInfo struct {
	Protocol    string
	NextHopList []NextHopInfo `DESCRIPTION: "List of next hops to reach this network"`
}
type IPv4Route struct {
	baseObj
	DestinationNw string `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"w", MULTIPLICITY:"*", ACCELERATED: "true", DESCRIPTION: "IP address of the route"`
	NetworkMask   string `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"w", MULTIPLICITY:"*", ACCELERATED: "true", DESCRIPTION: "mask of the route"`
	Protocol      string `DESCRIPTION :"Protocol type of the route", OPTIONAL, DEFAULT:"STATIC"`
	Cost          uint32 `DESCRIPTION :"Cost of this route", OPTIONAL, DEFAULT:0`
	NullRoute     bool   `DESCRIPTION : "Specify if this is a null route", OPTIONAL, DEFAULT:false`
	NextHop       []NextHopInfo
}

type IPv6Route struct {
	baseObj
	DestinationNw string `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"w", MULTIPLICITY:"*", ACCELERATED: "true", DESCRIPTION: "IP address of the route"`
	NetworkMask   string `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"w", MULTIPLICITY:"*", ACCELERATED: "true", DESCRIPTION: "mask of the route"`
	Protocol      string `DESCRIPTION :"Protocol type of the route", OPTIONAL, DEFAULT:"STATIC"`
	Cost          uint32 `DESCRIPTION :"Cost of this route", OPTIONAL, DEFAULT:0`
	NullRoute     bool   `DESCRIPTION : "Specify if this is a null route", OPTIONAL, DEFAULT:false`
	NextHop       []NextHopInfo
}
type IPv6RouteState struct {
	baseObj
	DestinationNw      string        `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "IP address of the route", USESTATEDB:"true"`
	Protocol           string        `DESCRIPTION :"Protocol type of the route"`
	IsNetworkReachable bool          `DESCRIPTION :"Indicates whether this network is reachable"`
	RouteCreatedTime   string        `DESCRIPTION :"Time when the route was added"`
	RouteUpdatedTime   string        `DESCRIPTION :"Time when the route was last updated"`
	NextHopList        []NextHopInfo `DESCRIPTION: "List of next hops to reach this network"`
	PolicyList         []string      `DESCRIPTION :"List of policies applied on this route"`
	NextBestRoute      NextBestRouteInfo
}

type IPv4RouteState struct {
	baseObj
	DestinationNw      string        `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "IP address of the route", USESTATEDB:"true"`
	Protocol           string        `DESCRIPTION :"Protocol type of the route"`
	IsNetworkReachable bool          `DESCRIPTION :"Indicates whether this network is reachable"`
	RouteCreatedTime   string        `DESCRIPTION :"Time when the route was added"`
	RouteUpdatedTime   string        `DESCRIPTION :"Time when the route was last updated"`
	NextHopList        []NextHopInfo `DESCRIPTION: "List of next hops to reach this network"`
	PolicyList         []string      `DESCRIPTION :"List of policies applied on this route"`
	NextBestRoute      NextBestRouteInfo
}
type RIBEventState struct {
	baseObj
	Index     uint32 `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "Event ID"`
	TimeStamp string `DESCRIPTION :"Time when the event occured"`
	EventInfo string `DESCRIPTION :"Detailed description of the event"`
}
type PolicyPrefix struct {
	Prefix          string
	MaskLengthRange string
}
type PolicyPrefixSet struct {
	baseObj
	Name       string         `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"w",MULTIPLICITY:"*",DESCRIPTION:"Policy Prefix set name.`
	PrefixList []PolicyPrefix `DESCRIPTION:"List of policy prefixes part of this prefix set."`
}
type PolicyPrefixSetState struct {
	baseObj
	Name                string         `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"r",MULTIPLICITY:"*",DESCRIPTION:"Policy Prefix set name.`
	PrefixList          []PolicyPrefix `DESCRIPTION:"List of policy prefixes part of this prefix set."`
	PolicyConditionList []string       `DESCRIPTION:"List of policy conditions using this prefix set"`
}
type PolicyCommunitySet struct {
	baseObj
	Name          string   `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"w",MULTIPLICITY:"*",DESCRIPTION:"Policy Community List name.`
	CommunityList []string `DESCRIPTION:"List of policy communities part of this community list."`
}
type PolicyCommunitySetState struct {
	baseObj
	Name                string   `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"r",MULTIPLICITY:"*",DESCRIPTION:"Policy Community list name.`
	CommunityList       []string `DESCRIPTION:"List of policy communities part of this community list."`
	PolicyConditionList []string `DESCRIPTION:"List of policy conditions using this community list"`
}
type PolicyExtendedCommunity struct {
	Type  string `DESCRIPTION: "Type of extended community",SELECTION:"Route-Target"/"Route-Origin"`
	Value string `DESCRIPTION: "A : separated value of the extended community, examples: 200:10 / 192.168.0.2:300 / 3000.200:210"`
}
type PolicyExtendedCommunitySet struct {
	baseObj
	Name                  string                    `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"w",MULTIPLICITY:"*",DESCRIPTION:"Policy Extended Community List name.`
	ExtendedCommunityList []PolicyExtendedCommunity `DESCRIPTION:"List of policy communities part of this community list."`
}
type PolicyExtendedCommunitySetState struct {
	baseObj
	Name                  string                    `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"r",MULTIPLICITY:"*",DESCRIPTION:"Policy Community list name.`
	ExtendedCommunityList []PolicyExtendedCommunity `DESCRIPTION:"List of policy extended communities part of this extended community list."`
	PolicyConditionList   []string                  `DESCRIPTION:"List of policy conditions using this extended community list"`
}
type PolicyASPathSet struct {
	baseObj
	Name       string   `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"w",MULTIPLICITY:"*",DESCRIPTION:"Policy ASPath List name.`
	ASPathList []string `DESCRIPTION:"List of ASPaths part of this list."`
}
type PolicyASPathSetState struct {
	baseObj
	Name                string   `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"r",MULTIPLICITY:"*",DESCRIPTION:"Policy ASPath list name.`
	ASPathList          []string `DESCRIPTION:"List of ASPaths part of this community list."`
	PolicyConditionList []string `DESCRIPTION:"List of policy conditions using this aspath list"`
}
type PolicyCondition struct {
	baseObj
	Name                   string `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"w", MULTIPLICITY:"*", DESCRIPTION: "PolicyConditionName"`
	ConditionType          string `DESCRIPTION: "Specifies the match criterion this condition defines", SELECTION: "MatchProtocol"/"MatchDstIpPrefix"/"MatchSrcIpPrefix"/"MatchCommunity"/"MatchExtendedCommunity"/"MatchLocalPref"/"MatchASPath"/"MatchMED"`
	Protocol               string `DESCRIPTION: "Protocol to match on if the ConditionType is set to MatchProtocol",SELECTION:"CONNECTED"/"STATIC"/"OSPF"/"BGP"`
	IpPrefix               string `DESCRIPTION: "Used in conjunction with MaskLengthRange to specify the IP Prefix to match on when the ConditionType is MatchDstIpPrefix/MatchSrcIpPrefix.", OPTIONAL, DEFAULT:""`
	MaskLengthRange        string `DESCRIPTION: "Used in conjuction with IpPrefix to specify specify the IP Prefix to match on when the ConditionType is MatchDstIpPrefix/MatchSrcIpPrefix.", OPTIONAL, DEFAULT:""`
	PrefixSet              string `DESCRIPTION: "Name of a pre-defined prefix set to be used as a condition qualifier.", OPTIONAL, DEFAULT:""`
	Community              string `DESCRIPTION: "BGP Community attrribute value to match on when the conditionType is MatchCommunity - based on RFC 1997. Can either specify the well-known communities or any other community value in the format AA:NN or 0x1234abcd format or a number.", OPTIONAL, DEFAULT:""`
	CommunitySet           string `DESCRIPTION: "List of BGP communities attribute to match on when the conditionType is MatchCommunity", OPTIONAL, DEFAULT:""`
	ExtendedCommunityType  string `DESCRIPTION: "Specifies BGP Extended Community type (used along with value)to match on when the conditionType is MatchExtendedCommunity - based on RFC 4360.",SELECTION:"Route-Target"/"Route-Origin", OPTIONAL, DEFAULT:""`
	ExtendedCommunityValue string `DESCRIPTION: "Specifies BGP Extended Community value (used along with type)to match on when the conditionType is MatchExtendedCommunity - based on RFC 4360.This is a ":" separated string.Examples: 200:10 / 192.168.0.2:300 / 3000.200:210", OPTIONAL, DEFAULT:""`
	ExtendedCommunitySet   string `DESCRIPTIONL "List of BGP Extended Community type/values to match on when the ConditionType is MatchExtendedCommunity", OPTIONAL, DEFAULT:""`
	LocalPref              uint32 `DESCRIPTION: "BGP LocalPreference attribute value to match on when the ConditionType is MatchLocalPref.", OPTIONAL, DEFAULT:0`
	ASPath                 string `DESCRIPTION: "BGP ASPath value (specified using regular expressions) to match on when ConditionType is MatchASPath.", OPTIONAL, DEFAULT:""`
	ASPathSet              string `DESCRIPTION: "List of ASPath values to match on when ConditionType is MATCHASPath", OPTIONAL, DEFAULT:""`
	MED                    uint32 `DESCRIPTION: "BGP MED value ro match on when ConditionType is MatchMED", OPTIONAL, DEFAULT:0`
}
type PolicyConditionState struct {
	baseObj
	Name           string `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "Condition name"`
	ConditionInfo  string
	PolicyStmtList []string `DESCRIPTION: "List of policy statements using this condition"`
}
type PolicyAction struct {
	Attr                   string `DESCRIPTION:"Attribute on which action is being applied",SELECTION:"Community"/"LocalPref"/"ExtendedCommunity"/"PrependASPath"/"MED"`
	Community              string `DESCRIPTION: "BGP Community attribute value when the action attr is Community.Can either specify the well-known communities or any other community value in the format AA:NN or 0x1234abcd format or a number.", OPTIONAL, DEFAULT:""`
	ExtendedCommunityType  string `DESCRIPTION: "Specifies BGP Extended Community type (used along with value)to set when the attr is ExtendedCommunity - based on RFC 4360.",SELECTION:"Route-Target"/"Route-Origin", OPTIONAL, DEFAULT:""`
	ExtendedCommunityValue string `DESCRIPTION: "Specifies BGP Extended Community value (used along with type)to set when the attr is ExtendedCommunity - based on RFC 4360.This is a ":" separated string.Examples: 200:10 / 192.168.0.2:300 / 3000.200:210", OPTIONAL, DEFAULT:""`
	LocalPref              uint32 `DESCRIPTION: "BGP LocalPreference attribute value when the action attr is LocalPref.", OPTIONAL, DEFAULT:0`
	PrependASPath          string `DESCRIPTION: "BGP ASPath Value (specified using regular expressions) to prepend when the attr is ASPath", OPTIONAL, DEFAULT:""`
	MED                    uint32 `DESCRIPTION: "BGP MED Value to set when attr is MED", OPTIONAL, DEFAULT:0`
}
type PolicyStmt struct {
	baseObj
	SetActions      []PolicyAction `DESCRIPTION : "A set of attr/value pairs to be set associatded with this statement.", OPTIONAL`
	Name            string         `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"w", MULTIPLICITY:"*", DESCRIPTION: "Policy Statement Name"`
	MatchConditions string         `DESCRIPTION :"Specifies whether to match all/any of the conditions of this policy statement",SELECTION:"any"/"all",DEFAULT:"all"`
	Conditions      []string       `DESCRIPTION :"List of conditions added to this policy statement"`
	Action          string         `DESCRIPTION :"Action for this policy statement", SELECTION:"permit"/"deny",DEFAULT: "deny"`
}
type PolicyStmtState struct {
	baseObj
	Name            string         `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "PolicyStmtState"`
	MatchConditions string         `DESCRIPTION :"Specifies whether to match all/any of the conditions of this policy statement"`
	Conditions      []string       `DESCRIPTION :"List of conditions added to this policy statement"`
	Action          string         `DESCRIPTION :"Action corresponding to this policy statement"`
	SetActions      []PolicyAction `DESCRIPTION : "A set of attr/value pairs to be set associatded with this statement."`
	PolicyList      []string       `DESCRIPTION :"List of policies using this policy statement"`
}
type PolicyDefinitionStmtPriority struct {
	Priority  int32 `DESCRIPTION:"Priority of the policy w.r.t other policies configured", MIN:0, MAX:255`
	Statement string
}
type PolicyDefinition struct {
	baseObj
	Name          string                         `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"w", MULTIPLICITY:"*", DESCRIPTION: "Policy Name"`
	Priority      int32                          `DESCRIPTION :"Priority of the policy w.r.t other policies configured", MIN: 0, MAX: 255`
	MatchType     string                         `DESCRIPTION :"Specifies whether to match all/any of the statements within this policy",SELECTION:"all"/"any",DEFAULT:"all"`
	PolicyType    string                         `DESCRIPTION : Specifies the intended protocol application for the policy", SELECTION: "BGP"/"OSPF"/"ALL", DEFAULT:"ALL"`
	StatementList []PolicyDefinitionStmtPriority `DESCRIPTION :"Specifies list of statements along with their precedence order."`
}
type PolicyDefinitionState struct {
	baseObj
	Name         string   `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "PolicyDefinitionState"`
	IpPrefixList []string `DESCRIPTION :"List of networks/IP Prefixes this policy has been applied on to."`
}

type RedistributionPolicy struct {
	baseObj
	Target string `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"w", MULTIPLICITY:"*", DESCRIPTION: "Target protocol for redistribution"`
	Source string `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"w", MULTIPLICITY:"*", DESCRIPTION: "Source Protocol for redistribution"`
	Policy string `DESCRIPTION:"Policy to be applied from source to Target"`
}
type RouteDistanceState struct {
	baseObj
	Protocol string `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "RouteDistanceState protocol"`
	Distance int32  `DESCRIPTION: "The current value of the admin distance of this protocol"`
}

type PerProtocolRouteCount struct {
	Protocol   string
	RouteCount int32
	EcmpCount  int32
}
type RouteStatState struct {
	baseObj
	Vrf                       string                  `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"r", MULTIPLICITY:"1", DESCRIPTION: "System Vrf", DEFAULT:"default"`
	TotalRouteCount           int32                   `DESCRIPTION: Total number of routes on the system`
	ECMPRouteCount            int32                   `DESCRIPTION: ECMP routes on the system`
	V4RouteCount              int32                   `DESCRIPTION: Total number of IPv4 routes on the system`
	V6RouteCount              int32                   `DESCRIPTION: Total number of IPv6 routes on the system`
	PerProtocolRouteCountList []PerProtocolRouteCount `DESCRIPTION: Per Protocol routes stats`
}
type RouteInfoSummary struct {
	DestinationNw   string        `DESCRIPTION: "IP address of the route"`
	IsInstalledInHw bool          `DESCRIPTION :"Indicates whether this route is installed in HW"`
	NextHopList     []NextHopInfo `DESCRIPTION: "List of next hops to reach this network"`
}
type RouteStatsPerProtocolState struct {
	baseObj
	Protocol string             `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"r", MULTIPLICITY:"1", DESCRIPTION :"Protocol type of the route"`
	V4Routes []RouteInfoSummary `DESCRIPTION: "Brief summary info of ipv4 routes of this protocol type"`
	V6Routes []RouteInfoSummary `DESCRIPTION: "Brief summary info of ipv6 routes of this protocol type"`
}
type RouteStatsPerInterfaceState struct {
	baseObj
	Intfref  string   `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"r", MULTIPLICITY:"1", DESCRIPTION :Interface of the next hop"`
	V4Routes []string `DESCRIPTION: "Brief summary info of ipv4 routes which have nexthop on this interface"`
	V6Routes []string `DESCRIPTION: "Brief summary info of ipv6 routes which have nexthop on this interface"`
}
