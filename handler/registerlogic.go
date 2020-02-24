package handler

import "shortmessage/logic"

var shortMessageLogic *logic.ShortMessageLogic

func RegLogic(l *logic.ShortMessageLogic) {
	shortMessageLogic = l
}
