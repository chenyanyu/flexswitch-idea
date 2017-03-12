package objects

type LaPortChannelIntfRefListState struct {
	baseObj
	IntfRef                    string `SNAPROUTE: "KEY", CATEGORY:"L2",  ACCESS:"r", DESCRIPTION: Reference to aggregate member interface`
	LagIntfRef                 string `DESCRIPTION: Id of the lag group to which this port is associated with`
	OperState                  string `DESCRIPTION: The operation state, typically UP IN BUNDLE, or DOWN`
	IfIndex                    int32  `DESCRIPTION: Interface member of the LACP aggregate`
	Activity                   int32  `DESCRIPTION: Indicates participant is active or passive, SELECTION: ACTIVE(0)/PASSIVE(1)`
	Timeout                    int32  `DESCRIPTION: The timeout type (short or long) used by the participant, SELECTION: SHORT(1)/LONG(0)`
	Synchronization            int32  `DESCRIPTION: Indicates whether the participant is in-sync or out-of-sync, SELECTION: OUT_SYNC(1)/IN_SYNC(0)`
	Aggregatable               bool   `DESCRIPTION: A true value indicates that the participant will allow the link to be used as part of the aggregate. A false value indicates the link should be used as an individual link`
	Collecting                 bool   `DESCRIPTION: If true, the participant is collecting incoming frames on the link, otherwise false`
	Distributing               bool   `DESCRIPTION: When true, the participant is distributing outgoing frames; when false, distribution is disabled`
	Defaulted                  bool   `DESCRIPTION: When no partner information is exchanged port will come up in a defaulted state`
	SystemId                   string `DESCRIPTION: MAC address that defines the local system ID for the aggregate interface, SELECTION: "[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}"`
	OperKey                    uint16 `DESCRIPTION: Current operational value of the key for the aggregate interface`
	DrniName                   string `DESCRIPTION: Defines the Lag is owned by the Distributed Relay Object`
	DrniSynced                 bool   `DESCRIPTION: Identify that the Distributed Relay is in sync with neighbor`
	PartnerId                  string `DESCRIPTION: MAC address representing the protocol partners interface system ID, SELECTION: "[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}"`
	PartnerKey                 uint16 `DESCRIPTION: Operational value of the protocol partners key`
	DebugId                    uint32 `DESCRIPTION: Debug Information Id`
	RxMachine                  int32  `DESCRIPTION: Current Rx Machine State, SELECTION: RX_CURRENT(0)/RX_PORT_DISABLE(5)/RX_DEFAULTED(2)/RX_LACP_DISABLED(4)/RX_EXPIRED(1)/RX_INITIALIZE(3)`
	RxTime                     uint32 `DESCRIPTION: Time at which the last LACPDU was received by a given port,  in terms of centiseconds since the system was last reset`
	MuxMachine                 int32  `DESCRIPTION: Current MUX Machine State, SELECTION: MUX_COLLECTING(3)/MUX_COLLECTING_DISTRIBUTING_DEFAULTED(7)/MUX_COLLECTING_DISTRIBUTING(5)/MUX_DISTRIBUTING_DEFAULTED(6)/MUX_ATTACHED(2)/MUX_DETACHED(0)/MUX_DISTRIBUTING(4)/MUX_WAITING(1)`
	MuxReason                  string `DESCRIPTION: Reason for the most recent MUX state change`
	ActorChurnMachine          int32  `DESCRIPTION: Actor Churn Detection Machine State, SELECTION: CHURN_NO_CHURN(0)/CHURN_CHURN(1)`
	PartnerChurnMachine        int32  `DESCRIPTION: Partner Churn Detection Machine State, SELECTION: CHURN_NO_CHURN(0)/CHURN_CHURN(1)`
	ActorChurnCount            uint64 `DESCRIPTION: Number of times the Actor State machine has entered the  ACTOR_CHURN state`
	PartnerChurnCount          uint64 `DESCRIPTION: Number of times the Partner State machine has entered the  ACTOR_CHURN state`
	ActorSyncTransitionCount   uint64 `DESCRIPTION: Number of times the Actors Mux state machine has entered the  IN_SYNC state.`
	PartnerSyncTransitionCount uint64 `DESCRIPTION: Number of times the Partners Mux state machine has entered the  IN_SYNC state.`
	ActorChangeCount           uint64 `DESCRIPTION: Number of times the Actors perception of the LAG ID for the  Aggregation Port has changed.`
	PartnerChangeCount         uint64 `DESCRIPTION: Number of times the Partners perception of the LAG ID for the  Aggregation Port has changed.`
	ActorCdsChurnMachine       int32  `DESCRIPTION: If supported Actor CDS Churn Machine State, SELECTION: CHURN_NO_CHURN(0)/CHURN_CHURN(1)`
	PartnerCdsChurnMachine     int32  `DESCRIPTION: If supported Partner CDS Churn Machine State, SELECTION: CHURN_NO_CHURN(0)/CHURN_CHURN(1)`
	ActorCdsChurnCount         uint64 `DESCRIPTION: If supported the number of times the Actor CDS Churn state has entered the ACTOR_CDS_CHURN state`
	PartnerCdsChurnCount       uint64 `DESCRIPTION: If supported the number of times the Actor CDS Churn state has entered the ACTOR_CDS_CHURN state`
	LacpInPkts                 uint64 `DESCRIPTION: Number of LACPDUs received`
	LacpOutPkts                uint64 `DESCRIPTION: Number of LACPDUs transmitted`
	LacpRxErrors               uint64 `DESCRIPTION: The number of frames received that carry the Slow Protocols Ethernet Type value (IEEE Std 802.3-2008, Annex 57A.4), but contain a badly formed PDU or an illegal value of Protocol Subtype (IEEE Std 802.3-2008, Annex 57A.3).`
	LacpTxErrors               uint64 `DESCRIPTION: Number of LACPDU transmit packet errors`
	LacpUnknownErrors          uint64 `DESCRIPTION: Carry the Slow Protocols Ethernet Type value (IEEE Std 802.3-2008, Annex 57A.4), but contain an unknown PDU or are addressed to the Slow Protocols group MAC Address (IEEE Std 802.3-2008, Annex 57A.3), but do not carry the Slow Protocols Ethernet Type.`
	LacpErrors                 uint64 `DESCRIPTION: Number of LACPDU Errors; sum of all Rx Errors`
	LacpInMissMatchPkts        uint64 `DESCRIPTION: Numeer of LACPDU which are received with mismatched info from Peer`
	LampInPdu                  uint64 `DESCRIPTION: Number of LAMPDU received`
	LampInResponsePdu          uint64 `DESCRIPTION: Number of LAMPDU Response received`
	LampOutPdu                 uint64 `DESCRIPTION: Number of LAMPDU transmited`
	LampOutResponsePdu         uint64 `DESCRIPTION: Number of LAMPDU Response received`
}

