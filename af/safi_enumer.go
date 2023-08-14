// Code generated by "enumer -type SAFI -trimprefix SAFI_"; DO NOT EDIT.

package af

import (
	"fmt"
	"strings"
)

const (
	_SAFIName_0      = "UNICASTMULTICAST"
	_SAFILowerName_0 = "unicastmulticast"
	_SAFIName_1      = "MPLSMCAST_VPNPLACEMENT_MSPW"
	_SAFILowerName_1 = "mplsmcast_vpnplacement_mspw"
	_SAFIName_2      = "MCAST_VPLSSFC"
	_SAFILowerName_2 = "mcast_vplssfc"
	_SAFIName_3      = "TUNNELVPLSMDT4OVER66OVER4L1VPN_DISCOVERYEVPNSLSLS_VPNSR_TE_POLICYSD_WAN_CAPABILITIESROUTING_POLICYCLASSFUL_TRANSPORTTUNNELED_FLOWSPECMCAST_TREEDPSLS_SPF"
	_SAFILowerName_3 = "tunnelvplsmdt4over66over4l1vpn_discoveryevpnslsls_vpnsr_te_policysd_wan_capabilitiesrouting_policyclassful_transporttunneled_flowspecmcast_treedpsls_spf"
	_SAFIName_4      = "CARVPN_CARMUP"
	_SAFILowerName_4 = "carvpn_carmup"
	_SAFIName_5      = "MPLS_VPNMULTICAST_VPNS"
	_SAFILowerName_5 = "mpls_vpnmulticast_vpns"
	_SAFIName_6      = "ROUTE_TARGETFLOWSPECL3VPN_FLOWSPEC"
	_SAFILowerName_6 = "route_targetflowspecl3vpn_flowspec"
	_SAFIName_7      = "VPN_DISCOVERY"
	_SAFILowerName_7 = "vpn_discovery"
)

var (
	_SAFIIndex_0 = [...]uint8{0, 7, 16}
	_SAFIIndex_1 = [...]uint8{0, 4, 13, 27}
	_SAFIIndex_2 = [...]uint8{0, 10, 13}
	_SAFIIndex_3 = [...]uint8{0, 6, 10, 13, 19, 25, 40, 45, 47, 53, 65, 84, 98, 116, 133, 143, 146, 152}
	_SAFIIndex_4 = [...]uint8{0, 3, 10, 13}
	_SAFIIndex_5 = [...]uint8{0, 8, 22}
	_SAFIIndex_6 = [...]uint8{0, 12, 20, 34}
	_SAFIIndex_7 = [...]uint8{0, 13}
)

