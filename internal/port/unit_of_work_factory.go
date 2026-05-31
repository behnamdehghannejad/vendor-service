package port

import "context"

type UnitOfWorkFactor interface {
	CreateInventoryUnitOfWork(context.Context) (InventoryUnitOfWork, error)
}
