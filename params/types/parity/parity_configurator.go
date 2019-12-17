// Copyright 2019 The multi-geth Authors
// This file is part of the multi-geth library.
//
// The multi-geth library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The multi-geth library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the multi-geth library. If not, see <http://www.gnu.org/licenses/>.


package parity

import (
	"encoding/json"
	"log"
	"math/big"
	"reflect"

	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params/types/ctypes"
	"github.com/ethereum/go-ethereum/params/vars"
)

var zero = uint64(0)

func (spec *ParityChainSpec) GetAccountStartNonce() *uint64 {
	return spec.Params.AccountStartNonce.Uint64P()
}

func (spec *ParityChainSpec) SetAccountStartNonce(i *uint64) error {
	if i == nil {
		return ctypes.ErrUnsupportedConfigFatal
	}
	spec.Params.AccountStartNonce = new(ParityU64).SetUint64(i)
	return nil
}

func (spec *ParityChainSpec) GetMaximumExtraDataSize() *uint64 {
	return spec.Params.MaximumExtraDataSize.Uint64P()
}

func (spec *ParityChainSpec) SetMaximumExtraDataSize(i *uint64) error {
	spec.Params.MaximumExtraDataSize = new(ParityU64).SetUint64(i)
	return nil
}

func (spec *ParityChainSpec) GetMinGasLimit() *uint64 {
	return spec.Params.MinGasLimit.Uint64P()
}

func (spec *ParityChainSpec) SetMinGasLimit(i *uint64) error {
	spec.Params.MinGasLimit = new(ParityU64).SetUint64(i)
	return nil
}

func (spec *ParityChainSpec) GetGasLimitBoundDivisor() *uint64 {
	return spec.Params.GasLimitBoundDivisor.Uint64P()
}

func (spec *ParityChainSpec) SetGasLimitBoundDivisor(i *uint64) error {
	spec.Params.GasLimitBoundDivisor = new(ParityU64).SetUint64(i)
	return nil
}

func (spec *ParityChainSpec) GetNetworkID() *uint64 {
	return spec.Params.NetworkID.Uint64P()
}

func (spec *ParityChainSpec) SetNetworkID(i *uint64) error {
	spec.Params.NetworkID = new(ParityU64).SetUint64(i)
	return nil
}

func (spec *ParityChainSpec) GetChainID() *big.Int {
	if chainid := spec.Params.ChainID.Big(); chainid == nil {
		return spec.Params.NetworkID.Big()
	} else {
		return chainid
	}
}

func (spec *ParityChainSpec) SetChainID(i *big.Int) error {
	if i == nil {
		return nil
	}
	u := i.Uint64()
	spec.Params.ChainID = new(ParityU64).SetUint64(&u)
	return nil
}

func (spec *ParityChainSpec) GetEIP7Transition() *uint64 {
	return spec.Engine.Ethash.Params.HomesteadTransition.Uint64P()
}

func (spec *ParityChainSpec) SetEIP7Transition(i *uint64) error {
	spec.Engine.Ethash.Params.HomesteadTransition = new(ParityU64).SetUint64(i)
	return nil
}

func (spec *ParityChainSpec) GetEIP150Transition() *uint64 {
	return spec.Params.EIP150Transition.Uint64P()
}

func (spec *ParityChainSpec) SetEIP150Transition(i *uint64) error {
	spec.Params.EIP150Transition = new(ParityU64).SetUint64(i)
	return nil
}

func (spec *ParityChainSpec) GetEIP160Transition() *uint64 {
	return spec.Params.EIP160Transition.Uint64P()
}

func (spec *ParityChainSpec) SetEIP160Transition(i *uint64) error {
	spec.Params.EIP160Transition = new(ParityU64).SetUint64(i)
	return nil
}

func (spec *ParityChainSpec) GetEIP161dTransition() *uint64 {
	return spec.Params.EIP161dTransition.Uint64P()
}

func (spec *ParityChainSpec) SetEIP161dTransition(i *uint64) error {
	spec.Params.EIP161dTransition = new(ParityU64).SetUint64(i)
	return nil
}

func (spec *ParityChainSpec) GetEIP161abcTransition() *uint64 {
	return spec.Params.EIP161abcTransition.Uint64P()
}

func (spec *ParityChainSpec) SetEIP161abcTransition(i *uint64) error {
	spec.Params.EIP161abcTransition = new(ParityU64).SetUint64(i)
	return nil
}

