package entry

/*
#include <stdlib.h>

#include "./entry.h"

void register_module(const char *);
*/
import "C"

import (
	"unsafe"

	napi "github.com/Xinlong-Wu/ohos-napi"
)

//export InitializeModule
func InitializeModule(cEnv C.napi_env, cExports C.napi_value) C.napi_value {
	env, exports := napi.Env(cEnv), napi.Value(cExports)
	napi.InitializeInstanceData(env)

	for _, export := range napiGoGlobalExports {
		cb, _ := napi.CreateFunction(env, export.Name, export.Callback)
		name, _ := napi.CreateStringUtf8(env, export.Name)
		napi.SetProperty(env, exports, name, cb)
	}

	return cExports
}

func RegisterModule(modname string) {
	cmodname := C.CString(modname)
	defer C.free(unsafe.Pointer(cmodname))
	C.register_module(cmodname)
}
