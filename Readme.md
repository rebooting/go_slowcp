slowcp.exe
==========

A simple utility to copy a file from source to destination.

I wrote it as I realized that Linux tends to cache large files when copying, has the effect of seeing the file hit 100% on copy but then the file is still being copied in the background. 

This is pretty annoying when you are copying a large file to a USB stick and then pull it out thinking it is done.

This utility will copy the file in chunks and display the progress as it goes.

It copies in small chunks, flushes it and reports the percentage giving it a better sense of progress.

Build
-----

You need Go installed.

If you have VScode and Podman you can trigger the devcontainer and run the build inside too


```make```


Usage
-----

```slowcp.exe <source> <destination>```
