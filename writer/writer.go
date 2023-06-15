package writer

import "mongo-oplog-populator/internal/app/populator/types"

type Writer interface {
	WriteData(personnelInfo types.PersonnelInfo)
}
