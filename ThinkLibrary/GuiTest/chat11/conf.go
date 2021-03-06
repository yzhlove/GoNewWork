package main

import "fyne.io/fyne/v2"

var (
	ListItems = []ListItem{
		{Name: "生成CSV文件"},
		{Name: "上传ZIP包"},
		{Name: "生成Gacha测试用例"},
	}
	Address = []Addr{
		{"1", "10.155.2.15", "本地服"},
		{"2", "10.155.2.16", "本地服"},
		{"3", "10.155.2.17", "内网服"},
		{"4", "10.155.2.18", "外网服"},
		{"5", "10.155.2.15", "本地服"},
		{"6", "10.155.2.16", "本地服"},
		{"7", "10.155.2.17", "内网服"},
		{"8", "10.155.2.18", "外网服"},
		{"9", "10.155.2.15", "本地服"},
		{"10", "10.155.2.16", "本地服"},
		{"11", "10.155.2.17", "内网服"},
		{"12", "10.155.2.18", "外网服"},
		{"13", "10.155.2.16", "本地服"},
		{"14", "10.155.2.17", "内网服"},
		{"15", "10.155.2.18", "外网服"},
	}
)

type Addr struct {
	Name string
	Ip   string
	Desc string
}

type ListItem struct {
	Name   string
	Widget fyne.CanvasObject
}