func (spec *ParityChainSpec) GetEIP155Transition() *uint64 {
	return spec.Params.EIP155Transition.Uint64P()
}

func (spec *ParityChainSpec) SetEIP155Transition(i *uint64) error {
	spec.Params.EIP155Transition = new(ParityU64).SetUint64(i)
	return nil
}

func (spec *ParityChainSpec) GetEIP140Transition() *uint64 {
	return spec.Params.EIP140Transition.Uint64P()
}

func (spec *ParityChainSpec) SetEIP140Transition(i *uint64) error {
	spec.Params.EIP140Transition = new(ParityU64).SetUint64(i)
	return nil
}

func (spec *ParityChainSpec) GetEIP211Transition() *uint64 {
	return spec.Params.EIP211Transition.Uint64P()
}

func (spec *ParityChainSpec) SetEIP211Transition(i *uint64) error {
	spec.Params.EIP211Transition = new(ParityU64).SetUint64(i)
	return nil
}

func (spec *ParityChainSpec) GetEIP214Transition() *uint64 {
	return spec.Params.EIP214Transition.Uint64P()
}

func (spec *ParityChainSpec) SetEIP214Transition(i *uint64) error {
	spec.Params.EIP214Transition = new(ParityU64).SetUint64(i)
	return nil
}

func (spec *ParityChainSpec) GetEIP658Transition() *uint64 {
	return spec.Params.EIP658Transition.Uint64P()
}

func (spec *ParityChainSpec) SetEIP658Transition(i *uint64) error {
	spec.Params.EIP658Transition = new(ParityU64).SetUint64(i)
	return nil
}

func (spec *ParityChainSpec) GetEIP145Transition() *uint64 {
	return spec.Params.EIP145Transition.Uint64P()
}

func (spec *ParityChainSpec) SetEIP145Transition(i *uint64) error {
	spec.Params.EIP145Transition = new(ParityU64).SetUint64(i)
	return nil
}

func (spec *ParityChainSpec) GetEIP1014Transition() *uint64 {
	return spec.Params.EIP1014Transition.Uint64P()
}

func (spec *ParityChainSpec) SetEIP1014Transition(i *uint64) error {
	spec.Params.EIP1014Transition = new(ParityU64).SetUint64(i)
	return nil
}

func (spec *ParityChainSpec) GetEIP1052Transition() *uint64 {
	return spec.Params.EIP1052Transition.Uint64P()
}

func (spec *ParityChainSpec) SetEIP1052Transition(i *uint64) error {
	spec.Params.EIP1052Transition = new(ParityU64).SetUint64(i)
	return nil
}

func (spec *ParityChainSpec) GetEIP1283Transition() *uint64 {
	return spec.Params.EIP1283Transition.Uint64P()
}

func (spec *ParityChainSpec) SetEIP1283Transition(i *uint64) error {
	spec.Params.EIP1283Transition = new(ParityU64).SetUint64(i)
	return nil
}

func (spec *ParityChainSpec) GetEIP1283DisableTransition() *uint64 {
	return spec.Params.EIP1283DisableTransition.Uint64P()
}

func (spec *ParityChainSpec) SetEIP1283DisableTransition(i *uint64) error {
	spec.Params.EIP1283DisableTransition = new(ParityU64).SetUint64(i)
	return nil
}

func (spec *ParityChainSpec) GetEIP2200Transition() *uint64 {
	return spec.Params.EIP1283ReenableTransition.Uint64P()
}

func (spec *ParityChainSpec) SetEIP2200Transition(i *uint64) error {
	spec.Params.EIP1283ReenableTransition = new(ParityU64).SetUint64(i)
	return nil
}

func (spec *ParityChainSpec) GetEIP1344Transition() *uint64 {
	return spec.Params.EIP1344Transition.Uint64P()
}

func (spec *ParityChainSpec) SetEIP1344Transition(i *uint64) error {
	spec.Params.EIP1344Transition = new(ParityU64).SetUint64(i)
	return nil
}

func (spec *ParityChainSpec) GetEIP1884Transition() *uint64 {
	return spec.Params.EIP1884Transition.Uint64P()
}

func (spec *ParityChainSpec) SetEIP1884Transition(i *uint64) error {
	spec.Params.EIP1884Transition = new(ParityU64).SetUint64(i)
	return nil
}

func (spec *ParityChainSpec) GetEIP2028Transition() *uint64 {
	return spec.Params.EIP2028Transition.Uint64P()
}

