package bwlist

import (
	"net"
	"testing"
)

func Test_BWList(t *testing.T) {
	testCases := []struct {
		subnet   string
		inRange  []string // первый IP из подсети, какой-нибудь из середины и последний
		outRange []string // на единицу меньше, чем первый, и на единицу больше, чем последний
	}{
		{
			subnet:   "192.168.0.0/25",
			inRange:  []string{"192.168.0.0", "192.168.0.60", "192.168.0.127"},
			outRange: []string{"192.167.255.255", "192.168.0.128"},
		},
		{
			subnet:   "10.1.1.0/16",
			inRange:  []string{"10.1.0.0", "10.1.1.60", "10.1.255.255"},
			outRange: []string{"10.0.255.255", "10.2.0.0"},
		},
	}

	for _, tc := range testCases {
		_, subnet, _ := net.ParseCIDR(tc.subnet)
		bw := New(subnet)

		for _, ip := range tc.inRange {
			if ok := bw.Contains(net.ParseIP(ip)); !ok {
				t.Fatalf("expected IP(%s) is in Subnet(%s)", ip, tc.subnet)
			}
		}
	}

	for _, tc := range testCases {
		bw := New()
		_, subnet, _ := net.ParseCIDR(tc.subnet)

		bw.Append(subnet)
		for _, ip := range tc.inRange {
			if ok := bw.Contains(net.ParseIP(ip)); !ok {
				t.Fatalf("expected IP(%s) is in Subnet(%s)", ip, tc.subnet)
			}
		}
		for _, ip := range tc.outRange {
			if ok := bw.Contains(net.ParseIP(ip)); ok {
				t.Fatalf("expected IP(%s) is not in Subnet(%s)", ip, tc.subnet)
			}
		}

		bw.Remove(subnet)
		for _, ip := range tc.inRange {
			if ok := bw.Contains(net.ParseIP(ip)); ok {
				t.Fatalf("expected IP(%s) is not in Subnet(%s)", ip, tc.subnet)
			}
		}
	}
}
