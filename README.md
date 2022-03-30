# procid

A multiplatform (Linux, Mac OS) command line tool to get process id for any given binary.

Unlike tools like `pidof` it matches exact binary name of the process.

Usage:
```shell
procid /bin/sh
```
It prints the first detected PID to `stdout` or nothing. All errors are dumped into `stderr`.