type LaPortChannel struct {
	baseObj
	IntfRef        string   `SNAPROUTE: "KEY", CATEGORY:"L2", ACCESS:"w",  MULTIPLICITY:"*", DESCRIPTION: Id of the lag group`
	LagType        int32    `DESCRIPTION: Sets the type of LAG, i.e., how it is configured / maintained, SELECTION: LACP(0)/STATIC(1), DEFAULT: 0`
	MinLinks       uint16   `DESCRIPTION: Specifies the mininum number of member interfaces that must be active for the aggregate interface to be available, DEFAULT: 1`
	Interval       int32    `DESCRIPTION: Set the period between LACP messages -- uses the lacp-period-type enumeration., SELECTION: SLOW(1)/FAST(0), DEFAULT: 1`
	LacpMode       int32    `DESCRIPTION: ACTIVE is to initiate the transmission of LACP packets. PASSIVE is to wait for peer to initiate the transmission of LACP packets., SELECTION: ACTIVE(0)/PASSIVE(1), DEFAULT: 0`
	SystemIdMac    string   `DESCRIPTION: The MAC address portion of the nodes System ID. This is combined with the system priority to construct the 8-octet system-id, DEFAULT: "00-00-00-00-00-00"`
	SystemPriority uint16   `DESCRIPTION: Sytem priority used by the node on this LAG interface. Lower value is higher priority for determining which node is the controlling system., DEFAULT: 32768`
	LagHash        int32    `DESCRIPTION: The tx hashing algorithm used by the lag group, SELECTION: LAYER2(0)/LAYER3_4(2)/LAYER2_3(1), DEFAULT: 0`
	AdminState     string   `DESCRIPTION: Convenient way to disable/enable a lag group.  The behaviour should be such that all traffic should stop.  LACP frames should continue to be processed, SELECTION: UP/DOWN, DEFAULT: "UP"`
	IntfRefList    []string `DESCRIPTION: List of current member interfaces for the aggregate, expressed as references to existing interfaces`
}

