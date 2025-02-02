// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package protocol

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type FilePushSegmentRequest struct {
	_tab flatbuffers.Table
}

func GetRootAsFilePushSegmentRequest(buf []byte, offset flatbuffers.UOffsetT) *FilePushSegmentRequest {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &FilePushSegmentRequest{}
	x.Init(buf, n+offset)
	return x
}

func FinishFilePushSegmentRequestBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsFilePushSegmentRequest(buf []byte, offset flatbuffers.UOffsetT) *FilePushSegmentRequest {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &FilePushSegmentRequest{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedFilePushSegmentRequestBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *FilePushSegmentRequest) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *FilePushSegmentRequest) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *FilePushSegmentRequest) StartId() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *FilePushSegmentRequest) MutateStartId(n uint64) bool {
	return rcv._tab.MutateUint64Slot(4, n)
}

func (rcv *FilePushSegmentRequest) End() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *FilePushSegmentRequest) MutateEnd(n bool) bool {
	return rcv._tab.MutateBoolSlot(6, n)
}

func (rcv *FilePushSegmentRequest) SegmentId() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *FilePushSegmentRequest) MutateSegmentId(n uint64) bool {
	return rcv._tab.MutateUint64Slot(8, n)
}

func (rcv *FilePushSegmentRequest) Data(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *FilePushSegmentRequest) DataLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *FilePushSegmentRequest) DataBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *FilePushSegmentRequest) MutateData(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func FilePushSegmentRequestStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func FilePushSegmentRequestAddStartId(builder *flatbuffers.Builder, startId uint64) {
	builder.PrependUint64Slot(0, startId, 0)
}
func FilePushSegmentRequestAddEnd(builder *flatbuffers.Builder, end bool) {
	builder.PrependBoolSlot(1, end, false)
}
func FilePushSegmentRequestAddSegmentId(builder *flatbuffers.Builder, segmentId uint64) {
	builder.PrependUint64Slot(2, segmentId, 0)
}
func FilePushSegmentRequestAddData(builder *flatbuffers.Builder, data flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(data), 0)
}
func FilePushSegmentRequestStartDataVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func FilePushSegmentRequestEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
