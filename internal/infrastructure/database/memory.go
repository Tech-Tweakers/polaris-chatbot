package database

import (
	"fmt"
	"polarisai/internal/domain/polaris"
	"sort"
)

func NewMemoryDatabase() polaris.Repository {
	return &memoryDatabase{
		records: make(map[string]*polaris.ChatPersistence),
	}
}

type memoryDatabase struct {
	records map[string]*polaris.ChatPersistence
}

func (m *memoryDatabase) Insert(ChatPersistence polaris.ChatPersistence) (*polaris.ChatPersistence, error) {
	m.records[fmt.Sprint(ChatPersistence.EntryID)] = &ChatPersistence
	return &ChatPersistence, nil
}

func (m *memoryDatabase) Upsert(applicationEntity polaris.ChatPersistence) (*polaris.ChatPersistence, error) {
	return m.Insert(applicationEntity)
}

func (m *memoryDatabase) List() (*[]polaris.ChatPersistence, error) {
	var records []polaris.ChatPersistence

	for _, record := range m.records {
		records = append(records, *record)
	}

	sort.Slice(records, func(i, j int) bool {
		return records[i].EntryID < records[j].EntryID
	})

	return &records, nil
}
