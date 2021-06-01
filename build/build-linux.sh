#!/bin/bash
source flags.sh

APP=${NAME_LOWER}
APPDIR=../bin/${VERSION}/${APP}

mkdir -p $APPDIR/usr/bin
mkdir -p $APPDIR/usr/share/applications
mkdir -p $APPDIR/usr/share/icons/hicolor/1024x1024/apps
mkdir -p $APPDIR/usr/share/icons/hicolor/256x256/apps
mkdir -p $APPDIR/DEBIAN

CC="gcc" CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o $APPDIR/usr/bin/$APP -ldflags="${LDFLAGS}" ../main.go

cp ../icon/icon.png $APPDIR/usr/share/icons/hicolor/1024x1024/apps/${APP}.png
cp ../icon/icon.png $APPDIR/usr/share/icons/hicolor/256x256/apps/${APP}.png

cat > $APPDIR/usr/share/applications/${APP}.desktop << EOF
[Desktop Entry]
Version=${VERSION}
Type=Application
Name=$APP
Exec=$APP
Icon=$APP
Terminal=false
StartupWMClass=ExampleTrayGUI
EOF

cat > $APPDIR/DEBIAN/control << EOF
Package: ${APP}
Version: ${VERSION}
Section: base
Priority: optional
Architecture: amd64
Maintainer: Grant Moore <grantmoore3d@gmail.com>
Description: Example Tray GUI Application
EOF

dpkg-deb --build $APPDIR
mv ${APPDIR}.deb ../bin/${VERSION}/${NAME}_${VERSION}_linux_amd64.deb
rm -rf ${APPDIR}