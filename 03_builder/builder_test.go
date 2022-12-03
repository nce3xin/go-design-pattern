package builder

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

type Args struct {
	Addr    string
	Port    int
	Options []Option
}

func TestNewServer(t *testing.T) {
	options1 := []Option{
		Protocol("ipv4"),
		Timeout(60 * time.Second),
		MaxConns(1000),
	}

	options2 := []Option{
		MaxConns(10),
	}

	testCases := []struct {
		name    string
		args    Args
		want    *Server
		wantErr bool
	}{
		{
			name: "all config set",
			args: Args{
				Addr:    "127.0.0.1",
				Port:    8080,
				Options: options1,
			},
			want: &Server{
				Addr:     "127.0.0.1",
				Port:     8080,
				Protocol: "ipv4",
				Timeout:  60 * time.Second,
				MaxConns: 1000,
			},
			wantErr: false,
		},
		{
			name: "set ip,port and maxconn",
			args: Args{
				Addr:    "0.0.0.0",
				Port:    80,
				Options: options2,
			},
			want: &Server{
				Addr:     "0.0.0.0",
				Port:     80,
				MaxConns: 10,
			},
			wantErr: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewServer(tt.args.Addr, tt.args.Port, tt.args.Options...)
			assert.Equal(t, tt.want, got)
			require.Equalf(t, tt.wantErr, err != nil, "WantErr: %v. GotErr: %v.", tt.wantErr, err)
		})
	}

}
