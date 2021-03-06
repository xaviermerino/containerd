// Code generated by protoc-gen-gogo.
// source: github.com/containerd/containerd/api/types/event/container.proto
// DO NOT EDIT!

/*
	Package event is a generated protocol buffer package.

	It is generated from these files:
		github.com/containerd/containerd/api/types/event/container.proto
		github.com/containerd/containerd/api/types/event/content.proto
		github.com/containerd/containerd/api/types/event/event.proto
		github.com/containerd/containerd/api/types/event/image.proto
		github.com/containerd/containerd/api/types/event/namespace.proto
		github.com/containerd/containerd/api/types/event/runtime.proto
		github.com/containerd/containerd/api/types/event/snapshot.proto
		github.com/containerd/containerd/api/types/event/task.proto

	It has these top-level messages:
		ContainerCreate
		ContainerUpdate
		ContainerDelete
		ContentDelete
		Envelope
		ImagePut
		ImageDelete
		NamespaceCreate
		NamespaceUpdate
		NamespaceDelete
		RuntimeIO
		RuntimeMount
		RuntimeCreate
		RuntimeEvent
		RuntimeDelete
		SnapshotPrepare
		SnapshotCommit
		SnapshotRemove
		TaskCreate
		TaskStart
		TaskDelete
*/
package event

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import strings "strings"
import reflect "reflect"
import github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type ContainerCreate struct {
	ContainerID string `protobuf:"bytes,1,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	Image       string `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
	Runtime     string `protobuf:"bytes,3,opt,name=runtime,proto3" json:"runtime,omitempty"`
}

func (m *ContainerCreate) Reset()                    { *m = ContainerCreate{} }
func (*ContainerCreate) ProtoMessage()               {}
func (*ContainerCreate) Descriptor() ([]byte, []int) { return fileDescriptorContainer, []int{0} }

type ContainerUpdate struct {
	ContainerID string            `protobuf:"bytes,1,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	Image       string            `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
	Labels      map[string]string `protobuf:"bytes,3,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	RootFS      string            `protobuf:"bytes,4,opt,name=rootfs,proto3" json:"rootfs,omitempty"`
}

func (m *ContainerUpdate) Reset()                    { *m = ContainerUpdate{} }
func (*ContainerUpdate) ProtoMessage()               {}
func (*ContainerUpdate) Descriptor() ([]byte, []int) { return fileDescriptorContainer, []int{1} }

type ContainerDelete struct {
	ContainerID string `protobuf:"bytes,1,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
}

func (m *ContainerDelete) Reset()                    { *m = ContainerDelete{} }
func (*ContainerDelete) ProtoMessage()               {}
func (*ContainerDelete) Descriptor() ([]byte, []int) { return fileDescriptorContainer, []int{2} }

func init() {
	proto.RegisterType((*ContainerCreate)(nil), "containerd.v1.types.ContainerCreate")
	proto.RegisterType((*ContainerUpdate)(nil), "containerd.v1.types.ContainerUpdate")
	proto.RegisterType((*ContainerDelete)(nil), "containerd.v1.types.ContainerDelete")
}
func (m *ContainerCreate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContainerCreate) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ContainerID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintContainer(dAtA, i, uint64(len(m.ContainerID)))
		i += copy(dAtA[i:], m.ContainerID)
	}
	if len(m.Image) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintContainer(dAtA, i, uint64(len(m.Image)))
		i += copy(dAtA[i:], m.Image)
	}
	if len(m.Runtime) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintContainer(dAtA, i, uint64(len(m.Runtime)))
		i += copy(dAtA[i:], m.Runtime)
	}
	return i, nil
}

func (m *ContainerUpdate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContainerUpdate) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ContainerID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintContainer(dAtA, i, uint64(len(m.ContainerID)))
		i += copy(dAtA[i:], m.ContainerID)
	}
	if len(m.Image) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintContainer(dAtA, i, uint64(len(m.Image)))
		i += copy(dAtA[i:], m.Image)
	}
	if len(m.Labels) > 0 {
		for k, _ := range m.Labels {
			dAtA[i] = 0x1a
			i++
			v := m.Labels[k]
			mapSize := 1 + len(k) + sovContainer(uint64(len(k))) + 1 + len(v) + sovContainer(uint64(len(v)))
			i = encodeVarintContainer(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintContainer(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintContainer(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	if len(m.RootFS) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintContainer(dAtA, i, uint64(len(m.RootFS)))
		i += copy(dAtA[i:], m.RootFS)
	}
	return i, nil
}

func (m *ContainerDelete) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContainerDelete) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ContainerID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintContainer(dAtA, i, uint64(len(m.ContainerID)))
		i += copy(dAtA[i:], m.ContainerID)
	}
	return i, nil
}

