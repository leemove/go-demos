package icons

import (
	"image"
)

func loadIcon(name string) image.Image

var icons = make(map[string]image.Image)

func Icon(name string) image.Image {
	icon, ok := icons[name]
	if !ok {
		icon = loadIcon(name)
		icons[name] = icon
	}
	return icon
}
