package fileflatdb

import (
	"fmt"
	"log"
	"os"

	"github.com/c3systems/go-substrate/db"
)

// FileFlatDB ...
type FileFlatDB struct {
	impl       *Impl
	cache      *Cache
	file       *File
	serializer *Serializer
}

// NewFileFlatDB ...
func NewFileFlatDB(base, file string) *FileFlatDB {
	return &FileFlatDB{
		impl:       NewImpl(),
		cache:      NewCache(base, file, nil),
		file:       NewFile(base, file, nil),
		serializer: NewSerializer(),
	}
}

// Open ...
func (f *FileFlatDB) Open() {
	f.file.AssertOpen(false)

	f.file.Open(f.file.path, false)
	f.cache.lruBranch = make(LruMap)
	f.cache.lruData = make(LruMap)
}

// Close ...
func (f *FileFlatDB) Close() {
	f.file.AssertOpen(false)

	f.file.Close()
	f.cache.lruBranch = make(LruMap)
	f.cache.lruData = make(LruMap)
}

// Drop ...
func (f *FileFlatDB) Drop() {
	f.file.AssertOpen(false)
	os.Remove(f.file.path)
}

// Empty ...
func (f *FileFlatDB) Empty() {
	f.file.AssertOpen(false)
	f.file.Open(f.file.path, true)
}

// Maintain ...
func (f *FileFlatDB) Maintain(fn *db.ProgressCB) error {
	f.file.AssertOpen(false)
	compactor := NewCompact(f.file.file)
	compactor.Maintain(fn)
	return nil
}

// Rename ...
func (f *FileFlatDB) Rename(base, file string) {
	f.file.AssertOpen(false)
	oldPath := f.file.path

	f.file.file = file
	f.file.path = fmt.Sprintf("%s/%s", base, file)
	os.Rename(oldPath, f.file.path)
}

// Size ...
func (f *FileFlatDB) Size() int {
	return int(f.file.fileSize)
}

// Del ...
func (f *FileFlatDB) Del(key []uint8) {
	log.Fatal("delete not implemented")
}

// Get ...
func (f *FileFlatDB) Get(key []uint8) []uint8 {
	f.file.AssertOpen(false)

	k := f.FindKey(f.serializer.SerializeKey(key), false)
	if k == nil {
		return nil
	}

	result := f.ReadValue(k)

	if result != nil && len(result.Value) > 0 {
		return f.serializer.DeserializeValue(result.Value)
	}

	return nil
}

// Put ...
func (f *FileFlatDB) Put(key, value []uint8) {
	f.file.AssertOpen(false)
	k := f.FindKey(f.serializer.SerializeKey(key), true)

	if k == nil {
		log.Fatal("Unable to create key")
	}

	f.WriteValue(k, f.serializer.SerializeValue(value))
}

// FindKey ...
func (f *FileFlatDB) FindKey(key *NibbleBuffer, doCreate bool) *Key {
	return f.impl.FindKey(key, doCreate, 0, 0)
}

// ReadValue ...
func (f *FileFlatDB) ReadValue(key *Key) *Value {
	keyValue := key.KeyValue
	return f.impl.ReadValue(keyValue)
}

// WriteValue ...
func (f *FileFlatDB) WriteValue(key *Key, value []byte) *Value {

	return f.impl.WriteValue(int(key.KeyAt), key.KeyValue, value)
}