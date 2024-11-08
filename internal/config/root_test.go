package config

import (
	"fmt"
	"testing"
)

func init() {
	InitConfig()
}

func TestInitConfig(t *testing.T) {
	fmt.Printf("GlobalConfig: %v\n", GlobalConfig)
}
