package mockrouting

import (
	context "context"
	ds "gx/ipfs/QmVSase1JP7cq9QkPT46oNwdp9pT6kBkG3oqS14y3QcZjG/go-datastore"
	sync "gx/ipfs/QmVSase1JP7cq9QkPT46oNwdp9pT6kBkG3oqS14y3QcZjG/go-datastore/sync"
	"gx/ipfs/QmYazuWNCrC7LXeRt3utQuRsCDufpNM176rD1efuvF2pWJ/go-testutil"
	dht "gx/ipfs/QmZRKQ5zeRxJHn9e1fP72nVXGydCzvMN6oz5Q3P8DtT7Jv/go-libp2p-kad-dht"
	mocknet "gx/ipfs/QmdYrd8ygc6EuRcQj4XsVr5ajx945v9PizJHXUoDdX7VSn/go-libp2p/p2p/net/mock"
)

type mocknetserver struct {
	mn mocknet.Mocknet
}

func NewDHTNetwork(mn mocknet.Mocknet) Server {
	return &mocknetserver{
		mn: mn,
	}
}

func (rs *mocknetserver) Client(p testutil.Identity) Client {
	return rs.ClientWithDatastore(context.TODO(), p, ds.NewMapDatastore())
}

func (rs *mocknetserver) ClientWithDatastore(ctx context.Context, p testutil.Identity, ds ds.Datastore) Client {

	// FIXME AddPeer doesn't appear to be idempotent

	host, err := rs.mn.AddPeer(p.PrivateKey(), p.Address())
	if err != nil {
		panic("FIXME")
	}
	return dht.NewDHT(ctx, host, sync.MutexWrap(ds))
}

var _ Server = &mocknetserver{}
