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

// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package rpc

import (
	"encoding/json"
	"errors"
	"math/big"
)

// MarshalJSON marshals as JSON.
func (t TransferMetadata) MarshalJSON() ([]byte, error) {
	type TransferMetadata struct {
		Balance *big.Int             `json:"balance" gencodec:"required"`
		Tx      *TransactionMetadata `json:"tx"`
	}
	var enc TransferMetadata
	enc.Balance = t.Balance
	enc.Tx = t.Tx
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (t *TransferMetadata) UnmarshalJSON(input []byte) error {
	type TransferMetadata struct {
		Balance *big.Int             `json:"balance" gencodec:"required"`
		Tx      *TransactionMetadata `json:"tx"`
	}
	var dec TransferMetadata
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Balance == nil {
		return errors.New("missing required field 'balance' for TransferMetadata")
	}
	t.Balance = dec.Balance
	if dec.Tx != nil {
		t.Tx = dec.Tx
	}
	return nil
}