func (spec *ParityChainSpec) SetEIP2028Transition(i *uint64) error {
	spec.Params.EIP2028Transition = new(ParityU64).SetUint64(i)
	return nil
}

func (spec *ParityChainSpec) GetEIP152Transition() *uint64 {
	return spec.GetPrecompile(common.BytesToAddress([]byte{9}), ParityChainSpecPricing{
		Blake2F: &ParityChainSpecBlakePricing{
			GasPerRound: 1,
		},
	}).Uint64P()
}

func (spec *ParityChainSpec) SetEIP152Transition(i *uint64) error {
	spec.SetPrecompile2(common.BytesToAddress([]byte{9}), "blake2_f", i, ParityChainSpecPricing{
		Blake2F: &ParityChainSpecBlakePricing{
			GasPerRound: 1,
		},
	})
	return nil
}

func (spec *ParityChainSpec) GetEIP170Transition() *uint64 {
	return spec.Params.MaxCodeSizeTransition.Uint64P()
}

func (spec *ParityChainSpec) SetEIP170Transition(i *uint64) error {
	spec.Params.MaxCodeSizeTransition = new(ParityU64).SetUint64(i)
	return nil
}

func (spec *ParityChainSpec) GetMaxCodeSize() *uint64 {
	return spec.Params.MaxCodeSize.Uint64P()
}

func (spec *ParityChainSpec) SetMaxCodeSize(i *uint64) error {
	spec.Params.MaxCodeSize = new(ParityU64).SetUint64(i)
	return nil
}

func (spec *ParityChainSpec) GetEIP198Transition() *uint64 {
	return spec.GetPrecompile(common.BytesToAddress([]byte{5}), ParityChainSpecPricing{
		ModExp: &ParityChainSpecModExpPricing{
			Divisor: 20,
		},
	}).Uint64P()
}

func (spec *ParityChainSpec) SetEIP198Transition(n *uint64) error {
	spec.SetPrecompile2(common.BytesToAddress([]byte{5}), "modexp", n, ParityChainSpecPricing{
		ModExp: &ParityChainSpecModExpPricing{
			Divisor: 20,
		},
	})
	return nil
}

func (spec *ParityChainSpec) GetEIP212Transition() *uint64 {
	f212 := spec.GetPrecompile(common.BytesToAddress([]byte{8}),
		ParityChainSpecPricing{
			AltBnPairing: &ParityChainSpecAltBnPairingPricing{
				Base: 100000,
				Pair: 80000,
			},
		}).Uint64P()

	if f212 != nil {
		return f212
	}
	return spec.GetEIP1108Transition()
}

func (spec *ParityChainSpec) SetEIP212Transition(n *uint64) error {
	spec.SetPrecompile2(common.BytesToAddress([]byte{8}), "alt_bn128_pairing", n, ParityChainSpecPricing{
		AltBnPairing: &ParityChainSpecAltBnPairingPricing{
			Base: 100000,
			Pair: 80000,
		},
	})
	return nil
}

func (spec *ParityChainSpec) GetEIP213Transition() *uint64 {
	x := spec.GetPrecompile(common.BytesToAddress([]byte{6}),
		ParityChainSpecPricing{
			AltBnConstOperation: &ParityChainSpecAltBnConstOperationPricing{
				Price: 500,
			},
		}).Uint64P()

	y := spec.GetPrecompile(common.BytesToAddress([]byte{7}),
		ParityChainSpecPricing{
			AltBnConstOperation: &ParityChainSpecAltBnConstOperationPricing{
				Price: 40000,
			},
		}).Uint64P()

	if x == nil && y == nil {
		return spec.GetEIP1108Transition()
	}
	if x == nil || y == nil {
		return nil
	}
	if *x != *y {
		return nil
	}
	return x
}

func (spec *ParityChainSpec) SetEIP213Transition(n *uint64) error {
	spec.SetPrecompile2(common.BytesToAddress([]byte{6}), "alt_bn128_add", n, ParityChainSpecPricing{
		AltBnConstOperation: &ParityChainSpecAltBnConstOperationPricing{
			Price: 500,
		},
	})
	spec.SetPrecompile2(common.BytesToAddress([]byte{7}), "alt_bn128_mul", n, ParityChainSpecPricing{
		AltBnConstOperation: &ParityChainSpecAltBnConstOperationPricing{
			Price: 40000,
		},
	})
	return nil
}

