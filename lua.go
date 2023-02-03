package cpu

import (
	"github.com/vela-ssoc/vela-kit/vela"
)

var (
	xEnv vela.Environment
)

func WithEnv(env vela.Environment) {
	xEnv = env
	sum := New()
	sum.Update()
	//mt := lua.NewUserKV()
	//mt.Set("cnt", lua.NewFunction(LookupCpuCntL))
	//mt.Set("info", lua.NewFunction(LookupCpuInfoL))
	xEnv.Set("cpu", sum)
}
