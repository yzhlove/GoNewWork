package main

import (
	"strings"
	"testing"
)

func Test_Str(t *testing.T) {

	str := "/Users/yostar/workSpace/GoNewWork/ThinkLibrary/GuiTest/chat11/ToolBox.app/Contents/MacOS/chat11"

	if strings.Contains(str, "/ToolBox.app/Contents/MacOS") {
		ret := strings.Replace(str, "/ToolBox.app/Contents/MacOS", "", 1)
		t.Log(ret)
	}

}
