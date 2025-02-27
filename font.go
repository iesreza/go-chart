package chart

import (
	"sync"

	"github.com/adnsv/freetype/truetype"
	"github.com/iesreza/go-chart/roboto"
)

var (
	_defaultFontLock sync.Mutex
	_defaultFont     *truetype.Font
)

// GetDefaultFont returns the default font (Roboto-Medium).
func GetDefaultFont() (*truetype.Font, error) {
	if _defaultFont == nil {
		_defaultFontLock.Lock()
		defer _defaultFontLock.Unlock()
		if _defaultFont == nil {
			font, err := truetype.Parse(roboto.Roboto)
			if err != nil {
				return nil, err
			}
			_defaultFont = font
		}
	}
	return _defaultFont, nil
}
