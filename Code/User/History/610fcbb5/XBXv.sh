#!/bin/bash
/usr/bin/code &
while ! hyprctl clients | grep -q "- Discord"; do
  sleep 0.2
done
hyprctl dispatch movetoworkspacesilent 2,Code