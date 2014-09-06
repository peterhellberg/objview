objview
=======

A simple QML interface to the [go-qml](https://github.com/go-qml/qml) wavefront parser.
(Requires [Qt 5](https://qt-project.org/wiki/Qt_5.0))

![tv4 play 3d logo](http://assets.c7.se/skitch/objview_-_tv4_play_3d_logo-20140906-155743.png)

## Dependencies

On Mac OS X you'll need [Qt 5](https://qt-project.org/wiki/Qt_5.0).
It's easiest to install with [Homebrew](http://brew.sh/).

Install the qt5 and pkg-config packages:

    $ brew install qt5 pkg-config

Then, force brew to "link" qt5 (this makes it available under /usr/local):

    $ brew link --force qt5

And finally, fetch and install go-qml and wavefront:

    $ go get -u gopkg.in/qml.v1
    $ go get -u github.com/peterhellberg/wavefront

## Installation

    $ go get -u github.com/peterhellberg/objview
