package models

type Deployment struct {
	Name                string `json:"name,omitempty"`
	RunningReplicas     int32  `json:"runningReplicas,omitempty"`
	UnavailableReplicas int32  `json:"unavailableReplicas,omitempty"`
	DesiredReplicas     int32  `json:"desiredReplicas,omitempty"`
}