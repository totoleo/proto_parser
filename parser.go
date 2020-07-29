package proto_parser

import (
	"fmt"
	"io"
)

type Field struct {
	FieldNum int32
	WireType int
	Data     []byte
}

func Decode(buff []byte) ([]Field, error) {
	next := 0
	var results = make([]Field, 0, 4)
	for next < len(buff) {
		wire, n := DecodeVariant(buff[next:])
		if n == 0 {
			return nil, io.ErrUnexpectedEOF
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		f, ok := mappers[wireType]
		if !ok {
			panic("unknown wire type")
		}
		next2, data, err := f(buff, n+next)
		if err != nil {
			return nil, err
		}
		next = next2
		results = append(results, Field{
			FieldNum: fieldNum,
			WireType: wireType,
			Data:     data,
		})
	}
	return results, nil
}

func DecodeVariant(buff []byte) (uint64, int) {
	var wire uint64
	l := len(buff)
	iNdEx := 0
	for shift := uint(0); ; shift += 7 {
		if shift >= 64 {
			return 0, 0
		}
		if iNdEx >= l {
			return 0, 0
		}
		b := buff[iNdEx]
		iNdEx++
		wire |= (uint64(b) & 0x7F) << shift
		if b < 0x80 {
			break
		}
	}
	return wire, iNdEx
}

var mappers = map[int]func(buff []byte, iNdEx int) (next int, data []byte, err error){
	0: decodeVariant,
	1: decodeFixed64,
	2: decodeLength,
	5: decodeFixed32,
}

func decodeLength(buff []byte, iNdEx int) (int, []byte, error) {
	l := len(buff)
	start := iNdEx

	var length uint64
	for shift := uint(0); ; shift += 7 {
		if shift >= 64 {
			return 0, nil, ErrIntOverflow
		}
		if iNdEx >= l {
			return 0, nil, io.ErrUnexpectedEOF
		}
		b := buff[iNdEx]
		iNdEx++
		length |= (uint64(b) & 0x7F) << shift
		if b < 0x80 {
			break
		}
	}
	valLength := int(length)
	if valLength < 0 {
		return 0, nil, ErrInvalidLength
	}
	postIndex := iNdEx + valLength
	if postIndex > l {
		return 0, nil, io.ErrUnexpectedEOF
	}
	iNdEx = postIndex
	return iNdEx, buff[start:iNdEx], nil
}

func decodeFixed32(buff []byte, iNdEx int) (next int, data []byte, err error) {
	l := len(buff)
	start := iNdEx
	if (iNdEx + 4) > l {
		return 0, nil, io.ErrUnexpectedEOF
	}
	iNdEx += 4
	return iNdEx, buff[start:iNdEx], nil
}

func decodeFixed64(buff []byte, iNdEx int) (next int, data []byte, err error) {
	l := len(buff)
	start := iNdEx
	if (iNdEx + 8) > l {
		return 0, nil, io.ErrUnexpectedEOF
	}
	iNdEx += 8
	return iNdEx, buff[start:iNdEx], nil
}

func decodeVariant(buff []byte, iNdEx int) (next int, data []byte, err error) {
	start := iNdEx
	l := len(buff)
	for shift := uint(0); ; shift += 7 {
		if shift >= 64 {
			return 0, nil, ErrIntOverflow
		}
		if iNdEx >= l {
			return 0, nil, io.ErrUnexpectedEOF
		}
		b := buff[iNdEx]
		iNdEx++
		if b < 0x80 {
			break
		}
	}
	return iNdEx, buff[start:iNdEx], nil
}

var (
	ErrInvalidLength = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflow   = fmt.Errorf("proto: integer overflow")
)
