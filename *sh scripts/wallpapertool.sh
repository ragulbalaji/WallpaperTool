#/bin/bash

wget -O ~/randWallpaper.jpg "https://unsplash.it/2560/1440/?random"
gsettings set org.gnome.desktop.background picture-uri ~/randWallpaper.jpg