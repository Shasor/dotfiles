#!/bin/bash
/usr/bin/discord &
while ! hyprctl clients | grep -q "- Discord"; do
  sleep 0.2
done
hyprctl dispatch movetoworkspacesilent special:magic,Discord