type LaPortChannelState struct {
	baseObj
	IntfRef               string   `SNAPROUTE: "KEY", CATEGORY:"L2",  ACCESS:"r", DESCRIPTION: Id of the lag group`
	IfIndex               int32    `DESCRIPTION: the ifindex of the bonded interface`
	LagType               int32    `DESCRIPTION: Sets the type of LAG, i.e., how it is configured / maintained, SELECTION: LACP(0)/STATIC(1)`
	MinLinks              uint16   `DESCRIPTION: Specifies the mininum number of member interfaces that must be active for the aggregate interface to be available`
	Interval              int32    `DESCRIPTION: Set the period between LACP messages -- uses the lacp-period-type enumeration., SELECTION: SLOW(1)/FAST(0), DEFAULT: 1`
	LacpMode              int32    `DESCRIPTION: ACTIVE is to initiate the transmission of LACP packets. PASSIVE is to wait for peer to initiate the transmission of LACP packets., SELECTION: ACTIVE(0)/PASSIVE(1), DEFAULT: 0`
	SystemIdMac           string   `DESCRIPTION: The MAC address portion of the nodes System ID. This is combined with the system priority to construct the 8-octet system-id"`
	SystemPriority        uint16   `DESCRIPTION: Sytem priority used by the node on this LAG interface. Lower value is higher priority for determining which node is the controlling system.`
	LagHash               int32    `DESCRIPTION: The tx hashing algorithm used by the lag group, SELECTION: LAYER2(0)/LAYER3_4(2)/LAYER2_3(1), DEFAULT: 0`
	AdminState            string   `DESCRIPTION: Convenient way to disable/enable a lag group.  The behaviour should be such that all traffic should stop.  LACP frames should continue to be processed`
	OperState             string   `DESCRIPTION: Operational status of the lag group.  If all ports are DOWN this will display DOWN.  If the group was admin disabled then will display DOWN.  No ports configured in group will display DOWN`
	IntfRefList           []string `DESCRIPTION: List of current member interfaces for the aggregate, expressed as references to existing interfaces`
	IntfRefListUpInBundle []string `DESCRIPTION: List of current member interfaces for the aggregate, expressed as references to existing interfaces`
	PresentInHw           bool     `DESCRIPTION: Boolean value indicating whether the LAG is created in H/W`
}

type IppLinkState struct {
	baseObj
	DrNameRef                    string  `SNAPROUTE: "KEY", CATEGORY:"L2", ACCESS:"r", DESCRIPTION: A human-readable text string containing the Distributed Relay Instance name to which this IPP is associated`
	IntfRef                      string  `SNAPROUTE: "KEY", CATEGORY:"L2", ACCESS:"r", DESCRIPTION: A human-readable text string containing a locally significant name for the Distributed Relay IPP port link`
	IPPID                        uint32  `DESCRIPTION: The unique identifier allocated to this IPP by the local Portal System. This attribute identifies an IPP instance among the subordinate managed objects of the containing object.`
	PortConversationPasses       []uint8 `DESCRIPTION: A read-only current operational vector of Boolean values, with one value for each possible Port Conversation ID. A 1 indicates that the Port Conversation ID is allowed to be transmitted through this IPP Intra-Portal Port, and a 0 indicates that it cannot. aIPPPortConversationPasses is referencing Ipp_Port_Conversation_Passes (9.3.4.3).`
	GatewayConversationDirection []uint8 `DESCRIPTION: A read-only current operational vector of Boolean values, with one value for each possible Port Conversation ID. A 1 indicates that the Port Conversation ID is allowed to be transmitted through this IPP Intra-Portal Port, and a 0 indicates that it cannot. aIPPPortConversationPasses is referencing Ipp_Port_Conversation_Passes (9.3.4.3).`
	DRCPDUsRx                    uint64  `DESCRIPTION: The number of valid DRCPDUs received on this IPP.`
	DRCPDUsTx                    uint64  `DESCRIPTION: The number of DRCPDUs transmitted on this IPP.`
	DRCPRxState                  int32   `DESCRIPTION: This attribute holds the value current if the DRCPDU Receive state machine for the IPP is in the CURRENT state, expired if the DRCPDU Receive state machine is in the EXPIRED state, defaulted if the DRCPDU Receive state machine is in the DEFAULTED state, initialize if the DRCPDU Receive state machine is in the INITIALIZE state, or reportToManagement if the Receive state machine is in the REPORT_TO_MANAGEMENT state, SELECTION: CURRENT(0)/INITIALIZE(3)/EXPIRED(1)/DEFAULTED(2)/REPORT_TO_MANAGEMENT(4)`
	LastRxTime                   uint32  `DESCRIPTION: The time at which the last DRCPDU was received by this IPP, in terms of centiseconds since the system was last reset. The ifLastChange object in the Interfaces MIB defined in IETF RFC 2863 is a suitable object for supplying a value for aIPPDebugLastRxTime`
	DiffPortalReason             string  `DESCRIPTION: A human-readable text string indicating the most recent set of variables that are responsible for setting the variable Differ_Portal or Differ_Conf_Portal (9.4.8) on this IPP to TRUE`
}

