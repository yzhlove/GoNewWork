package main

import (
	"fmt"
	"github.com/beevik/etree"
)

func main() {

	path := "/Users/yostar/workSpace/gowork/src/GoNewWork/TestChat/enums.xml"

	doc := etree.NewDocument()
	if err := doc.ReadFromFile(path); err != nil {
		panic(err)
	}

	test2(doc)

}

func test1(doc *etree.Document) {
	for _, element := range doc.Root().ChildElements() {
		fmt.Println("tag -> ", element.Tag)
		for _, attr := range element.Attr {
			fmt.Println("1.", attr.Key, "2.", attr.Value, "3.", attr.Space, "4.", attr.FullKey(), "5.", attr.NamespaceURI())
		}
	}
}

func test2(doc *etree.Document) {
	root := doc.SelectElement("config")
	if root == nil {
		panic("set root error")
	}
	for _, res := range root.ChildElements() {
		fmt.Println(res.Tag)
		for _, attr := range res.Attr {
			fmt.Println("  ", attr.Key, attr.Value)
		}
		for _, element := range res.ChildElements() {
			fmt.Println(element.Tag)
			for _, attr := range element.Attr {
				fmt.Println("    ", attr.Key, attr.Value)
			}
		}
	}
}
