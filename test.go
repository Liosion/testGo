package main

import (
	"fmt"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

func NewKeyManager(bitSize int, mnemonic string) (*KeyManager, error) {

	if mnemonic == "" {
		entropy, err := bip39.NewEntropy(bitSize)
		if err != nil {
			return nil, err
		}
		mnemonic, err = bip39.NewMnemonic(entropy)
		if err != nil {
			return nil, err
		}
	}

	km := &KeyManager{
		mnemonic:   mnemonic,
		keys:       make(map[string]*bip32.Key, 0),
	}
	return km, nil
}

func (km *KeyManager) GetMnemonic() string {
	return km.mnemonic
}

func main() {
	mnemonic := "abandon ability able about above absent abandon ability able about above absent"
	km, err := NewKeyManager(128, mnemonic)
	fmt.Println("Awnser: %s\nError: %s\n", km, err)


}



//"btcsuite/btcd/btcec"
//"btcsuite/btcd/chaincfg"
//"btcsuite/btcd/txscript"
//"btcsuite/btcutil"
