package models

type Node struct {
	Name string `json:"name,omitempty"`
	MemoryPressure string `json:"memoryPressure,omitempty"`
	DiskPressure string `json:"diskPressure,omitempty"`
	PIDPressure string `json:"pidPressure,omitempty"`
	Ready string `json:"ready,omitempty"`
}