#!/bin/bash
/usr/bin/discord &
while hyprctl clients | grep -q "\- Discord"; do
  echo ##############################
  echo ### WINDOWS AND WORKSPACES ###
  echo ##############################
  hyprctl dispatch movetoworkspacesilent special:magic,discord
  echo ##############################
  echo ### WINDOWS AND WORKSPACES ###
  echo ##############################
done
