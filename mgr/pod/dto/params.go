package dto

// This is a compile-time assertion to ensure that this generated file
// is compatible with the protoc-gen-go-gframe package it is being compiled against.
import (
	"github.com/brunowang/easypaas/pbgen/pod"
	"github.com/golang/protobuf/jsonpb"
)

var jspb = jsonpb.Marshaler{OrigName: true, EmitDefaults: true}

type PodInfo struct {
	pod.PodInfo `json:",inline"`

	Err error `json:"err"`
}

func (m *PodInfo) IsValid() bool {

	if m.GetId() == 0 {
		return false
	}

	if m.GetPodNamespace() == "" {
		return false
	}

	if m.GetPodName() == "" {
		return false
	}

	if m.GetPodTeamId() == "" {
		return false
	}

	if m.GetPodCpuMax() == 0 {
		return false
	}

	if m.GetPodReplicas() == 0 {
		return false
	}

	if m.GetPodMemoryMax() == 0 {
		return false
	}

	if len(m.GetPodPort()) == 0 {
		return false
	}

	if len(m.GetPodEnv()) == 0 {
		return false
	}

	if m.GetPodPullPolicy() == "" {
		return false
	}

	if m.GetPodRestart() == "" {
		return false
	}

	if m.GetPodType() == "" {
		return false
	}

	if m.GetPodImage() == "" {
		return false
	}

	return true
}

func (m *PodInfo) Fill(pb *pod.PodInfo) {
	if pb == nil {
		return
	}
	m.PodInfo = *pb
	return
}

func (m *PodInfo) ToPb() *pod.PodInfo {
	return &m.PodInfo
}

func (m *PodInfo) ToJson() []byte {
	js, _ := jspb.MarshalToString(m)
	return []byte(js)
}

type PodPort struct {
	pod.PodPort `json:",inline"`
}

type PodEnv struct {
	pod.PodEnv `json:",inline"`
}

type PodId struct {
	pod.PodId `json:",inline"`
}

func (m *PodId) IsValid() bool {

	if m.GetId() == 0 {
		return false
	}

	return true
}

func (m *PodId) Fill(pb *pod.PodId) {
	if pb == nil {
		return
	}
	m.PodId = *pb
	return
}

type Response struct {
	pod.Response `json:",inline"`

	Err error `json:"err"`
}

func (m *Response) ToPb() *pod.Response {
	return &m.Response
}

func (m *Response) ToJson() []byte {
	js, _ := jspb.MarshalToString(m)
	return []byte(js)
}

type AllPod struct {
	pod.AllPod `json:",inline"`

	Err error `json:"err"`
}

func (m *AllPod) ToPb() *pod.AllPod {
	return &m.AllPod
}

func (m *AllPod) ToJson() []byte {
	js, _ := jspb.MarshalToString(m)
	return []byte(js)
}

type Empty struct {
	pod.Empty `json:",inline"`
}

func (m *Empty) IsValid() bool {

	return true
}

func (m *Empty) Fill(pb *pod.Empty) {
	if pb == nil {
		return
	}
	m.Empty = *pb
	return
}
