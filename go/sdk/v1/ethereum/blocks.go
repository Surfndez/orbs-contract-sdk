// Copyright 2019 the orbs-contract-sdk authors
// This file is part of the orbs-contract-sdk library in the Orbs project.
//
// This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
// The above notice should be included in all copies or substantial portions of the software.

package ethereum

import "github.com/orbs-network/orbs-contract-sdk/go/context"

// returns the current Ethereum block number with some safety margin to avoid forks (about 25 minutes)
func GetBlockNumber() uint64 {
	contextId, handler, permissionScope := context.GetContext()
	return handler.SdkEthereumGetBlockNumber(contextId, permissionScope)
}
