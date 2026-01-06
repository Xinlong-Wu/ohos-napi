package js

import napi "github.com/Xinlong-Wu/ohos-napi"

type Value struct {
	Env   Env
	Value napi.Value
}

func (v Value) GetEnv() Env {
	return v.Env
}
