package internal

import "sync"

type Index struct {
	shards   map[string]*IndexShard
	shardNum int64
	lock     sync.RWMutex
}

// CreateDocument 在索引之上新建一个文档
func (index *Index) CreateDocument(docID string, doc map[string]interface{}, update bool) error {
	return nil
}

func (index *Index) GetShardByDocID(docID string) *IndexShard {
	return index.shards[docID]
}
