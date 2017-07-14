package iputils

import "testing"

type wantContainsIP struct {
	ok      bool
	network string
	err     bool
}

var tcContainsIP = []struct {
	name             string
	skip, skipAlways bool
	ip               string
	networks         []string
	want             wantContainsIP
}{
	{
		name: "positive",
		ip:   "192.168.0.1",
		networks: []string{
			"192.168.0.0/24",
		},
		want: wantContainsIP{
			ok:      true,
			network: "192.168.0.0/24",
			err:     false,
		},
	},
	{
		name: "positiveMoreThanOne",
		ip:   "192.168.0.1",
		networks: []string{
			"10.10.10.0/24",
			"192.168.0.0/24",
		},
		want: wantContainsIP{
			ok:      true,
			network: "192.168.0.0/24",
			err:     false,
		},
	},
	{
		name: "noContains",
		ip:   "192.168.0.1",
		networks: []string{
			"10.10.10.0/24",
		},
		want: wantContainsIP{
			ok:      false,
			network: "",
			err:     false,
		},
	},
	{
		name: "wrongIP",
		ip:   "x.x.x.x",
		networks: []string{
			"10.10.10.0/24",
		},
		want: wantContainsIP{
			ok:      false,
			network: "",
			err:     true,
		},
	},
	{
		name: "wrongNet",
		ip:   "192.168.0.1",
		networks: []string{
			"y.y.y.y/24",
		},
		want: wantContainsIP{
			ok:      false,
			network: "",
			err:     true,
		},
	},
}

func TestContainsIP(t *testing.T) {
	for _, tc := range tcContainsIP {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			if testing.Short() && tc.skip || tc.skipAlways {
				t.SkipNow()
			}

			ok, network, err := ContainsIP(tc.ip, tc.networks...)
			if ok != tc.want.ok {
				t.Errorf("ok: want %v, have %v", tc.want.ok, ok)
			}
			if network != tc.want.network {
				t.Errorf("network: want %v, have %v", tc.want.network, network)
			}
			if tc.want.err && err == nil {
				t.Error("error expected, got nil")
			}
			if !tc.want.err && err != nil {
				t.Errorf("error not expected, got %v", err)
			}
		})
	}
}
