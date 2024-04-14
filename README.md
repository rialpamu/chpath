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

Tested under OpenBSD 7.2, Linux (Ubuntu 22.04) and Windows 10. \


## Example Usage

### Windows

Windows CMD:

    chpath %HOMEPATH%\bin %HOMEPATH%\other\bin > %TEMP%\set_my_path.cmd
    call %TEMP%\set_my_path.cmd

### Unix (OpenBSD, Linux)

ksh or bash:

    export PATH=$(chpath ~/bin ~/other/bin)

## Install

    go install github.com/rialpamu/chpath@latest

