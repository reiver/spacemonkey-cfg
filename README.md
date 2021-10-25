# spacemonkey-cfg

Get spacemonkey configuration information.

## Usage

This program is meant to be used both my humans, and by machines.

Humans will probably run it like this:
```
spacemonkey cfg --list
```
Or this:
```
spacemonkey cfg -eol <name>
```
Or this:
```
spacemonkey cfg -eol -v <name>
```

Where machines will probably run it like:
```
spacemonkey cfg <name>
```

## Example Usage

To get the path to the data-directory for the gemini-protocol:
```
spacemonkey cfg protocol/gemini/data-path
```

To get the path to the data-directory for the rook-protocol:
```
spacemonkey cfg protocol/rook/data-path
```

To get a complete list of all the (registered) configuration names.
```
spacemonkey cfg --list
```

## See Also
* https://github.com/reiver/spacemonkey
