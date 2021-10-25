# datadir

Package `datadir` provides tools for working with data-directories.

With Linux, Unix, and other Unix-like systems, one concept is — where should a user-application store user-data‽

Over the decades, the conventions for where user-applications should store their user-data has changed.

## Usage

Here is how it can be used:
```go
path, err := datadir.Path()

// path might, for example, have a value similar to "/home/username/.local/share"

// ...

const appDataDir string = "myapp"

var appDataPath string = filepath.Join(path, appDataDir)

// appDataPath might, for example, have a value similar to "/home/username/.local/share/myapp"
//
// And then the user-application would store this user's user-data at "/home/username/.local/share/myapp"

```

## Algorithm

One of the (at the time of writing this) current conventions for determining where a user-application should store its user-data is as follows:

### Step №1:

Look at the `$XDG_DATA_HOME` environment variable.

If the `$XDG_DATA_HOME` environment variable is set then its value is the base-directory under which applications should store their user-data, and we are done!

For example, the value might be:
* `~/.local/share`
* `/usr/local/share`
* etc

It could also be some random path a user decides on. For example: `~/apple/banana/cherry/share`, or `~/once/twice/thrice/fource/here`, etc.

(The `$XDG_DATA_HOME` environment variable is probably what the user should set if they want to change where user-applications they use store their user-data.)

However, if the `$XDG_DATA_HOME` environment variable is _not_ set then we have to go to the next step.

### Step №2:

Look at the `$XDG_DATA_DIRS` environment variable.

If the `$XDG_DATA_DIRS` envrionment variable is set then its value will contain a list of paths.

It might look something like the following:
* `~/.local/share:/usr/local/share:/usr/share`

(On my current system there are more paths in `$XDG_DATA_DIRS` than what I've listed there. So don't assume it will be the same for everyone.)

If the `$XDG_DATA_DIRS` envrionment variable is set then use the first directory as the base-directory under which the user-applications should store its user-data, and we are done!

However, if the `$XDG_DATA_DIRS` environment variable is _not_ set then we have to go to the next step.

### Step №3:

Default to using:
* `~/.local/share`

And we are done!

## ~/

One detail I've glazed-over was the use of `~/` in paths such as `~/.local/share`, `~/once/twice/thrice/fource/here`, etc.

The `~` part is a short-cut for the user's home-directory.

So, for example, if the user's home-directory is at `/home/daruish` —

Then `~/once/twice/thrice/fource/here` would become `/home/daruish/once/twice/thrice/fource/here` —

And `~/.local/share` would become `/home/daruish/.local/share —`

Etc.

Or course, the the user's home-directory was different, then the tilde-expansion would be different.

For example, if the user's home-directory was instead `/root` —

Then `~/once/twice/thrice/fource/here` would become `/root/once/twice/thrice/fource/here` —

And `~/.local/share` would become `/root/.local/share` —

Etc.

## References

One of the references this document references is:

* freedesktop.org's “XDG Base Directory Specification” https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html

Here are some relevant quotes from freedesktop.org's “XDG Base Directory Specification”:

> “There is a single base directory relative to which user-specific data files should be written.”

> “$XDG_DATA_HOME defines the base directory relative to which user-specific data files should be stored. If $XDG_DATA_HOME is either not set or empty, a default equal to $HOME/.local/share should be used.”

> “$XDG_DATA_DIRS defines the preference-ordered set of base directories to search for data files in addition to the $XDG_DATA_HOME base directory. The directories in $XDG_DATA_DIRS should be seperated with a colon ':'.”

> “The order of base directories denotes their importance; the first directory listed is the most important. When the same information is defined in multiple places the information defined relative to the more important base directory takes precedent. The base directory defined by $XDG_DATA_HOME is considered more important than any of the base directories defined by $XDG_DATA_DIRS.”
