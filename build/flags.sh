#!/bin/sh
PKGCONFIG="github.com/ctrlshiftmake/example-tray-gui/config"

LD_FLAG_MESSAGE="-X '${PKGCONFIG}.ApplicationVersion=${VERSION}'"

LDFLAGS="${LD_FLAG_MESSAGE}"