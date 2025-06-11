package contracts

import (
	"math/big"
	"math/bits"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// OBLevel represents an orderbook level with tick and volume
type OBLevel struct {
	Tick uint32
	Vol  *big.Int
}

// BitScanHelper provides bit scanning utilities
type BitScanHelper struct{}

// LSB finds the index of the least significant bit (rightmost 1-bit)
func (BitScanHelper) LSB(x *big.Int) uint8 {
	if x.Sign() == 0 {
		return 0
	}
	// Convert to uint64 for trailing zeros operation
	return uint8(bits.TrailingZeros64(x.Uint64()))
}

// MSB finds the index of the most significant bit (leftmost 1-bit)
func (BitScanHelper) MSB(x *big.Int) uint8 {
	if x.Sign() == 0 {
		return 0
	}
	// Convert to uint64 for leading zeros operation
	return uint8(63 - bits.LeadingZeros64(x.Uint64()))
}

// Join combines three 8-bit values into a 24-bit tick
func (BitScanHelper) Join(l1, l2, l3 uint8) uint32 {
	return uint32(l1)<<16 | uint32(l2)<<8 | uint32(l3)
}

// Split breaks a 24-bit tick into three 8-bit components
func (BitScanHelper) Split(tick uint32) (uint8, uint8, uint8) {
	l1 := uint8((tick >> 16) & 0xFF)
	l2 := uint8((tick >> 8) & 0xFF)
	l3 := uint8(tick & 0xFF)
	return l1, l2, l3
}

// GetOrderbook replicates the dumpBook functionality from Solidity
func GetOrderbook(client *ethclient.Client, marketAddress common.Address, isBid bool, isPut bool) ([]OBLevel, error) {
	market, err := NewMarket(marketAddress, client)
	if err != nil {
		return nil, err
	}

	bitScan := BitScanHelper{}
	var levels []OBLevel

	// Get summary index (replicating _summaryIx internal function)
	summaryIdx := getSummaryIndex(isBid, isPut)

	// Get summary bitmap
	summary, err := market.Summaries(nil, big.NewInt(int64(summaryIdx)))
	if err != nil {
		return nil, err
	}

	// Pass 1: Count levels (replicating the first pass of dumpBook)
	levelCount := 0
	tempSummary := new(big.Int).Set(summary)

	for tempSummary.Sign() != 0 {
		l1 := bitScan.LSB(tempSummary)

		// Get mid-level bitmap
		mid, err := getMidBitmap(market, isBid, isPut, l1)
		if err != nil {
			return nil, err
		}

		tempMid := new(big.Int).Set(mid)
		for tempMid.Sign() != 0 {
			l2 := bitScan.LSB(tempMid)
			detKey := uint16(l1)<<8 | uint16(l2)

			// Get detail-level bitmap
			det, err := getDetailBitmap(market, isBid, isPut, detKey)
			if err != nil {
				return nil, err
			}

			tempDet := new(big.Int).Set(det)
			for tempDet.Sign() != 0 {
				levelCount++
				// Clear lowest set bit: d &= d - 1
				tempDet.And(tempDet, new(big.Int).Sub(tempDet, big.NewInt(1)))
			}
			// Clear lowest set bit: m &= m - 1
			tempMid.And(tempMid, new(big.Int).Sub(tempMid, big.NewInt(1)))
		}
		// Clear lowest set bit: s &= s - 1
		tempSummary.And(tempSummary, new(big.Int).Sub(tempSummary, big.NewInt(1)))
	}

	// Pass 2: Collect the actual data (replicating the second pass of dumpBook)
	levels = make([]OBLevel, 0, levelCount)
	tempSummary = new(big.Int).Set(summary)

	for tempSummary.Sign() != 0 {
		l1 := bitScan.LSB(tempSummary)

		// Get mid-level bitmap
		mid, err := getMidBitmap(market, isBid, isPut, l1)
		if err != nil {
			return nil, err
		}

		tempMid := new(big.Int).Set(mid)
		for tempMid.Sign() != 0 {
			l2 := bitScan.LSB(tempMid)
			detKey := uint16(l1)<<8 | uint16(l2)

			// Get detail-level bitmap
			det, err := getDetailBitmap(market, isBid, isPut, detKey)
			if err != nil {
				return nil, err
			}

			tempDet := new(big.Int).Set(det)
			for tempDet.Sign() != 0 {
				l3 := bitScan.LSB(tempDet)
				tick := bitScan.Join(l1, l2, l3)
				tickKey := getTickKey(tick, isPut, isBid)

				// Get level data
				level, err := market.Levels(nil, tickKey)
				if err != nil {
					return nil, err
				}

				levels = append(levels, OBLevel{
					Tick: tick,
					Vol:  level.Vol,
				})

				// Clear lowest set bit: d &= d - 1
				tempDet.And(tempDet, new(big.Int).Sub(tempDet, big.NewInt(1)))
			}
			// Clear lowest set bit: m &= m - 1
			tempMid.And(tempMid, new(big.Int).Sub(tempMid, big.NewInt(1)))
		}
		// Clear lowest set bit: s &= s - 1
		tempSummary.And(tempSummary, new(big.Int).Sub(tempSummary, big.NewInt(1)))
	}

	return levels, nil
}

// getSummaryIndex replicates the _summaryIx internal function
func getSummaryIndex(isBid bool, isPut bool) uint8 {
	result := uint8(0)
	if isBid {
		result |= 1
	}
	if isPut {
		result |= 2
	}
	return result
}

// getTickKey replicates the _key internal function
func getTickKey(tick uint32, isPut bool, isBid bool) uint32 {
	key := tick & 0x00FFFFFF // Take only rightmost 24 bits
	if isPut {
		key |= 1 << 31
	}
	if isBid {
		key |= 1 << 30
	}
	return key
}

// getMidBitmap gets the appropriate mid-level bitmap based on side
func getMidBitmap(market *Market, isBid bool, isPut bool, l1 uint8) (*big.Int, error) {
	if isBid {
		if isPut {
			return market.MidPB(nil, l1)
		} else {
			return market.MidCB(nil, l1)
		}
	} else {
		if isPut {
			return market.MidPA(nil, l1)
		} else {
			return market.MidCA(nil, l1)
		}
	}
}

// getDetailBitmap gets the appropriate detail-level bitmap based on side
func getDetailBitmap(market *Market, isBid bool, isPut bool, detKey uint16) (*big.Int, error) {
	if isBid {
		if isPut {
			return market.DetPB(nil, detKey)
		} else {
			return market.DetCB(nil, detKey)
		}
	} else {
		if isPut {
			return market.DetPA(nil, detKey)
		} else {
			return market.DetCA(nil, detKey)
		}
	}
}
