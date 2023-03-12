package utils

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"net"
	"time"

	"github.com/containernetworking/plugins/pkg/ns"
	"github.com/coreos/go-iptables/iptables"
	"github.com/vishvananda/netlink"
)
func CreateBridge(bridgeName, gw string, mtu int) (*netlink.Bridge, error) {
	//find existed bridge
	link, err := netlink.LinkByName(bridgeName)
	if err != nil {
		Write("", "find bridge by name err: %s", err.Error())
		return nil, err
	}
	br, ok := link.(*netlink.Bridge)
	if ok && br != nil {
		return br, nil
	}
	//begin to create bridge
	br = &netlink.Bridge{
		LinkAttrs: netlink.LinkAttrs{
			Name: bridgeName,
			MTU: mtu,
			TxQLen: -1,
		},
	}
	netlink.LinkAdd(br)

	//check bridge information
	l, err := netlink.LinkByName(bridgeName)
	if err != nil {
		return nil ,err
	}
	br, ok = l.(*netlink.Bridge)
	if (! ok) {
		Write("", "the device is not bridge type")
		return nil, fmt.Errorf("found the device %s, but not bridge", bridgeName)
	}
	ip, ipNet, err := net.ParseCIDR(gw)
	if err != nil {
		Write("", "parse gw cidr failed, the gw: %s", gw)
		return nil, err
	}
	ipNet.IP = ip
	addr := &netlink.Addr{
		IPNet: ipNet,
	}
	err = netlink.AddrAdd(br, addr)
	if err != nil {
		Write("", "add ip: %s to bridge: %s failed, error : %s", ip.String(), bridgeName, err.Error())
		return nil, err
	}
	if err = netlink.LinkSetUp(br); err != nil {
		Write("","bootstrap bridge failed, err : %s", err.Error())
		return nil ,err
	}
	return br , nil

}

func CreateVerthPair (ifName, brName string, mtu int) (netlink.Veth, netlink.Veth){
	peerName = ifName + "-" + brName + "-" + genRandomStr()
	
}

func genRandomStr() (string) {
	rand.Seed(time.Now().UnixMicro())
	bytes := make([]byte, 6)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}