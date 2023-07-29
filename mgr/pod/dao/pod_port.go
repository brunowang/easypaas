package dao

type PodPort struct {
	Id            *int64  `db:"id"`
	PodId         *int64  `db:"pod_id"`
	ContainerPort *int32  `db:"container_port"`
	Protocol      *string `db:"protocol"`
}

func (m *PodPort) GetId() int64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *PodPort) GetPodId() int64 {
	if m != nil && m.PodId != nil {
		return *m.PodId
	}
	return 0
}

func (m *PodPort) GetContainerPort() int32 {
	if m != nil && m.ContainerPort != nil {
		return *m.ContainerPort
	}
	return 0
}

func (m *PodPort) GetProtocol() string {
	if m != nil && m.Protocol != nil {
		return *m.Protocol
	}
	return ""
}

func (m *PodPort) SetId(v int64) *PodPort {
	if m == nil {
		return m
	}
	m.Id = &v
	return m
}

func (m *PodPort) SetPodId(v int64) *PodPort {
	if m == nil {
		return m
	}
	m.PodId = &v
	return m
}

func (m *PodPort) SetContainerPort(v int32) *PodPort {
	if m == nil {
		return m
	}
	m.ContainerPort = &v
	return m
}

func (m *PodPort) SetProtocol(v string) *PodPort {
	if m == nil {
		return m
	}
	m.Protocol = &v
	return m
}
