# mswpdl
Microsoft Bing Daily Image Downloader

This program downloads the Bing image of the day from Microsoft's API, and spits the output to the terminal.
It is designed to be used in conjunction with other command line image editing programs so that the user can
edit the image before displaying it.

While this program was designed for use in Linux, it should work on any platform that the Go language supports.

## Features
* Directly write image to file instead of through shell's `stdout`.
* Download images from other locales.
* Download old images from previous days.

## Planned Features
* Download images from Windows Spotlight.
* Implement a `-query` switch that allows user to get the image name or URL.

## Known Issues
* PowerShell messes with the encoding of the file when it is dumped to stdout, yielding a corrupted image.
  Use the `-outfile` switch instead.
