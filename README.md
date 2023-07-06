# Swayidle Brightness Manager

This tiny Go application is a utility for Hyprland or Sway that get your actual brightness and set it to 10% after X seconds of idle time. The bash script also locks the screen before the system goes to sleep. It is designed to work with `swayidle` and `swaylock`.


## Features

- Decreases the screen brightness to 10% after 30 seconds of idle time.
- Restores the brightness to its original level upon resuming activity.
- Ensures that the screen is locked before the system goes to sleep.

## Pre-requisites

- `go 1.20` or higher
- `Sway or Hyprland`
- `swayidle` 
- `swaylock` 
- `brightnessctl`
- `playerctl`