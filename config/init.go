package config

import (
	"github.com/inconshreveable/log15"
	"os"
	"github.com/kooksee/kdb"
	"path/filepath"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/accounts"
	"fmt"
	"bufio"
	"time"
	"quorum/crypto"
	"bytes"
)

var (
	instance *Config
)

type Config struct {
	Home     string
	IsDev    bool
	LogLevel string
	Seeds    []string
	Adds     []string

	l        log15.Logger
	db       *kdb.KDB
	keyStore *keystore.KeyStore
	account  accounts.Account
}

func (t *Config) InitLog() {
	t.l = log15.New()
	ll, err := log15.LvlFromString(t.LogLevel)
	if err != nil {
		panic(err.Error())
	}
	t.l.SetHandler(log15.LvlFilterHandler(ll, log15.StreamHandler(os.Stdout, log15.TerminalFormat())))
}

func (t *Config) InitKeyStore() {
	t.keyStore = keystore.NewKeyStore(t.Home, keystore.LightScryptN, keystore.LightScryptP)
	if len(t.keyStore.Accounts()) != 1 {
		panic("please contain one account keystore ")
	}
	t.account = t.keyStore.Accounts()[0]
}

func (t *Config) GetKeyStore() *keystore.KeyStore {
	if t.keyStore == nil {
		panic("please init keystore")
	}
	return t.keyStore
}

func (t *Config) Sign(hash []byte) ([]byte, error) {
	return t.GetKeyStore().SignHash(t.account, hash)
}

func (t *Config) Unlock() {
	fmt.Print("your passwd: ")
	passwd, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		panic(fmt.Sprintf("The input was: %s\n", err.Error()))
	}

	fmt.Print("unlock time(300ms, -1.5h or 2h45m): ")
	tm, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		panic(fmt.Sprintf("The input was: %s\n", err.Error()))
	}

	dur, err := time.ParseDuration(tm)
	if err != nil {
		panic(err.Error())
	}
	if err := t.keyStore.TimedUnlock(t.account, passwd, dur); err != nil {
		panic(err.Error())
	}
	return
}

func (t *Config) CheckSign(hash []byte, sig []byte) ([]byte, error) {
	return crypto.Ecrecover(hash, sig)
}

func (t *Config) PubkeyToAddress(pubkey []byte) []byte {
	return crypto.PubkeyToAddress(*crypto.ToECDSAPub(pubkey)).Bytes()
}

func (t *Config) CheckLocalAddress(hash []byte, sig []byte) bool {
	pubk, err := t.CheckSign(hash, sig)
	if err != nil {
		t.Log().Error("CheckSign Error", "err", err.Error())
		return false
	}

	return bytes.Equal(t.account.Address.Bytes(), t.PubkeyToAddress(pubk))
}

func (t *Config) GetAccount() accounts.Account {
	return t.GetKeyStore().Accounts()[0]
}

func (t *Config) InitDb() {
	kdb.InitKdb(filepath.Join(t.Home, "db"))
	t.db = kdb.GetKdb()
}