func (spec *ParityChainSpec) GetEIP1108Transition() *uint64 {
	x := spec.GetPrecompile(common.BytesToAddress([]byte{6}),
		ParityChainSpecPricing{
			AltBnConstOperation: &ParityChainSpecAltBnConstOperationPricing{
				Price: 150,
			},
		}).Uint64P()

	y := spec.GetPrecompile(common.BytesToAddress([]byte{7}),
		ParityChainSpecPricing{
			AltBnConstOperation: &ParityChainSpecAltBnConstOperationPricing{
				Price: 6000,
			},
		}).Uint64P()

	z := spec.GetPrecompile(common.BytesToAddress([]byte{8}),
		ParityChainSpecPricing{
			AltBnPairing: &ParityChainSpecAltBnPairingPricing{
				Base: 45000,
				Pair: 34000,
			},
		}).Uint64P()

	if x == nil || y == nil || z == nil {
		return nil
	}

	if *x != *y || *y != *z {
		return nil
	}
	return x
}

func (spec *ParityChainSpec) SetEIP1108Transition(n *uint64) error {
	spec.SetPrecompile2(common.BytesToAddress([]byte{6}), "alt_bn128_add", n, ParityChainSpecPricing{
		AltBnConstOperation: &ParityChainSpecAltBnConstOperationPricing{
			Price: 150,
		},
	})
	spec.SetPrecompile2(common.BytesToAddress([]byte{7}), "alt_bn128_mul", n, ParityChainSpecPricing{
		AltBnConstOperation: &ParityChainSpecAltBnConstOperationPricing{
			Price: 6000,
		},
	})
	spec.SetPrecompile2(common.BytesToAddress([]byte{8}), "alt_bn128_pairing", n, ParityChainSpecPricing{
		AltBnPairing: &ParityChainSpecAltBnPairingPricing{
			Base: 45000,
			Pair: 34000,
		},
	})
	return nil
}

func (spec *ParityChainSpec) IsForked(fn func() *uint64, n *big.Int) bool {
	f := fn()
	if f == nil || n == nil {
		return false
	}
	return big.NewInt(int64(*f)).Cmp(n) <= 0
}

func (spec *ParityChainSpec) GetForkCanonHash(n uint64) common.Hash {
	if spec.Params.ForkBlock == nil || spec.Params.ForkCanonHash == nil {
		return common.Hash{}
	}
	if spec.Params.ForkBlock.Big().Uint64() == n {
		return *spec.Params.ForkCanonHash
	}
	return common.Hash{}
}

func (spec *ParityChainSpec) SetForkCanonHash(n uint64, h common.Hash) error {
	spec.Params.ForkBlock = new(ParityU64).SetUint64(&n)
	spec.Params.ForkCanonHash = &h
	return nil
}

func (spec *ParityChainSpec) GetForkCanonHashes() map[uint64]common.Hash {
	if spec.Params.ForkBlock == nil || spec.Params.ForkCanonHash == nil {
		return nil
	}
	return map[uint64]common.Hash{
		spec.Params.ForkBlock.Big().Uint64(): *spec.Params.ForkCanonHash,
	}
}

func (spec *ParityChainSpec) GetConsensusEngineType() ctypes.ConsensusEngineT {
	if !reflect.DeepEqual(spec.Engine.Ethash, reflect.Zero(reflect.TypeOf(spec.Engine.Ethash)).Interface()) {
		return ctypes.ConsensusEngineT_Ethash
	}
	if !reflect.DeepEqual(spec.Engine.Clique, reflect.Zero(reflect.TypeOf(spec.Engine.Clique)).Interface()) {
		return ctypes.ConsensusEngineT_Clique
	}
	return ctypes.ConsensusEngineT_Unknown
}

func (spec *ParityChainSpec) MustSetConsensusEngineType(t ctypes.ConsensusEngineT) error {
	switch t {
	case ctypes.ConsensusEngineT_Ethash:
		return nil
	case ctypes.ConsensusEngineT_Clique:
		return nil
	default:
		return ctypes.ErrUnsupportedConfigFatal
	}
}

func (spec *ParityChainSpec) GetEthashMinimumDifficulty() *big.Int {
	return spec.Engine.Ethash.Params.MinimumDifficulty.ToInt()
}

