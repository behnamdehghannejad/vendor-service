package port

import "context"

type UnitOfWorkFactor interface {
	CreateInventoryUnitOfWork(context.Context) (ReserveInventoryUnitOfWork, error)
}
