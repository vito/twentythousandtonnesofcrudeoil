# twentythousandtonnesofcrudeoil

[![the front fell off](https://img.youtube.com/vi/3m5qxZm_JqM/0.jpg)](https://www.youtube.com/watch?v=3m5qxZm_JqM)

This hokey little package will modify an option parser for
[go-flags](https://godoc.org/github.com/jessevdk/go-flags) to support pulling
all flags from the environment, optionally with a prefix. All env var names
will be derived from the fully namespaced flag name.

After parsing, the values are towed outside the environment.
