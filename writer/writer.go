package writer

import "mongo-oplog-populator/internal/app/populator/generator"

type Writer interface {
	WriteData(personnelInfo generator.FakeData)
}
