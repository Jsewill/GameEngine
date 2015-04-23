## GameEngine ##
### Rapid Game Prototyping Library, Written in Go ###

Using OpenGL and GLFW, this library makes it easy to get that first window open and start working on your game loop. It even supplies a native mesh type (as if we need another one of these) in case you don't have one to work with. See examples for details.

Admittedly, this library is far from finished. This being my very first Go project, and having worked on several since, there are many things I would have done differently (e.g., using interfaces, better abstraction, etc) to meet this library's goal. My hope is that by publishing it on github it might spark some ideas or get improved/forked. There is no documentation yet, but if someone finds this library interesting I may write some--perhaps even continue development.  So, give provide some feedback!

This library utilizes the legacy GL package, simply because it was written before the updated OpenGL bindings at https://github.com/go-gl/gl were updated. This could be changed in the future.

GameEngine is released under the MIT license; see LICENSE for more details.

## Licenses and Copyright Acknowledgements ##

* Go Bindings for OpenGL (https://github.com/go-gl-legacy/gl): Copyright (c) 2012 The go-gl Authors. All rights reserved.
* Go Bindings for GLFW 3.0 (github.com/go-gl/glfw3): Copyright (c) 2012 The glfw3-go Authors. All rights reserved.
* MathGL (https://github.com/go-gl/mathgl): Copyright ©2013 The go-gl Authors. All rights reserved.
* go-xsd (https://github.com/metaleap/go-xsd): Copyright (c) 2014 Phil Schumann
