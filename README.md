# shot-capture

Shot Capture is a tool designed to remotely control a DSLR camera and record single hi-resolution RAW images to disk.  It is meant to be used in conjunction with a motion controller rig.  The recorded images can later be processed and combined to create HD video sequences.

## Features

Timelapse mode
UI
    change focus?
    capture
    live preview
project config files?


## gphoto2

The go bindings for gphoto2 are sourced from https://github.com/aqiank/go-gphoto2
Additional bindings are sourced from https://github.com/micahwedemeyer/gphoto2go

The bindings are copied directly into the project rather than imported as vendor module so that the cgo settings can be tailored to the project and any additional binding features can be easily added as necessary.