func (i SAFI) String() string {
	switch {
	case 1 <= i && i <= 2:
		i -= 1
		return _SAFIName_0[_SAFIIndex_0[i]:_SAFIIndex_0[i+1]]
	case 4 <= i && i <= 6:
		i -= 4
		return _SAFIName_1[_SAFIIndex_1[i]:_SAFIIndex_1[i+1]]
	case 8 <= i && i <= 9:
		i -= 8
		return _SAFIName_2[_SAFIIndex_2[i]:_SAFIIndex_2[i+1]]
	case 64 <= i && i <= 80:
		i -= 64
		return _SAFIName_3[_SAFIIndex_3[i]:_SAFIIndex_3[i+1]]
	case 83 <= i && i <= 85:
		i -= 83
		return _SAFIName_4[_SAFIIndex_4[i]:_SAFIIndex_4[i+1]]
	case 128 <= i && i <= 129:
		i -= 128
		return _SAFIName_5[_SAFIIndex_5[i]:_SAFIIndex_5[i+1]]
	case 132 <= i && i <= 134:
		i -= 132
		return _SAFIName_6[_SAFIIndex_6[i]:_SAFIIndex_6[i+1]]
	case i == 140:
		return _SAFIName_7
	default:
		return fmt.Sprintf("SAFI(%d)", i)
	}
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _SAFINoOp() {
	var x [1]struct{}
	_ = x[SAFI_UNICAST-(1)]
	_ = x[SAFI_MULTICAST-(2)]
	_ = x[SAFI_MPLS-(4)]
	_ = x[SAFI_MCAST_VPN-(5)]
	_ = x[SAFI_PLACEMENT_MSPW-(6)]
	_ = x[SAFI_MCAST_VPLS-(8)]
	_ = x[SAFI_SFC-(9)]
	_ = x[SAFI_TUNNEL-(64)]
	_ = x[SAFI_VPLS-(65)]
	_ = x[SAFI_MDT-(66)]
	_ = x[SAFI_4OVER6-(67)]
	_ = x[SAFI_6OVER4-(68)]
	_ = x[SAFI_L1VPN_DISCOVERY-(69)]
	_ = x[SAFI_EVPNS-(70)]
	_ = x[SAFI_LS-(71)]
	_ = x[SAFI_LS_VPN-(72)]
	_ = x[SAFI_SR_TE_POLICY-(73)]
	_ = x[SAFI_SD_WAN_CAPABILITIES-(74)]
	_ = x[SAFI_ROUTING_POLICY-(75)]
	_ = x[SAFI_CLASSFUL_TRANSPORT-(76)]
	_ = x[SAFI_TUNNELED_FLOWSPEC-(77)]
	_ = x[SAFI_MCAST_TREE-(78)]
	_ = x[SAFI_DPS-(79)]
	_ = x[SAFI_LS_SPF-(80)]
	_ = x[SAFI_CAR-(83)]
	_ = x[SAFI_VPN_CAR-(84)]
	_ = x[SAFI_MUP-(85)]
	_ = x[SAFI_MPLS_VPN-(128)]
	_ = x[SAFI_MULTICAST_VPNS-(129)]
	_ = x[SAFI_ROUTE_TARGET-(132)]
	_ = x[SAFI_FLOWSPEC-(133)]
	_ = x[SAFI_L3VPN_FLOWSPEC-(134)]
	_ = x[SAFI_VPN_DISCOVERY-(140)]
}

var _SAFIValues = []SAFI{SAFI_UNICAST, SAFI_MULTICAST, SAFI_MPLS, SAFI_MCAST_VPN, SAFI_PLACEMENT_MSPW, SAFI_MCAST_VPLS, SAFI_SFC, SAFI_TUNNEL, SAFI_VPLS, SAFI_MDT, SAFI_4OVER6, SAFI_6OVER4, SAFI_L1VPN_DISCOVERY, SAFI_EVPNS, SAFI_LS, SAFI_LS_VPN, SAFI_SR_TE_POLICY, SAFI_SD_WAN_CAPABILITIES, SAFI_ROUTING_POLICY, SAFI_CLASSFUL_TRANSPORT, SAFI_TUNNELED_FLOWSPEC, SAFI_MCAST_TREE, SAFI_DPS, SAFI_LS_SPF, SAFI_CAR, SAFI_VPN_CAR, SAFI_MUP, SAFI_MPLS_VPN, SAFI_MULTICAST_VPNS, SAFI_ROUTE_TARGET, SAFI_FLOWSPEC, SAFI_L3VPN_FLOWSPEC, SAFI_VPN_DISCOVERY}

var _SAFINameToValueMap = map[string]SAFI{
	_SAFIName_0[0:7]:          SAFI_UNICAST,
	_SAFILowerName_0[0:7]:     SAFI_UNICAST,
	_SAFIName_0[7:16]:         SAFI_MULTICAST,
	_SAFILowerName_0[7:16]:    SAFI_MULTICAST,
	_SAFIName_1[0:4]:          SAFI_MPLS,
	_SAFILowerName_1[0:4]:     SAFI_MPLS,
	_SAFIName_1[4:13]:         SAFI_MCAST_VPN,
	_SAFILowerName_1[4:13]:    SAFI_MCAST_VPN,
	_SAFIName_1[13:27]:        SAFI_PLACEMENT_MSPW,
	_SAFILowerName_1[13:27]:   SAFI_PLACEMENT_MSPW,
	_SAFIName_2[0:10]:         SAFI_MCAST_VPLS,
	_SAFILowerName_2[0:10]:    SAFI_MCAST_VPLS,
	_SAFIName_2[10:13]:        SAFI_SFC,
	_SAFILowerName_2[10:13]:   SAFI_SFC,
	_SAFIName_3[0:6]:          SAFI_TUNNEL,
	_SAFILowerName_3[0:6]:     SAFI_TUNNEL,
	_SAFIName_3[6:10]:         SAFI_VPLS,
	_SAFILowerName_3[6:10]:    SAFI_VPLS,
	_SAFIName_3[10:13]:        SAFI_MDT,
	_SAFILowerName_3[10:13]:   SAFI_MDT,
	_SAFIName_3[13:19]:        SAFI_4OVER6,
	_SAFILowerName_3[13:19]:   SAFI_4OVER6,
	_SAFIName_3[19:25]:        SAFI_6OVER4,
	_SAFILowerName_3[19:25]:   SAFI_6OVER4,
	_SAFIName_3[25:40]:        SAFI_L1VPN_DISCOVERY,
	_SAFILowerName_3[25:40]:   SAFI_L1VPN_DISCOVERY,
	_SAFIName_3[40:45]:        SAFI_EVPNS,
	_SAFILowerName_3[40:45]:   SAFI_EVPNS,
	_SAFIName_3[45:47]:        SAFI_LS,
	_SAFILowerName_3[45:47]:   SAFI_LS,
	_SAFIName_3[47:53]:        SAFI_LS_VPN,
	_SAFILowerName_3[47:53]:   SAFI_LS_VPN,
	_SAFIName_3[53:65]:        SAFI_SR_TE_POLICY,
	_SAFILowerName_3[53:65]:   SAFI_SR_TE_POLICY,
	_SAFIName_3[65:84]:        SAFI_SD_WAN_CAPABILITIES,
	_SAFILowerName_3[65:84]:   SAFI_SD_WAN_CAPABILITIES,
	_SAFIName_3[84:98]:        SAFI_ROUTING_POLICY,
	_SAFILowerName_3[84:98]:   SAFI_ROUTING_POLICY,
	_SAFIName_3[98:116]:       SAFI_CLASSFUL_TRANSPORT,
	_SAFILowerName_3[98:116]:  SAFI_CLASSFUL_TRANSPORT,
	_SAFIName_3[116:133]:      SAFI_TUNNELED_FLOWSPEC,
	_SAFILowerName_3[116:133]: SAFI_TUNNELED_FLOWSPEC,
	_SAFIName_3[133:143]:      SAFI_MCAST_TREE,
	_SAFILowerName_3[133:143]: SAFI_MCAST_TREE,
	_SAFIName_3[143:146]:      SAFI_DPS,
	_SAFILowerName_3[143:146]: SAFI_DPS,
	_SAFIName_3[146:152]:      SAFI_LS_SPF,
	_SAFILowerName_3[146:152]: SAFI_LS_SPF,
	_SAFIName_4[0:3]:          SAFI_CAR,
	_SAFILowerName_4[0:3]:     SAFI_CAR,
	_SAFIName_4[3:10]:         SAFI_VPN_CAR,
	_SAFILowerName_4[3:10]:    SAFI_VPN_CAR,
	_SAFIName_4[10:13]:        SAFI_MUP,
	_SAFILowerName_4[10:13]:   SAFI_MUP,
	_SAFIName_5[0:8]:          SAFI_MPLS_VPN,
	_SAFILowerName_5[0:8]:     SAFI_MPLS_VPN,
	_SAFIName_5[8:22]:         SAFI_MULTICAST_VPNS,
	_SAFILowerName_5[8:22]:    SAFI_MULTICAST_VPNS,
	_SAFIName_6[0:12]:         SAFI_ROUTE_TARGET,
	_SAFILowerName_6[0:12]:    SAFI_ROUTE_TARGET,
	_SAFIName_6[12:20]:        SAFI_FLOWSPEC,
	_SAFILowerName_6[12:20]:   SAFI_FLOWSPEC,
	_SAFIName_6[20:34]:        SAFI_L3VPN_FLOWSPEC,
	_SAFILowerName_6[20:34]:   SAFI_L3VPN_FLOWSPEC,
	_SAFIName_7[0:13]:         SAFI_VPN_DISCOVERY,
	_SAFILowerName_7[0:13]:    SAFI_VPN_DISCOVERY,
}

var _SAFINames = []string{
	_SAFIName_0[0:7],
	_SAFIName_0[7:16],
	_SAFIName_1[0:4],
	_SAFIName_1[4:13],
	_SAFIName_1[13:27],
	_SAFIName_2[0:10],
	_SAFIName_2[10:13],
	_SAFIName_3[0:6],
	_SAFIName_3[6:10],
	_SAFIName_3[10:13],
	_SAFIName_3[13:19],
	_SAFIName_3[19:25],
	_SAFIName_3[25:40],
	_SAFIName_3[40:45],
	_SAFIName_3[45:47],
	_SAFIName_3[47:53],
	_SAFIName_3[53:65],
	_SAFIName_3[65:84],
	_SAFIName_3[84:98],
	_SAFIName_3[98:116],
	_SAFIName_3[116:133],
	_SAFIName_3[133:143],
	_SAFIName_3[143:146],
	_SAFIName_3[146:152],
	_SAFIName_4[0:3],
	_SAFIName_4[3:10],
	_SAFIName_4[10:13],
	_SAFIName_5[0:8],
	_SAFIName_5[8:22],
	_SAFIName_6[0:12],
	_SAFIName_6[12:20],
	_SAFIName_6[20:34],
	_SAFIName_7[0:13],
}

// SAFIString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func SAFIString(s string) (SAFI, error) {
	if val, ok := _SAFINameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _SAFINameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to SAFI values", s)
}

// SAFIValues returns all values of the enum
func SAFIValues() []SAFI {
	return _SAFIValues
}

// SAFIStrings returns a slice of all String values of the enum
func SAFIStrings() []string {
	strs := make([]string, len(_SAFINames))
	copy(strs, _SAFINames)
	return strs
}

// IsASAFI returns "true" if the value is listed in the enum definition. "false" otherwise
func (i SAFI) IsASAFI() bool {
	for _, v := range _SAFIValues {
		if i == v {
			return true
		}
	}
	return false
}
