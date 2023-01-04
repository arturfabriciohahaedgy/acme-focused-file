# acme-focused-file
This is a fork/rewrite of [fhs](https://github.com/fhs)'s [acmefocused](https://github.com/fhs/acme-lsp/tree/master/cmd/acmefocused) which uses a simple text file instead of a socket. Made because the original acmefocused wasn't working for me.

# Setup and installation
You can install `acme-focused-file` by running the following commands

``` shell
$ git clone https://github.com/arturfabriciohahaedgy/acme-focused-file.git
$ cd acme-focused-file
$ go install
```
# Usage
To just generate the file and check the current window you can use it like this:

``` shell
$ acme & # Your acme startup script
$ acme-focused &
$ cat /tmp/acme-focused
   2 # Window id
```

As shown, by default `acme-focused` writes to `/tmp/`, creating the file `/tmp/acme-focused`.

You can also pass a parameter to `acme-focused` which will determine the folder where it will store the file. For an example, like this:

``` shell
$ acme &
$ acme-focused $XDG_CACHE_HOME &
$ cat $XDG_CACHE_HOME/acme-focused
   1 # Window id
```

# Issues
Since `acme-focused-file` works by writing acme's `$winid` every two seconds to a temporary file, it isn't exactly the most optimal way to check the window's id. Feel free to make a fork or a PR request with changes to the source code which you think would make the program better.

# Credits and Acknowledgments
- [fhs](https://github.com/fhs)'s [acmefocused](https://github.com/fhs/acme-lsp/tree/master/cmd/acmefocused) and [acme-lsp](https://github.com/fhs/acme-lsp), where a part of the source code was borrowed and modified from.
