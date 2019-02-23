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
	Name  string `json:"name", yaml:"name"`
	Image string `json:"image", yaml:"image"`
}

// Application defines the application container going to run
type Application struct {
	Name        string              `json:"name" yaml:"name"`
	Type        string              `json:"type" yaml:"type"`
	Description string              `json:"description" yaml:"description"`
	Labels      []keyvalue.KeyValue `json:"labels" yaml:"labels"`
	Containers  []Container         `json:"containers" yaml:"containers"`
	EnvVars     []keyvalue.KeyValue `json:"envVars" yaml:"envVars"`
	Ports       []int               `json:"ports" yaml:"ports"`
	Skip        bool                `json:"skip" yaml:"skip"`
	K8sExtra    string              `json:"k8sExtra" yaml:"k8sExtra"`
	Stacks      []string            `json:"stacks" yaml:"stacks"`
}

type Credential struct {
	Name     string `json:"name" yaml:"name"`
	Provider string `json:"provider" yaml:"provider"`
}

type WorkerGroup struct {
	Type    string `json:"type" yaml:"type"`
	Min     string `json:"min" yaml:"min"`
	Max     string `json:"max" yaml:"max"`
	LowCost bool   `json:"lowCost" yaml:"lowCost"`
}

// Stack describes the infrastructures where the applications going to run
type Stack struct {
	Name         string        `json:"name" yaml:"name"`
	Type         string        `json:"type" yaml:"type"`
	Credential   string        `json:"credential" yaml:"credential"`
	Region       string        `json:"region" yaml:"region"`
	KeyPair      bool          `json:"keyPair" yaml:"keyPair"`
	WebShell     bool          `json:"webShell" yaml:"webShell"`
	HA           bool          `json:"highlyAvailable" yaml:"highlyAvailable"`
	WorkerGroups []WorkerGroup `json:"workerGroups" yaml:"workerGroups"`
	Skip         bool          `json:"skip" yaml:"skip"`
	CfnExtra     string        `json:"cfnExtra" yaml:"cfnExtra"`
	DmExtra      string        `json:"dmExtra" yaml:"dmExtra"`
	ArmExtra     string        `json:"armExtra" yaml:"armExtra"`
}
