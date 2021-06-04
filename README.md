# ExampleTrayGUI
An example cross-platform (Mac, Windows, Linux) system tray application that can launch HTML5 windows, developed in Go including functional build process. This repository is intended as a quick reference to help others start similar projects using the referenced libraries and will not be actively maintained.

## Requirements

The build process requires a `.env` at the root of your repo file, defining the following:

```
VERSION=1.0.0
NAME=ExampleTrayGUI
```

Additionally, you'll need to install the following to build for all platforms:

```
go get github.com/akavel/rsrc
npm install --global create-dmg
brew install graphicsmagick imagemagick
```

https://www.docker.com/get-started