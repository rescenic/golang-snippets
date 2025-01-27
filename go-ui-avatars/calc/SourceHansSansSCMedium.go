package calc

import (
	"regexp"

	"github.com/lordbasex/golang-examples/go-ui-avatars/avatarbuilderMod"
)

type SourceHansSansSCMedium struct{}

func (f *SourceHansSansSCMedium) CalculateCenterLocation(s string, ab *avatarbuilderMod.AvatarBuilder) (x int, y int) {
	cr := regexp.MustCompile("[\u4e00-\u9FA5]{1}")
	er := regexp.MustCompile("[a-zA-Z]{1}")
	nr := regexp.MustCompile("[0-9]{1}")

	cCount := len(cr.FindAllStringSubmatch(s, -1))
	eCount := len(er.FindAllStringSubmatch(s, -1))
	nCount := len(nr.FindAllStringSubmatch(s, -1))

	x = ab.W/2 - (cCount*ab.GetFontWidth()+eCount*ab.GetFontWidth()*3/5+nCount*ab.GetFontWidth()*3/5)/2
	y = ab.H/2 + ab.GetFontWidth()*4/11

	return
}
