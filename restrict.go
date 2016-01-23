package main

import (
	"github.com/yuin/gopher-lua"
)

func filter(configFile string, request string) string {
	L := lua.NewState()
	defer L.Close()
	if err := L.DoFile(configFile); err != nil {
		panic(err)
	}

	if err := L.CallByParam(lua.P{
		Fn:      L.GetGlobal("filter"),
		NRet:    1,
		Protect: true,
	}, lua.LString(request)); err != nil {
		panic(err)
	}
	ret := L.Get(-1)
	L.Pop(1)
	return ret.String()
}
