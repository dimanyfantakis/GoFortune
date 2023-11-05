# GoFortune

A [Fortune](https://en.wikipedia.org/wiki/Fortune_(Unix)) clone using Go!

## Prerequisites

- [Go](https://go.dev/doc/install)

- fortune
```bash
sudo apt-get install fortune
```

## Installation

Compile the code into an executable.
```go
go build -o gofortune
```

You can discover the install path by running the go list command, as in the following example:

```go
go list -f '{{.Target}}'
```

Add the Go install directory to your system's shell path.

As an alternative, you can change the install target by setting the GOBIN variable using the go env command:

```go
go env -w GOBIN=/path/to/your/bin
```

Compile and install the package.

```go
go install github.com/dimanyfantakis/gofortune
```

## Usage

Run the application
```go
gofortune
```

Or add `gofortune` to your `~/.bashrc` file.

If you don't want to get a fortune on every new terminal session but rather once per login then you can do the following.

In the `~/.bashrc` file add the following.

```bash
# Define the flag file path
flag_file="/tmp/.firstrun_$USER"

# If the flag file doesn't exist
if [ ! -f "$flag_file" ]; then
    # Execute the gofortune command
    gofortune

    # Create the flag file in /tmp
    touch "$flag_file"
fi
```

In the `~/.bash_logout` file add the following.

```bash
flag_file="/tmp/.firstrun_$USER"

if [ -f "$flag_file" ]; then
    rm -f "$flag_file"
fi
```
