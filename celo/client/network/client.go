// Copyright 2020 Celo Org
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package network

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/rpc"
)

// NetworkClient defines typed wrappers for the Network RPC API.
type NetworkClient struct {
	c *rpc.Client
}

// NewClient creates a client that uses the given RPC client.
func NewClient(c *rpc.Client) *NetworkClient {
	return &NetworkClient{c}
}

func (nc *NetworkClient) ChainId(ctx context.Context) (*big.Int, error) {
	version := new(big.Int)
	var ver string
	if err := nc.c.CallContext(ctx, &ver, "net_version"); err != nil {
		return nil, err
	}
	if _, ok := version.SetString(ver, 10); !ok {
		return nil, fmt.Errorf("invalid net_version result %q", ver)
	}
	return version, nil
}
