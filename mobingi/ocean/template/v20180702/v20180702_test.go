package v20180702

import (
	"testing"

	"gopkg.in/yaml.v2"
)

func TestTemplate1(t *testing.T) {
	tmpl := `
---
version: 20180702
name: oceanGcpTest
description: test for ocean-gcp
applications:
- name: app1
  k8sExtra: |
    apiVersion: apps/v1
    kind: Deployment
    metadata:
      name: hello1
    spec:
      selector:
        matchLabels:
          app: hello1
      replicas: 1
      revisionHistoryLimit: 5
      template:
        metadata:
          labels:
            app: hello1
        spec:
          containers:
          - name: hello1
            image: hashicorp/http-echo
            args: ["-text=Hello world (1)"]
            resources:
              requests:
                cpu: 100m
                memory: 300Mi
            imagePullPolicy: Always
            ports:
            - containerPort: 5678
    ---
    apiVersion: v1
    kind: Service
    metadata:
      name: hello1
    spec:
      type: LoadBalancer
      ports:
      - port: 80
        targetPort: 5678
        protocol: TCP
      selector:
        app: hello1
  stacks:
  - cluster1
- name: app2
  k8sExtra: |
    apiVersion: apps/v1
    kind: Deployment
    metadata:
      name: hello2
    spec:
      selector:
        matchLabels:
          app: hello2
      replicas: 1
      revisionHistoryLimit: 5
      template:
        metadata:
          labels:
            app: hello2
        spec:
          containers:
          - name: hello2
            image: hashicorp/http-echo
            args: ["-text=Hello world (2)"]
            resources:
              requests:
                cpu: 100m
                memory: 300Mi
            imagePullPolicy: Always
            ports:
            - containerPort: 5678
    ---
    apiVersion: v1
    kind: Service
    metadata:
      name: hello2
    spec:
      type: NodePort
      ports:
      - protocol: TCP
        port: 80
        targetPort: 5678
      selector:
        app: hello2
  stacks:
  - cluster1

credentials:
- name: oceandevgcp
  provider: gcp

stacks:
- name: cluster1
  type: k8s
  credential: oceandevgcp
  region: asia-northeast1
  keyPair: true
  master:
    zones:
    - asia-northeast1-b
    - asia-northeast1-a
    - asia-northeast1-c
    nodeCount: 2
  workerGroups:
  - type: n1-standard-1
    min: 2
    max: 3
  - type: n1-standard-1
    min: 1
    max: 3
  dmExtra:
    sample.yaml: |
      resources:
          - type: pubsub.v1.topic
            name: oceangcptest-dmextra
            properties:
                topic: oceangcptest-dmextra-pubsubtopic
`

	var tp Template
	err := yaml.Unmarshal([]byte(tmpl), &tp)
	if err != nil {
		t.Fatal(err)
	}

	// Check for Kimura=kun.
	if len(tp.Stacks[0].Master.Zones) == 0 {
		t.Fatal("should not be zero")
	}
}
