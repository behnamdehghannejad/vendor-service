package port

import (
	"context"
)

type UnitOfWorkFactor interface {
	CreateInventoryUnitOfWork(context.Context) (ReserveInventoryUnitOfWork, error)
	AcceptReserveInventoryUnitOfWork(context.Context) (AcceptInventoryUnitOfWork, error)
	RejectReserveInventoryUnitOfWork(context.Context) (RejectInventoryUnitOfWork, error)
}
