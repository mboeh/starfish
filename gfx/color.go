/*
   Copyright 2011-2012 starfish authors

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
package gfx

import b "github.com/mboeh/starfish/backend"

//An RGB color representation.
type Color struct {
	Red, Green, Blue, Alpha byte
}

func (me *Color) bColor() (c b.Color) {
	c.Red = me.Red
	c.Green = me.Green
	c.Blue = me.Blue
	c.Alpha = me.Alpha
	return
}
