package store

import (
	dbm "github.com/creatachain/tm-db"

	"github.com/creatachain/creata-sdk/store/cache"
	"github.com/creatachain/creata-sdk/store/rootmulti"
	"github.com/creatachain/creata-sdk/store/types"
)

func NewCommitMultiStore(db dbm.DB) types.CommitMultiStore {
	return rootmulti.NewStore(db)
}

func NewCommitKVStoreCacheManager() types.MultiStorePersistentCache {
	return cache.NewCommitKVStoreCacheManager(cache.DefaultCommitKVStoreCacheSize)
}
