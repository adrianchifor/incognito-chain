package netsync

import (
	"encoding/json"
	"fmt"
	"github.com/incognitochain/incognito-chain/blockchain"
	"github.com/incognitochain/incognito-chain/common"
	"github.com/incognitochain/incognito-chain/mempool"
	libp2p "github.com/libp2p/go-libp2p-peer"
	"testing"
)

var (
	senderID                   = "QmSoLer265NRgSp2LA3dPaeykiS1J6DifTC88f5uVQKNAd"
	hash0                      = common.HashH([]byte{0})
	hash1                      = common.HashH([]byte{1})
	hash2                      = common.HashH([]byte{2})
	peerID                     libp2p.ID
	shardBlockByteNoCrossShard = []byte{123, 34, 65, 103, 103, 114, 101, 103, 97, 116, 101, 100, 83, 105, 103, 34, 58, 34, 49, 50, 109, 90, 114, 76, 120, 74, 80, 55, 90, 113, 88, 78, 68, 68, 51, 67, 113, 100, 121, 104, 112, 66, 106, 102, 117, 74, 103, 105, 70, 56, 122, 121, 99, 80, 122, 111, 90, 116, 114, 77, 49, 51, 103, 113, 98, 57, 66, 104, 57, 112, 54, 49, 56, 88, 113, 74, 109, 71, 122, 97, 55, 70, 98, 120, 106, 52, 119, 81, 69, 105, 82, 80, 106, 80, 100, 68, 109, 121, 74, 111, 82, 75, 82, 118, 100, 121, 57, 102, 71, 98, 83, 52, 105, 34, 44, 34, 82, 34, 58, 34, 49, 56, 53, 111, 72, 89, 56, 115, 90, 78, 116, 81, 54, 98, 84, 56, 83, 80, 77, 72, 98, 104, 69, 113, 98, 122, 80, 76, 117, 87, 65, 88, 74, 89, 111, 85, 78, 87, 74, 81, 114, 56, 74, 85, 89, 122, 71, 109, 56, 65, 82, 34, 44, 34, 86, 97, 108, 105, 100, 97, 116, 111, 114, 115, 73, 100, 120, 34, 58, 91, 91, 48, 44, 49, 44, 50, 44, 51, 93, 44, 91, 48, 44, 49, 44, 50, 93, 93, 44, 34, 80, 114, 111, 100, 117, 99, 101, 114, 83, 105, 103, 34, 58, 34, 49, 90, 53, 104, 111, 50, 72, 52, 90, 82, 107, 80, 105, 65, 82, 76, 115, 50, 82, 86, 70, 55, 51, 71, 106, 86, 103, 104, 110, 52, 70, 113, 65, 65, 81, 103, 117, 82, 109, 51, 86, 83, 118, 115, 117, 116, 83, 50, 89, 49, 87, 111, 74, 121, 103, 105, 71, 111, 106, 114, 57, 49, 97, 97, 77, 119, 99, 65, 119, 78, 57, 83, 66, 82, 70, 121, 77, 69, 51, 51, 66, 89, 122, 119, 116, 56, 83, 77, 111, 69, 70, 80, 65, 120, 34, 44, 34, 66, 111, 100, 121, 34, 58, 123, 34, 73, 110, 115, 116, 114, 117, 99, 116, 105, 111, 110, 115, 34, 58, 91, 93, 44, 34, 67, 114, 111, 115, 115, 84, 114, 97, 110, 115, 97, 99, 116, 105, 111, 110, 115, 34, 58, 123, 125, 44, 34, 84, 114, 97, 110, 115, 97, 99, 116, 105, 111, 110, 115, 34, 58, 91, 93, 125, 44, 34, 72, 101, 97, 100, 101, 114, 34, 58, 123, 34, 80, 114, 111, 100, 117, 99, 101, 114, 65, 100, 100, 114, 101, 115, 115, 34, 58, 123, 34, 80, 107, 34, 58, 34, 65, 53, 90, 117, 78, 52, 97, 116, 66, 110, 105, 70, 89, 97, 52, 120, 117, 122, 48, 66, 53, 101, 103, 47, 70, 80, 51, 105, 68, 77, 108, 57, 55, 101, 90, 43, 55, 97, 90, 73, 48, 106, 115, 65, 34, 44, 34, 84, 107, 34, 58, 34, 65, 116, 54, 98, 72, 84, 100, 47, 107, 83, 67, 57, 78, 47, 78, 116, 78, 118, 70, 108, 67, 89, 47, 116, 47, 82, 90, 118, 119, 85, 108, 113, 112, 98, 52, 90, 43, 102, 77, 105, 101, 70, 74, 66, 34, 125, 44, 34, 83, 104, 97, 114, 100, 73, 68, 34, 58, 48, 44, 34, 86, 101, 114, 115, 105, 111, 110, 34, 58, 49, 44, 34, 80, 114, 101, 118, 66, 108, 111, 99, 107, 72, 97, 115, 104, 34, 58, 34, 55, 56, 57, 53, 51, 101, 56, 49, 57, 52, 101, 51, 49, 48, 48, 51, 56, 99, 99, 56, 53, 49, 52, 100, 51, 51, 98, 53, 48, 53, 55, 50, 56, 99, 51, 97, 57, 55, 56, 55, 98, 100, 55, 51, 50, 52, 50, 48, 97, 55, 56, 56, 48, 51, 51, 53, 55, 54, 100, 102, 51, 99, 99, 52, 34, 44, 34, 72, 101, 105, 103, 104, 116, 34, 58, 50, 44, 34, 82, 111, 117, 110, 100, 34, 58, 49, 44, 34, 69, 112, 111, 99, 104, 34, 58, 49, 44, 34, 84, 105, 109, 101, 115, 116, 97, 109, 112, 34, 58, 49, 53, 54, 49, 55, 57, 50, 53, 53, 49, 44, 34, 84, 120, 82, 111, 111, 116, 34, 58, 34, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 34, 44, 34, 83, 104, 97, 114, 100, 84, 120, 82, 111, 111, 116, 34, 58, 34, 50, 54, 101, 55, 100, 50, 49, 102, 98, 54, 100, 102, 102, 101, 53, 57, 56, 97, 97, 57, 49, 101, 55, 56, 48, 57, 57, 53, 56, 56, 53, 100, 97, 101, 56, 50, 49, 50, 97, 97, 99, 57, 56, 51, 57, 52, 102, 98, 57, 56, 52, 50, 99, 99, 55, 97, 98, 54, 52, 102, 48, 53, 52, 51, 34, 44, 34, 67, 114, 111, 115, 115, 84, 114, 97, 110, 115, 97, 99, 116, 105, 111, 110, 82, 111, 111, 116, 34, 58, 34, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 34, 44, 34, 73, 110, 115, 116, 114, 117, 99, 116, 105, 111, 110, 115, 82, 111, 111, 116, 34, 58, 34, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 34, 44, 34, 67, 111, 109, 109, 105, 116, 116, 101, 101, 82, 111, 111, 116, 34, 58, 34, 57, 50, 98, 102, 97, 57, 50, 57, 98, 99, 50, 97, 54, 97, 49, 54, 57, 100, 99, 52, 102, 53, 99, 97, 52, 56, 50, 52, 102, 56, 52, 51, 53, 55, 99, 99, 57, 53, 56, 56, 55, 50, 56, 52, 55, 56, 53, 56, 49, 101, 54, 49, 97, 53, 56, 98, 54, 49, 56, 56, 53, 98, 50, 51, 34, 44, 34, 80, 101, 110, 100, 105, 110, 103, 86, 97, 108, 105, 100, 97, 116, 111, 114, 82, 111, 111, 116, 34, 58, 34, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 34, 44, 34, 67, 114, 111, 115, 115, 83, 104, 97, 114, 100, 115, 34, 58, 34, 34, 44, 34, 66, 101, 97, 99, 111, 110, 72, 101, 105, 103, 104, 116, 34, 58, 49, 44, 34, 66, 101, 97, 99, 111, 110, 72, 97, 115, 104, 34, 58, 34, 55, 101, 48, 101, 97, 98, 101, 101, 49, 99, 97, 51, 52, 54, 48, 56, 101, 97, 50, 98, 55, 57, 53, 48, 101, 99, 53, 99, 98, 48, 98, 48, 53, 50, 55, 99, 51, 54, 49, 57, 54, 53, 102, 51, 101, 53, 102, 49, 100, 57, 56, 54, 56, 54, 53, 50, 51, 52, 57, 100, 100, 52, 101, 54, 34, 44, 34, 84, 111, 116, 97, 108, 84, 120, 115, 70, 101, 101, 34, 58, 123, 125, 125, 125}
	shardBlockByteCrossShard01 = []byte{123, 34, 65, 103, 103, 114, 101, 103, 97, 116, 101, 100, 83, 105, 103, 34, 58, 34, 49, 51, 66, 66, 120, 83, 50, 99, 84, 89, 85, 105, 85, 106, 109, 113, 55, 88, 90, 85, 115, 117, 81, 111, 82, 98, 88, 111, 65, 56, 87, 87, 74, 99, 70, 81, 56, 105, 52, 54, 97, 76, 88, 76, 109, 86, 54, 69, 101, 119, 89, 122, 122, 77, 53, 111, 99, 111, 113, 112, 56, 57, 116, 103, 87, 55, 114, 122, 118, 97, 54, 69, 118, 110, 119, 68, 100, 114, 106, 78, 102, 81, 98, 97, 80, 110, 110, 83, 121, 82, 97, 102, 68, 116, 85, 34, 44, 34, 82, 34, 58, 34, 49, 56, 75, 110, 111, 85, 65, 55, 105, 69, 75, 105, 86, 86, 65, 121, 90, 99, 49, 80, 52, 81, 67, 49, 100, 72, 105, 115, 98, 57, 77, 70, 121, 111, 101, 114, 69, 110, 75, 66, 119, 88, 106, 112, 102, 50, 50, 114, 68, 70, 117, 34, 44, 34, 86, 97, 108, 105, 100, 97, 116, 111, 114, 115, 73, 100, 120, 34, 58, 91, 91, 48, 44, 49, 44, 50, 93, 44, 91, 48, 44, 49, 44, 50, 93, 93, 44, 34, 80, 114, 111, 100, 117, 99, 101, 114, 83, 105, 103, 34, 58, 34, 49, 69, 119, 111, 49, 119, 75, 77, 81, 111, 50, 106, 67, 110, 118, 71, 71, 82, 53, 100, 113, 75, 74, 74, 85, 117, 71, 111, 111, 83, 84, 103, 90, 77, 85, 76, 81, 104, 55, 115, 116, 72, 74, 56, 85, 78, 71, 52, 65, 74, 68, 113, 101, 99, 56, 53, 121, 56, 114, 118, 89, 57, 97, 119, 90, 49, 78, 98, 90, 99, 80, 83, 121, 119, 81, 106, 55, 111, 118, 68, 55, 116, 83, 83, 113, 83, 74, 122, 99, 101, 120, 89, 49, 99, 34, 44, 34, 66, 111, 100, 121, 34, 58, 123, 34, 73, 110, 115, 116, 114, 117, 99, 116, 105, 111, 110, 115, 34, 58, 91, 93, 44, 34, 67, 114, 111, 115, 115, 84, 114, 97, 110, 115, 97, 99, 116, 105, 111, 110, 115, 34, 58, 123, 125, 44, 34, 84, 114, 97, 110, 115, 97, 99, 116, 105, 111, 110, 115, 34, 58, 91, 123, 34, 86, 101, 114, 115, 105, 111, 110, 34, 58, 49, 44, 34, 84, 121, 112, 101, 34, 58, 34, 110, 34, 44, 34, 76, 111, 99, 107, 84, 105, 109, 101, 34, 58, 49, 53, 54, 49, 55, 57, 51, 55, 51, 55, 44, 34, 70, 101, 101, 34, 58, 51, 48, 44, 34, 73, 110, 102, 111, 34, 58, 34, 34, 44, 34, 83, 105, 103, 80, 117, 98, 75, 101, 121, 34, 58, 34, 65, 54, 122, 109, 70, 113, 73, 108, 84, 75, 103, 115, 86, 50, 51, 81, 107, 57, 106, 122, 50, 114, 111, 111, 51, 86, 104, 105, 115, 86, 121, 53, 70, 108, 103, 54, 69, 71, 117, 79, 75, 97, 81, 65, 34, 44, 34, 83, 105, 103, 34, 58, 34, 121, 82, 81, 70, 56, 101, 47, 79, 81, 81, 89, 102, 79, 119, 117, 119, 75, 52, 67, 79, 114, 68, 90, 87, 48, 48, 72, 77, 66, 109, 122, 86, 111, 121, 118, 103, 107, 81, 65, 85, 88, 115, 112, 68, 82, 55, 54, 80, 49, 69, 111, 82, 51, 105, 107, 112, 65, 112, 77, 78, 109, 83, 73, 99, 110, 68, 100, 106, 86, 71, 90, 88, 86, 102, 50, 53, 55, 118, 105, 72, 89, 112, 100, 106, 115, 65, 61, 61, 34, 44, 34, 80, 114, 111, 111, 102, 34, 58, 34, 49, 49, 49, 54, 52, 121, 74, 115, 113, 89, 120, 69, 65, 114, 75, 78, 97, 81, 56, 100, 54, 84, 82, 109, 84, 110, 50, 84, 107, 106, 113, 78, 77, 86, 104, 109, 66, 110, 81, 77, 80, 109, 74, 116, 88, 85, 111, 50, 118, 67, 86, 112, 70, 106, 109, 86, 70, 112, 81, 72, 102, 80, 115, 53, 84, 119, 55, 97, 87, 51, 78, 49, 80, 111, 119, 115, 104, 102, 75, 71, 110, 113, 83, 98, 54, 106, 112, 66, 83, 98, 83, 111, 84, 118, 69, 75, 84, 56, 52, 82, 54, 56, 75, 85, 49, 100, 66, 55, 85, 103, 100, 66, 49, 104, 65, 80, 101, 109, 112, 53, 51, 100, 101, 112, 84, 99, 65, 52, 102, 55, 69, 101, 75, 52, 121, 119, 49, 106, 51, 99, 122, 74, 72, 65, 113, 101, 51, 50, 114, 50, 97, 103, 105, 113, 90, 70, 107, 87, 105, 109, 107, 56, 51, 90, 80, 100, 75, 107, 97, 52, 107, 72, 71, 76, 102, 117, 120, 106, 117, 85, 107, 83, 66, 102, 69, 75, 111, 105, 121, 72, 114, 106, 80, 69, 112, 119, 56, 85, 86, 90, 82, 112, 103, 71, 100, 100, 76, 85, 106, 103, 100, 56, 98, 71, 90, 55, 110, 77, 102, 88, 52, 85, 72, 57, 120, 53, 70, 83, 67, 106, 54, 119, 118, 82, 66, 78, 112, 109, 74, 65, 105, 83, 69, 50, 85, 72, 112, 49, 115, 99, 83, 75, 57, 76, 83, 102, 99, 85, 69, 65, 100, 49, 55, 70, 51, 100, 109, 116, 66, 70, 117, 80, 101, 99, 86, 67, 74, 115, 83, 67, 100, 113, 109, 82, 55, 68, 75, 115, 54, 78, 116, 75, 89, 49, 51, 118, 71, 113, 105, 78, 83, 66, 80, 82, 116, 118, 99, 74, 76, 115, 86, 119, 66, 111, 84, 76, 52, 75, 86, 74, 70, 106, 50, 118, 85, 107, 112, 84, 69, 110, 78, 99, 118, 106, 65, 119, 106, 90, 104, 90, 52, 51, 111, 51, 50, 71, 50, 101, 114, 88, 49, 98, 99, 71, 97, 105, 81, 49, 119, 49, 68, 66, 110, 112, 84, 54, 88, 67, 68, 88, 100, 81, 119, 57, 101, 120, 66, 74, 74, 115, 86, 69, 75, 111, 121, 51, 88, 101, 98, 109, 117, 53, 105, 55, 50, 75, 76, 74, 72, 68, 54, 105, 101, 70, 51, 105, 88, 82, 117, 100, 113, 70, 66, 102, 104, 65, 81, 65, 57, 66, 105, 113, 71, 109, 78, 57, 76, 109, 65, 120, 82, 71, 83, 85, 113, 109, 121, 78, 77, 97, 84, 89, 119, 109, 116, 68, 77, 85, 114, 70, 119, 89, 68, 101, 116, 106, 98, 88, 116, 121, 101, 97, 107, 78, 50, 70, 97, 120, 88, 114, 76, 76, 98, 50, 109, 97, 119, 114, 122, 89, 66, 74, 68, 84, 89, 99, 106, 78, 119, 78, 57, 98, 115, 50, 82, 102, 89, 98, 81, 106, 100, 76, 57, 57, 101, 102, 83, 99, 104, 66, 81, 112, 53, 122, 110, 99, 74, 80, 71, 53, 89, 102, 116, 52, 101, 104, 121, 112, 107, 104, 97, 66, 84, 114, 103, 112, 81, 113, 97, 53, 113, 66, 114, 69, 105, 111, 54, 122, 118, 78, 88, 98, 56, 106, 55, 120, 86, 77, 68, 55, 77, 69, 113, 98, 56, 80, 109, 104, 66, 76, 50, 87, 56, 84, 117, 114, 113, 50, 55, 88, 70, 68, 77, 85, 51, 71, 86, 103, 109, 70, 113, 118, 104, 117, 107, 69, 82, 100, 86, 101, 111, 68, 107, 51, 104, 121, 116, 99, 120, 109, 112, 115, 102, 88, 114, 111, 99, 105, 90, 68, 97, 99, 88, 50, 89, 76, 80, 53, 85, 113, 101, 90, 53, 117, 72, 109, 122, 78, 50, 70, 71, 88, 55, 101, 105, 84, 49, 70, 120, 105, 81, 87, 57, 55, 105, 83, 70, 89, 101, 98, 69, 121, 107, 100, 57, 52, 115, 102, 103, 104, 100, 51, 57, 89, 111, 53, 55, 119, 67, 81, 83, 71, 86, 50, 85, 99, 102, 76, 78, 68, 116, 56, 84, 57, 77, 87, 68, 90, 72, 113, 83, 78, 118, 102, 99, 74, 99, 52, 81, 65, 78, 57, 117, 77, 121, 116, 67, 122, 101, 122, 74, 100, 53, 66, 68, 114, 51, 80, 116, 50, 57, 84, 104, 74, 99, 114, 87, 69, 101, 115, 75, 97, 97, 107, 77, 69, 100, 74, 102, 99, 121, 77, 110, 51, 74, 86, 83, 97, 122, 120, 85, 51, 90, 54, 115, 71, 111, 84, 112, 55, 116, 67, 102, 51, 72, 85, 74, 72, 111, 85, 116, 103, 113, 83, 75, 103, 109, 50, 103, 75, 109, 53, 112, 115, 75, 117, 117, 83, 121, 103, 116, 114, 56, 99, 57, 122, 98, 87, 105, 102, 72, 55, 98, 75, 112, 80, 82, 50, 119, 54, 90, 113, 99, 115, 97, 103, 110, 49, 77, 103, 50, 85, 86, 90, 82, 71, 49, 114, 65, 82, 56, 88, 100, 72, 100, 77, 106, 87, 122, 78, 77, 107, 104, 120, 83, 109, 98, 121, 80, 107, 107, 82, 55, 106, 87, 109, 105, 70, 88, 57, 80, 107, 85, 103, 101, 118, 111, 122, 99, 110, 116, 89, 65, 106, 114, 122, 101, 71, 116, 88, 100, 113, 74, 70, 65, 115, 70, 88, 102, 120, 52, 112, 57, 105, 112, 66, 87, 118, 68, 56, 50, 53, 118, 68, 80, 67, 34, 44, 34, 80, 117, 98, 75, 101, 121, 76, 97, 115, 116, 66, 121, 116, 101, 83, 101, 110, 100, 101, 114, 34, 58, 48, 44, 34, 77, 101, 116, 97, 100, 97, 116, 97, 34, 58, 110, 117, 108, 108, 125, 93, 125, 44, 34, 72, 101, 97, 100, 101, 114, 34, 58, 123, 34, 80, 114, 111, 100, 117, 99, 101, 114, 65, 100, 100, 114, 101, 115, 115, 34, 58, 123, 34, 80, 107, 34, 58, 34, 65, 53, 90, 117, 78, 52, 97, 116, 66, 110, 105, 70, 89, 97, 52, 120, 117, 122, 48, 66, 53, 101, 103, 47, 70, 80, 51, 105, 68, 77, 108, 57, 55, 101, 90, 43, 55, 97, 90, 73, 48, 106, 115, 65, 34, 44, 34, 84, 107, 34, 58, 34, 65, 116, 54, 98, 72, 84, 100, 47, 107, 83, 67, 57, 78, 47, 78, 116, 78, 118, 70, 108, 67, 89, 47, 116, 47, 82, 90, 118, 119, 85, 108, 113, 112, 98, 52, 90, 43, 102, 77, 105, 101, 70, 74, 66, 34, 125, 44, 34, 83, 104, 97, 114, 100, 73, 68, 34, 58, 48, 44, 34, 86, 101, 114, 115, 105, 111, 110, 34, 58, 49, 44, 34, 80, 114, 101, 118, 66, 108, 111, 99, 107, 72, 97, 115, 104, 34, 58, 34, 49, 51, 57, 97, 56, 49, 102, 100, 53, 101, 55, 53, 101, 102, 54, 57, 98, 53, 48, 101, 48, 49, 57, 100, 51, 52, 51, 53, 48, 48, 51, 52, 57, 56, 99, 97, 102, 51, 99, 55, 53, 55, 53, 100, 49, 100, 50, 51, 102, 100, 49, 50, 102, 55, 101, 52, 49, 55, 101, 102, 48, 53, 101, 48, 34, 44, 34, 72, 101, 105, 103, 104, 116, 34, 58, 51, 54, 44, 34, 82, 111, 117, 110, 100, 34, 58, 50, 44, 34, 69, 112, 111, 99, 104, 34, 58, 49, 44, 34, 84, 105, 109, 101, 115, 116, 97, 109, 112, 34, 58, 49, 53, 54, 49, 55, 57, 51, 55, 53, 50, 44, 34, 84, 120, 82, 111, 111, 116, 34, 58, 34, 51, 100, 102, 101, 52, 99, 101, 101, 53, 56, 54, 51, 53, 50, 48, 56, 55, 50, 52, 98, 48, 98, 101, 100, 101, 99, 100, 101, 50, 49, 56, 57, 48, 56, 53, 55, 100, 52, 52, 56, 55, 55, 98, 99, 100, 102, 52, 98, 50, 50, 53, 55, 101, 50, 57, 54, 50, 55, 50, 98, 99, 54, 48, 97, 34, 44, 34, 83, 104, 97, 114, 100, 84, 120, 82, 111, 111, 116, 34, 58, 34, 97, 57, 53, 53, 97, 54, 55, 102, 55, 100, 50, 98, 54, 51, 57, 55, 97, 49, 98, 49, 54, 98, 56, 57, 53, 98, 97, 100, 98, 57, 49, 50, 49, 99, 98, 55, 54, 50, 97, 100, 98, 53, 53, 52, 98, 100, 51, 102, 50, 49, 53, 48, 102, 57, 51, 53, 53, 101, 51, 54, 100, 52, 99, 100, 34, 44, 34, 67, 114, 111, 115, 115, 84, 114, 97, 110, 115, 97, 99, 116, 105, 111, 110, 82, 111, 111, 116, 34, 58, 34, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 34, 44, 34, 73, 110, 115, 116, 114, 117, 99, 116, 105, 111, 110, 115, 82, 111, 111, 116, 34, 58, 34, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 34, 44, 34, 67, 111, 109, 109, 105, 116, 116, 101, 101, 82, 111, 111, 116, 34, 58, 34, 102, 49, 57, 51, 101, 53, 102, 51, 101, 53, 98, 57, 97, 97, 100, 52, 57, 57, 49, 101, 48, 53, 101, 99, 50, 52, 51, 101, 56, 50, 49, 54, 54, 52, 99, 98, 53, 102, 53, 101, 102, 54, 53, 49, 101, 49, 99, 49, 55, 98, 55, 48, 101, 54, 101, 99, 101, 48, 98, 51, 102, 54, 48, 102, 34, 44, 34, 80, 101, 110, 100, 105, 110, 103, 86, 97, 108, 105, 100, 97, 116, 111, 114, 82, 111, 111, 116, 34, 58, 34, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 34, 44, 34, 67, 114, 111, 115, 115, 83, 104, 97, 114, 100, 115, 34, 58, 34, 65, 81, 61, 61, 34, 44, 34, 66, 101, 97, 99, 111, 110, 72, 101, 105, 103, 104, 116, 34, 58, 49, 44, 34, 66, 101, 97, 99, 111, 110, 72, 97, 115, 104, 34, 58, 34, 55, 101, 48, 101, 97, 98, 101, 101, 49, 99, 97, 51, 52, 54, 48, 56, 101, 97, 50, 98, 55, 57, 53, 48, 101, 99, 53, 99, 98, 48, 98, 48, 53, 50, 55, 99, 51, 54, 49, 57, 54, 53, 102, 51, 101, 53, 102, 49, 100, 57, 56, 54, 56, 54, 53, 50, 51, 52, 57, 100, 100, 52, 101, 54, 34, 44, 34, 84, 111, 116, 97, 108, 84, 120, 115, 70, 101, 101, 34, 58, 123, 34, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 52, 34, 58, 51, 48, 125, 125, 125}
	shardBlockNoCrossShard     = &blockchain.ShardBlock{}
	shardBlockCrossShard01     = &blockchain.ShardBlock{Body: blockchain.ShardBody{}}
)

