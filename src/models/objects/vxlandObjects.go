package objects

type VxlanInstance struct {
	baseObj
	Vni            uint32   `SNAPROUTE: "KEY", CATEGORY:"Tunnel",  ACCESS:"w", MULTIPLICITY:"*", DESCRIPTION: VXLAN Network Id, MIN: "1" ,  MAX: "16777215"`
	AdminState     string   `DESCRIPTION: Administrative state of VXLAN layer, UP will allow for traffic to be processed in the VNI, DOWN will drop traffic within this layer, STRLEN:"4", SELECTION: UP/DOWN, DEFAULT: "UP"`
	UntaggedVlanId []uint16 `DESCRIPTION: Vlan associated with the untagged traffic.  Used in conjunction with a given VTEP inner-vlan-handling-mode, MIN: "1" ,  MAX: "4094"`
	VlanId         []uint16 `DESCRIPTION: Vlan associated with the Access targets.  Used in conjunction with a given VTEP inner-vlan-handling-mode, MIN: "1" ,  MAX: "4094"`
}

type VxlanInstanceState struct {
	baseObj
	Vni       uint32   `SNAPROUTE: "KEY", CATEGORY:"Tunnel",  ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: VXLAN Network Id, MIN: "1" ,  MAX: "16777215"`
	OperState string   `DESCRIPTION: Operational state of VXLAN layer, UP will allow for traffic to be processed in the VNI, DOWN will drop traffic within this layer, STRLEN:"4", SELECTION: UP/DOWN, DEFAULT: "UP"`
	VlanId    []uint16 `DESCRIPTION: Vlan associated with the Access targets.  Used in conjunction with a given VTEP inner-vlan-handling-mode, MIN: "1" ,  MAX: "4094"`
}

type VxlanVtepInstance struct {
	baseObj
	Intf                  string   `SNAPROUTE: "KEY", CATEGORY:"Tunnel", ACCESS:"w", MULTIPLICITY:"*", DESCRIPTION: VTEP instance identifier name. should be defined as either vtep<id#> or <id#> if the later then 'vtep' will be prepended to the <id#> example: vtep100 or 100`
	Vni                   uint32   `SNAPROUTE: "KEY", CATEGORY:"Tunnel", DESCRIPTION: VXLAN Network ID, MIN: "1" ,  MAX: "16777215"`
	IntfRef               string   `DESCRIPTION: Source interface where the source ip will be derived from.  If an interface is not supplied the src-ip will be used. This attribute takes presedence over src-ip attribute.`
	DstUDP                uint16   `DESCRIPTION: vxlan udp port.  Deafult is the iana default udp port, DEFAULT: 4789`
	TTL                   uint16   `DESCRIPTION: TTL of the Vxlan tunnel, MIN:0, MAX:255, DEFAULT: 64`
	TOS                   uint16   `DESCRIPTION: Type of Service, MIN:0, MAX:255, DEFAULT: 0`
	InnerVlanHandlingMode int32    `DESCRIPTION: The inner vlan tag handling mode., SELECTION: DISCARD_INNER_VLAN(0)/NO_DISCARD_INNER_VLAN(1), DEFAULT: 0`
	SrcIp                 string   `DESCRIPTION: Source IP address for the VxLAN tunnel, if this is supplied it is assumed that the intf-ref is this vtep.  This  attribute will be ignored if intf-ref is set", DEFAULT: "0.0.0.0"`
	VlanId                uint16   `DESCRIPTION: Vlan Id to encapsulate with the vtep tunnel ethernet header`
	Mtu                   uint32   `DESCRIPTION: Set the MTU to be applied to all VTEP within this VxLAN, DEFAULT: 1450`
	AdminState            string   `DESCRIPTION: Administrative state of VXLAN MAC/IP layer, UP will allow for traffic to be processed in the VNI, DOWN will drop traffic within this layer, STRLEN:"4", SELECTION: UP/DOWN, DEFAULT: "UP"`
	DstIp                 []string `DESCRIPTION: Destination IP address list for the VxLAN tunnel"`
}

