#!/bin/sh
PKGCONFIG="github.com/ctrlshiftmake/example-go-desktop-traygui/config"

LD_FLAG_MESSAGE="-X '${PKGCONFIG}.ApplicationVersion=${VERSION}'"

LDFLAGS="${LD_FLAG_MESSAGE}"