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

type DWDMModuleState struct {
	baseObj
	ModuleId               uint8   `SNAPROUTE: "KEY", CATEGORY:"Optical", ACCESS:"r", MULTIPLICITY: "*", DESCRIPTION: "DWDM Module identifier"`
	ModuleState            string  `DESCRIPTION: "Current MSA state of dwdm module"`
	ModuleVoltage          float64 `DESCRIPTION: "Module power supply voltage in Volts"`
	ModuleTemp             float64 `DESCRIPTION: "Module temperature in deg Celsius"`
	Populated              bool    `DESCRIPTION: "Is module populated"`
	VendorName             string  `DESCRIPTION: "Vendor name of dwdm module"`
	VendorPartNum          string  `DESCRIPTION: "Vendor assigned part number of dwdm module"`
	VendorSerialNum        string  `DESCRIPTION: "Vendor assigned serial number of dwdm module "`
	VendorDateCode         string  `DESCRIPTION: "Device manufacture data code of dwdm module"`
	ModuleHWVersion        string  `DESCRIPTION: "HW version of dwdm module"`
	ModuleActiveFWVersion  string  `DESCRIPTION: "Firmware version of active partition of dwdm module"`
	ModuleStandByFWVersion string  `DESCRIPTION: "Firmware version of standby partition of dwdm module"`
	ModuleActiveFWStatus   string  `DESCRIPTION: "Firmware image status of active partition of dwdm module"`
	ModuleStandByFWStatus  string  `DESCRIPTION: "Firmware image status of standby partition of dwdm module"`
}

type DWDMModule struct {
	baseObj
	ModuleId            uint8  `SNAPROUTE: "KEY", CATEGORY:"Optical", ACCESS:"rw", MULTIPLICITY: "*", AUTODISCOVER: "true", DESCRIPTION: "DWDM Module identifier"`
	AdminState          string `DESCRIPTION: "Reset state of this dwdm module (false (Reset deasserted), true (Reset asserted))", SELECTION: "UP"/"DOWN", DEFAULT:"DOWN"`
	IndependentLaneMode bool   `DESCRIPTION: "Network lane configuration for the DWDM Module. true-Independent lanes, false-Coupled lanes, DEFAULT:true"`
	PMInterval          uint8  `DESCRIPTION: "Performance monitoring interval, i.e. time interval between successive PM ticks in seconds", DEFAULT:1`
	EnableExtPMTickSrc  bool   `DESCRIPTION:"Enable/Disable external tick source for performance monitoring", DEFAULT:false`
}

type DWDMModuleNwIntf struct {
	baseObj
	ModuleId                  uint8   `SNAPROUTE: "KEY", CATEGORY:"Optical", ACCESS:"rw", MULTIPLICITY: "*", AUTODISCOVER: "true", DESCRIPTION: "DWDM Module identifier"`
	NwIntfId                  uint8   `SNAPROUTE: "KEY", CATEGORY:"Optical", DESCRIPTION: "DWDM Module network interface identifier"`
	ModulationFmt             string  `DESCRIPTION: "Modulation format to use for this network interface", SELECTION: "QPSK"/"8QAM/"16QAM", DEFAULT:"16QAM"`
	TxPower                   float64 `DESCRIPTION: "Transmit output power for this network interface in dBm, MIN:0, MAX:4294967295", DEFAULT:0, UNIT:dBm`
	ChannelNumber             uint8   `DESCRIPTION: "TX Channel number to use for this network interface", MIN:1, MAX:100, DEFAULT:48`
	FECMode                   string  `DESCRIPTION: "DWDM Module network interface FEC mode", SELECTION: "15%SDFEC"/"15%OvrHeadSDFEC"/"25%OvrHeadSDFEC", DEFAULT:"15%SDFEC"`
	DiffEncoding              bool    `DESCRIPTION: "Control to enable/disable DWDM Module network interface encoding type", DEFAULT: true`
	TxPulseShapeFltrType      string  `DESCRIPTION: "TX pulse shaping filter type", SELECTION: "RootRaisedCos"/"RaisedCos", DEFAULT:RootRaisedCos"`
	TxPulseShapeFltrRollOff   float64 `DESCRIPTION: "TX pulse shape filter roll off factor, MIN:0.004, MAX:0.8", DEFAULT:0.301`
	AdminState                string  `DESCRIPTION: "Administrative state of this network interface", SELECTION: "UP"/"DOWN", DEFAULT: "UP"`
	EnableTxPRBS              bool    `DESCRIPTION: "Enable TX PRBS generation on this network interface", DEFAULT: false`
	TxPRBSPattern             string  `DESCRIPTION: "Pattern to use for TX PRBS generation", SELECTION:"2^7"/"2^15"/"2^31", DEFAULT:"2^31"`
	TxPRBSInvertPattern       bool    `DESCRIPTION: "Generate inverted PRBS polynomial pattern", DEFAULT:true`
	EnableRxPRBSChecker       bool    `DESCRIPTION: "Enable RX PRBS checker", DEFAULT: false`
	RxPRBSPattern             string  `DESCRIPTION: "PRBS pattern to use for checker", SELECTION:"2^7"/"2^15"/"2^31", DEFAULT:"2^31"`
	RxPRBSInvertPattern       bool    `DESCRIPTION: "Check against inverted PRBS polynomial pattern", DEFAULT:true`
	ClntIntfIdToTributary0Map uint8   `DESCRIPTION: "Client interface ID to map to network interface tributary 0", MIN:0, MAX:3`
	ClntIntfIdToTributary1Map uint8   `DESCRIPTION: "Client interface ID to map to network interface tributary 1", MIN:0, MAX:3`
	TxPowerRampdBmPerSec      float64 `DESCRIPTION: "Rate of change of tx power on this network interface", UNIT: dBm/s, DEFAULT:1`
}