type DistributedRelay struct {
	baseObj
	DrniName                  string   `SNAPROUTE: "KEY", CATEGORY:"L2", ACCESS:"w", DESCRIPTION: The unique identifier allocated to this Distributed Relay by the local System. This attribute identifies a Distributed Relay instance among the subordinate managed objects of the containing object.`
	PortalAddress             string   `DESCRIPTION: A read-write identifier of a particular Portal. Portal-Addr has to be unique among at least all of the potential Portal Systems to which a given Portal System might be attached via an IPL Intra-Portal Link. Also used as the Actors System ID (6.3.2) for the emulated system`
	PortalPriority            uint16   `DESCRIPTION: A 2octet read-write value indicating the priority value associated with the Portals System ID. Also used as the Actors System Priority (6.3.2) for the emulated system., MIN: "1" ,  MAX: "65535", DEFAULT: 32768`
	ThreePortalSystem         bool     `DESCRIPTION: A read-write Boolean value indicating whether this Portal System is part of a Portal consisting of three Portal Systems or not. Value 1 stands for a Portal of three Portal Systems, value 0 stands for a Portal of two or one Portal Systems. The default value is 0, DEFAULT: "false"`
	PortalSystemNumber        uint8    `DESCRIPTION: A read-write identifier of this particular Portal System within a Portal. It is the responsibility of the network administrator to ensure that these numbers are unique among the Portal Systems with the same aDrniPortalAddr (7.4.1.1.4), MIN: "1" ,  MAX: "2"`
	IntfReflist               []string `DESCRIPTION: Read-write list of the Interface Identifiers of the Ports to the Intra-Portal Links assigned to this Distributed Relay. Each Interface Identifier, a Port ID (6.3.4), has the two least significant bits of its Port Priority (7.3.2.1.15) configured to match the Portal System Number of the attached Portal System. The number of IPLs in the list depends on the Portal topology. For a Portal of three Portal Systems two or three IPLs can be used, for a Portal of two Portal Systems a single IPL is required and for a single Portal System no IPL is required`
	IntfRef                   string   `DESCRIPTION: Read-write Interface Identifier of the Aggregator Port assigned to this Distributed Relay`
	GatewayAlgorithm          string   `DESCRIPTION: This object identifies the algorithm used by the DR Function to assign frames to a Gateway Conversation ID. Table 9-7 provides the IEEE 802.1 OUI (00:80:C2) Gateway Algorithm encodings, DEFAULT: "00-80-C2-01"`
	NeighborGatewayAlgorithm  string   `DESCRIPTION: TThis object identifies the value for the Gateway algorithm of the Neighbor Portal System, assigned by administrator or System policy for use when the Neighbor Portal Systems information is unknown. Table 9-7 provides the IEEE 802.1 OUI (00-80-C2) Gateway Algorithm encodings, DEFAULT: "00-80-C2-01"`
	NeighborPortAlgorithm     string   `DESCRIPTION: This object identifies the value for the Port Algorithm of the Neighbor Portal System, assigned by administrator or System policy for use when the Neighbor Portal System’s information is unknown. Table 6-4 provides the IEEE 802.1 OUI (00-80-C2) Port Algorithm encodings., DEFAULT: "00-80-C2-01"`
	NeighborAdminDRCPState    string   `DESCRIPTION: A string of 8 bits, corresponding to the administrative values of DRCP_State [item s) in 9.4.3.2] as transmitted by this Portal System in DRCPDUs. The first bit corresponds to bit 0 of DRCP_State (Home_Gateway), the second bit corresponds to bit 1 (Neighbor_Gateway), the third bit corresponds to bit 2 (Other_Gateway), the fourth bit corresponds to bit 3 (IPP_Activity), the fifth bit corresponds to bit 4 (DRCP_Timeout), the sixth bit corresponds to bit 5 (Gateway_Sync), the seventh bit corresponds to bit 6 (Port_Sync), and the eighth bit corresponds to bit 7 (Expired). These values allow administrative control over the values of Home_Gateway, Neighbor_Gateway, Other_Gateway, IPP_Activity, and DRCP_Timeout., DEFAULT: 00000000`
	EncapMethod               string   `DESCRIPTION: This managed object is applicable only when Network / IPL sharing by time (9.3.2.1) or Network / IPL sharing by tag (9.3.2.2) or Network / IPL sharing by encapsulation (9.3.2.3) is supported. The object identifies the value representing the encapsulation method that is used to transport IPL frames to the Neighbor Portal System when the IPL and network link are sharing the same physical link. It consists of the 3-octet OUI or CID identifying the organization that is responsible for this encapsulation and one following octet used to identify the encapsulation method defined by that organization. Table 9-11 provides the IEEE 802.1 OUI (00-80-C2) encapsulation method encodings. A Default value of 0x00-80-C2-00 indicates that the IPL is using a separate physical or Aggregation link. A value of 1 indicates that Network / IPL sharing by time (9.3.2.1) is used. A value of 2 indicates that the encapsulation method used is the same as the one used by network frames and that Network / IPL sharing by tag (9.3.2.2) is used, DEFAULT: "00-80-C2-01"`
	IntraPortalPortProtocolDA string   `DESCRIPTION: A 6-octet read-write MAC Address value specifying the DA to be used when sending DRCPDUs, corresponding to the value of DRCP_Protocol_DA in 9.4.4.1.3. Its values is one of the addresses selected from Table 9-6 and its default shall be the IEEE 802.1 Nearest non-TPMR Bridge group address (01-80-C2-00-00-03)., DEFAULT: "01-80-C2-00-00-03"`
}

