package idgenerator

import (
	"hash/fnv"
	"hash/maphash"
)

type IdGenerator func(bookName string) uint32

func MapHashIdGenerator(bookName string) uint32 {
	var h maphash.Hash
	h.WriteString(bookName)
	return uint32(h.Sum64())
}

func FnvIdGenerator(bookName string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(bookName))
	return h.Sum32()
}
