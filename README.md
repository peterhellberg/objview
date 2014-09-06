objview
=======

A simple QML interface to the go-qml wavefront parser. (Requires Qt 5)

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
