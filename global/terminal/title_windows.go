package terminal

import (
	"fmt"
	"github.com/Mrs4s/go-cqhttp/internal/base"
	"time"
)

// SetTitle 设置标题为 go-cqhttp `版本` `版权`
func SetTitle() {
	fmt.Sprintf("go-cqhttp "+base.Version+" © 2020 - %d Mrs4s", time.Now().Year())
}
