package protocol

import (
	"fmt"

	"github.com/kaspanet/kaspad/domain"

	"github.com/kaspanet/kaspad/domain/consensus/model/externalapi"

	"github.com/kaspanet/kaspad/app/protocol/flowcontext"
	peerpkg "github.com/kaspanet/kaspad/app/protocol/peer"
	"github.com/kaspanet/kaspad/infrastructure/config"
	"github.com/kaspanet/kaspad/infrastructure/network/addressmanager"
	"github.com/kaspanet/kaspad/infrastructure/network/connmanager"
	"github.com/kaspanet/kaspad/infrastructure/network/netadapter"
)

// Manager manages the p2p protocol
type Manager struct {
	context *flowcontext.FlowContext
}

// NewManager creates a new instance of the p2p protocol manager
func NewManager(cfg *config.Config, domain domain.Domain, netAdapter *netadapter.NetAdapter, addressManager *addressmanager.AddressManager,
	connectionManager *connmanager.ConnectionManager) (*Manager, error) {

	manager := Manager{
		context: flowcontext.New(cfg, domain, addressManager, netAdapter, connectionManager),
	}

	netAdapter.SetP2PRouterInitializer(manager.routerInitializer)
	return &manager, nil
}

// Peers returns the currently active peers
func (m *Manager) Peers() []*peerpkg.Peer {
	return m.context.Peers()
}

// AddTransaction adds transaction to the mempool and propagates it.
func (m *Manager) AddTransaction(tx *externalapi.DomainTransaction) error {
	return m.context.AddTransaction(tx)
}

// AddBlock adds the given block to the DAG and propagates it.
func (m *Manager) AddBlock(block *externalapi.DomainBlock) error {
	return m.context.AddBlock(block)
}

func (m *Manager) runFlows(flows []*flow, peer *peerpkg.Peer, errChan <-chan error) error {
	for _, flow := range flows {
		executeFunc := flow.executeFunc // extract to new variable so that it's not overwritten
		spawn(fmt.Sprintf("flow-%s", flow.name), func() {
			executeFunc(peer)
		})
	}

	return <-errChan
}

// SetOnBlockAddedToDAGHandler sets the onBlockAddedToDAG handler
func (m *Manager) SetOnBlockAddedToDAGHandler(onBlockAddedToDAGHandler flowcontext.OnBlockAddedToDAGHandler) {
	m.context.SetOnBlockAddedToDAGHandler(onBlockAddedToDAGHandler)
}

// SetOnTransactionAddedToMempoolHandler sets the onTransactionAddedToMempool handler
func (m *Manager) SetOnTransactionAddedToMempoolHandler(onTransactionAddedToMempoolHandler flowcontext.OnTransactionAddedToMempoolHandler) {
	m.context.SetOnTransactionAddedToMempoolHandler(onTransactionAddedToMempoolHandler)
}

// ShouldMine returns whether it's ok to use block template from this node
// for mining purposes.
func (m *Manager) ShouldMine() (bool, error) {
	return m.context.ShouldMine()
}
