package types

import (
	"time"
	"io/ioutil"
	"encoding/hex"
	"os"
	"io"
	"math/big"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ipfs/go-ipfs-chunker"
	"github.com/kooksee/kdb"
)

func CreateFileMeta(db *kdb.KHash, f string) (IMetadata, error) {
	fm := &FileMeta{}

	d, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, err
	}
	rootHash := crypto.Keccak256(d)

	fm.ContentHash = hex.EncodeToString(rootHash)
	fm.ID = fm.ContentHash
	fm.Status = "create"

	f1, _ := os.Open(f)
	stat, err := f1.Stat()
	if err != nil {
		return nil, err
	}

	fm.Title = stat.Name()
	fm.Size = stat.Size()
	fm.Mode = stat.Mode().String()
	fm.ModTime = stat.ModTime().Unix()
	fm.IsDir = stat.IsDir()
	fm.CreateTime = time.Now().Unix()
	fm.UpdateTime = fm.CreateTime

	r := chunk.NewRabin(f1, 1024*256)

	for i := 1; ; i++ {
		ck, err := r.NextBytes()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		fm.ChunkNum = uint64(i)

		if err := db.Set(crypto.Keccak256(append(rootHash, big.NewInt(int64(i)).Bytes()...)), ck); err != nil {
			return nil, err
		}
	}

	return fm, nil
}