func encodeFixed64Container(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Container(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintContainer(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ContainerCreate) Size() (n int) {
	var l int
	_ = l
	l = len(m.ContainerID)
	if l > 0 {
		n += 1 + l + sovContainer(uint64(l))
	}
	l = len(m.Image)
	if l > 0 {
		n += 1 + l + sovContainer(uint64(l))
	}
	l = len(m.Runtime)
	if l > 0 {
		n += 1 + l + sovContainer(uint64(l))
	}
	return n
}

func (m *ContainerUpdate) Size() (n int) {
	var l int
	_ = l
	l = len(m.ContainerID)
	if l > 0 {
		n += 1 + l + sovContainer(uint64(l))
	}
	l = len(m.Image)
	if l > 0 {
		n += 1 + l + sovContainer(uint64(l))
	}
	if len(m.Labels) > 0 {
		for k, v := range m.Labels {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovContainer(uint64(len(k))) + 1 + len(v) + sovContainer(uint64(len(v)))
			n += mapEntrySize + 1 + sovContainer(uint64(mapEntrySize))
		}
	}
	l = len(m.RootFS)
	if l > 0 {
		n += 1 + l + sovContainer(uint64(l))
	}
	return n
}

func (m *ContainerDelete) Size() (n int) {
	var l int
	_ = l
	l = len(m.ContainerID)
	if l > 0 {
		n += 1 + l + sovContainer(uint64(l))
	}
	return n
}

func sovContainer(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozContainer(x uint64) (n int) {
	return sovContainer(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ContainerCreate) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ContainerCreate{`,
		`ContainerID:` + fmt.Sprintf("%v", this.ContainerID) + `,`,
		`Image:` + fmt.Sprintf("%v", this.Image) + `,`,
		`Runtime:` + fmt.Sprintf("%v", this.Runtime) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ContainerUpdate) String() string {
	if this == nil {
		return "nil"
	}
	keysForLabels := make([]string, 0, len(this.Labels))
	for k, _ := range this.Labels {
		keysForLabels = append(keysForLabels, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForLabels)
	mapStringForLabels := "map[string]string{"
	for _, k := range keysForLabels {
		mapStringForLabels += fmt.Sprintf("%v: %v,", k, this.Labels[k])
	}
	mapStringForLabels += "}"
	s := strings.Join([]string{`&ContainerUpdate{`,
		`ContainerID:` + fmt.Sprintf("%v", this.ContainerID) + `,`,
		`Image:` + fmt.Sprintf("%v", this.Image) + `,`,
		`Labels:` + mapStringForLabels + `,`,
		`RootFS:` + fmt.Sprintf("%v", this.RootFS) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ContainerDelete) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ContainerDelete{`,
		`ContainerID:` + fmt.Sprintf("%v", this.ContainerID) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringContainer(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ContainerCreate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowContainer
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ContainerCreate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContainerCreate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContainerID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContainer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthContainer
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContainerID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Image", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContainer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthContainer
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Image = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Runtime", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContainer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthContainer
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Runtime = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipContainer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthContainer
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ContainerUpdate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowContainer
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ContainerUpdate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContainerUpdate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContainerID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContainer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthContainer
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContainerID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Image", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContainer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthContainer
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Image = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Labels", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContainer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthContainer
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var keykey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContainer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				keykey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			var stringLenmapkey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContainer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLenmapkey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLenmapkey := int(stringLenmapkey)
			if intStringLenmapkey < 0 {
				return ErrInvalidLengthContainer
			}
			postStringIndexmapkey := iNdEx + intStringLenmapkey
			if postStringIndexmapkey > l {
				return io.ErrUnexpectedEOF
			}
			mapkey := string(dAtA[iNdEx:postStringIndexmapkey])
			iNdEx = postStringIndexmapkey
			if m.Labels == nil {
				m.Labels = make(map[string]string)
			}
			if iNdEx < postIndex {
				var valuekey uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowContainer
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					valuekey |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				var stringLenmapvalue uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowContainer
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLenmapvalue |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLenmapvalue := int(stringLenmapvalue)
				if intStringLenmapvalue < 0 {
					return ErrInvalidLengthContainer
				}
				postStringIndexmapvalue := iNdEx + intStringLenmapvalue
				if postStringIndexmapvalue > l {
					return io.ErrUnexpectedEOF
				}
				mapvalue := string(dAtA[iNdEx:postStringIndexmapvalue])
				iNdEx = postStringIndexmapvalue
				m.Labels[mapkey] = mapvalue
			} else {
				var mapvalue string
				m.Labels[mapkey] = mapvalue
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RootFS", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContainer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthContainer
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RootFS = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipContainer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthContainer
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ContainerDelete) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowContainer
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ContainerDelete: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContainerDelete: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContainerID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContainer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthContainer
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContainerID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipContainer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthContainer
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipContainer(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowContainer
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowContainer
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowContainer
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthContainer
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowContainer
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipContainer(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthContainer = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowContainer   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/containerd/containerd/api/types/event/container.proto", fileDescriptorContainer)
}

var fileDescriptorContainer = []byte{
	// 322 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x91, 0xbf, 0x4e, 0xf3, 0x30,
	0x14, 0xc5, 0xeb, 0xe6, 0xfb, 0x82, 0x70, 0x90, 0x8a, 0x42, 0x07, 0xab, 0x43, 0x5a, 0x65, 0xea,
	0xe4, 0x40, 0x59, 0x80, 0x09, 0xf5, 0x0f, 0xa2, 0x12, 0x53, 0x10, 0x33, 0x4a, 0x9b, 0x4b, 0xb0,
	0x48, 0xe3, 0xc8, 0x75, 0x2a, 0x75, 0x63, 0xe0, 0xe1, 0x3a, 0x32, 0x32, 0x55, 0xd4, 0x4f, 0x82,
	0xe2, 0xb4, 0x69, 0x84, 0x98, 0x10, 0xdb, 0xbd, 0xbe, 0xbf, 0xa3, 0x73, 0x74, 0x8c, 0xaf, 0x23,
	0x26, 0x9f, 0xb3, 0x09, 0x9d, 0xf2, 0x99, 0x37, 0xe5, 0x89, 0x0c, 0x58, 0x02, 0x22, 0xac, 0x8e,
	0x41, 0xca, 0x3c, 0xb9, 0x4c, 0x61, 0xee, 0xc1, 0x02, 0x12, 0xb9, 0x3f, 0xd1, 0x54, 0x70, 0xc9,
	0xed, 0x93, 0x3d, 0x4b, 0x17, 0x67, 0x54, 0xa3, 0xad, 0x66, 0xc4, 0x23, 0xae, 0xef, 0x5e, 0x3e,
	0x15, 0xa8, 0x9b, 0xe1, 0xc6, 0x60, 0x07, 0x0f, 0x04, 0x04, 0x12, 0xec, 0x1e, 0x3e, 0x2a, 0xf5,
	0x8f, 0x2c, 0x24, 0xa8, 0x83, 0xba, 0x87, 0xfd, 0x86, 0x5a, 0xb7, 0xad, 0x12, 0x1d, 0x0f, 0x7d,
	0xab, 0x84, 0xc6, 0xa1, 0xdd, 0xc4, 0xff, 0xd9, 0x2c, 0x88, 0x80, 0xd4, 0x73, 0xd8, 0x2f, 0x16,
	0x9b, 0xe0, 0x03, 0x91, 0x25, 0x92, 0xcd, 0x80, 0x18, 0xfa, 0x7d, 0xb7, 0xba, 0x6f, 0xf5, 0x8a,
	0xef, 0x43, 0x1a, 0xfe, 0xad, 0xef, 0x2d, 0x36, 0xe3, 0x60, 0x02, 0xf1, 0x9c, 0x18, 0x1d, 0xa3,
	0x6b, 0xf5, 0x4e, 0xe9, 0x0f, 0x85, 0xd0, 0x6f, 0xfe, 0xf4, 0x4e, 0x4b, 0x46, 0x89, 0x14, 0x4b,
	0x7f, 0xab, 0xb7, 0x5d, 0x6c, 0x0a, 0xce, 0xe5, 0xd3, 0x9c, 0xfc, 0xd3, 0x69, 0xb0, 0x5a, 0xb7,
	0x4d, 0x9f, 0x73, 0x79, 0x73, 0xef, 0x6f, 0x2f, 0xad, 0x4b, 0x6c, 0x55, 0xa4, 0xf6, 0x31, 0x36,
	0x5e, 0x60, 0x59, 0xa4, 0xf7, 0xf3, 0x31, 0x0f, 0xb9, 0x08, 0xe2, 0xac, 0x0c, 0xa9, 0x97, 0xab,
	0xfa, 0x05, 0x72, 0x47, 0x95, 0x16, 0x86, 0x10, 0xc3, 0xef, 0x5a, 0xe8, 0x93, 0xd5, 0xc6, 0xa9,
	0x7d, 0x6c, 0x9c, 0xda, 0xab, 0x72, 0xd0, 0x4a, 0x39, 0xe8, 0x5d, 0x39, 0xe8, 0x53, 0x39, 0x68,
	0x62, 0xea, 0x5f, 0x3e, 0xff, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x3f, 0xf2, 0x95, 0xd3, 0x54, 0x02,
	0x00, 0x00,
}
