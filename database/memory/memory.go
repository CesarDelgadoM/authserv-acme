package memory

type MemoryDB struct {
	db map[string]string
}

func NewMemoryDB() *MemoryDB {

	users := map[string]string{
		"IU123": "password123",
		"IU456": "password456",
		"IU789": "password789",
	}

	return &MemoryDB{
		db: users,
	}
}

func (m *MemoryDB) Find(key string) (string, bool) {

	pass, exist := m.db[key]
	return pass, exist
}
