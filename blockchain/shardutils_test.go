package blockchain

import (
	"reflect"
	"testing"

	"github.com/incognitochain/incognito-chain/common"
	"github.com/incognitochain/incognito-chain/incognitokey"
	"github.com/incognitochain/incognito-chain/metadata"
	"github.com/incognitochain/incognito-chain/metadata/mocks"
)

var (
	validPrivateKeys = []string{
		"112t8rq19Uu7UGbTApZzZwCAvVszAgRNAzHzr3p8Cu75jPH3h5AUtRXMKiqF3hw8NbEfeLcjtbpeUvJfw4tGj7pbqwDYngc8wB13Gf77o33f",
		"112t8rrEW3NPNgU8xzbeqE7cr4WTT8JvyaQqSZyczA5hBJVvpQMTBVqNfcCdzhvquWCHH11jHihZtgyJqbdWPhWYbmmsw5aV29WSXBEsgbVX",
		"112t8rnky34vDfgZp4MizVfNS1EttVucnwFPL4ULFwtiS4nwvEJryyWASWG5MrVv5hvFn6nAqMqcVorguTjKjiHJmjD55jCTvEjWPEUmnMP1",
	}
	validCommitteePublicKeys = []string{
		"121VhftSAygpEJZ6i9jGk4fj81FpWVTwe3wWDzRZjzdjaQXk9QtGbwNWNwjt3p8zi3p2LRug8m78TDeq4LCAiQT2shDLSrK9sSHBX4DrNgnqsRbkEazrnWapvs7F5CMTPj5kT859WHJV26Wm1P8hwHXpxLwbeMM9n2kJXznTgRJGzdBZ4iY2CTF28s7ADyknqcBJ1RBfEUT9GVeixKC3AKDAna2QqQfdcdFiJaps5PixjJznk7CcTgcYgfPcnysdUgRuygAcbDikvw35KF9jzmeTZWZtbXhbXePhyPP8MuaGwDY75hCiDn1iDEvNHBGMqKJtENq8mfkQTW9GrGu2kkDBmNsmDVannjsbxUuoHU9MT5hYftTcsvyVi4s2S73JbGDNnWD7e3cVwXF8rgYGMFNyYBm3qWB3jobBkGwTPNh5Tpb7",
		"121VhftSAygpEJZ6i9jGkCFHRkD4yhxxccAqVjQTWR9gy7skM1KcNf3uGLpX1NvojmHqs9bWwsPfvyBmer39YNBPwBHpgXg1Qku4EDhtUBZnGw2PZGMF7DMCrYa27GNS97uA9WC5z55YuCDA4WsnKfoEEuCFDNUN3iSCeUyrQ4SF5smx9CwBYX6AWAMAvNDPKf4tCuc7Wiafv9xkLKuHSFr7jaxBfg4rdaxtwXzR5eMpFDDpiXz6hQmdcee8xSXQRKceiafg9RMiuqLxDzx9tmLKvBD5TJq4G76LB3rrVmsYwMo1fY4RZLpiYn6AstAfca5EVnMeexueSAE5sam3Lsq8mq5poJfsW6KXzAbsmFPSsSjhmQ4wGhSXoKSap331gBMuuy7KtmVwQAPpwuFPo9hi7RBgrrn1ssdCdjYSwE226Ekc",
		"121VhftSAygpEJZ6i9jGkRSiFCZCAU54hEi1ZgW1insUMxeB4DgKpTZKejgVi2D7ENHC6XfRwsAcyiEaeiuis9XRU7YTMXUUGi29SzByMnXfGVRsGAb2hew9W32QMi23QDvYjoSVgUH6rSdWX9wGaPyaUV9SoyHng63Ee9zDc8AVFv1xgqbKNE7BquQzYR22j3AypirG2MmYDSUMLe2HJHBkF9Y7UphmFABNeVKhtZTXVQP78SKpfrEHigg4Gzm595EGFWLLekn6Gcs9HZb7B6gusrMfYbACsRSbCXZ6UcpaYEDx91xReAE3SDktmUHdLh2U7JhJpxgKXK4jjtjNbXwjFAbJqi1eATG8oCA2tEtaubNB9aDQMJjnK5if9KUbt92RGk4d94Ff9Gnr9CG7jVFTfem8UNUzK8KiXHvumziwaoiX",
	}
	validCommitteePublicKeyStructs = []incognitokey.CommitteePublicKey{}
	validPaymentAddresses          = []string{
		"12S42qYc9pzsfWoxPZ21sVihEHJxYfNzEp1SXNnxvr7CGYMHNWX12ZaQkzcwvTYKAnhiVsDWwSqz5jFo6xuwzXZmz7QX1TnJaWnwEyX",
		"12RrjUWjyCNPXoCChrpEVLxucs3WEw9KyFxzP3UrdRzped2UouDzBM9gNugySqt4RpmgkqL1H7xxE8PfNmDwAatnSXPUVdNomBK1yYC",
		"12S2xCenBEHuyyZQ3VVqfMUvEEwcKL1UawNEkfSX8BL8HpPwSPu3yaYptvRYfuPzr1GUsyGBtUoet5B6VT1nGMLL8xTErYgZr6uuY52",
	}
)

