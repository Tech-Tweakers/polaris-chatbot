package database

import (
	"fmt"
	"polarisai/internal/domain/ecatrom"
	"sort"
)

func NewMemoryDatabase() ecatrom.Repository {
	return &memoryDatabase{
		records: make(map[string]*ecatrom.ChatPersistence),
	}
}

type memoryDatabase struct {
	records map[string]*ecatrom.ChatPersistence
}

func (m *memoryDatabase) Insert(ChatPersistence ecatrom.ChatPersistence) (*ecatrom.ChatPersistence, error) {
	m.records[fmt.Sprint(ChatPersistence.EntryID)] = &ChatPersistence
	return &ChatPersistence, nil
}

func (m *memoryDatabase) Upsert(applicationEntity ecatrom.ChatPersistence) (*ecatrom.ChatPersistence, error) {
	return m.Insert(applicationEntity)
}

func (m *memoryDatabase) List() (*[]ecatrom.ChatPersistence, error) {
	var records []ecatrom.ChatPersistence

	for _, record := range m.records {
		records = append(records, *record)
	}

	sort.Slice(records, func(i, j int) bool {
		return records[i].EntryID < records[j].EntryID
	})

	return &records, nil
}
