package handler

import "shortmessage/logic"

var shortMessageLogic *logic.ShortMessageLogic

func RegLogic(logic *logic.ShortMessageLogic) {
	shortMessageLogic = logic
}
