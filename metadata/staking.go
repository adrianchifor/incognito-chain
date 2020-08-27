package metadata

import (
	"bytes"
	"errors"

	"github.com/incognitochain/incognito-chain/common"
	"github.com/incognitochain/incognito-chain/dataaccessobject/statedb"
	"github.com/incognitochain/incognito-chain/incognitokey"
	"github.com/incognitochain/incognito-chain/wallet"
)

type StakingMetadata struct {
	MetadataBase
	FunderPaymentAddress         string
	RewardReceiverPaymentAddress string
	StakingAmountShard           uint64
	AutoReStaking                bool
	CommitteePublicKey           string
	// CommitteePublicKey PublicKeys of a candidate who join consensus, base58CheckEncode
	// CommitteePublicKey string <= encode byte <= mashal struct
}

func NewStakingMetadata(
	stakingType int,
	funderPaymentAddress string,
	rewardReceiverPaymentAddress string,
	// candidatePaymentAddress string,
	stakingAmountShard uint64,
	committeePublicKey string,
	autoReStaking bool,
) (
	*StakingMetadata,
	error,
) {
	if stakingType != ShardStakingMeta && stakingType != BeaconStakingMeta {
		return nil, errors.New("invalid staking type")
	}
	metadataBase := NewMetadataBase(stakingType)
	return &StakingMetadata{
		MetadataBase:                 *metadataBase,
		FunderPaymentAddress:         funderPaymentAddress,
		RewardReceiverPaymentAddress: rewardReceiverPaymentAddress,
		StakingAmountShard:           stakingAmountShard,
		CommitteePublicKey:           committeePublicKey,
		AutoReStaking:                autoReStaking,
	}, nil
}

/*
 */
// REVIEW: @hung
// - REDUCE_CHECK:
//		+ no need to IsInBase58ShortFormat because error is already check below by FromString
//		+ what IsInBase58ShortFormat does is the same as FromString does but for an array
//		+ should we check FunderPaymentAddress?
func (sm *StakingMetadata) ValidateMetadataByItself() bool {
	rewardReceiverPaymentAddress := sm.RewardReceiverPaymentAddress
	rewardReceiverWallet, err := wallet.Base58CheckDeserialize(rewardReceiverPaymentAddress)
	if err != nil || rewardReceiverWallet == nil {
		return false
	}
	///// REDUCE_CHECK START
	if !incognitokey.IsInBase58ShortFormat([]string{sm.CommitteePublicKey}) {
		return false
	}
	CommitteePublicKey := new(incognitokey.CommitteePublicKey)
	if err := CommitteePublicKey.FromString(sm.CommitteePublicKey); err != nil {
		return false
	}
	///// REDUCE_CHECK END

	if !CommitteePublicKey.CheckSanityData() {
		return false
	}
	//return (sm.Type == ShardStakingMeta || sm.Type == BeaconStakingMeta)
	// only stake to shard
	return sm.Type == ShardStakingMeta
}

// REVIEW: @hung
// - This func GetAllCommitteeValidatorCandidate doesn't return any err (check implementation at retrieverformetadata.go)
func (stakingMetadata StakingMetadata) ValidateTxWithBlockChain(txr Transaction, bcr BlockchainRetriever, b byte, stateDB *statedb.StateDB) (bool, error) {
	SC, SPV, BC, BPV, CBWFCR, CBWFNR, CSWFCR, CSWFNR, err := bcr.GetAllCommitteeValidatorCandidate()
	if err != nil {
		return false, err
	}
	tempStaker, err := incognitokey.CommitteeBase58KeyListToStruct([]string{stakingMetadata.CommitteePublicKey})
	if err != nil {
		return false, err
	}
	for _, committees := range SC {
		tempStaker = incognitokey.GetValidStakeStructCommitteePublicKey(committees, tempStaker)
	}
	for _, validators := range SPV {
		tempStaker = incognitokey.GetValidStakeStructCommitteePublicKey(validators, tempStaker)
	}
	tempStaker = incognitokey.GetValidStakeStructCommitteePublicKey(BC, tempStaker)
	tempStaker = incognitokey.GetValidStakeStructCommitteePublicKey(BPV, tempStaker)
	tempStaker = incognitokey.GetValidStakeStructCommitteePublicKey(CBWFCR, tempStaker)
	tempStaker = incognitokey.GetValidStakeStructCommitteePublicKey(CBWFNR, tempStaker)
	tempStaker = incognitokey.GetValidStakeStructCommitteePublicKey(CSWFCR, tempStaker)
	tempStaker = incognitokey.GetValidStakeStructCommitteePublicKey(CSWFNR, tempStaker)
	if len(tempStaker) == 0 {
		return false, errors.New("invalid Staker, This pubkey may staked already")
	}
	return true, nil
}

