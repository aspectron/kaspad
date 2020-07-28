package flowcontext

import "github.com/kaspanet/kaspad/concensus/blockdag"

// DAG returns the DAG associated to the flow context.
func (f *FlowContext) DAG() *blockdag.BlockDAG {
	return f.dag
}
