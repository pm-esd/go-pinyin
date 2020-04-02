package gojieba

import (
	"github.com/pm-esd/go-pinyin/gojieba/deps/cppjieba"
	"github.com/pm-esd/go-pinyin/gojieba/deps/limonp"
	"github.com/pm-esd/go-pinyin/gojieba/dict"
)

func init() {
	dict.Init()
	limonp.Init()
	cppjieba.Init()
}
