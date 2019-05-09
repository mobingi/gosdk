package v20180702

import "github.com/mobingi/gosdk/mobingi/types/keyvalue"

// Template defines the template structure v2018-07-02, this struct is work in progress
type Template struct {
	Version      string              `json:"version" yaml:"version"`
	Name         string              `json:"name" yaml:"name"`
	Description  string              `json:"description" yaml:"description"`
	Labels       []keyvalue.KeyValue `json:"labels" yaml:"labels"`
	Applications []Application       `json:"applications" yaml:"applications"`
	Credentials  []Credential        `json:"credentials" yaml:"credentials"`
	Stacks       []Stack             `json:"stacks" yaml:"stacks"`
}

type Container struct {
	Name    string              `json:"name", yaml:"name"`
	Image   string              `json:"image", yaml:"image"`
	EnvVars []keyvalue.KeyValue `json:"envVars" yaml:"envVars"`
	Ports   []int               `json:"ports" yaml:"ports"`
}

type Service struct {
	Type       string `json:"type", yaml:"type"`
	Port       int    `json:"port", yaml:"port"`
	TargetPort int    `json:"targetPort", yaml:"targetPort"`
}

// Application defines the application container going to run
type Application struct {
	Name        string              `json:"name" yaml:"name"`
	Type        string              `json:"type" yaml:"type"`
	Credential  string              `json:"credential" yaml:"credential"`
	Description string              `json:"description" yaml:"description"`
	Labels      []keyvalue.KeyValue `json:"labels" yaml:"labels"`
	Replicas    uint64              `json:"replicas" yaml:"replicas"`
	Min         uint64              `json:"min" yaml:"min"`
	Max         uint64              `json:"max" yaml:"max"`
	Containers  []Container         `json:"containers" yaml:"containers"`
	Service     Service             `json:"service" yaml:"service"`
	Skip        bool                `json:"skip" yaml:"skip"`
	K8sExtra    string              `json:"k8sExtra" yaml:"k8sExtra"`
	Stacks      []string            `json:"stacks" yaml:"stacks"`
}

type Credential struct {
	Name     string `json:"name" yaml:"name"`
	Provider string `json:"provider" yaml:"provider"`
}

type Master struct {
	Zones     []string `json:"zones" yaml:"zones"`
	NodeCount int      `json:"nodeCount" yaml:"nodeCount"`
}

type WorkerGroup struct {
	Type    string   `json:"type" yaml:"type"`
	Zones   []string `json:"zones" yaml:"zones"`
	Min     string   `json:"min" yaml:"min"`
	Max     string   `json:"max" yaml:"max"`
	LowCost bool     `json:"lowCost" yaml:"lowCost"`
}

// Stack describes the infrastructures where the applications going to run
type Stack struct {
	Name         string            `json:"name" yaml:"name"`
	Type         string            `json:"type" yaml:"type"`
	Credential   string            `json:"credential" yaml:"credential"`
	Region       string            `json:"region" yaml:"region"`
	KeyPair      bool              `json:"keyPair" yaml:"keyPair"`
	WebShell     bool              `json:"webShell" yaml:"webShell"`
	Master       Master            `json:"master" yaml:"master"`
	WorkerGroups []WorkerGroup     `json:"workerGroups" yaml:"workerGroups"`
	Skip         bool              `json:"skip" yaml:"skip"`
	CfnExtra     string            `json:"cfnExtra" yaml:"cfnExtra"`
	DmExtra      map[string]string `json:"dmExtra" yaml:"dmExtra"` // key = filename, value = contents
	ArmExtra     string            `json:"armExtra" yaml:"armExtra"`
	AliExtra     string            `json:"aliExtra" yaml:"aliExtra"`
}
