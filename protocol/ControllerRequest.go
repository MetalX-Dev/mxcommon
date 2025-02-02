// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package protocol

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ControllerRequest struct {
	_tab flatbuffers.Table
}

func GetRootAsControllerRequest(buf []byte, offset flatbuffers.UOffsetT) *ControllerRequest {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ControllerRequest{}
	x.Init(buf, n+offset)
	return x
}

func FinishControllerRequestBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsControllerRequest(buf []byte, offset flatbuffers.UOffsetT) *ControllerRequest {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &ControllerRequest{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedControllerRequestBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *ControllerRequest) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ControllerRequest) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ControllerRequest) Version() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 1
}

func (rcv *ControllerRequest) MutateVersion(n byte) bool {
	return rcv._tab.MutateByteSlot(4, n)
}

func (rcv *ControllerRequest) Id() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *ControllerRequest) MutateId(n uint64) bool {
	return rcv._tab.MutateUint64Slot(6, n)
}

func (rcv *ControllerRequest) PayloadType() ControllerRequestPayload {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return ControllerRequestPayload(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *ControllerRequest) MutatePayloadType(n ControllerRequestPayload) bool {
	return rcv._tab.MutateByteSlot(8, byte(n))
}

func (rcv *ControllerRequest) Payload(obj *flatbuffers.Table) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		rcv._tab.Union(obj, o)
		return true
	}
	return false
}

func ControllerRequestStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func ControllerRequestAddVersion(builder *flatbuffers.Builder, version byte) {
	builder.PrependByteSlot(0, version, 1)
}
func ControllerRequestAddId(builder *flatbuffers.Builder, id uint64) {
	builder.PrependUint64Slot(1, id, 0)
}
func ControllerRequestAddPayloadType(builder *flatbuffers.Builder, payloadType ControllerRequestPayload) {
	builder.PrependByteSlot(2, byte(payloadType), 0)
}
func ControllerRequestAddPayload(builder *flatbuffers.Builder, payload flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(payload), 0)
}
func ControllerRequestEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
