package factory

import "github.com/jeisonborba/full-cycle-gateway/domain/repository"

type RepositoryFactory interface {
	CreateTransactionRepository() repository.TransactionRepository
}