var _ = func() (_ struct{}) {
	fmt.Println("This runs before init()!")
	bc.Init(&blockchain.Config{})
	bc.IsTest = true
	txPool.Init(&mempool.Config{
		PubSubManager: pb,
	})
	txPool.IsTest = true
	for i := 0; i < 255; i++ {
		crossShardPool[byte(i)] = mempool.GetCrossShardPool(byte(i))
	}
	peerID, _ = libp2p.IDB58Decode(senderID)
	json.Unmarshal(shardBlockByteNoCrossShard, shardBlockNoCrossShard)
	shardBlockCrossShard01.UnmarshalJSON(shardBlockByteCrossShard01)
	Logger.Init(common.NewBackend(nil).Logger("test", true))
	return
}()

// Just test flow of this function, other type with be test in later case
func TestNetSyncGetBlkShardByHashAndSend(t *testing.T) {
	netSync := NetSync{}
	netSync.Init(&NetSyncConfig{
		BlockChain:        bc,
		PubSubManager:     pb,
		Server:            server,
		TxMemPool:         txPool,
		Consensus:         consensus,
		CrossShardPool:    crossShardPool,
		ShardToBeaconPool: shardToBeaconPool,
	})
	// type 0: shard block
	netSync.getBlkShardByHashAndSend(peerID, 0, []common.Hash{hash0, hash1, hash2}, 1)
}

