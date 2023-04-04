package domain

import (
	"time"
)

// enum
type ClusterStatus int32

const (
	ClusterStatus_PENDING ClusterStatus = iota
	ClusterStatus_INSTALLING
	ClusterStatus_RUNNING
	ClusterStatus_DELETING
	ClusterStatus_DELETED
	ClusterStatus_ERROR
)

var clusterStatus = [...]string{
	"PENDING",
	"INSTALLING",
	"RUNNING",
	"DELETING",
	"DELETED",
	"ERROR",
}

func (m ClusterStatus) String() string { return clusterStatus[(m)] }
func (m ClusterStatus) FromString(s string) ClusterStatus {
	for _, v := range clusterStatus {
		if v == s {
			return ClusterStatus_ERROR
		}
	}
	return ClusterStatus_ERROR
}

// model
type Cluster = struct {
	ID             string      `json:"id"`
	OrganizationId string      `json:"organizationId"`
	Name           string      `json:"name"`
	Description    string      `json:"description"`
	Status         string      `json:"status"`
	StatusDesc     string      `json:"statusDesc"`
	Conf           ClusterConf `json:"conf"`
	Creator        string      `json:"creator"`
	CreatedAt      time.Time   `json:"createdAt"`
	UpdatedAt      time.Time   `json:"updatedAt"`
}

type ClusterConf = struct {
	SshKeyName      string `json:"sshKeyName"`
	Region          string `json:"region"`
	MachineType     string `json:"machineType"`
	NumOfAz         int    `json:"numOfAz"`
	MinSizePerAz    int    `json:"minSizePerAz"`
	MaxSizePerAz    int    `json:"maxSizePerAz"`
	MachineReplicas int    `json:"machineReplicas"`
}

type ClusterCapacity = struct {
	Max     int
	Current int
}

type ClusterKubeInfo = struct {
	Version        string          `json:"version"`
	TotalResources int             `json:"totalResources"`
	Nodes          int             `json:"nodes"`
	Namespaces     int             `json:"namespaces"`
	Services       int             `json:"services"`
	Pods           int             `json:"pods"`
	Cores          ClusterCapacity `json:"cores"`
	Memory         ClusterCapacity `json:"memory"`
	Updated        time.Time       `json:"updated"`
}

type Event = struct {
	ID        string    `json:"id"`
	Namespace string    `json:"namespace"`
	Type      string    `json:"type"`
	Reason    string    `json:"reason"`
	Message   string    `json:"message"`
	Updated   time.Time `json:"updated"`
}

type Node = struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Status       string    `json:"status"`
	InstanceType string    `json:"instanceType"`
	Role         string    `json:"role"`
	Updated      time.Time `json:"updated"`
}

type CreateClusterRequest struct {
	OrganizationId  string `json:"organizationId"`
	TemplateId      string `json:"templateId"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	NumberOfAz      int    `json:"numberOfAz"`
	MachineType     string `json:"machineType"`
	Region          string `json:"region"`
	MachineReplicas int    `json:"machineReplicas"`
	Creator         string `json:"creator"`
}
