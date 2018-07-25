package cmd

import (
	"github.com/urfave/cli"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/FactomProject/go-bip44"
	"github.com/FactomProject/go-bip39"
	"github.com/FactomProject/go-bip32"
	"github.com/ethereum/go-ethereum/crypto"
	"bufio"
	"os"
	"fmt"
	"io/ioutil"
)

func AccountCmd() cli.Command {
	return cli.Command{
		Name:    "newAccount",
		Aliases: []string{"nc"},
		Usage:   "create account",
		Flags:   []cli.Flag{},
		Action: func(c *cli.Context) error {

			entropy, err := bip39.NewEntropy(128)
			if err != nil {
				panic(err.Error())
			}

			mnemonic, err := bip39.NewMnemonic(entropy)
			if err != nil {
				panic(err.Error())
			}

			seed := bip39.NewSeed(mnemonic, "")

			masterKey, err := bip32.NewMasterKey(seed)
			if err != nil {
				panic(err)
			}

			fKey, err := bip44.NewKeyFromMasterKey(masterKey, bip44.TypeEther, bip32.FirstHardenedChild, 0, 0)
			if err != nil {
				panic(err)
			}

			p1, err := crypto.ToECDSA(fKey.Key)
			if err != nil {
				panic(err.Error())
			}

			fmt.Print("please enter your passwd: ")
			passwd, err := bufio.NewReader(os.Stdin).ReadString('\n')
			if err != nil {
				panic(fmt.Sprintf("The input was: %s\n", err.Error()))
			}

			ks := keystore.NewKeyStore("", keystore.LightScryptN, keystore.LightScryptP)
			if aa, err := ks.ImportECDSA(p1, passwd); err != nil {
				panic(err.Error())
			} else {
				p := fmt.Sprintf("%s.mnemonic", aa.Address.String())
				if err := ioutil.WriteFile(p, []byte(mnemonic), 0755); err != nil {
					panic(err.Error())
				}
			}
			return nil
		},
	}
}
