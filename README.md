# Sway Brightness Manager

This application is a utility for Hyprland or Sway that automatically manages the brightness of your screen. It is designed to work with `swayidle`, the idle management daemon of the Sway window manager.

It also locks the screen before the system goes to sleep.

## Features

- Decreases the screen brightness to 10% after 30 seconds of idle time.
- Restores the brightness to its original level upon resuming activity.
- Ensures that the screen is locked before the system goes to sleep.

## Pre-requisites

- `go 1.20` or higher
- `Sway or Hyprland`
- `swayidle` 
- `swaylock` 

## Installation

- clone this repository
- build the application with `go build`
- copy the binary to `/usr/local/bin` or any other location in your `$PATH`
- copy the `agent.sh in your $HOME/.config/swayidle/agent-brightness-manager.sh`

### Configuration
- In your Sway/Hyprland config add `exec $HOME/.config/swayidle/agent-brightness-manager.sh` 