// TODO: @lam
// TESTCASE
// 1. RETURN 1 STAKING INSTRUCTION, NO-ERROR
//	INPUT: 3 Transaction with Valid Shard Stake Metadata (no duplicate committee publickey)
// 2. RETURN 1 STOP AUTO STAKE INSTRUCTION, NO-ERROR
//	INPUT: 3 Transaction with Valid Stop Auto Stake Metadata (no duplicate committee publickey)
// 3. COMBINE 2 cases above
func TestCreateShardInstructionsFromTransactionAndInstruction(t *testing.T) {
	type args struct {
		transactions []metadata.Transaction
		bc           *BlockChain
		shardID      byte
	}

	happyCaseBC := &BlockChain{}

	//staking happy case args start
	stakingTx1 := &mocks.Transaction{}
	stakingTx1Hash := common.BytesToHash([]byte{1})
	var stakingTx1Meta metadata.Metadata
	stakingTx1Meta = &metadata.StakingMetadata{
		MetadataBase: metadata.MetadataBase{
			metadata.ShardStakingMeta,
		},
		FunderPaymentAddress:         validPaymentAddresses[0],
		RewardReceiverPaymentAddress: validPaymentAddresses[0],
		StakingAmountShard:           1750000000000,
		AutoReStaking:                false,
		CommitteePublicKey:           validCommitteePublicKeys[0],
	}
	stakingTx1.On("GetMetadataType").Return(metadata.ShardStakingMeta)
	stakingTx1.On("GetMetadata").Once().Return(nil)
	stakingTx1.On("GetMetadata").Twice().Return(stakingTx1Meta)
	stakingTx1.On("GetMetadata").Times(3).Return(stakingTx1Meta)

	stakingTx1.On("Hash").Return(&stakingTx1Hash)

	stakingTx2 := &mocks.Transaction{}
	stakingTx2Hash := common.BytesToHash([]byte{2})
	var stakingTx2Meta metadata.Metadata
	stakingTx2Meta = &metadata.StakingMetadata{
		MetadataBase: metadata.MetadataBase{
			metadata.ShardStakingMeta,
		},
		FunderPaymentAddress:         validPaymentAddresses[1],
		RewardReceiverPaymentAddress: validPaymentAddresses[1],
		StakingAmountShard:           1750000000000,
		AutoReStaking:                false,
		CommitteePublicKey:           validCommitteePublicKeys[1],
	}
	stakingTx2.On("GetMetadataType").Return(metadata.ShardStakingMeta)
	stakingTx2.On("GetMetadata").Once().Return(nil)
	stakingTx2.On("GetMetadata").Twice().Return(stakingTx2Meta)
	stakingTx2.On("GetMetadata").Times(3).Return(stakingTx2Meta)
	stakingTx2.On("Hash").Return(&stakingTx2Hash)

	stakingTx3 := &mocks.Transaction{}
	stakingTx3Hash := common.BytesToHash([]byte{3})
	var stakingTx3Meta metadata.Metadata
	stakingTx3Meta = &metadata.StakingMetadata{
		MetadataBase: metadata.MetadataBase{
			metadata.ShardStakingMeta,
		},
		FunderPaymentAddress:         validPaymentAddresses[2],
		RewardReceiverPaymentAddress: validPaymentAddresses[2],
		StakingAmountShard:           1750000000000,
		AutoReStaking:                false,
		CommitteePublicKey:           validCommitteePublicKeys[2],
	}
	stakingTx3.On("GetMetadataType").Return(metadata.ShardStakingMeta)
	stakingTx3.On("GetMetadata").Once().Return(nil)
	stakingTx3.On("GetMetadata").Twice().Return(stakingTx3Meta)
	stakingTx3.On("GetMetadata").Times(3).Return(stakingTx3Meta)
	stakingTx3.On("Hash").Return(&stakingTx3Hash)
	//staking happy case args end

	//stop auto staking case args start
	stopStakeTx1 := &mocks.Transaction{}
	var stopStakeTx1Meta metadata.Metadata
	stopStakeTx1Meta = &metadata.StopAutoStakingMetadata{
		MetadataBase: metadata.MetadataBase{
			metadata.StopAutoStakingMeta,
		},
		CommitteePublicKey: validCommitteePublicKeys[0],
	}
	stopStakeTx1.On("GetMetadataType").Return(metadata.StopAutoStakingMeta)
	stopStakeTx1.On("GetMetadata").Once().Return(nil)
	stopStakeTx1.On("GetMetadata").Twice().Return(stopStakeTx1Meta)
	stopStakeTx1.On("GetMetadata").Times(3).Return(stopStakeTx1Meta)

	stopStakeTx2 := &mocks.Transaction{}
	var stopStakeTx2Meta metadata.Metadata
	stopStakeTx2Meta = &metadata.StopAutoStakingMetadata{
		MetadataBase: metadata.MetadataBase{
			metadata.StopAutoStakingMeta,
		},
		CommitteePublicKey: validCommitteePublicKeys[1],
	}
	stopStakeTx2.On("GetMetadataType").Return(metadata.StopAutoStakingMeta)
	stopStakeTx2.On("GetMetadata").Once().Return(nil)
	stopStakeTx2.On("GetMetadata").Twice().Return(stopStakeTx2Meta)
	stopStakeTx2.On("GetMetadata").Times(3).Return(stopStakeTx2Meta)

	stopStakeTx3 := &mocks.Transaction{}
	var stopStakeTx3Meta metadata.Metadata
	stopStakeTx3Meta = &metadata.StopAutoStakingMetadata{
		MetadataBase: metadata.MetadataBase{
			metadata.StopAutoStakingMeta,
		},
		CommitteePublicKey: validCommitteePublicKeys[2],
	}
	stopStakeTx3.On("GetMetadataType").Return(metadata.StopAutoStakingMeta)
	stopStakeTx3.On("GetMetadata").Once().Return(nil)
	stopStakeTx3.On("GetMetadata").Twice().Return(stopStakeTx3Meta)
	stopStakeTx3.On("GetMetadata").Times(3).Return(stopStakeTx3Meta)
	//stop auto staking case args end
	tests := []struct {
		name             string
		args             args
		wantInstructions [][]string
		wantErr          bool
	}{
		{
			name: "staking happy case",
			args: args{
				transactions: []metadata.Transaction{
					stakingTx1, stakingTx2, stakingTx3,
				},
				bc:      happyCaseBC,
				shardID: 0,
			},
			wantInstructions: [][]string{{"stake", "121VhftSAygpEJZ6i9jGk4fj81FpWVTwe3wWDzRZjzdjaQXk9QtGbwNWNwjt3p8zi3p2LRug8m78TDeq4LCAiQT2shDLSrK9sSHBX4DrNgnqsRbkEazrnWapvs7F5CMTPj5kT859WHJV26Wm1P8hwHXpxLwbeMM9n2kJXznTgRJGzdBZ4iY2CTF28s7ADyknqcBJ1RBfEUT9GVeixKC3AKDAna2QqQfdcdFiJaps5PixjJznk7CcTgcYgfPcnysdUgRuygAcbDikvw35KF9jzmeTZWZtbXhbXePhyPP8MuaGwDY75hCiDn1iDEvNHBGMqKJtENq8mfkQTW9GrGu2kkDBmNsmDVannjsbxUuoHU9MT5hYftTcsvyVi4s2S73JbGDNnWD7e3cVwXF8rgYGMFNyYBm3qWB3jobBkGwTPNh5Tpb7,121VhftSAygpEJZ6i9jGkCFHRkD4yhxxccAqVjQTWR9gy7skM1KcNf3uGLpX1NvojmHqs9bWwsPfvyBmer39YNBPwBHpgXg1Qku4EDhtUBZnGw2PZGMF7DMCrYa27GNS97uA9WC5z55YuCDA4WsnKfoEEuCFDNUN3iSCeUyrQ4SF5smx9CwBYX6AWAMAvNDPKf4tCuc7Wiafv9xkLKuHSFr7jaxBfg4rdaxtwXzR5eMpFDDpiXz6hQmdcee8xSXQRKceiafg9RMiuqLxDzx9tmLKvBD5TJq4G76LB3rrVmsYwMo1fY4RZLpiYn6AstAfca5EVnMeexueSAE5sam3Lsq8mq5poJfsW6KXzAbsmFPSsSjhmQ4wGhSXoKSap331gBMuuy7KtmVwQAPpwuFPo9hi7RBgrrn1ssdCdjYSwE226Ekc,121VhftSAygpEJZ6i9jGkRSiFCZCAU54hEi1ZgW1insUMxeB4DgKpTZKejgVi2D7ENHC6XfRwsAcyiEaeiuis9XRU7YTMXUUGi29SzByMnXfGVRsGAb2hew9W32QMi23QDvYjoSVgUH6rSdWX9wGaPyaUV9SoyHng63Ee9zDc8AVFv1xgqbKNE7BquQzYR22j3AypirG2MmYDSUMLe2HJHBkF9Y7UphmFABNeVKhtZTXVQP78SKpfrEHigg4Gzm595EGFWLLekn6Gcs9HZb7B6gusrMfYbACsRSbCXZ6UcpaYEDx91xReAE3SDktmUHdLh2U7JhJpxgKXK4jjtjNbXwjFAbJqi1eATG8oCA2tEtaubNB9aDQMJjnK5if9KUbt92RGk4d94Ff9Gnr9CG7jVFTfem8UNUzK8KiXHvumziwaoiX", "shard", "0000000000000000000000000000000000000000000000000000000000000000,0000000000000000000000000000000000000000000000000000000000000000,0000000000000000000000000000000000000000000000000000000000000000", "12S42qYc9pzsfWoxPZ21sVihEHJxYfNzEp1SXNnxvr7CGYMHNWX12ZaQkzcwvTYKAnhiVsDWwSqz5jFo6xuwzXZmz7QX1TnJaWnwEyX,12RrjUWjyCNPXoCChrpEVLxucs3WEw9KyFxzP3UrdRzped2UouDzBM9gNugySqt4RpmgkqL1H7xxE8PfNmDwAatnSXPUVdNomBK1yYC,12S2xCenBEHuyyZQ3VVqfMUvEEwcKL1UawNEkfSX8BL8HpPwSPu3yaYptvRYfuPzr1GUsyGBtUoet5B6VT1nGMLL8xTErYgZr6uuY52", "false,false,false"}},
			wantErr:          false,
		},
		{
			name: "stop auto stake happy case",
			args: args{
				transactions: []metadata.Transaction{
					stopStakeTx1, stopStakeTx2, stopStakeTx3,
				},
				bc:      happyCaseBC,
				shardID: 0,
			},
			wantInstructions: [][]string{{"stopautostake", "121VhftSAygpEJZ6i9jGk4fj81FpWVTwe3wWDzRZjzdjaQXk9QtGbwNWNwjt3p8zi3p2LRug8m78TDeq4LCAiQT2shDLSrK9sSHBX4DrNgnqsRbkEazrnWapvs7F5CMTPj5kT859WHJV26Wm1P8hwHXpxLwbeMM9n2kJXznTgRJGzdBZ4iY2CTF28s7ADyknqcBJ1RBfEUT9GVeixKC3AKDAna2QqQfdcdFiJaps5PixjJznk7CcTgcYgfPcnysdUgRuygAcbDikvw35KF9jzmeTZWZtbXhbXePhyPP8MuaGwDY75hCiDn1iDEvNHBGMqKJtENq8mfkQTW9GrGu2kkDBmNsmDVannjsbxUuoHU9MT5hYftTcsvyVi4s2S73JbGDNnWD7e3cVwXF8rgYGMFNyYBm3qWB3jobBkGwTPNh5Tpb7,121VhftSAygpEJZ6i9jGkCFHRkD4yhxxccAqVjQTWR9gy7skM1KcNf3uGLpX1NvojmHqs9bWwsPfvyBmer39YNBPwBHpgXg1Qku4EDhtUBZnGw2PZGMF7DMCrYa27GNS97uA9WC5z55YuCDA4WsnKfoEEuCFDNUN3iSCeUyrQ4SF5smx9CwBYX6AWAMAvNDPKf4tCuc7Wiafv9xkLKuHSFr7jaxBfg4rdaxtwXzR5eMpFDDpiXz6hQmdcee8xSXQRKceiafg9RMiuqLxDzx9tmLKvBD5TJq4G76LB3rrVmsYwMo1fY4RZLpiYn6AstAfca5EVnMeexueSAE5sam3Lsq8mq5poJfsW6KXzAbsmFPSsSjhmQ4wGhSXoKSap331gBMuuy7KtmVwQAPpwuFPo9hi7RBgrrn1ssdCdjYSwE226Ekc,121VhftSAygpEJZ6i9jGkRSiFCZCAU54hEi1ZgW1insUMxeB4DgKpTZKejgVi2D7ENHC6XfRwsAcyiEaeiuis9XRU7YTMXUUGi29SzByMnXfGVRsGAb2hew9W32QMi23QDvYjoSVgUH6rSdWX9wGaPyaUV9SoyHng63Ee9zDc8AVFv1xgqbKNE7BquQzYR22j3AypirG2MmYDSUMLe2HJHBkF9Y7UphmFABNeVKhtZTXVQP78SKpfrEHigg4Gzm595EGFWLLekn6Gcs9HZb7B6gusrMfYbACsRSbCXZ6UcpaYEDx91xReAE3SDktmUHdLh2U7JhJpxgKXK4jjtjNbXwjFAbJqi1eATG8oCA2tEtaubNB9aDQMJjnK5if9KUbt92RGk4d94Ff9Gnr9CG7jVFTfem8UNUzK8KiXHvumziwaoiX"}},
			wantErr:          false,
		},
		{
			name: "staking & stop auto stake happy case",
			args: args{
				transactions: []metadata.Transaction{
					stakingTx1, stakingTx2, stakingTx3, stopStakeTx1, stopStakeTx2, stopStakeTx3,
				},
				bc:      happyCaseBC,
				shardID: 0,
			},
			wantInstructions: [][]string{{"stake", "121VhftSAygpEJZ6i9jGk4fj81FpWVTwe3wWDzRZjzdjaQXk9QtGbwNWNwjt3p8zi3p2LRug8m78TDeq4LCAiQT2shDLSrK9sSHBX4DrNgnqsRbkEazrnWapvs7F5CMTPj5kT859WHJV26Wm1P8hwHXpxLwbeMM9n2kJXznTgRJGzdBZ4iY2CTF28s7ADyknqcBJ1RBfEUT9GVeixKC3AKDAna2QqQfdcdFiJaps5PixjJznk7CcTgcYgfPcnysdUgRuygAcbDikvw35KF9jzmeTZWZtbXhbXePhyPP8MuaGwDY75hCiDn1iDEvNHBGMqKJtENq8mfkQTW9GrGu2kkDBmNsmDVannjsbxUuoHU9MT5hYftTcsvyVi4s2S73JbGDNnWD7e3cVwXF8rgYGMFNyYBm3qWB3jobBkGwTPNh5Tpb7,121VhftSAygpEJZ6i9jGkCFHRkD4yhxxccAqVjQTWR9gy7skM1KcNf3uGLpX1NvojmHqs9bWwsPfvyBmer39YNBPwBHpgXg1Qku4EDhtUBZnGw2PZGMF7DMCrYa27GNS97uA9WC5z55YuCDA4WsnKfoEEuCFDNUN3iSCeUyrQ4SF5smx9CwBYX6AWAMAvNDPKf4tCuc7Wiafv9xkLKuHSFr7jaxBfg4rdaxtwXzR5eMpFDDpiXz6hQmdcee8xSXQRKceiafg9RMiuqLxDzx9tmLKvBD5TJq4G76LB3rrVmsYwMo1fY4RZLpiYn6AstAfca5EVnMeexueSAE5sam3Lsq8mq5poJfsW6KXzAbsmFPSsSjhmQ4wGhSXoKSap331gBMuuy7KtmVwQAPpwuFPo9hi7RBgrrn1ssdCdjYSwE226Ekc,121VhftSAygpEJZ6i9jGkRSiFCZCAU54hEi1ZgW1insUMxeB4DgKpTZKejgVi2D7ENHC6XfRwsAcyiEaeiuis9XRU7YTMXUUGi29SzByMnXfGVRsGAb2hew9W32QMi23QDvYjoSVgUH6rSdWX9wGaPyaUV9SoyHng63Ee9zDc8AVFv1xgqbKNE7BquQzYR22j3AypirG2MmYDSUMLe2HJHBkF9Y7UphmFABNeVKhtZTXVQP78SKpfrEHigg4Gzm595EGFWLLekn6Gcs9HZb7B6gusrMfYbACsRSbCXZ6UcpaYEDx91xReAE3SDktmUHdLh2U7JhJpxgKXK4jjtjNbXwjFAbJqi1eATG8oCA2tEtaubNB9aDQMJjnK5if9KUbt92RGk4d94Ff9Gnr9CG7jVFTfem8UNUzK8KiXHvumziwaoiX", "shard", "0000000000000000000000000000000000000000000000000000000000000000,0000000000000000000000000000000000000000000000000000000000000000,0000000000000000000000000000000000000000000000000000000000000000", "12S42qYc9pzsfWoxPZ21sVihEHJxYfNzEp1SXNnxvr7CGYMHNWX12ZaQkzcwvTYKAnhiVsDWwSqz5jFo6xuwzXZmz7QX1TnJaWnwEyX,12RrjUWjyCNPXoCChrpEVLxucs3WEw9KyFxzP3UrdRzped2UouDzBM9gNugySqt4RpmgkqL1H7xxE8PfNmDwAatnSXPUVdNomBK1yYC,12S2xCenBEHuyyZQ3VVqfMUvEEwcKL1UawNEkfSX8BL8HpPwSPu3yaYptvRYfuPzr1GUsyGBtUoet5B6VT1nGMLL8xTErYgZr6uuY52", "false,false,false"}, {"stopautostake", "121VhftSAygpEJZ6i9jGk4fj81FpWVTwe3wWDzRZjzdjaQXk9QtGbwNWNwjt3p8zi3p2LRug8m78TDeq4LCAiQT2shDLSrK9sSHBX4DrNgnqsRbkEazrnWapvs7F5CMTPj5kT859WHJV26Wm1P8hwHXpxLwbeMM9n2kJXznTgRJGzdBZ4iY2CTF28s7ADyknqcBJ1RBfEUT9GVeixKC3AKDAna2QqQfdcdFiJaps5PixjJznk7CcTgcYgfPcnysdUgRuygAcbDikvw35KF9jzmeTZWZtbXhbXePhyPP8MuaGwDY75hCiDn1iDEvNHBGMqKJtENq8mfkQTW9GrGu2kkDBmNsmDVannjsbxUuoHU9MT5hYftTcsvyVi4s2S73JbGDNnWD7e3cVwXF8rgYGMFNyYBm3qWB3jobBkGwTPNh5Tpb7,121VhftSAygpEJZ6i9jGkCFHRkD4yhxxccAqVjQTWR9gy7skM1KcNf3uGLpX1NvojmHqs9bWwsPfvyBmer39YNBPwBHpgXg1Qku4EDhtUBZnGw2PZGMF7DMCrYa27GNS97uA9WC5z55YuCDA4WsnKfoEEuCFDNUN3iSCeUyrQ4SF5smx9CwBYX6AWAMAvNDPKf4tCuc7Wiafv9xkLKuHSFr7jaxBfg4rdaxtwXzR5eMpFDDpiXz6hQmdcee8xSXQRKceiafg9RMiuqLxDzx9tmLKvBD5TJq4G76LB3rrVmsYwMo1fY4RZLpiYn6AstAfca5EVnMeexueSAE5sam3Lsq8mq5poJfsW6KXzAbsmFPSsSjhmQ4wGhSXoKSap331gBMuuy7KtmVwQAPpwuFPo9hi7RBgrrn1ssdCdjYSwE226Ekc,121VhftSAygpEJZ6i9jGkRSiFCZCAU54hEi1ZgW1insUMxeB4DgKpTZKejgVi2D7ENHC6XfRwsAcyiEaeiuis9XRU7YTMXUUGi29SzByMnXfGVRsGAb2hew9W32QMi23QDvYjoSVgUH6rSdWX9wGaPyaUV9SoyHng63Ee9zDc8AVFv1xgqbKNE7BquQzYR22j3AypirG2MmYDSUMLe2HJHBkF9Y7UphmFABNeVKhtZTXVQP78SKpfrEHigg4Gzm595EGFWLLekn6Gcs9HZb7B6gusrMfYbACsRSbCXZ6UcpaYEDx91xReAE3SDktmUHdLh2U7JhJpxgKXK4jjtjNbXwjFAbJqi1eATG8oCA2tEtaubNB9aDQMJjnK5if9KUbt92RGk4d94Ff9Gnr9CG7jVFTfem8UNUzK8KiXHvumziwaoiX"}},
			wantErr:          false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotInstructions, err := CreateShardInstructionsFromTransactionAndInstruction(tt.args.transactions, tt.args.bc, tt.args.shardID)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateShardInstructionsFromTransactionAndInstruction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInstructions, tt.wantInstructions) {
				t.Errorf("CreateShardInstructionsFromTransactionAndInstruction() gotInstructions = %v, want %v", gotInstructions, tt.wantInstructions)
			}
		})
	}
}
