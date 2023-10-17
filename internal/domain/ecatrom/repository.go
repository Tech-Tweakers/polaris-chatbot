package ecatrom

type RepositoryReader interface {
	List() (*[]ChatPersistence, error)
}

type RepositoryWriter interface {
	Insert(ChatPersistence ChatPersistence) (*ChatPersistence, error)
}

type Repository interface {
	RepositoryReader
	RepositoryWriter
}
