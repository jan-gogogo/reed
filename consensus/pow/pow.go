// Copyright 2020 The Reed Developers
// Distributed under the MIT software license, see the accompanying
// file COPYING or http://www.opensource.org/licenses/mit-license.php.

package pow

import (
	"github.com/reed/types"
	"math/big"
)

func CheckProofOfWork(target big.Int, hash types.Hash) bool {
	var hashIntVal big.Int

	hashIntVal.SetBytes(hash.Bytes())
	return target.Cmp(&hashIntVal) == -1
}
