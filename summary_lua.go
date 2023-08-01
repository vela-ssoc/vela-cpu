package cpu

import (
	"encoding/json"
	"github.com/vela-ssoc/vela-kit/kind"
	"github.com/vela-ssoc/vela-kit/lua"
)

func (sum *summary) String() string                         { return lua.B2S(sum.Byte()) }
func (sum *summary) Type() lua.LValueType                   { return lua.LTObject }
func (sum *summary) AssertFloat64() (float64, bool)         { return 0, false }
func (sum *summary) AssertString() (string, bool)           { return "", false }
func (sum *summary) AssertFunction() (*lua.LFunction, bool) { return nil, false }
func (sum *summary) Peek() lua.LValue                       { return sum }

func (sum *summary) Byte() []byte {
	enc := kind.NewJsonEncoder()
	enc.Tab("")
	enc.KV("arch", sum.Arch)
	enc.KV("num", sum.Num)
	enc.KV("pcnt", sum.PCnt)
	enc.KV("vendor", sum.Vendor)
	enc.KV("model_name", sum.ModelName)

	raw, _ := json.Marshal(sum.Info)
	enc.Raw("info", raw)
	enc.End("}")
	return enc.Bytes()
}

func (sum *summary) Index(L *lua.LState, key string) lua.LValue {
	switch key {

	case "total":
		return sum.Time()

	case "num":
		return lua.LInt(sum.Num)

	case "cnt":
		return lua.Slice{lua.LInt(sum.PCnt), lua.LInt(sum.Num)}

	case "model":
		return lua.S2L(sum.ModelName)

	case "vendor":
		return lua.S2L(sum.Vendor)

	case "sample":
		return sum.LoadAverage().Pct()

	case "update":
		return lua.NewFunction(func(_ *lua.LState) int {
			sum.Update()
			return 0
		})

	default:
		return lua.LNil
	}

}
