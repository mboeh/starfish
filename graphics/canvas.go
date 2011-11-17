/*
   Copyright 2011 gtalent2@gmail.com

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/
package graphics

/*
#cgo LDFLAGS: -lSDL -lSDL_image
#include "SDL/SDL.h"
#include "SDL/SDL_rotozoom.h"
#include "SDL/SDL_image.h"

*/
import "C"
import (
	"wombat/core/util"
)

//Used to draw and to hold data for the drawing context.
type Canvas struct {
	viewport    viewport
	pane        *C.SDL_Surface
	color       uint32
	translation util.Point
	origin      util.Point
}

func newCanvas(surface *C.SDL_Surface) (p Canvas) {
	p.pane = surface
	p.viewport.X = 0
	p.viewport.Y = 0
	p.viewport.Width = 65000
	p.viewport.Height = 65000
	return
}

//Loads the settings for this Pane onto the SDL Surface.
func (me *Canvas) load() {
	me.viewport.calcBounds()
	b := me.viewport.Bounds
	r := toSDL_Rect(b)
	C.SDL_SetClipRect(me.pane, &r)
}

//Returns the bounds of this Canvas
func (me *Canvas) GetViewport() util.Bounds {
	return me.viewport.Bounds
}

//Pushs a viewport to limit the drawing space to the given bounds within the current drawing space.
func (me *Canvas) PushViewport(x, y, width, height int) {
	me.viewport.push(util.Bounds{util.Point{x, y}, util.Size{width, height}})
	b := me.viewport.Bounds
	r := toSDL_Rect(b)
	C.SDL_SetClipRect(me.pane, &r)
	me.origin = me.translation.AddOf(me.viewport.Point)
}

//Exits the current viewport, unless there is no viewport.
func (me *Canvas) PopViewport() {
	if me.viewport.pt != 0 {
		me.viewport.pop()
		r := toSDL_Rect(me.viewport.Bounds)
		C.SDL_SetClipRect(me.pane, &r)
		me.origin = me.translation.AddOf(me.viewport.Point)
	}
}

//Sets the color that the Canvas will draw with.
func (me *Canvas) SetColor(color Color) {
	me.color = color.toUint32()
}

//Fills a rectangle at the given coordinates and size on this Canvas.
func (me *Canvas) FillRect(x, y, width, height int) {
	r := sdl_Rect(x + me.origin.X, y + me.origin.Y, width, height)
	C.SDL_FillRect(me.pane, &r, C.Uint32(me.color))
}

//Draws the text at the given coordinates.
func (me *Canvas) DrawText(text *Text, x, y int) {
	var dest C.SDL_Rect
	dest.x = C.Sint16(x + me.origin.X)
	dest.y = C.Sint16(y + me.origin.Y)
	C.SDL_BlitSurface(me.pane, &dest, text.text, nil)
}

//Draws the image at the given coordinates.
func (me *Canvas) DrawAnimation(animation *Animation, x, y int) {
	me.DrawImage(animation.GetImage(), x, y)
}

//Draws the image at the given coordinates.
func (me *Canvas) DrawImage(img *Image, x, y int) {
	var dest C.SDL_Rect
	dest.x = C.Sint16(x + me.origin.X)
	dest.y = C.Sint16(y + me.origin.Y)
	C.SDL_BlitSurface(me.pane, &dest, img.img, nil)
}
