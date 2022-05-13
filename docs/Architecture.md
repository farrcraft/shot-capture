# Application Architecture

The application is written in Go. It uses CGo bindings for gtk3 to provide UI features and libghoto2 for interfacing with the camera.

## Project Configuration

```json
{
    "projectName": "",
    "type": "timelapse",
    "frequency": 60, // seconds between captures
    "frames": 24 // total number of captures
}
```

## UI

Button - Detect Cameras
Dropdown - Camera list
