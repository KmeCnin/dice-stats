#! /bin/bash

# List of supported OS
OSLIST="linux windows darwin"
# List of supported architectures
ARCHLIST="386 amd64"

# Clean old versions
rm -fR ./misc/build
rm -fR ./misc/app
mkdir ./misc/app

for os in $OSLIST; do
    for arch in $ARCHLIST; do
        echo "Package for os $os and arch $arch"
        # Build
        echo "Building..."
        BUILDPATH="./misc/build/dice-stats-$os-$arch"
        env GOOS=$os GOARCH=$arch go build -o $BUILDPATH/dice-stats-$os-$arch
        # Package
        echo "Exporting..."
        # Do some OS specific stuff
        if [ $os = "linux" ]; then
            # Create a .desktop file on Linux
            cat << EOF > $BUILDPATH/DiceStats.desktop
[Desktop Entry]
Version=1.0
Name=Dice Stats
Comment=Dice Probability Distribution Software
Exec=bash -c 'cd "\$(dirname %k)" && ./dice-stats-$os-$arch;\$SHELL'
Icon=utilities-terminal
Terminal=true 
Type=Application
Categories=Application;
EOF
            chmod +x $BUILDPATH/DiceStats.desktop
        fi
        if [ $os = "darwin" ]; then
            # Rename to .command on OSX
            mv $BUILDPATH/dice-stats-$os-$arch $BUILDPATH/dice-stats-$os-$arch.command
        fi
        if [ $os = "windows" ]; then
            # Rename to .exe on Windows
            mv $BUILDPATH/dice-stats-$os-$arch $BUILDPATH/dice-stats-$os-$arch.exe
        fi
        tar czf ./misc/app/dice-stats-$os-$arch.tar.gz -C $BUILDPATH/ .
    done
done