package objects

type OspfAreaEntry struct {
	ConfigObj
	AreaId          string `SNAPROUTE: "KEY", CATEGORY:"L3",  ACCESS:"w", MULTIPLICITY:"*", DESCRIPTION: A 32-bit integer uniquely identifying an area. Area ID 0.0.0.0 is used for the OSPF backbone.`
	AuthType        int32  `DESCRIPTION: The authentication type specified for an area., SELECTION: none(0)/simplePassword(1)/md5(2)`
	ImportAsExtern  int32  `DESCRIPTION: Indicates if an area is a stub area, NSSA, or standard area.  Type-5 AS-external LSAs and type-11 Opaque LSAs are not imported into stub areas or NSSAs.  NSSAs import AS-external data as type-7 LSAs, SELECTION: importExternal(1)/importNoExternal(2)/importNssa(3)`
	AreaSummary     int32  `DESCRIPTION: The variable ospfAreaSummary controls the import of summary LSAs into stub and NSSA areas. It has no effect on other areas.  If it is noAreaSummary, the router will not originate summary LSAs into the stub or NSSA area. It will rely entirely on its default route.  If it is sendAreaSummary, the router will both summarize and propagate summary LSAs., SELECTION: sendAreaSummary(2)/noAreaSummary(1)`
	StubDefaultCost int32  `DESCRIPTION: "For ABR this cost indicates default cost for summary LSA.", DEFAULT:"10" `
}

type OspfAreaEntryState struct {
	ConfigObj
	AreaId          string `SNAPROUTE: "KEY", CATEGORY:"L3",  ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: A 32-bit integer uniquely identifying an area. Area ID 0.0.0.0 is used for the OSPF backbone.`
	SpfRuns         uint32 `DESCRIPTION: The number of times that the intra-area route table has been calculated using this area's link state database.  This is typically done using Dijkstra's algorithm.  Discontinuities in the value of this counter can occur at re-initialization of the management system, and at other times as indicated by the value of ospfDiscontinuityTime.`
	AreaBdrRtrCount uint32 `DESCRIPTION: The total number of Area Border Routers reachable within this area.  This is initially zero and is calculated in each Shortest Path First (SPF) pass.`
	AsBdrRtrCount   uint32 `DESCRIPTION: The total number of Autonomous System Border Routers reachable within this area.  This is initially zero and is calculated in each SPF pass.`
	AreaLsaCount    uint32 `DESCRIPTION: The total number of link state advertisements in this area's link state database, excluding AS-external LSAs.`
}

type OspfLsdbEntryState struct {
	baseObj
	LsdbType          int32  `SNAPROUTE: "KEY", CATEGORY:"L3",  ACCESS:"r",  MULTIPLICITY:"*", DESCRIPTION: "The type of the link state advertisement. Each link state type has a separate advertisement format.  Note: External link state advertisements are permitted for backward compatibility, but should be displayed in the AsLsdbTable rather than here., SELECTION: routerLink(1)/asSummaryLink(4)/asExternalLink(5)/nssaExternalLink(7)/networkLink(2)/multicastLink(6)/summaryLink(3)/areaOpaqueLink(10)", USESTATEDB:"true"`
	LsdbLsid          string `SNAPROUTE: "KEY", CATEGORY:"L3",  DESCRIPTION: The Link State ID is an LS Type Specific field containing either a Router ID or an IP address; it identifies the piece of the routing domain that is being described by the advertisement.`
	LsdbAreaId        string `SNAPROUTE: "KEY", CATEGORY:"L3",  DESCRIPTION: The 32-bit identifier of the area from which the LSA was received.`
	LsdbRouterId      string `SNAPROUTE: "KEY", CATEGORY:"L3",  DESCRIPTION: The 32-bit number that uniquely identifies the originating router in the Autonomous System.`
	LsdbSequence      int32  `DESCRIPTION: The sequence number field is a signed 32-bit integer.  It starts with the value '80000001'h, or -'7FFFFFFF'h, and increments until '7FFFFFFF'h. Thus, a typical sequence number will be very negative. It is used to detect old and duplicate Link State Advertisements.  The space of sequence numbers is linearly ordered.  The larger the sequence number, the more recent the advertisement.`
	LsdbAge           int32  `DESCRIPTION: This field is the age of the link state advertisement in seconds.`
	LsdbChecksum      int32  `DESCRIPTION: This field is the checksum of the complete contents of the advertisement, excepting the age field.  The age field is excepted so that an advertisement's age can be incremented without updating the checksum.  The checksum used is the same that is used for ISO connectionless  datagrams; it is commonly referred to as the Fletcher checksum.`
	LsdbAdvertisement string `DESCRIPTION: The entire link state advertisement, including its header.  Note that for variable length LSAs, SNMP agents may not be able to return the largest string size.`
}

