package cpu

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/vela-ssoc/vela-kit/lua"
	"runtime"
	"strings"
	"time"
)

type snapshot struct {
	time int64
	data Stat
}

type summary struct {
	Arch      string
	Num       int
	PCnt      int //physical
	Vendor    string
	ModelName string
	Info      []cpu.InfoStat
	Err       error
	snap      snapshot
}

func New() *summary {
	sum := &summary{Arch: runtime.GOARCH, Num: runtime.NumCPU()}
	return sum

	return sum
}

func (sum *summary) updateModeName(name string) {
	if sum.ModelName == "" {
		sum.ModelName = name
		return
	}

	if !strings.Contains(sum.ModelName, name) {
		sum.ModelName = sum.ModelName + "," + name
	}
}

func (sum *summary) updateVendor(name string) {
	if sum.Vendor == "" {
		sum.Vendor = name
		return
	}

	if !strings.Contains(sum.Vendor, name) {
		sum.Vendor = sum.Vendor + "," + name
	}
}

func (sum *summary) updateCnt() bool {
	cnt, err := cpu.Counts(false)
	if err != nil {
		sum.Err = err
		return false
	}

	sum.PCnt = cnt
	return true
}

func (sum *summary) Percent(L *lua.LState) int {
	pre := L.IsTrue(1)
	pct, err := cpu.Percent(1*time.Second, pre)
	if err != nil {
		L.RaiseError("call pct fail %v", err)
	}

	n := len(pct)
	s := lua.NewSlice(n)

	for i := 0; i < n; i++ {
		s.Set(i, lua.LNumber(pct[i]))
	}
	L.Push(s)
	return 1
}
