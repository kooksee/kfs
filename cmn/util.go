package cmn

import (
	"os"
	"fmt"
	"errors"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/common"
	"github.com/json-iterator/go"
)

func EnsureDir(dir string, mode os.FileMode) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, mode)
		if err != nil {
			return fmt.Errorf("Could not create directory %v. %v", dir, err)
		}
	}
	return nil
}

func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}

func MustNotErr(errs ... error) {
	for _, err := range errs {
		if err != nil {
			panic(err.Error())
		}
	}
}

func Err(err string, params ...interface{}) error {
	return errors.New(fmt.Sprintf(err, params...))
}

func PubkeykToAddress(pubkey []byte) common.Address {
	return crypto.PubkeyToAddress(*crypto.ToECDSAPub(pubkey))
}

func HexToAddress(hash string) common.Address {
	return common.HexToAddress(hash)
}

func JsonGet(data []byte, params ... interface{}) jsoniter.Any {
	return json.Get(data, params...)
}

func shorten(str string) string {
	if len(str) <= 8 {
		return str
	}
	return str[:3] + ".." + str[len(str)-3:]
}

var bunits = [...]string{"", "Ki", "Mi", "Gi", "Ti"}

func shortenb(bytes int) string {
	i := 0
	for ; bytes > 1024 && i < 4; i++ {
		bytes /= 1024
	}
	return fmt.Sprintf("%d%sB", bytes, bunits[i])
}

func sshortenb(bytes int) string {
	if bytes == 0 {
		return "~"
	}
	sign := "+"
	if bytes < 0 {
		sign = "-"
		bytes *= -1
	}
	i := 0
	for ; bytes > 1024 && i < 4; i++ {
		bytes /= 1024
	}
	return fmt.Sprintf("%s%d%sB", sign, bytes, bunits[i])
}

func sint(x int) string {
	if x == 0 {
		return "~"
	}
	sign := "+"
	if x < 0 {
		sign = "-"
		x *= -1
	}
	return fmt.Sprintf("%s%d", sign, x)
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func ensureBuffer(b []byte, n int) []byte {
	if cap(b) < n {
		return make([]byte, n)
	}
	return b[:n]
}

func BloomHash(key []byte) uint32 {
	return Hash(key, 0xbc9f1d34)
}

func StructSortMarshal(s interface{}) []byte {
	s1, _ := json.Marshal(s)
	b := map[string]interface{}{}
	json.Unmarshal(s1, &b)
	b1, _ := json.Marshal(b)
	return b1
}