type OspfIfEntry struct {
	ConfigObj
	IfIpAddress       string `SNAPROUTE: "KEY", CATEGORY:"L3",  ACCESS:"w", MULTIPLICITY:"*", DESCRIPTION: The IP address of this OSPF interface., RELTN:"DEP:[Vlan, Port]`
	AddressLessIf     int32  `SNAPROUTE: "KEY", CATEGORY:"L3",  DESCRIPTION: For the purpose of easing the instancing of addressed and addressless interfaces; this variable takes the value 0 on interfaces with IP addresses and the corresponding value of ifIndex for interfaces having no IP address., MIN: 0, MAX: 2147483647`
	IfAdminStat       int32  `DESCRIPTION: Indiacates if OSPF is enabled on this interface`
	IfAreaId          string `DESCRIPTION: A 32-bit integer uniquely identifying the area to which the interface connects.  Area ID 0.0.0.0 is used for the OSPF backbone.`
	IfType            string `DESCRIPTION: The OSPF interface type. By way of a default, this field may be intuited from the corresponding value of ifType. Broadcast LANs, such as Ethernet and IEEE 802.5, take the value 'broadcast', X.25 and similar technologies take the value 'nbma', and links that are definitively point to point take the value 'pointToPoint'., SELECTION: broadcast/nbma/pointToPoint/pointToMultipoint`
	IfRtrPriority     int32  `DESCRIPTION: The priority of this interface.  Used in multi-access networks, this field is used in the designated router election algorithm.  The value 0 signifies that the router is not eligible to become the designated router on this particular network.  In the event of a tie in this value, routers will use their Router ID as a tie breaker., MIN: 0, MAX: 255`
	IfTransitDelay    int32  `DESCRIPTION: The estimated number of seconds it takes to transmit a link state update packet over this interface.  Note that the minimal value SHOULD be 1 second., MIN: 0, MAX: 3600`
	IfRetransInterval int32  `DESCRIPTION: The number of seconds between link state advertisement retransmissions, for adjacencies belonging to this interface.  This value is also used when retransmitting  database description and Link State request packets. Note that minimal value SHOULD be 1 second., MIN: 0, MAX:3600`
	IfHelloInterval   int32  `DESCRIPTION: The length of time, in seconds, between the Hello packets that the router sends on the interface.  This value must be the same for all routers attached to a common network., MIN: 1, MAX: 65535, DEFAULT:"10"`
	IfRtrDeadInterval int32  `DESCRIPTION: The number of seconds that a router's Hello packets have not been seen before its neighbors declare the router down. This should be some multiple of the Hello interval.  This value must be the same for all routers attached to a common network., MIN: 0, MAX: 2147483647, DEFAULT:"40"`
	IfPollInterval    int32  `DESCRIPTION: The larger time interval, in seconds, between the Hello packets sent to an inactive non-broadcast multi-access neighbor., MIN: 0, MAX: 2147483647`
}

type OspfIfEntryState struct {
	ConfigObj
	IfIpAddress                string `SNAPROUTE: "KEY", CATEGORY:"L3",   ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: The IP address of this OSPF interface.`
	AddressLessIf              int32  `SNAPROUTE: "KEY", CATEGORY:"L3",  DESCRIPTION: For the purpose of easing the instancing of addressed and addressless interfaces; this variable takes the value 0 on interfaces with IP addresses and the corresponding value of ifIndex for interfaces having no IP address., MIN: 0, MAX: 2147483647`
	IfState                    int32  `DESCRIPTION: The OSPF Interface State., SELECTION: otherDesignatedRouter(7)/backupDesignatedRouter(6)/loopback(2)/down(1)/designatedRouter(5)/waiting(3)/pointToPoint(4)`
	IfDesignatedRouter         string `DESCRIPTION: The IP address of the designated router.`
	IfBackupDesignatedRouter   string `DESCRIPTION: The IP address of the backup designated router.`
	IfEvents                   uint32 `DESCRIPTION: The number of times this OSPF interface has changed its state or an error has occurred.  Discontinuities in the value of this counter can occur at re-initialization of the management system, and at other times as indicated by the value of ospfDiscontinuityTime.`
	IfLsaCount                 uint32 `DESCRIPTION: The total number of link-local link state advertisements in this interface's link-local link state database.`
	IfDesignatedRouterId       string `DESCRIPTION: The Router ID of the designated router.`
	IfBackupDesignatedRouterId string `DESCRIPTION: The Router ID of the backup designated router.`
}

