package core

import (
	"github.com/tenderly/net-celo/common"
	"github.com/tenderly/net-celo/contracts"
	"github.com/tenderly/net-celo/core/types"
	"github.com/tenderly/net-celo/core/vm"
	"github.com/tenderly/net-celo/log"
	"github.com/tenderly/net-celo/params"
)

func GetFeeCurrencyContext(header *types.Header, config *params.ChainConfig, statedb vm.StateDB) *common.FeeCurrencyContext {
	if !config.IsCel2(header.Time) {
		return &common.FeeCurrencyContext{}
	}

	caller := &contracts.CeloBackend{ChainConfig: config, State: statedb, BlockNumber: header.Number, Time: header.Time}

	feeCurrencyContext, err := contracts.GetFeeCurrencyContext(caller)
	if err != nil {
		log.Error("Error fetching exchange rates!", "err", err)
	}
	return &feeCurrencyContext
}
