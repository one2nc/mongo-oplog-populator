package reader

import "mongo-oplog-populator/internal/app/populator/generator"

type Reader interface {
	ReadData() generator.FakeData
}
