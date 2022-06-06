package dataMapper

import (
	"testing"

	cmp "github.com/google/go-cmp/cmp"
)

func TestDeployKeyMapping(t *testing.T) {
	src := map[string]interface{}{
		"id":         100,
		"key":        "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQConScVc7ouWWgwcjneNnJ4PScDkkwEjuDL5leLIUU5aIg13dH55/f4aqKUSvfcLUOKJ0a8073tFqMbR9rfvLAhLGeStKxmYApJXpzVkphauu7kfNW8kQNi1fI4kmHyOpQ+dKtoonzjZAT4L9AV3FlVTOfRq3U8wJ2RPwU+4EtOpMKUF+wcoDJ5ONlKBOW6uAeBt/guBiu6r3awDClDGRo4Q2YCmMceiAyoiuXcr2mFNSyzTqU1f20fftFwucV/VqnxlJjZvZ/zhlfB+v+UgQN11pJJ5vChZ7bzyRtIRRsjxbTReyWxqVZ5hEle5sm1oAR97abW9zTWfwIABgClKo+z",
		"url":        "https://api.github.com/repos/Codertocat/Hello-World/keys/100",
		"title":      "hey-its-a-deploy-key",
		"verified":   true,
		"created_at": "2019-04-02T17:37:07Z",
		"read_only":  true,
	}
	expected := map[string]interface{}{
		"createdAt":  "2019-04-02T17:37:07Z",
		"id":         100,
		"isReadOnly": true,
		"isVerified": true,
		"key":        "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQConScVc7ouWWgwcjneNnJ4PScDkkwEjuDL5leLIUU5aIg13dH55/f4aqKUSvfcLUOKJ0a8073tFqMbR9rfvLAhLGeStKxmYApJXpzVkphauu7kfNW8kQNi1fI4kmHyOpQ+dKtoonzjZAT4L9AV3FlVTOfRq3U8wJ2RPwU+4EtOpMKUF+wcoDJ5ONlKBOW6uAeBt/guBiu6r3awDClDGRo4Q2YCmMceiAyoiuXcr2mFNSyzTqU1f20fftFwucV/VqnxlJjZvZ/zhlfB+v+UgQN11pJJ5vChZ7bzyRtIRRsjxbTReyWxqVZ5hEle5sm1oAR97abW9zTWfwIABgClKo+z",
		"title":      "hey-its-a-deploy-key",
		"url":        "https://api.github.com/repos/Codertocat/Hello-World/keys/100",
	}

	dest := mapDeployKey(src)

	if diff := cmp.Diff(expected, dest); diff != "" {
		t.Errorf("Returned value did not match expected. Difference (-want +got):\n%s", diff)
	}
}
