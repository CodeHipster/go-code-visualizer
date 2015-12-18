==================
go-code-visualizer
==================

This project aims to visualize the go code inside a location on your file system.
It will show the dependencies between packages and content of those packages
- Public Types
- Public Variables
- Public Functions

===
What does it do?
===

The go-code-visualizer will generate a gv(graphviz) file that can be visualized by a dot language visualizer.

==
How can i use it?
==

Build the go code into an executable.

Run the executable in the folder of the project you want to visualize.

It will put a .gv file in the same folder

download a dot graph visualizer (e.g. http://www.graphviz.org/Download..php)

pass the .gv file to the dot program. (e.g. > dot -Tsvgz dot-visual.gv -o go-code_dot.svgz)

![graph_image](https://github.com/ThijsOostdam/go-code-visualizer/blob/master/example/go-code_dot.png)
