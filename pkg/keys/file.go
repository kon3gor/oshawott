package keys

import (
	"os"
	"slices"
	"strings"

	"github.com/kon3gor/oshawott/internal/core"
)

type fileKeyProvider struct {
	fileName  string
	separator string
}

func NewFileKeyProvider(fileName string, separator string) KeyProvider {
	return fileKeyProvider{fileName: fileName, separator: separator}
}

// GetKey implements KeyProvider.
func (f fileKeyProvider) GetKey(ctx core.OshawottContext, url string) (Key, error) {
	cb, err := os.ReadFile(f.fileName)
	if err != nil {
		return NoKey, nil
	}

	c := string(cb)
	keys := strings.Split(c, f.separator)
	key := keys[0]

	keys = slices.Delete(keys, 0, 1)
	os.WriteFile(f.fileName, []byte(strings.Join(keys, f.separator)), 644)
	return Key(key), nil
}
