package models

type Service struct {
	Name string `json:"name,omitempty"`
	ApplicationGroup string `json:"applicationGroup,omitempty"`
	RunningPodsCount int32 `json:"runningPodsCount,omitempty"`
}