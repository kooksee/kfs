package config

import (
	"github.com/inconshreveable/log15"
	"os"
	"github.com/kooksee/kdb"
	"path/filepath"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/accounts"
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
	return t.GetKeyStore().SignHashWithPassphrase(t.account, "", hash)
}

func (t *Config) GetAccount() accounts.Account {
	return t.GetKeyStore().Accounts()[0]
}

func (t *Config) InitDb() {
	kdb.InitKdb(filepath.Join(t.Home, "db"))
	t.db = kdb.GetKdb()
}
