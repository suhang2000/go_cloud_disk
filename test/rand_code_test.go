package test

import (
	"go_cloud_disk/core/helper"
	"testing"
)

func TestRandCode(t *testing.T) {
	for i := 0; i < 5; i++ {
		println(helper.RandCode())
	}
}
