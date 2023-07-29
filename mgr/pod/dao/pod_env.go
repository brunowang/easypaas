package dao

type PodEnv struct {
	Id       *int64  `db:"id"`
	PodId    *int64  `db:"pod_id"`
	EnvKey   *string `db:"env_key"`
	EnvValue *string `db:"env_value"`
}

func (m *PodEnv) GetId() int64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *PodEnv) GetPodId() int64 {
	if m != nil && m.PodId != nil {
		return *m.PodId
	}
	return 0
}

func (m *PodEnv) GetEnvKey() string {
	if m != nil && m.EnvKey != nil {
		return *m.EnvKey
	}
	return ""
}

func (m *PodEnv) GetEnvValue() string {
	if m != nil && m.EnvValue != nil {
		return *m.EnvValue
	}
	return ""
}

func (m *PodEnv) SetId(v int64) *PodEnv {
	if m == nil {
		return m
	}
	m.Id = &v
	return m
}

func (m *PodEnv) SetPodId(v int64) *PodEnv {
	if m == nil {
		return m
	}
	m.PodId = &v
	return m
}

func (m *PodEnv) SetEnvKey(v string) *PodEnv {
	if m == nil {
		return m
	}
	m.EnvKey = &v
	return m
}

func (m *PodEnv) SetEnvValue(v string) *PodEnv {
	if m == nil {
		return m
	}
	m.EnvValue = &v
	return m
}
