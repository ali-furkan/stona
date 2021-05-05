#!/bin/sh

set -eu

mkdir releases

Release() {
    echo "Creating '$2'"
    
    mkdir release
    cd release && mkdir bin && cd ..

    sudo cp -f ./README.md ./release/README.md
    sudo cp -f ./bin/$1 ./release/bin/$2

    cd release
    sudo zip -r ../$2.zip ./
    cd ..
    sudo rm -fr release
}

echo "📦 Preaparing Releases"

make release

# Darwin
Release darwin/stona-amd64 stona_darwin_amd64

#Linux
Release linux/stona-amd64 stona_linux_amd64

Release linux/stona-arm64 stona_linux_arm64

Release linux/stona-arm stona_linux_arm

# Windows
Release win/stona-amd64.exe stona_win_amd64

Release win/stona-arm.exe stona_win_arm

# Pack all built packages & Removes bin folder
sudo zip -r ./stona_all_platforms.zip ./bin

sudo rm -fr bin

echo "✅ Successfully Created Release files"