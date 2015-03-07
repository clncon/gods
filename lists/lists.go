/*
Copyright (c) Emir Pasic, All rights reserved.

This library is free software; you can redistribute it and/or
modify it under the terms of the GNU Lesser General Public
License as published by the Free Software Foundation; either
version 3.0 of the License, or (at your option) any later version.

This library is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public
License along with this library. See the file LICENSE included
with this distribution for more information.
*/

package lists

import "github.com/emirpasic/gods/containers"

type Interface interface {
	Get(index int) (interface{}, bool)
	Remove(index int)
	Add(elements ...interface{})
	Contains(elements ...interface{}) bool

	containers.Interface
	// Empty() bool
	// Size() int
	// Clear()
	// Values() []interface{}
}
