package main

type IList interface {
	Init()
	Add(value int)
	AddOnIndex(index int, value int)
	Remove()
	RemoveOnIndex(index int)
	Get(index int)
	Set(index int, value int)
	Size()
	ShowData()
}