type OspfIfMetricEntry struct {
	ConfigObj
	IfMetricAddressLessIf int32  `SNAPROUTE: "KEY", CATEGORY:"L3",  ACCESS:"w", MULTIPLICITY:"*", DESCRIPTION: For the purpose of easing the instancing of addressed and addressless interfaces; this variable takes the value 0 on interfaces with IP addresses and the value of ifIndex for interfaces having no IP address.  On row creation, this can be derived from the instance., MIN: 0, MAX: 2147483647`
	IfMetricTOS           int32  `SNAPROUTE: "KEY", CATEGORY:"L3",  DESCRIPTION: The Type of Service metric being referenced. On row creation, this can be derived from the instance., MIN: 0, MAX: 30`
	IfMetricIpAddress     string `SNAPROUTE: "KEY", CATEGORY:"L3",  DESCRIPTION: The IP address of this OSPF interface.  On row creation, this can be derived from the instance.`
	IfMetricValue         int32  `DESCRIPTION: The metric of using this Type of Service on this interface.  The default value of the TOS 0 metric is 10^8 / ifSpeed., MIN: 0, MAX: 65535`
}

type OspfNbrEntryState struct {
	ConfigObj
	NbrIpAddr           string `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"r", MULTIPLICITY:"*",  DESCRIPTION: The IP address this neighbor is using in its IP source address.  Note that, on addressless links, this will not be 0.0.0.0 but the  address of another of the neighbor's interfaces.`
	NbrAddressLessIndex int32  `SNAPROUTE: "KEY", CATEGORY:"L3",  DESCRIPTION: On an interface having an IP address, zero. On addressless interfaces, the corresponding value of ifIndex in the Internet Standard MIB. On row creation, this can be derived from the instance., MIN:0, MAX: 2147483647`
	NbrRtrId            string `DESCRIPTION: A 32-bit integer (represented as a type IpAddress) uniquely identifying the neighboring router in the Autonomous System.`
	NbrOptions          int32  `DESCRIPTION: A bit mask corresponding to the neighbor's options field.  Bit 0, if set, indicates that the system will operate on Type of Service metrics other than TOS 0.  If zero, the neighbor will ignore all metrics except the TOS 0 metric.  Bit 1, if set, indicates that the associated area accepts and operates on external information; if zero, it is a stub area.  Bit 2, if set, indicates that the system is capable of routing IP multicast datagrams, that is that it implements the multicast extensions to OSPF.  Bit 3, if set, indicates that the associated area is an NSSA.  These areas are capable of carrying type-7 external advertisements, which are translated into type-5 external advertisements at NSSA borders.`
	NbrState            string `DESCRIPTION: The state of the relationship with this neighbor., SELECTION: exchangeStart(5)/loading(7)/attempt(2)/exchange(6)/down(1)/init(3)/full(8)/twoWay(4)`
	NbrEvents           uint32 `DESCRIPTION: The number of times this neighbor relationship has changed state or an error has occurred.  Discontinuities in the value of this counter can occur at re-initialization of the management system, and at other times as indicated by the value of ospfDiscontinuityTime.`
	NbrHelloSuppressed  bool   `DESCRIPTION:  *** This element is added for future use. *** Indicates whether Hellos are being suppressed to the neighbor.`
}

type OspfLsaKey struct {
	LSType    uint8  `DESCRIPTION: Link state type`
	LSId      uint32 `DESCRIPTION: Link state id`
	AdvRouter uint32 `DESCRIPTION: Advertising router`
}

type OspfNextHop struct {
	IfIPAddr  string `DESCRIPTION: O/P interface IP address`
	IfIdx     uint32 `DESCRIPTION: Interface index `
	NextHopIP string `DESCRIPTION: Nexthop ip address`
	AdvRtr    string `DESCRIPTION: Advertising router id`
}

