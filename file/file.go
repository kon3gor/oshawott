package file

import (
	"bufio"
	"os"

	"github.com/kon3gor/oshawott"
)

type fileKeyProvider struct {
	fileName string
	sep      byte

	offset int
}

func NewFileKeyProvider(fileName string, separator byte, skip int) oshawott.KeyProvider {
	return &fileKeyProvider{fileName, separator, countByteOffset(fileName, separator, skip)}
}

func countByteOffset(fn string, sep byte, skip int) int {
	fd, err := os.Open(fn)
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	r := bufio.NewReader(fd)
	var offset int
	for i := 0; i < skip; i++ {
		key, _ := r.ReadString(sep)
		offset += len(key)
	}

	return offset
}

// GetKey implements KeyProvider.
func (fkp *fileKeyProvider) GetKey(url string) (oshawott.Key, error) {
	fd, err := os.Open(fkp.fileName)
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	r := bufio.NewReader(fd)
	r.Discard(fkp.offset)
	//todo: do not ignore error
	key, _ := r.ReadString(fkp.sep)

	fkp.offset += len(key)

	return oshawott.Key(key[:len(key)-1]), nil
}
