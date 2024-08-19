#!/bin/bash
/usr/bin/discord &
while ! hyprctl clients | grep -q "\- Discord"; do
  sleep 0.3
done
sleep 5
hyprctl dispatch movetoworkspacesilent special:magic,discord