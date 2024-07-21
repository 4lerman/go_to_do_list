package db

import "sync"

func NewSyncMapStorage() *sync.Map {
	return &sync.Map{}
}