/*
	// Have only one receiver
	// Have only one amount corresponding to receiver
	// Receiver Is Burning Address
	//
*/
// REVIEW: @hung
// - bcr.GetStakingAmountShard()*3, 3 should be set as a constants elsewhere (same for method GetBeaconStakeAmount of stakingMetadata - below) or implement a seperate method for Beacon, reason: easy to find and change
// - only one of these 2 combinations of 'true, true' and 'false, false' is return instead of 4 possible combinations -> only return true or false and error is enough
// - stakingMetadata.Type is ShardStakingMeta or BeaconStakingMeta before other
func (stakingMetadata StakingMetadata) ValidateSanityData(
	bcr BlockchainRetriever,
	txr Transaction,
	beaconHeight uint64,
) (
	bool,
	bool,
	error,
) {
	if txr.IsPrivacy() {
		return false, false, errors.New("staking Transaction Is No Privacy Transaction")
	}
	onlyOne, pubkey, amount := txr.GetUniqueReceiver()
	if !onlyOne {
		return false, false, errors.New("staking Transaction Should Have 1 Output Amount crossponding to 1 Receiver")
	}

	// get burning address
	burningAddress := bcr.GetBurningAddress(beaconHeight)
	keyWalletBurningAdd, err := wallet.Base58CheckDeserialize(burningAddress)
	if err != nil {
		return false, false, errors.New("burning address is invalid")
	}
	if !bytes.Equal(pubkey, keyWalletBurningAdd.KeySet.PaymentAddress.Pk) {
		return false, false, errors.New("receiver Should be Burning Address")
	}

	if stakingMetadata.Type == ShardStakingMeta && amount != bcr.GetStakingAmountShard() {
		return false, false, errors.New("invalid Stake Shard Amount")
	}
	if stakingMetadata.Type == BeaconStakingMeta && amount != bcr.GetStakingAmountShard()*3 {
		return false, false, errors.New("invalid Stake Beacon Amount")
	}

	rewardReceiverPaymentAddress := stakingMetadata.RewardReceiverPaymentAddress
	rewardReceiverWallet, err := wallet.Base58CheckDeserialize(rewardReceiverPaymentAddress)
	if err != nil || rewardReceiverWallet == nil {
		return false, false, errors.New("Invalid Candidate Payment Address, Failed to Deserialized Into Key Wallet")
	}
	if len(rewardReceiverWallet.KeySet.PaymentAddress.Pk) != common.PublicKeySize {
		return false, false, errors.New("Invalid Public Key of Candidate Payment Address")
	}

	funderPaymentAddress := stakingMetadata.FunderPaymentAddress
	funderWallet, err := wallet.Base58CheckDeserialize(funderPaymentAddress)
	if err != nil || funderWallet == nil {
		return false, false, errors.New("Invalid Funder Payment Address, Failed to Deserialized Into Key Wallet")
	}

	CommitteePublicKey := new(incognitokey.CommitteePublicKey)
	err = CommitteePublicKey.FromString(stakingMetadata.CommitteePublicKey)
	if err != nil {
		return false, false, err
	}
	if !CommitteePublicKey.CheckSanityData() {
		return false, false, errors.New("Invalid Commitee Public Key of Candidate who join consensus")
	}
	return true, true, nil
}
func (stakingMetadata StakingMetadata) GetType() int {
	return stakingMetadata.Type
}

func (stakingMetadata *StakingMetadata) CalculateSize() uint64 {
	return calculateSize(stakingMetadata)
}

// REVIEW: @hung
// 3 should be set as a constants elsewhere
func (stakingMetadata StakingMetadata) GetBeaconStakeAmount() uint64 {
	return stakingMetadata.StakingAmountShard * 3
}

func (stakingMetadata StakingMetadata) GetShardStateAmount() uint64 {
	return stakingMetadata.StakingAmountShard
}