type DWDMModuleNwIntfState struct {
	baseObj
	ModuleId                        uint8   `SNAPROUTE: "KEY", CATEGORY:"Optical", ACCESS:"r", MULTIPLICITY: "*", DESCRIPTION: "DWDM Module identifier"`
	NwIntfId                        uint8   `SNAPROUTE: "KEY", CATEGORY:"Optical", DESCRIPTION: "DWDM Module network interface identifier"`
	TxChanGridSpacing               string  `DESCRIPTION: "The channel grid spacing used for this network interface in GHz, UNIT: GHz"`
	CurrentBER                      float64 `DESCRIPTION: "Current value of BER on the DWDM module network interface"`
	MinBEROverPMInterval            float64 `DESCRIPTION: "Minimum value of BER over the last PM interval for the DWDM module network interface"`
	AvgBEROverPMInterval            float64 `DESCRIPTION: "Average value of BER over the last PM interval for the DWDM module network interface"`
	MaxBEROverPMInterval            float64 `DESCRIPTION: "Maximum value of BER over the last PM interval for the DWDM module network interface"`
	CurrUncorrectableFECBlkCnt      float64 `DESCRIPTION: "Current value of uncorrectable FEC code block count"`
	UncorrectableFECBlkCntOverPMInt float64 `DESCRIPTION: "Average value of uncorrectable FEC code block count over the last PM interval"`
	PRBSRxErrCnt                    float64 `DESCRIPTION: "RX PRBS error count for network interface"`
	RxPower                         float64 `DESCRIPTION: "Current RX power for the DWDM module network interface", UNIT: dBm`
	ChanFrequency                   float64 `DESCRIPTION: "Channel frequency corresponding to selected channel number for the DWDM module network interface", UNIT: GHz`
	CurrChromDisp                   int32   `DESCRIPTION: "Current RX chromatic dispersion for the DWDM module network interface", UNIT:ps/nm`
	AvgChromDispOverPMInt           int32   `DESCRIPTION: "Average RX chromatic dispersion over the last PM interval for the DWDM module network interface", UNIT:ps/nm`
	MinChromDispOverPMInt           int32   `DESCRIPTION: "Current RX chromatic dispersion over the last PM interval for the DWDM module network interface", UNIT:ps/nm`
	MaxChromDispOverPMInt           int32   `DESCRIPTION: "Current RX chromatic dispersion over the last PM interval for the DWDM module network interface", UNIT:ps/nm`
}

