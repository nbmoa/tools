// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: types.proto

package generated

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// DeepCopyInto supports using TagType within kubernetes types, where deepcopy-gen is used.
func (in *TagType) DeepCopyInto(out *TagType) {
	p := proto.Clone(in).(*TagType)
	*out = *p
}

// DeepCopyInto supports using SeparatedTagType within kubernetes types, where deepcopy-gen is used.
func (in *SeparatedTagType) DeepCopyInto(out *SeparatedTagType) {
	p := proto.Clone(in).(*SeparatedTagType)
	*out = *p
}

// DeepCopyInto supports using RepeatedFieldType within kubernetes types, where deepcopy-gen is used.
func (in *RepeatedFieldType) DeepCopyInto(out *RepeatedFieldType) {
	p := proto.Clone(in).(*RepeatedFieldType)
	*out = *p
}
