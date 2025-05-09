package config

import (
	"os"
	"path/filepath"

	"github.com/goexl/gox"
)

type Runtime struct {
	// 是否启用默认配置
	Default *bool `default:"true"`
	// 是否显示详细信息
	Verbose bool `default:"false"`
	// 是否在出错时打印输出
	Pwe *bool `default:"true"`

	_ gox.Pointerized
}

func newRuntime(wrapper *Wrapper) *Runtime {
	return wrapper.Runtime
}

func (*Runtime) Path(required string, optionals ...string) (final string) {
	if home, uhe := os.UserHomeDir(); nil == uhe {
		finals := make([]string, 0, 1+len(optionals)+1)
		finals = append(finals, home)
		finals = append(finals, required)
		finals = append(finals, optionals...)
		final = filepath.Join(finals...)
	}

	return
}
