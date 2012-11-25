Background
==========
I bought the 2x2 Rubik's cube and needed a way to experiment with it.
Actually, I wanted to find interesting sequences through brute force.
So I wrote this simple simulator in Go.

Description
===========
The logic is simple:
The cube is composed of 8 cubelets.
Each cubelet has an ID and 6 sides with a color each (CubeletColors), 3 of which are visible and the other 3 are internal.
You can rotate any the cube sides (using the rotate func) and see the resulting cube (printAllFaces).
The rotation is implemented as follows:
- get the 4 cubelets on the relevant side
- swap the cubelets to simulate a rotation
- rotate each one of the cubelets (which is performed by swapping the 6 cube sides)

Installation instructions
=========================
Install Go - I used go version go1.0.2 (appengine-1.7.2) on Windows 7
Run: dev_appserver.py C:\google_appengine\2d-cube

Wishlist
========
Implement logic to find interesting sequences
Configure sequences from the GUI
Operate the cube from the GUI
Move to Google Engine
Add an HTML5 frontend
Any other cool features...

Contribution guidelines
=======================
Any contribution is welcome as long as you keep the code clean

Enjoy!
