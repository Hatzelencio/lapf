package application

import (
	"encoding/binary"
	"net"
	"net/netip"
	provider "overlapping-finder/cloud"
	"sync"
)

type ResultDescribeAllNetwork struct {
	Networks []provider.CloudNetwork
	Region   string
	Err      error
}

type ResultIsContainsOverlap struct {
	CloudNetwork provider.CloudNetwork
	CurrentCidr  string
	IsOverlap    bool
	Err          error `json:"-"`
}

type ResultIsPrivateCIDRBlock struct {
	CIDR         string   `json:"cidr"`
	IsPrivate    bool     `json:"isPrivate"`
	PublicIPList []string `json:"publicIPList,omitempty"`
	Err          error    `json:"-"`
}

type ipv4CIDR struct {
	Network net.IPNet
	FirstIP net.IP
	LastIP  net.IP
}

func parseCIDRBlockToIpv4(cidr string) (*ipv4CIDR, error) {
	_, ipv4Net, _ := net.ParseCIDR(cidr)

	start := binary.BigEndian.Uint32(ipv4Net.IP)
	mask := binary.BigEndian.Uint32(ipv4Net.Mask)

	last := (start & mask) | (mask ^ 0xffffffff)

	firstIP := make(net.IP, 4)
	binary.BigEndian.PutUint32(firstIP, start)

	lastIP := make(net.IP, 4)
	binary.BigEndian.PutUint32(lastIP, last)

	return &ipv4CIDR{
		*ipv4Net,
		firstIP,
		lastIP,
	}, nil
}

func ensureCIDRBlock(networks []provider.CloudNetwork, cidrBlocks []string) ([]*ResultIsContainsOverlap, error) {
	var wg sync.WaitGroup
	var ch = make(chan *ResultIsContainsOverlap)
	var results []*ResultIsContainsOverlap

	for _, network := range networks {
		for _, currentCidr := range cidrBlocks {
			wg.Add(1)

			go func(cloudNetwork provider.CloudNetwork, currentCidr string) {
				defer wg.Done()

				parsedCidr, err := parseCIDRBlockToIpv4(currentCidr)
				cidr, err := parseCIDRBlockToIpv4(cloudNetwork.CidrBlock)
				isOverlap := cidr.contains(parsedCidr)
				ch <- &ResultIsContainsOverlap{
					CloudNetwork: cloudNetwork,
					CurrentCidr:  parsedCidr.Network.String(),
					IsOverlap:    isOverlap,
					Err:          err,
				}
			}(network, currentCidr)
		}
	}

	for range networks {
		for range cidrBlocks {
			results = append(results, <-ch)
		}
	}

	wg.Wait()

	return results, nil
}

func ensureIsPrivateCIDRBlock(cidrBlocks []string, showIPList bool) ([]*ResultIsPrivateCIDRBlock, error) {
	var wg sync.WaitGroup
	var ch = make(chan *ResultIsPrivateCIDRBlock)
	var results []*ResultIsPrivateCIDRBlock

	for _, currentCidr := range cidrBlocks {
		wg.Add(1)

		go func(currentCidr string, showIPList bool) {
			defer wg.Done()

			var isPrivate = true
			var publicIPList []string

			prefix, err := netip.ParsePrefix(currentCidr)
			for addr := prefix.Addr(); prefix.Contains(addr); addr = addr.Next() {
				if !addr.IsPrivate() {
					isPrivate = false

					if showIPList {
						publicIPList = append(publicIPList, addr.String())
					}
				}
			}

			ch <- &ResultIsPrivateCIDRBlock{
				CIDR:         currentCidr,
				IsPrivate:    isPrivate,
				PublicIPList: publicIPList,
				Err:          err,
			}
		}(currentCidr, showIPList)
	}

	for range cidrBlocks {
		results = append(results, <-ch)
	}

	wg.Wait()

	return results, nil
}

func (ip *ipv4CIDR) contains(net *ipv4CIDR) bool {
	return net.Network.Contains(ip.FirstIP) || net.Network.Contains(ip.LastIP) || ip.Network.Contains(net.FirstIP) || ip.Network.Contains(net.LastIP)
}
