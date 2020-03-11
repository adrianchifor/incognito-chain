package statedb

import (
	"github.com/incognitochain/incognito-chain/common"
)

func StorePrivacyToken(stateDB *StateDB, tokenID common.Hash, name string, symbol string, tokenType int, mintable bool, amount uint64, info []byte, txHash common.Hash) error {
	key := GenerateTokenObjectKey(tokenID)
	_, has, err := stateDB.getTokenState(key)
	if err != nil {
		return NewStatedbError(StorePrivacyTokenError, err)
	}
	if has {
		return nil
	}
	value := NewTokenStateWithValue(tokenID, name, symbol, tokenType, mintable, amount, info, txHash)
	err = stateDB.SetStateObject(TokenObjectType, key, value)
	if err != nil {
		return NewStatedbError(StorePrivacyTokenError, err)
	}
	return nil
}

func StorePrivacyTokenTx(stateDB *StateDB, tokenID common.Hash, txHash common.Hash) error {
	keyToken := GenerateTokenObjectKey(tokenID)
	_, has, err := stateDB.getTokenState(keyToken)
	if err != nil {
		return NewStatedbError(GetPrivacyTokenError, err)
	}
	if !has {
		err := StorePrivacyToken(stateDB, tokenID, "", "", UnknownToken, false, 0, []byte{}, common.Hash{})
		if err != nil {
			return err
		}
	}
	keyTokenTx := GenerateTokenTransactionObjectKey(tokenID, txHash)
	tokenTransactionState := NewTokenTransactionStateWithValue(txHash)
	err = stateDB.SetStateObject(TokenTransactionObjectType, keyTokenTx, tokenTransactionState)
	if err != nil {
		return NewStatedbError(StorePrivacyTokenTransactionError, err)
	}
	return nil
}

func ListPrivacyToken(stateDB *StateDB) map[common.Hash]*TokenState {
	return stateDB.getAllToken()
}

func GetPrivacyTokenTxs(stateDB *StateDB, tokenID common.Hash) []common.Hash {
	txs := stateDB.getTokenTxs(tokenID)
	return txs
}

func PrivacyTokenIDExisted(stateDB *StateDB, tokenID common.Hash) bool {
	key := GenerateTokenObjectKey(tokenID)
	tokenState, has, err := stateDB.getTokenState(key)
	if err != nil {
		return false
	}
	tempTokenID := tokenState.TokenID()
	if has && !tempTokenID.IsEqual(&tokenID) {
		panic("same key wrong value")
	}
	return has
}