type DistributedRelayState struct {
	baseObj
	DrniName                           string   `SNAPROUTE: "KEY", CATEGORY:"L2", ACCESS:"r",  DESCRIPTION: The unique identifier allocated to this Distributed Relay by the local System. This attribute identifies a Distributed Relay instance among the subordinate managed objects of the containing object.`
	Description                        string   `DESCRIPTION: A human-readable text string containing information about the Distribute Relay. This string is read-only. The contents are vendor specific`
	PortalAddress                      string   `DESCRIPTION: A read-write identifier of a particular Portal. Portal-Addr has to be unique among at least all of the potential Portal Systems to which a given Portal System might be attached via an IPL Intra-Portal Link. Also used as the Actors System ID (6.3.2) for the emulated system, SELECTION: "[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}"`
	PortalPriority                     uint16   `DESCRIPTION: A 2 octet read-write value indicating the priority value associated with the Portals System ID. Also used as the Actors System Priority (6.3.2) for the emulated system.`
	ThreePortalSystem                  bool     `DESCRIPTION: A read-write Boolean value indicating whether this Portal System is part of a Portal consisting of three Portal Systems or not. Value 1 stands for a Portal of three Portal Systems, value 0 stands for a Portal of two or one Portal Systems. The default value is 0`
	PortalSystemNumber                 uint8    `DESCRIPTION: A read-write identifier of this particular Portal System within a Portal. It is the responsibility of the network administrator to ensure that these numbers are unique among the Portal Systems with the same aDrniPortalAddr (7.4.1.1.4)`
	IntfRefList                        []string `DESCRIPTION: Read-write list of the Interface Identifiers of the Ports to the Intra-Portal Links assigned to this Distributed Relay. Each Interface Identifier, a Port ID (6.3.4), has the two least significant bits of its Port Priority (7.3.2.1.15) configured to match the Portal System Number of the attached Portal System. The number of IPLs in the list depends on the Portal topology. For a Portal of three Portal Systems two or three IPLs can be used, for a Portal of two Portal Systems a single IPL is required and for a single Portal System no IPL is required`
	IntfRef                            string   `DESCRIPTION: Read-write Interface Identifier of the Aggregator Port assigned to this Distributed Relay`
	ConvAdminGateway                   []uint32 `DESCRIPTION: There are 4096 aDrniConvAdminGateway[] variables, aDrniConvAdminGateway[0] through aDrniConvAdminGateway[4095], indexed by Gateway Conversation ID. Each contains administrative values of the Gateway selection priority list for the Distributed Relay for the referenced Gateway Conversation ID. This selection priority list, a sequence of integers for each Gateway Conversation ID, is a list of Portal System Numbers in the order of preference, highest to lowest, for the corresponding preferred Portal Systems Gateway to carry that Conversation.`
	NeighborAdminConvGatewayListDigest []uint8  `DESCRIPTION: The value for the digest of the prioritized Gateway Conversation ID-to-Gateway assignments of the Neighbor Portal System, assigned by administrator or System policy for use when the Neighbor Portal System’s information is unknown`
	NeighborAdminConvPortListDigest    []uint8  `DESCRIPTION: The value for the digest of the prioritized Port Conversation ID-to-Aggregation Port assignments of the Neighbor Portal System, assigned by administrator or System policy for use when the Neighbor Portal System’s information is unknown`
	GatewayAlgorithm                   string   `DESCRIPTION: This object identifies the algorithm used by the DR Function to assign frames to a Gateway Conversation ID. Table 9-7 provides the IEEE 802.1 OUI (00:80:C2) Gateway Algorithm encodings`
	NeighborGatewayAlgorithm           string   `DESCRIPTION: TThis object identifies the value for the Gateway algorithm of the Neighbor Portal System, assigned by administrator or System policy for use when the Neighbor Portal System’s information is unknown. Table 9-7 provides the IEEE 802.1 OUI (00-80-C2) Gateway Algorithm encodings`
	NeighborPortAlgorithm              string   `DESCRIPTION: This object identifies the value for the Port Algorithm of the Neighbor Portal System, assigned by administrator or System policy for use when the Neighbor Portal System’s information is unknown. Table 6-4 provides the IEEE 802.1 OUI (00-80-C2) Port Algorithm encodings.`
	NeighborAdminDRCPState             string   `DESCRIPTION: A string of 8 bits, corresponding to the administrative values of DRCP_State [item s) in 9.4.3.2] as transmitted by this Portal System in DRCPDUs. The first bit corresponds to bit 0 of DRCP_State (Home_Gateway), the second bit corresponds to bit 1 (Neighbor_Gateway), the third bit corresponds to bit 2 (Other_Gateway), the fourth bit corresponds to bit 3 (IPP_Activity), the fifth bit corresponds to bit 4 (DRCP_Timeout), the sixth bit corresponds to bit 5 (Gateway_Sync), the seventh bit corresponds to bit 6 (Port_Sync), and the eighth bit corresponds to bit 7 (Expired). These values allow administrative control over the values of Home_Gateway, Neighbor_Gateway, Other_Gateway, IPP_Activity, and DRCP_Timeout.`
	EncapMethod                        string   `DESCRIPTION: This managed object is applicable only when Network / IPL sharing by time (9.3.2.1) or Network / IPL sharing by tag (9.3.2.2) or Network / IPL sharing by encapsulation (9.3.2.3) is supported. The object identifies the value representing the encapsulation method that is used to transport IPL frames to the Neighbor Portal System when the IPL and network link are sharing the same physical link. It consists of the 3-octet OUI or CID identifying the organization that is responsible for this encapsulation and one following octet used to identify the encapsulation method defined by that organization. Table 9-11 provides the IEEE 802.1 OUI (00-80-C2) encapsulation method encodings. A Default value of 0x00-80-C2-00 indicates that the IPL is using a separate physical or Aggregation link. A value of 1 indicates that Network / IPL sharing by time (9.3.2.1) is used. A value of 2 indicates that the encapsulation method used is the same as the one used by network frames and that Network / IPL sharing by tag (9.3.2.2) is used`
	IPLEncapMap                        []uint32 `DESCRIPTION: This managed object is applicable only when Network / IPL sharing by tag (9.3.2.2) or Network / IPL sharing by encapsulation (9.3.2.3) is supported. Each entry represents the value of the identifier used for an IPL frame associated with that Gateway Conversation ID for the encapsulation method specified in 7.4.1.1.17. There are 1024 possible Conversation Ids in a three portal system`
	NetEncapMap                        []uint32 `DESCRIPTION: This managed object is applicable only when Network / IPL sharing by tag (9.3.2.2) is supported. Each entry represents the translated value of the identifier used for a network frame associated with that Gateway Conversation ID when the method specified in 7.4.1.1.17 is the Network / IPL sharing by tag method specified in 9.3.2.2 and the network frames need to share the tag space used by IPL frames`
	DRGatewayConversationPasses        []uint8  `DESCRIPTION: A read-only current operational vector of Boolean values, with one value for each possible Port Conversation ID. A 1 indicates that the Port Conversation ID is allowed to be distributed through this DR Function’s Aggregator, and a 0 indicates that it cannot. aDrniDRPortConversationPasses is referencing the current value of  Drni_Portal_System_Port_Conversation ()`
	PSI                                bool     `DESCRIPTION: A read-only Boolean value providing the value of PSI, which indicates whether this Portal System is isolated from the other Portal Systems within the same Portal (“TRUE”) or not (“FALSE”)`
	PortConversationControl            []bool   `DESCRIPTION: A read-write Boolean value that controls the operation of the updateDRFHomeState (9.4.11). When set to TRUE the Home Gateway Vector is set equal to Drni_Portal_System_Port_Conversation. Setting this object to TRUE is only possible when the Gateway algorithm and the Port algorithm use the same distributions methods. The default is FALSE, indicating that the Home Gateway Vector is controlled by the network control protocol`
	IntraPortalPortProtocolDA          string   `DESCRIPTION: A 6-octet read-write MAC Address value specifying the DA to be used when sending DRCPDUs, corresponding to the value of DRCP_Protocol_DA in 9.4.4.1.3. Its values is one of the addresses selected from Table 9-6 and its default shall be the IEEE 802.1 Nearest non-TPMR Bridge group address (01-80-C2-00-00-03)., SELECTION: "[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}"`
}

