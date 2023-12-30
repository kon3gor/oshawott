package oshawott

type Engine struct {
	keyProvider KeyProvider
	storage     Storage
}

func NewEngine(kp KeyProvider, s Storage) Engine {
	return Engine{kp, s}
}

func (e Engine) SaveUrl(url string) (Key, error) {
	existing, found := e.storage.IsUrlSaved(url)
	if found {
		return existing, nil
	}

	key, err := e.keyProvider.GetKey(url)
	if err != nil {
		return NoKey, err
	}

	go e.storage.SaveUrl(key, url)

	return key, nil
}

func (e Engine) GetUrl(key Key) (string, error) {
	return e.storage.GetUrl(key)
}