type OspfIPv4RouteState struct {
	baseObj
	DestId          string        `SNAPROUTE: "KEY", CATEGORY:"L3",  ACCESS:"r",  MULTIPLICITY:"*", DESCRIPTION: "Dest ip" , USESTATEDB:"true"`
	AddrMask        string        ` SNAPROUTE: "KEY", CATEGORY:"L3", DESCRIPTION: "netmask"`
	DestType        string        `SNAPROUTE: "KEY", CATEGORY:"L3", DESCRIPTION: destination type`
	OptCapabilities int32         `DESCRIPTION: "capabilities", MIN: 0, MAX:2147483647`
	AreaId          string        `DESCRIPTION: area id for the route`
	PathType        string        `DESCRIPTION: "Path type such as direct / connected / ext"`
	Cost            int32         `DESCRIPTION: "Cost to reach the destination", MIN: 0, MAX:2147483647`
	Type2Cost       int32         `DESCRIPTION: "Type2 cost used for external routes.", MIN: 0, MAX:2147483647`
	NumOfPaths      int32         `DESCRIPTION: "Total number of paths", MIN: 0, MAX: 2147483647`
	NextHops        []OspfNextHop `DESCRIPTION: "Nexthops for this route"`
	LSOrigin        OspfLsaKey    `DESCRIPTION: "Ls dabatase key"`
}

type OspfGlobal struct {
	ConfigObj
	Vrf                string `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"w", MULTIPLICITY:"1", AUTOCREATE: "true", DESCRIPTION: "VRF id for OSPF global config", DEFAULT:"default"`
	RouterId           string `DESCRIPTION: A 32-bit integer uniquely identifying the router in the Autonomous System. By convention, to ensure uniqueness, this should default to the value of one of the router's IP interface addresses.  This object is persistent and when written the entity SHOULD save the change to non-volatile storage., DEFAULT:"0.0.0.0"`
	AdminStat          int32  `DESCRIPTION: Indicates if OSPF is enabled globally`
	ASBdrRtrStatus     bool   `DESCRIPTION: A flag to note whether this router is configured as an Autonomous System Border Router.  This object is persistent and when written the entity SHOULD save the change to non-volatile storage.`
	TOSSupport         bool   `DESCRIPTION: *** This element is added for future use. *** The router's support for type-of-service routing. This object is persistent and when written the entity SHOULD save the change to non-volatile storage., DEFAULT: false`
	RestartSupport     int32  `DESCRIPTION: *** This element is added for future use. *** The router's support for OSPF graceful restart. Options include: no restart support, only planned restarts, or both planned and unplanned restarts.  This object is persistent and when written the entity SHOULD save the change to non-volatile storage., SELECTION: plannedAndUnplanned(3)/none(1)/plannedOnly(2)`
	RestartInterval    int32  `DESCRIPTION: *** This element is added for future use. *** Configured OSPF graceful restart timeout interval. This object is persistent and when written the entity SHOULD save the change to non-volatile storage., MAN: 1, MAX: 1800`
	ReferenceBandwidth uint32 `DESCRIPTION: "Reference bandwidth in kilobits/second for calculating default interface metrics. Unit: Mbps", MIN: 100, MAX: 2147483647, DEFAULT: 100`
}

type OspfGlobalState struct {
	ConfigObj
	RouterId          string `SNAPROUTE: "KEY", CATEGORY:"L3",   ACCESS:"r", MULTIPLICITY:"1", DESCRIPTION: A 32-bit integer uniquely identifying the router in the Autonomous System. By convention, to ensure uniqueness, this should default to the value of one of the router's IP interface addresses.  This object is persistent and when written the entity SHOULD save the change to non-volatile storage.`
	VersionNumber     int32  `DESCRIPTION: The current version number of the OSPF protocol is 2., SELECTION: version2(2)`
	AreaBdrRtrStatus  bool   `DESCRIPTION: A flag to note whether this router is an Area Border Router.`
	ExternLsaCount    uint32 `DESCRIPTION: The number of external (LS type-5) link state advertisements in the link state database.`
	OpaqueLsaSupport  bool   `DESCRIPTION: The router's support for Opaque LSA types.`
	RestartExitReason int32  `DESCRIPTION: Describes the outcome of the last attempt at a graceful restart.  If the value is 'none', no restart has yet been attempted.  If the value is 'inProgress', a restart attempt is currently underway., SELECTION: inProgress(2)/none(1)/timedOut(4)/completed(3)/topologyChanged(5)`
}

type OspfEventState struct {
	baseObj
	Index     uint32 `SNAPROUTE: "KEY", CATEGORY:"L3", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "Event ID",  USESTATEDB:"true",`
	TimeStamp string `DESCRIPTION :"Time when the event occured"`
	EventType string `DESCRIPTION : "Type of the event"`
	EventInfo string `DESCRIPTION :"Detailed description of the event"`
}
