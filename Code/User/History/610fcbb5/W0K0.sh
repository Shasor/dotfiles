#!/bin/bash
/usr/bin/discord &
while hyprctl clients | grep -q "\- Discord"; do
  hyprctl dispatch movetoworkspacesilent special:magic,discord
done