type LacpGlobal struct {
	baseObj
	Vrf        string `SNAPROUTE: "KEY", CATEGORY:"L2", ACCESS:"w",  MULTIPLICITY:"1", AUTOCREATE: "true", DEFAULT: "default", DESCRIPTION: global system object defining the global state of LACPD.`
	AdminState string `DESCRIPTION: Administrative state of LACPD, UP will allow for lacp configuration to be applied, DOWN will disallow and de-provision from daemon, STRLEN:"4", SELECTION: UP/DOWN, DEFAULT: "DOWN"`
}

type LacpGlobalState struct {
	baseObj
	Vrf                          string   `SNAPROUTE: "KEY", CATEGORY:"L2", ACCESS:"r",  MULTIPLICITY:"1", DEFAULT: "default", DESCRIPTION: global system object defining the global state of LACPD.`
	AdminState                   string   `DESCRIPTION: Administrative state of LACPD, UP will allow for lacp configuration to be applied, DOWN will disallow and de-provision from daemon, STRLEN:"4", SELECTION: UP/DOWN, DEFAULT: "DOWN"`
	AggList                      []string `DESCRIPTION: List of Aggregation objects that have been created`
	DistributedRelayList         []string `DESCRIPTION: List of Distributed Relay objects`
	DistributedRelayAttachedList []string `DESCRIPTION: List of Distributed Relay objects who are attached to an Aggregator in the form of "DRName-AggName"`
	AggOperStateUpList           []string `DESCRIPTION: List of Aggregation objects which have at least one port in Distributed State`
	DistributedRelayUpList       []string `DESCRIPTION: List of Distributed Relay objects whose Aggregator is in Oper State Up`
	LacpErrorsInPkts             uint64   `DESCRIPTION: Total number of Error pkts received; Unknown or Illegal`
	LacpMissMatchPkts            uint64   `DESCRIPTION: Total number of pkts received which have received missmatched info from peer`
	LacpTotalRxPkts              uint64   `DESCRIPTION: Total Rx pkts accross all LACP Ports`
	LacpTotalTxPkts              uint64   `DESCRIPTION: Total Tx pkts accross all LACP Ports`
}