func (spec *ParityChainSpec) SetEthashMinimumDifficulty(n *big.Int) error {
	if n == nil {
		return nil
	}
	spec.Engine.Ethash.Params.MinimumDifficulty = math.NewHexOrDecimal256(n.Int64())
	return nil
}

func (spec *ParityChainSpec) GetEthashDifficultyBoundDivisor() *big.Int {
	return spec.Engine.Ethash.Params.DifficultyBoundDivisor.ToInt()
}

func (spec *ParityChainSpec) SetEthashDifficultyBoundDivisor(n *big.Int) error {
	if n == nil {
		return nil
	}
	spec.Engine.Ethash.Params.DifficultyBoundDivisor = math.NewHexOrDecimal256(n.Int64())
	return nil
}

func (spec *ParityChainSpec) GetEthashDurationLimit() *big.Int {
	return spec.Engine.Ethash.Params.DurationLimit.ToInt()
}

func (spec *ParityChainSpec) SetEthashDurationLimit(n *big.Int) error {
	if n == nil {
		return nil
	}
	spec.Engine.Ethash.Params.DurationLimit = math.NewHexOrDecimal256(n.Int64())
	return nil
}

func (spec *ParityChainSpec) GetEthashHomesteadTransition() *uint64 {
	return spec.Engine.Ethash.Params.HomesteadTransition.Uint64P()
}

func (spec *ParityChainSpec) SetEthashHomesteadTransition(n *uint64) error {
	spec.Engine.Ethash.Params.HomesteadTransition = new(ParityU64).SetUint64(n)
	return nil
}

func (spec *ParityChainSpec) GetEthashEIP2Transition() *uint64 {
	return spec.Engine.Ethash.Params.HomesteadTransition.Uint64P()
}

func (spec *ParityChainSpec) SetEthashEIP2Transition(n *uint64) error {
	spec.Engine.Ethash.Params.HomesteadTransition = new(ParityU64).SetUint64(n)
	return nil
}

func (spec *ParityChainSpec) GetEthashEIP779Transition() *uint64 {
	return spec.Engine.Ethash.Params.DaoHardforkTransition.Uint64P()
}

func (spec *ParityChainSpec) SetEthashEIP779Transition(n *uint64) error {
	spec.Engine.Ethash.Params.DaoHardforkTransition = new(ParityU64).SetUint64(n)
	spec.Engine.Ethash.Params.DaoHardforkBeneficiary = &vars.DAORefundContract
	spec.Engine.Ethash.Params.DaoHardforkAccounts = vars.DAODrainList()
	return nil
}

func (spec *ParityChainSpec) GetEthashEIP649Transition() *uint64 {
	if spec.Engine.Ethash.Params.eip649inferred {
		return spec.Engine.Ethash.Params.eip649Transition.Uint64P()
	}

	var diffN *uint64
	defer func() {
		spec.Engine.Ethash.Params.eip649Transition = new(ParityU64).SetUint64(diffN)
		spec.Engine.Ethash.Params.eip649inferred = true
	}()

	diffN = ctypes.ExtractHostageSituationN(
		spec.Engine.Ethash.Params.DifficultyBombDelays,
		ctypes.Uint64BigMapEncodesHex(spec.Engine.Ethash.Params.BlockReward),
		vars.EIP649DifficultyBombDelay,
		vars.EIP649FBlockReward,
	)
	return diffN
}

func (spec *ParityChainSpec) SetEthashEIP649Transition(n *uint64) error {
	spec.Engine.Ethash.Params.eip649Transition = new(ParityU64).SetUint64(n)
	spec.Engine.Ethash.Params.eip649inferred = true
	if n == nil {
		return nil
	}
	if spec.Engine.Ethash.Params.BlockReward == nil {
		spec.Engine.Ethash.Params.BlockReward = ctypes.Uint64BigValOrMapHex{}
	}
	if spec.Engine.Ethash.Params.DifficultyBombDelays == nil {
		spec.Engine.Ethash.Params.DifficultyBombDelays = ctypes.Uint64BigMapEncodesHex{}
	}
	spec.Engine.Ethash.Params.BlockReward[*n] = vars.EIP649FBlockReward

	eip1234N := spec.Engine.Ethash.Params.eip1234Transition
	if eip1234N == nil || *eip1234N.Uint64P() != *n {
		spec.Engine.Ethash.Params.DifficultyBombDelays[*n] = vars.EIP649DifficultyBombDelay
	}
	// Else EIP1234 has been set to equal activation value, which means the map contains a sum value (eg 5m),
	// so the EIP649 difficulty adjustment is already accounted for.
	return nil
}

