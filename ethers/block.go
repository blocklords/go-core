package ethers

import (
	"encoding/json"
	jsoniter "github.com/json-iterator/go"
	"os"
	"sync"
)

type (
	BlockData struct {
		BlockNumber int64 `json:"blockNumber" redis:"blockNumber"`
	}
	BlockInterface interface {
		Read() BlockData
		Write(data BlockData) error
	}

	FileBlock struct {
		mu   sync.RWMutex
		path string
	}
)

func NewFileBlock(path string) *FileBlock {
	return &FileBlock{
		mu:   sync.RWMutex{},
		path: path,
	}
}
func (f *FileBlock) Read() BlockData {
	f.mu.RLock()
	defer f.mu.RUnlock()
	bytes, err := os.ReadFile(f.path)
	if err != nil {
		panic("file not exist!")
	}
	var data BlockData
	if err = jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(bytes, &data); err != nil {
		panic("parse json config Fail!")
	}
	return data
}

func (f *FileBlock) Write(data BlockData) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	body, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}
	if err = os.WriteFile(f.path, body, 0777); err != nil {
		return err
	}
	return nil
}
