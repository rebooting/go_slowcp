slowcp.exe
==========

A simple utility to copy a file from source to destination.

I wrote it as I realized that Linux tends to cache large files when copying, has the effect of seeing the file hit 100% on copy and then had to wait for the write to complete with no feedback how much was done.

This is pretty annoying when I was copying some large files to a USB stick and I had to leave in a hurry, had to wait for the i/o to complete. Wrote this to solve that little problem without having to touch the caching config of my OS. Thus this little tool.

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