func (spec *ParityChainSpec) GetEthashEIP1234Transition() *uint64 {
	if spec.Engine.Ethash.Params.eip1234inferred {
		return spec.Engine.Ethash.Params.eip1234Transition.Uint64P()
	}

	var diffN *uint64
	defer func() {
		spec.Engine.Ethash.Params.eip1234Transition = new(ParityU64).SetUint64(diffN)
		spec.Engine.Ethash.Params.eip1234inferred = true
	}()

	diffN = ctypes.ExtractHostageSituationN(
		spec.Engine.Ethash.Params.DifficultyBombDelays,
		ctypes.Uint64BigMapEncodesHex(spec.Engine.Ethash.Params.BlockReward),
		vars.EIP1234DifficultyBombDelay,
		vars.EIP1234FBlockReward,
	)
	return diffN
}

func (spec *ParityChainSpec) SetEthashEIP1234Transition(n *uint64) error {
	spec.Engine.Ethash.Params.eip1234Transition = new(ParityU64).SetUint64(n)
	spec.Engine.Ethash.Params.eip1234inferred = true
	if n == nil {
		return nil
	}
	if spec.Engine.Ethash.Params.BlockReward == nil {
		spec.Engine.Ethash.Params.BlockReward = ctypes.Uint64BigValOrMapHex{}
	}
	if spec.Engine.Ethash.Params.DifficultyBombDelays == nil {
		spec.Engine.Ethash.Params.DifficultyBombDelays = ctypes.Uint64BigMapEncodesHex{}
	}
	// Block reward is a simple lookup; doesn't matter if overwrite or not.
	spec.Engine.Ethash.Params.BlockReward[*n] = vars.EIP1234FBlockReward

	eip649N := spec.Engine.Ethash.Params.eip649Transition
	if eip649N == nil || *eip649N.Uint64P() == *n {
		// EIP649 has NOT been set, OR has been set to identical block, eg. 0 for testing
		// Overwrite key with total delay (5m)
		spec.Engine.Ethash.Params.DifficultyBombDelays[*n] = vars.EIP1234DifficultyBombDelay
		return nil
	}

	spec.Engine.Ethash.Params.DifficultyBombDelays[*n] = new(big.Int).Sub(vars.EIP1234DifficultyBombDelay, vars.EIP649DifficultyBombDelay)

	return nil
}

func (spec *ParityChainSpec) GetEthashECIP1010PauseTransition() *uint64 {
	return spec.Engine.Ethash.Params.ECIP1010PauseTransition.Uint64P()
}

func (spec *ParityChainSpec) SetEthashECIP1010PauseTransition(n *uint64) error {
	spec.Engine.Ethash.Params.ECIP1010PauseTransition = new(ParityU64).SetUint64(n)
	return nil
}

func (spec *ParityChainSpec) GetEthashECIP1010ContinueTransition() *uint64 {
	return spec.Engine.Ethash.Params.ECIP1010ContinueTransition.Uint64P()
}

func (spec *ParityChainSpec) SetEthashECIP1010ContinueTransition(n *uint64) error {
	spec.Engine.Ethash.Params.ECIP1010ContinueTransition = new(ParityU64).SetUint64(n)
	return nil
}

// NOTE: Uses rounds as equivalence to transition.
// This is not per spec, but per implementation (it just so happened that the
// ETC fork happened at block 5m and rounds are 5m.
func (spec *ParityChainSpec) GetEthashECIP1017Transition() *uint64 {
	return spec.Engine.Ethash.Params.ECIP1017EraRounds.Uint64P()
}

func (spec *ParityChainSpec) SetEthashECIP1017Transition(n *uint64) error {
	// Even though this feature is not explicitly supported,
	// we'll follow the ad hoc logic as above.
	spec.Engine.Ethash.Params.ECIP1017EraRounds = new(ParityU64).SetUint64(n)
	return nil
}

func (spec *ParityChainSpec) GetEthashECIP1017EraRounds() *uint64 {
	return spec.Engine.Ethash.Params.ECIP1017EraRounds.Uint64P()
}

func (spec *ParityChainSpec) SetEthashECIP1017EraRounds(n *uint64) error {
	spec.Engine.Ethash.Params.ECIP1017EraRounds = new(ParityU64).SetUint64(n)
	return nil
}

