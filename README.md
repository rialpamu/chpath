# chpath

A simple helper program to adapt the PATH environment variable.

- Prepends its arguments to PATH
- Removes duplicates
- Removes non-existing directories from PATH
- Converts elements to absolute path
- Dereferences symbolic links (may be suppressed with `-keepsymlinks`)

Called without arguments it will clean the PATH
(remove duplicates, remove non-existing directories etc).

Will not actually modify the PATH environment variable,
but write the modified PATH variable to stdout.

Tested under OpenBSD 6.6, Linux (Ubuntu 18.04) and Windows 7.
Built with go1.13.

I find it useful for Windows commandline, where my PATH variable
has the tendency to grow unrestrained.


## Example Usage

### Windows

Windows CMD:

    chpath %HOMEPATH%\bin %HOMEPATH%\other\bin > %TEMP%\set_my_path.cmd
    call %TEMP%\set_my_path.cmd

### Unix (OpenBSD, Linux)

ksh or bash:

    export PATH=$(chpath ~/bin ~/other/bin)

