slowcp.exe
==========

A simple utility to copy a file from source to destination.

I wrote it as I realized that Linux tends to cache large files when copying, has the effect of seeing the file hit 100% on copy and then the file is still being written in the background. 

This is pretty annoying when you are copying a large file to a USB stick and then pull it out thinking it is done.

This utility will copy the file in chunks and display the progress as it goes.

It copies in small chunks, flushes it and reports the percentage giving it a better sense of progress.


![image](https://github.com/rebooting/go_slowcp/assets/487900/e3aa2f86-22de-46a7-b945-f2b2418e976b)


Build
-----

You need Go installed.

If you have VScode and Podman you can trigger the devcontainer and run the build inside too


```make```


Usage
-----

```slowcp.exe <source> <destination>```