type VtepStateEntry struct {
	SubVtepName        string `DESCRIPTION: Internal name used to represent the vtep dev in linux, Vxland uses veth interface to represent vxlan thus this name plus Int will also be created`
	DstIp              string `DESCRIPTION: Destination IP address for the static VxLAN tunnel"`
	RxSwPkts           uint64 `DESCRIPTION: Rx Packets`
	RxSwBytes          uint64 `DESCRIPTION: Rx Bytes of Customer Frames after VTEP decap`
	RxSwDropPkts       uint64 `DESCRIPTION: Rx Dropped Packets`
	RxSwDropBytes      uint64 `DESCRIPTION: Rx Dropped Bytes of Customer Frames after VTEP decap`
	RxSwFwdPkts        uint64 `DESCRIPTION: Rx Forwaded Packets`
	RxSwFwdBytes       uint64 `DESCRIPTION: Rx Forwaded Bytes of Customer Frames after VTEP decap`
	TxSwPkts           uint64 `DESCRIPTION: Tx Packets`
	TxSwBytes          uint64 `DESCRIPTION: Tx Bytes of Customer Frames before VTEP encap`
	TxSwDropPkts       uint64 `DESCRIPTION: Tx Dropped Packets`
	TxSwDropBytes      uint64 `DESCRIPTION: Tx Dropped Bytes of Customer Frames before VTEP encap`
	TxSwFwdPkts        uint64 `DESCRIPTION: Tx Forwaded Packets`
	TxSwFwdBytes       uint64 `DESCRIPTION: Tx Forwaded Bytes of Customer Frames before VTEP encap`
	LastSwRxDropReason string `DESCRIPTION: Rx Drop Reason`
	LastSwTxDropReason string `DESCRIPTION: Tx Drop Reason`
	RxHwPkts           uint64 `DESCRIPTION: Rx Packets from HW`
	RxHwBytes          uint64 `DESCRIPTION: Rx Bytes from HW`
	TxHwPkts           uint64 `DESCRIPTION: Rx Packets from HW`
	TxHwBytes          uint64 `DESCRIPTION: Rx Bytes from HW`
	VtepFsmState       string `DESCRIPTION: Current state of the VTEP FSM UNINITIALIZED/DISABLED/INIT/DETACHED/INTERFACE/NEXT HOP INFO/RESOLVE NEXT HOP INFO/HW CONFIG/LISTENER`
	VtepFsmPrevState   string `DESCRIPTION: Previous state of the VTEP FSM UNINITIALIZED/DISABLED/INIT/DETACHED/INTERFACE/NEXT HOP INFO/RESOLVE NEXT HOP INFO/HW CONFIG/LISTENER`
	OperState          string `DESCRIPTION: Operational state of VXLAN MAC/IP layer, UP will allow for traffic to be processed in the VNI, DOWN will drop traffic within this layer, STRLEN:"4", SELECTION: UP/DOWN, DEFAULT: "DOWN"`
}

type VxlanVtepInstanceState struct {
	baseObj
	Intf                  string           `SNAPROUTE: "KEY", CATEGORY:"Tunnel", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: VTEP instance identifier name. should be defined as either vtep<id#> or <id#> if the later then 'vtep' will be prepended to the <id#> example: vtep100 or 100`
	Vni                   uint32           `SNAPROUTE: "KEY", CATEGORY:"Tunnel", DESCRIPTION: VXLAN Network ID, MIN: "1" ,  MAX: "16777215"`
	IntfRef               string           `DESCRIPTION: Source interface where the source ip will be derived from.  If an interface is not supplied the src-ip will be used. This attribute takes presedence over src-ip attribute.`
	IfIndex               int32            `DESCRIPTION: Vtep IfIndex`
	DstUDP                uint16           `DESCRIPTION: vxlan udp port.  Deafult is the iana default udp port, DEFAULT: 4789`
	TTL                   uint16           `DESCRIPTION: TTL of the Vxlan tunnel, MIN:0, MAX:255, DEFAULT: 64`
	TOS                   uint16           `DESCRIPTION: Type of Service, MIN:0, MAX:255, DEFAULT: 0`
	InnerVlanHandlingMode int32            `DESCRIPTION: The inner vlan tag handling mode., SELECTION: DISCARD_INNER_VLAN(0)/NO_DISCARD_INNER_VLAN(1), DEFAULT: 0`
	SrcIp                 string           `DESCRIPTION: Source IP address for the VxLAN tunnel, if this is supplied it is assumed that the intf-ref is this vtep.  This  attribute will be ignored if intf-ref is set", DEFAULT: "0.0.0.0"`
	VlanId                uint16           `DESCRIPTION: Vlan Id to encapsulate with the vtep tunnel ethernet header`
	Mtu                   uint32           `DESCRIPTION: Set the MTU to be applied to all VTEP within this VxLAN, DEFAULT: 1550`
	PerDstIpState         []VtepStateEntry `DESCRIPTION: Data associated with each unique vtep associated with this vni/ srcip`
}

type VxlanGlobal struct {
	baseObj
	Vrf        string `SNAPROUTE: "KEY", CATEGORY:"Tunnel", ACCESS:"w",  MULTIPLICITY:"1", AUTOCREATE: "true", DEFAULT: "default", DESCRIPTION: global system object defining the global state of VXLAND.`
	AdminState string `DESCRIPTION: Administrative state of VXLAND, UP will allow for vxlan configuration to be applied, DOWN will disallow and de-provision from daemon, STRLEN:"4", SELECTION: UP/DOWN, DEFAULT: "UP"`
}

type VxlanGlobalState struct {
	baseObj
	Vrf              string `SNAPROUTE: "KEY", CATEGORY:"Tunnel", ACCESS:"r",  MULTIPLICITY:"1", DEFAULT: "default", DESCRIPTION: global system object defining the global state of VXLAND.`
	OperState        string `DESCRIPTION: Oper state of VXLAND, UP will allow for vxlan configuration to be applied, DOWN will disallow and de-provision from daemon, STRLEN:"4", SELECTION: UP/DOWN, DEFAULT: DOWN`
	RxInvalidVtepCnt uint64 `DESCRIPTION: Number of invalid VXLAN VTEP frames received`
	NumVteps         uint64 `DESCRIPTION: Number of Vteps provisioned`
}
