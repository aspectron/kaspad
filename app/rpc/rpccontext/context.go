package rpccontext

import (
	"github.com/kaspanet/kaspad/app/protocol"
	"github.com/kaspanet/kaspad/app/wallet/walletnotification"
	"github.com/kaspanet/kaspad/domain/blockdag"
	"github.com/kaspanet/kaspad/domain/blockdag/indexers"
	"github.com/kaspanet/kaspad/domain/mempool"
	"github.com/kaspanet/kaspad/domain/mining"
	"github.com/kaspanet/kaspad/infrastructure/config"
	"github.com/kaspanet/kaspad/infrastructure/network/addressmanager"
	"github.com/kaspanet/kaspad/infrastructure/network/connmanager"
	"github.com/kaspanet/kaspad/infrastructure/network/netadapter"
)

// Context represents the RPC context
type Context struct {
	Config                 *config.Config
	NetAdapter             *netadapter.NetAdapter
	DAG                    *blockdag.BlockDAG
	ProtocolManager        *protocol.Manager
	ConnectionManager      *connmanager.ConnectionManager
	BlockTemplateGenerator *mining.BlkTmplGenerator
	Mempool                *mempool.TxPool
	AddressManager         *addressmanager.AddressManager
	AcceptanceIndex        *indexers.AcceptanceIndex
	ShutDownChan           chan<- struct{}

	BlockTemplateState  *BlockTemplateState
	NotificationManager *walletnotification.Manager
}

// NewContext creates a new RPC context
func NewContext(cfg *config.Config,
	netAdapter *netadapter.NetAdapter,
	dag *blockdag.BlockDAG,
	protocolManager *protocol.Manager,
	connectionManager *connmanager.ConnectionManager,
	blockTemplateGenerator *mining.BlkTmplGenerator,
	mempool *mempool.TxPool,
	addressManager *addressmanager.AddressManager,
	acceptanceIndex *indexers.AcceptanceIndex,
	shutDownChan chan<- struct{}) *Context {

	context := &Context{
		Config:                 cfg,
		NetAdapter:             netAdapter,
		DAG:                    dag,
		ProtocolManager:        protocolManager,
		ConnectionManager:      connectionManager,
		BlockTemplateGenerator: blockTemplateGenerator,
		Mempool:                mempool,
		AddressManager:         addressManager,
		AcceptanceIndex:        acceptanceIndex,
		ShutDownChan:           shutDownChan,
	}
	context.BlockTemplateState = NewBlockTemplateState(context)
	context.NotificationManager = walletnotification.NewNotificationManager()
	return context
}
