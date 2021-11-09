package main

import webserver "go_example_2/web-server"

type Set map[string]struct{}

func main() {
	// basicBlocks.Index()
	// basicslices.Basicslices()
	// methodsinterfaces.Test()

	// list := SSL.NewSingleLinkedList()
	// list.Add(3)
	// list.Add(7)
	// list.Add(22)
	// list.Add(11)
	// fmt.Println(list)

	// chanel.Chanel()
	// chanel.BufferChanel()
	// chanel.RangeAndClose()
	// selectGo.SelectEx()
	// errorsR.ErrosRecovers()
	// packages.Scans()
	webserver.WebServer()
}