func (spec *ParityChainSpec) GetEthashEIP100BTransition() *uint64 {
	return spec.Engine.Ethash.Params.EIP100bTransition.Uint64P()
}

func (spec *ParityChainSpec) SetEthashEIP100BTransition(n *uint64) error {
	spec.Engine.Ethash.Params.EIP100bTransition = new(ParityU64).SetUint64(n)
	return nil
}

func (spec *ParityChainSpec) GetEthashECIP1041Transition() *uint64 {
	return spec.Engine.Ethash.Params.BombDefuseTransition.Uint64P()
}

func (spec *ParityChainSpec) SetEthashECIP1041Transition(n *uint64) error {
	spec.Engine.Ethash.Params.BombDefuseTransition = new(ParityU64).SetUint64(n)
	return nil
}

func (spec *ParityChainSpec) GetEthashDifficultyBombDelaySchedule() ctypes.Uint64BigMapEncodesHex {
	if reflect.DeepEqual(spec.Engine.Ethash, reflect.Zero(reflect.TypeOf(spec.Engine.Ethash)).Interface()) {
		return nil
	}
	return spec.Engine.Ethash.Params.DifficultyBombDelays
}

func (spec *ParityChainSpec) SetEthashDifficultyBombDelaySchedule(input ctypes.Uint64BigMapEncodesHex) error {
	spec.Engine.Ethash.Params.DifficultyBombDelays = input
	return nil
}

func (spec *ParityChainSpec) GetEthashBlockRewardSchedule() ctypes.Uint64BigMapEncodesHex {
	if reflect.DeepEqual(spec.Engine.Ethash, reflect.Zero(reflect.TypeOf(spec.Engine.Ethash)).Interface()) {
		return nil
	}
	return ctypes.Uint64BigMapEncodesHex(spec.Engine.Ethash.Params.BlockReward)
}

func (spec *ParityChainSpec) SetEthashBlockRewardSchedule(input ctypes.Uint64BigMapEncodesHex) error {
	spec.Engine.Ethash.Params.BlockReward = ctypes.Uint64BigValOrMapHex(input)
	return nil
}

func (spec *ParityChainSpec) GetCliquePeriod() uint64 {
	p := spec.Engine.Clique.Params.Period.Uint64P()
	if p == nil {
		return 0
	}
	return *p
}

func (spec *ParityChainSpec) SetCliquePeriod(i uint64) error {
	spec.Engine.Clique.Params.Period = new(ParityU64).SetUint64(&i)
	return nil
}

func (spec *ParityChainSpec) GetCliqueEpoch() uint64 {
	p := spec.Engine.Clique.Params.Epoch.Uint64P()
	if p == nil {
		return 0
	}
	return *p
}

func (spec *ParityChainSpec) SetCliqueEpoch(i uint64) error {
	spec.Engine.Clique.Params.Epoch = new(ParityU64).SetUint64(&i)
	return nil
}

func (spec *ParityChainSpec) GetSealingType() ctypes.BlockSealingT {
	if !reflect.DeepEqual(spec.Genesis.Seal.Ethereum, reflect.Zero(reflect.TypeOf(spec.Genesis.Seal.Ethereum)).Interface()) {
		return ctypes.BlockSealing_Ethereum
	}
	log.Println(spew.Sdump(spec.Genesis))
	log.Println(spew.Sdump(reflect.Zero(reflect.TypeOf(spec.Genesis.Seal.Ethereum)).Interface()))
	b, _ := json.MarshalIndent(spec, "", "    ")
	log.Println(string(b))
	return ctypes.BlockSealing_Unknown
}

func (spec *ParityChainSpec) SetSealingType(in ctypes.BlockSealingT) error {
	switch in {
	case ctypes.BlockSealing_Ethereum:
		return nil
	}
	return ctypes.ErrUnsupportedConfigFatal
}

func (spec *ParityChainSpec) GetGenesisSealerEthereumNonce() uint64 {
	return spec.Genesis.Seal.Ethereum.Nonce.Uint64()
}

func (spec *ParityChainSpec) SetGenesisSealerEthereumNonce(i uint64) error {
	spec.Genesis.Seal.Ethereum.Nonce = types.EncodeNonce(i)
	return nil
}

func (spec *ParityChainSpec) GetGenesisSealerEthereumMixHash() common.Hash {
	return common.BytesToHash(spec.Genesis.Seal.Ethereum.MixHash)
}

