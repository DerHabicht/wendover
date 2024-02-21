package repositories

type RepositoryObject interface {
}

type Repository interface {
	Create(object RepositoryObject)
}
