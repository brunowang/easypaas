package dao

type Pod struct {
	Id            *int64   `db:"id"`
	PodName       *string  `db:"pod_name"`
	PodNamespace  *string  `db:"pod_namespace"`
	PodTeamId     *string  `db:"pod_team_id"`
	PodCpuMin     *float64 `db:"pod_cpu_min"`
	PodCpuMax     *float64 `db:"pod_cpu_max"`
	PodReplicas   *int32   `db:"pod_replicas"`
	PodMemoryMin  *float64 `db:"pod_memory_min"`
	PodMemoryMax  *float64 `db:"pod_memory_max"`
	PodPullPolicy *string  `db:"pod_pull_policy"`
	PodRestart    *string  `db:"pod_restart"`
	PodType       *string  `db:"pod_type"`
	PodImage      *string  `db:"pod_image"`
}

func (m *Pod) GetId() int64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Pod) GetPodName() string {
	if m != nil && m.PodName != nil {
		return *m.PodName
	}
	return ""
}

func (m *Pod) GetPodNamespace() string {
	if m != nil && m.PodNamespace != nil {
		return *m.PodNamespace
	}
	return ""
}

func (m *Pod) GetPodTeamId() string {
	if m != nil && m.PodTeamId != nil {
		return *m.PodTeamId
	}
	return ""
}

func (m *Pod) GetPodCpuMin() float64 {
	if m != nil && m.PodCpuMin != nil {
		return *m.PodCpuMin
	}
	return 0.0
}

func (m *Pod) GetPodCpuMax() float64 {
	if m != nil && m.PodCpuMax != nil {
		return *m.PodCpuMax
	}
	return 0.0
}

func (m *Pod) GetPodReplicas() int32 {
	if m != nil && m.PodReplicas != nil {
		return *m.PodReplicas
	}
	return 0
}

func (m *Pod) GetPodMemoryMin() float64 {
	if m != nil && m.PodMemoryMin != nil {
		return *m.PodMemoryMin
	}
	return 0.0
}

func (m *Pod) GetPodMemoryMax() float64 {
	if m != nil && m.PodMemoryMax != nil {
		return *m.PodMemoryMax
	}
	return 0.0
}

func (m *Pod) GetPodPullPolicy() string {
	if m != nil && m.PodPullPolicy != nil {
		return *m.PodPullPolicy
	}
	return ""
}

func (m *Pod) GetPodRestart() string {
	if m != nil && m.PodRestart != nil {
		return *m.PodRestart
	}
	return ""
}

func (m *Pod) GetPodType() string {
	if m != nil && m.PodType != nil {
		return *m.PodType
	}
	return ""
}

func (m *Pod) GetPodImage() string {
	if m != nil && m.PodImage != nil {
		return *m.PodImage
	}
	return ""
}

func (m *Pod) SetId(v int64) *Pod {
	if m == nil {
		return m
	}
	m.Id = &v
	return m
}

func (m *Pod) SetPodName(v string) *Pod {
	if m == nil {
		return m
	}
	m.PodName = &v
	return m
}

func (m *Pod) SetPodNamespace(v string) *Pod {
	if m == nil {
		return m
	}
	m.PodNamespace = &v
	return m
}

func (m *Pod) SetPodTeamId(v string) *Pod {
	if m == nil {
		return m
	}
	m.PodTeamId = &v
	return m
}

func (m *Pod) SetPodCpuMin(v float64) *Pod {
	if m == nil {
		return m
	}
	m.PodCpuMin = &v
	return m
}

func (m *Pod) SetPodCpuMax(v float64) *Pod {
	if m == nil {
		return m
	}
	m.PodCpuMax = &v
	return m
}

func (m *Pod) SetPodReplicas(v int32) *Pod {
	if m == nil {
		return m
	}
	m.PodReplicas = &v
	return m
}

func (m *Pod) SetPodMemoryMin(v float64) *Pod {
	if m == nil {
		return m
	}
	m.PodMemoryMin = &v
	return m
}

func (m *Pod) SetPodMemoryMax(v float64) *Pod {
	if m == nil {
		return m
	}
	m.PodMemoryMax = &v
	return m
}

func (m *Pod) SetPodPullPolicy(v string) *Pod {
	if m == nil {
		return m
	}
	m.PodPullPolicy = &v
	return m
}

func (m *Pod) SetPodRestart(v string) *Pod {
	if m == nil {
		return m
	}
	m.PodRestart = &v
	return m
}

func (m *Pod) SetPodType(v string) *Pod {
	if m == nil {
		return m
	}
	m.PodType = &v
	return m
}

func (m *Pod) SetPodImage(v string) *Pod {
	if m == nil {
		return m
	}
	m.PodImage = &v
	return m
}