// Just test flow of this function, other type with be test in later case
func TestNetSyncCreateBlkShardMsgByType(t *testing.T) {
	netSync := NetSync{}
	netSync.Init(&NetSyncConfig{
		BlockChain:        bc,
		PubSubManager:     pb,
		Server:            server,
		TxMemPool:         txPool,
		Consensus:         consensus,
		CrossShardPool:    crossShardPool,
		ShardToBeaconPool: shardToBeaconPool,
	})
	// type 0: shard block
	_, err := netSync.createBlockShardMsgByType(shardBlockNoCrossShard, 0, 1)
	if err != nil {
		t.Error("Error create shard block msg", err)
	}
	// type 1: shard block
	_, err = netSync.createBlockShardMsgByType(shardBlockNoCrossShard, 1, 1)
	if err == nil {
		t.Error("Should have no outcoin", err)
	}
	// type 1: shard block
	_, err = netSync.createBlockShardMsgByType(shardBlockCrossShard01, 1, 1)
	if err == nil {
		// no cross output coin
		t.Error("Should NOT create cross shard block ", err)
	}
	//// type 2: shard block
	_, err = netSync.createBlockShardMsgByType(shardBlockNoCrossShard, 2, 1)
	if err != nil {
		t.Error("Error create shard block msg", err)
	}
}

func TestNetSyncGetBlkBeaconByHashAndSend(t *testing.T) {
	netSync := NetSync{}
	netSync.Init(&NetSyncConfig{
		BlockChain:        bc,
		PubSubManager:     pb,
		Server:            server,
		TxMemPool:         txPool,
		Consensus:         consensus,
		CrossShardPool:    crossShardPool,
		ShardToBeaconPool: shardToBeaconPool,
	})
	netSync.getBlockBeaconByHashAndSend(peerID, []common.Hash{hash0, hash1, hash2})
}
