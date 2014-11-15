# watchmaker

Simple program to watch the current directory for changes and run a command (make by default). Like facebook's [watchman](https://github.com/facebook/watchman) only simpler and not *a service*.

watchmaker attempts to batch up changes, waiting for 100ms of silence (no changes to files) before running the build step. Any changes while building (including from the build) will trigger another build.

watchmaker isn't (yet) designed for projects which always modify files when building. If running `make` on your project a second time outputs `make: Nothing to be done for 'all'` watchmaker is ready for you.

## Usage

Invocation is very simple for now. It is likely to change and be more feature rich in later versions.

```
watchmaker
```

Runs `make` any time a file is created, deleted, moved, or changed.

```
watchmaker go test
```

Runs `go test` any time file is created, deleted, moved, or changed.

