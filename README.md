Run a command. If the command fails (non-zero exit), retry it `-n` times with delay of `-d` seconds.

This is not meant to be robust, but a simple utility for daily shell usage.

# Install

```
go get github.com/hayeah/goforwin
```

# Example

Retry `a-possibly-failing-command` 3 times, with 10 seconds between the attempts.

```
goforwin -n 3 -d 10 a-possibly-failing-command arg1 arg2
```

# Multitasking

Bash's job control already has multitasking baked in. To run goforwin in backgground, just append an `&`:

```
goforwin -n 3 -d 10 a-possibly-failing-command arg1 arg2 &
```

If a goforwin process is in the foreground, you can send it to background to start another:

1. ^z to signal the foreground process to stop.
2. Use the `bg` builtin command to continue it in background.
3. Start another process.

# Completion with Notification

To get desktop notification (for Mac):

```
goforwin longcommand && osascript -e 'display notification "whatever"'
```