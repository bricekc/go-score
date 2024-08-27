package player

import (
	"but3/go-score/base"
)

var Controller = base.Controller[Player]{
	Repository: Repository(),
}
