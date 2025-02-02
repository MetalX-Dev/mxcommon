// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package protocol

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type UUID struct {
	_tab flatbuffers.Struct
}

func (rcv *UUID) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *UUID) Table() flatbuffers.Table {
	return rcv._tab.Table
}

func (rcv *UUID) Lower() uint64 {
	return rcv._tab.GetUint64(rcv._tab.Pos + flatbuffers.UOffsetT(0))
}
func (rcv *UUID) MutateLower(n uint64) bool {
	return rcv._tab.MutateUint64(rcv._tab.Pos+flatbuffers.UOffsetT(0), n)
}

func (rcv *UUID) Upper() uint64 {
	return rcv._tab.GetUint64(rcv._tab.Pos + flatbuffers.UOffsetT(8))
}
func (rcv *UUID) MutateUpper(n uint64) bool {
	return rcv._tab.MutateUint64(rcv._tab.Pos+flatbuffers.UOffsetT(8), n)
}

func CreateUUID(builder *flatbuffers.Builder, lower uint64, upper uint64) flatbuffers.UOffsetT {
	builder.Prep(8, 16)
	builder.PrependUint64(upper)
	builder.PrependUint64(lower)
	return builder.Offset()
}
