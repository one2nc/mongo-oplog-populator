package reader

import "mongo-oplog-populator/internal/app/populator/types"

type Reader interface {
	ReadData() types.Attributes
}
