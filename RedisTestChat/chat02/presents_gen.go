package main

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// MarshalMsg implements msgp.Marshaler
func (z V) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 3
	o = append(o, 0x93)
	o = msgp.AppendInt32(o, z.Star)
	o = msgp.AppendInt32(o, z.Left)
	o = msgp.AppendInt32(o, z.Right)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *V) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 3 {
		err = msgp.ArrayError{Wanted: 3, Got: zb0001}
		return
	}
	z.Star, bts, err = msgp.ReadInt32Bytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Star")
		return
	}
	z.Left, bts, err = msgp.ReadInt32Bytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Left")
		return
	}
	z.Right, bts, err = msgp.ReadInt32Bytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Right")
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z V) Msgsize() (s int) {
	s = 1 + msgp.Int32Size + msgp.Int32Size + msgp.Int32Size
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Values) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendArrayHeader(o, uint32(6))
	for za0001 := range z {
		// array header, size 3
		o = append(o, 0x93)
		o = msgp.AppendInt32(o, z[za0001].Star)
		o = msgp.AppendInt32(o, z[za0001].Left)
		o = msgp.AppendInt32(o, z[za0001].Right)
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Values) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != uint32(6) {
		err = msgp.ArrayError{Wanted: uint32(6), Got: zb0001}
		return
	}
	for za0001 := range z {
		var zb0002 uint32
		zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
		if err != nil {
			err = msgp.WrapError(err, za0001)
			return
		}
		if zb0002 != 3 {
			err = msgp.ArrayError{Wanted: 3, Got: zb0002}
			return
		}
		z[za0001].Star, bts, err = msgp.ReadInt32Bytes(bts)
		if err != nil {
			err = msgp.WrapError(err, za0001, "Star")
			return
		}
		z[za0001].Left, bts, err = msgp.ReadInt32Bytes(bts)
		if err != nil {
			err = msgp.WrapError(err, za0001, "Left")
			return
		}
		z[za0001].Right, bts, err = msgp.ReadInt32Bytes(bts)
		if err != nil {
			err = msgp.WrapError(err, za0001, "Right")
			return
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Values) Msgsize() (s int) {
	s = msgp.ArrayHeaderSize + (6 * (10 + msgp.Int32Size + msgp.Int32Size + msgp.Int32Size))
	return
}
