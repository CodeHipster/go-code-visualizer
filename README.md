go-code-visualizer
==================

This project aims to visualize the go code inside a location on your file system.
It will show the dependencies between packages and content of those packages
- Public Types
- Public Variables
- Public Functions

What does it do?
================

The go-code-visualizer will generate a gv(graphviz) file that can be visualized by a dot language visualizer.

How can i use it?
=================

`go run go-code-visualizer.go <path>`

Alternatively, you can also build the go code into an executable and place it in the folder of the project you want to visualize.

Program will produce a .gv file in the analyzed folder

Place the contents of the .gv file in this(http://mdaines.github.io/viz.js/) online visualizer.

Or

Download a dot graph visualizer (e.g. http://www.graphviz.org/Download..php)

Pass the .gv file to the dot program. (e.g. > dot -Tsvgz dot-visual.gv -o go-code_dot.svgz)

![GitHub Logo](https://github.com/ThijsOostdam/go-code-visualizer/blob/master/example/go-code_dot.png)