func (spec *ParityChainSpec) SetGenesisSealerEthereumMixHash(input common.Hash) error {
	spec.Genesis.Seal.Ethereum.MixHash = input[:]
	return nil
}

func (spec *ParityChainSpec) GetGenesisDifficulty() *big.Int {
	return spec.Genesis.Difficulty.ToInt()
}

func (spec *ParityChainSpec) SetGenesisDifficulty(i *big.Int) error {
	spec.Genesis.Difficulty = math.NewHexOrDecimal256(i.Int64())
	return nil
}

func (spec *ParityChainSpec) GetGenesisAuthor() common.Address {
	return spec.Genesis.Author
}

func (spec *ParityChainSpec) SetGenesisAuthor(input common.Address) error {
	spec.Genesis.Author = input
	return nil
}

func (spec *ParityChainSpec) GetGenesisTimestamp() uint64 {
	return uint64(spec.Genesis.Timestamp)
}

func (spec *ParityChainSpec) SetGenesisTimestamp(i uint64) error {
	spec.Genesis.Timestamp = math.HexOrDecimal64(i)
	return nil
}

func (spec *ParityChainSpec) GetGenesisParentHash() common.Hash {
	return spec.Genesis.ParentHash
}

func (spec *ParityChainSpec) SetGenesisParentHash(input common.Hash) error {
	spec.Genesis.ParentHash = input
	return nil
}

func (spec *ParityChainSpec) GetGenesisExtraData() []byte {
	return spec.Genesis.ExtraData
}

func (spec *ParityChainSpec) SetGenesisExtraData(input []byte) error {
	spec.Genesis.ExtraData = input
	return nil
}

func (spec *ParityChainSpec) GetGenesisGasLimit() uint64 {
	return uint64(spec.Genesis.GasLimit)
}

func (spec *ParityChainSpec) SetGenesisGasLimit(i uint64) error {
	spec.Genesis.GasLimit = math.HexOrDecimal64(i)
	return nil
}

func (spec *ParityChainSpec) ForEachAccount(fn func(address common.Address, bal *big.Int, nonce uint64, code []byte, storage map[common.Hash]common.Hash) error) error {
	var err error
	for k, v := range spec.Accounts {
		if v.Builtin != nil && (v.Balance.ToInt() == nil || v.Balance.ToInt().Cmp(new(big.Int)) == 0) {
			continue
		}
		err = fn(common.Address(k), v.Balance.ToInt(), uint64(v.Nonce), v.Code, v.Storage)
		if err != nil {
			return err
		}
	}
	return nil
}

func (spec *ParityChainSpec) UpdateAccount(address common.Address, bal *big.Int, nonce uint64, code []byte, storage map[common.Hash]common.Hash) error {
	addr := common.UnprefixedAddress(address)
	if spec.Accounts == nil {
		spec.Accounts = make(map[common.UnprefixedAddress]*ParityChainSpecAccount)
	}
	_, ok := spec.Accounts[addr]
	if !ok {
		spec.Accounts[addr] = &ParityChainSpecAccount{}
	}
	spec.Accounts[addr].Balance = *math.NewHexOrDecimal256(bal.Int64())
	spec.Accounts[addr].Nonce = math.HexOrDecimal64(nonce)

	zero := uint64(0)
	switch address {
	case common.BytesToAddress([]byte{1}):
		spec.SetPrecompile2(common.BytesToAddress([]byte{1}), "ecrecover", &zero, ParityChainSpecPricing{
			Linear: &ParityChainSpecLinearPricing{
				Base: 3000,
			},
		})
	case common.BytesToAddress([]byte{2}):
		spec.SetPrecompile2(common.BytesToAddress([]byte{2}), "sha256", &zero, ParityChainSpecPricing{
			Linear: &ParityChainSpecLinearPricing{
				Base: 60,
				Word: 2,
			},
		})
	case common.BytesToAddress([]byte{3}):
		spec.SetPrecompile2(common.BytesToAddress([]byte{3}), "ripemd160", &zero, ParityChainSpecPricing{
			Linear: &ParityChainSpecLinearPricing{
				Base: 600,
				Word: 1,
			},
		})
	case common.BytesToAddress([]byte{4}):
		spec.SetPrecompile2(common.BytesToAddress([]byte{4}), "identity", &zero, ParityChainSpecPricing{
			Linear: &ParityChainSpecLinearPricing{
				Base: 15,
				Word: 3,
			},
		})
	}
	return nil
}