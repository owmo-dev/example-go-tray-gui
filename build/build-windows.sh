#!/bin/sh
source flags.sh

APP="${NAME}.exe"
LDFLAGS="${LDFLAGS} -H windowsgui"
BUILD_DIR="../bin/${VERSION}/"

rm -rf ${BUILD_DIR}/"${APP}"

rsrc -arch amd64 -ico ../icon/iconwin.ico -manifest "./windows/ExampleTrayGUI.exe.manifest" -o ../ExampleTrayGUI.syso

GOOS=windows GOARCH=amd64 go build -o ${BUILD_DIR}/"${APP}" -ldflags="${LDFLAGS}" ../main.go

ditto -c -k --sequesterRsrc ${BUILD_DIR}/"${APP}" ${BUILD_DIR}/${NAME}_${VERSION}_amd64.zip

rm -rf ${BUILD_DIR}/"${APP}"
rm -rf ../ExampleTrayGUI.syso
