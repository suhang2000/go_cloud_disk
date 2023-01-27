package test

import (
	"fmt"
	"go_cloud_disk/core/define"
	"os"
	"strings"
	"testing"
)

func TestEnv(t *testing.T) {
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0], pair[1])
	}
	fmt.Println(define.RedisPassword)
}