type DWDMModuleClntIntf struct {
	baseObj
	ModuleId                     uint8  `SNAPROUTE: "KEY", CATEGORY:"Optical", ACCESS:"rw", MULTIPLICITY: "*", AUTODISCOVER: "true", DESCRIPTION: "DWDM Module identifier"`
	ClntIntfId                   uint8  `SNAPROUTE: "KEY", CATEGORY:"Optical", DESCRIPTION: "DWDM Module client interface identifier"`
	TXFECDecDisable              bool   `DESCRIPTION: "802.3bj FEC decoder enable/disable state for traffic from Host to DWDM Module", DEFAULT: false`
	RXFECDecDisable              bool   `DESCRIPTION: "802.3bj FEC decoder enable/disable state for traffic from DWDM module to Host", DEFAULT: false`
	HostTxEqLfCtle               uint8  `DESCRIPTION: "Host interface TX deserializer equalization. LELPZRC LF-CTLE LFPZ gain code.", MIN:0, MAX:8, DEFAULT:0`
	HostTxEqCtle                 uint8  `DESCRIPTION: "Host interface TX deserializer equalization. LELRC CTLE LE gain code.", MIN:0, MAX:20, DEFAULT:18`
	HostTxEqDfe                  uint8  `DESCRIPTION: "Host interface TX deserializer equalization. s-DFE, DFE tap coefficient", MIN:0, MAX:63, DEFAULT:0`
	HostRxSerializerTap0Gain     uint8  `DESCRIPTION: "Host RX Serializer tap 0 control, gain for equalization filter tap", DEFAULT:7, MIN:0, MAX:7`
	HostRxSerializerTap0Delay    uint8  `DESCRIPTION: "Host RX Serializer tap 0 control, delay for equalization filter tap", DEFAULT:7, MIN:0, MAX:7`
	HostRxSerializerTap1Gain     uint8  `DESCRIPTION: "Host RX Serializer tap 1 control, gain for equalization filter tap", DEFAULT:7, MIN:0, MAX:7`
	HostRxSerializerTap2Gain     uint8  `DESCRIPTION: "Host RX Serializer tap 2 control, gain for equalization filter tap", DEFAULT:15, MIN:0, MAX:15`
	HostRxSerializerTap2Delay    uint8  `DESCRIPTION: "Host RX Serializer tap 2 control, delay for equalization filter tap", DEFAULT:5, MIN:0, MAX:7`
	AdminState                   string `DESCRIPTION: "Administrative state of this client interface", SELECTION: "UP"/"DOWN", DEFAULT: "UP"`
	EnableTxPRBSChecker          bool   `DESCRIPTION: "Enable/Disable TX PRBS checker for all lanes of this client interface", DEFAULT:false`
	TxPRBSPattern                string `DESCRIPTION: "PRBS pattern to use for checker", SELECTION:"2^7"/"2^15"/"2^31", DEFAULT:"2^31"`
	EnableRxPRBS                 bool   `DESCRIPTION: "Enable/Disable RX PRBS generation for all lanes of this client interface", DEFAULT:false`
	RxPRBSPattern                string `DESCRIPTION: "RX PRBS generator pattern", SELECTION:"2^7"/"2^15"/"2^31", DEFAULT:"2^31"`
	EnableIntSerdesNWLoopback    bool   `DESCRIPTION: "Enable/Disable serdes internal loopback, N/W RX is looped back to N/W TX", DEFAULT: false`
	EnableHostLoopback           bool   `DESCRIPTION: "Enable/Disable loopback on all host lanes of this client interface", DEFAULT: false`
	NwLaneTributaryToClntIntfMap uint8  `DESCRIPTION: "Network lane/tributary id to map to client interface", MIN:0, MAX:3`
}

type DWDMModuleClntIntfState struct {
	baseObj
	ModuleId          uint8   `SNAPROUTE: "KEY", CATEGORY:"Optical", ACCESS:"r", MULTIPLICITY: "*", DESCRIPTION: "DWDM Module identifier"`
	ClntIntfId        uint8   `SNAPROUTE: "KEY", CATEGORY:"Optical", DESCRIPTION: "DWDM Module client interface identifier"`
	PRBSTxErrCntLane0 float64 `DESCRIPTION: "Client interface host lane 0 PRBS TX Error count"`
	PRBSTxErrCntLane1 float64 `DESCRIPTION: "Client interface host lane 1 PRBS TX Error count"`
	PRBSTxErrCntLane2 float64 `DESCRIPTION: "Client interface host lane 2 PRBS TX Error count"`
	PRBSTxErrCntLane3 float64 `DESCRIPTION: "Client interface host lane 3 PRBS TX Error count"`
}

type DWDMModulePMData struct {
	TimeStamp string  `DESCRIPTION: "Timestamp at which data is collected"`
	Value     float64 `DESCRIPTION: "PM Data Value"`
}

type DWDMModuleNwIntfPMState struct {
	baseObj
	ModuleId uint8  `SNAPROUTE: "KEY", CATEGORY:"Performance", ACCESS:"r", MULTIPLICITY: "*", DESCRIPTION: "DWDM Module identifier"`
	NwIntfId uint8  `SNAPROUTE: "KEY", CATEGORY:"Performance", DESCRIPTION: "DWDM Module network interface identifier"`
	Resource string `SNAPROUTE: "KEY", CATEGORY:"Performance", DESCRIPTION: "Opticd resource name for which PM Data is required"`
	Type     string `SNAPROUTE: "KEY", CATEGORY:"Performance", DESCRIPTION: "Min/Max/Avg"`
	Class    string `SNAPROUTE: "KEY", CATEGORY:"Performance", DESCRIPTION: "Class of PM Data", SELECTION: CLASS-A/CLASS-B/CLASS-B, DEFAULT: CLASS-A`
	Data     []DWDMModulePMData
}
