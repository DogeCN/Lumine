//go:build !windows && !linux

package lumine

import (
	"net"
	"time"

	"github.com/elastic/go-freelru"
	E "github.com/moi-si/lumine/internal/errors"
	"github.com/moi-si/lumine/internal/singleflight"
	log "github.com/moi-si/mylog"
)

var errTTLDNotSupported = E.New("`ttl-d` is not supported on current system")

var (
	ttlCache        *freelru.ShardedLRU[string, int]
	ttlCacheTTL     time.Duration
	ttlSingleflight *singleflight.Group[int]
)

func loadTTLRules(string) error {
	return nil
}

func getFakeTTL(*log.Logger, *Policy, string, bool) (ttl int, err error) {
	return unsetInt, errTTLDNotSupported
}

func desyncSend(net.Conn, bool, []byte, int, int, int, time.Duration) error {
	return errTTLDNotSupported
}
