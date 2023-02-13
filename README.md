# Ascii-Art-Fs

This program takes two arguments (String) and a (banner) and return written string in font you have chosen.

## Installation

```bash

  git clone https://01.kood.tech/git/suzoagba/Ascii-Art-Fs.git

```

## Run This Locally

Go to the project directory

```bash

  cd ascii-art-fs

```

Run the program

```bash

  go run . [String] [Banner]

```

These are the three banners available

- shadow
- standard
- thinkertoy

**How to run the program accurately:**

-while in the project directory-

```bash

go run . "Hello world!" standard

```

Output:

``` _    _          _   _                                           _       _   _
| |  | |        | | | |                                         | |     | | | |
| |__| |   ___  | | | |   ___         __      __   ___    _ __  | |   __| | | |
|  __  |  / _ \ | | | |  / _ \        \ \ /\ / /  / _ \  | '__| | |  / _` | | |
| |  | | |  __/ | | | | | (_) |        \ V  V /  | (_) | | |    | | | (_| | |_|
|_|  |_|  \___| |_| |_|  \___/          \_/\_/    \___/  |_|    |_|  \__,_| (_)

```

```bash

go run . "Hello world!" shadow

```

Result:

```*
_|    _|          _| _|                                                     _|       _| _|
_|    _|   _|_|   _| _|   _|_|         _|      _|      _|   _|_|   _|  _|_| _|   _|_|_| _|
_|_|_|_| _|_|_|_| _| _| _|    _|       _|      _|      _| _|    _| _|_|     _| _|    _| _|
_|    _| _|       _| _| _|    _|         _|  _|  _|  _|   _|    _| _|       _| _|    _|
_|    _|   _|_|_| _| _|   _|_|             _|      _|       _|_|   _|       _|   _|_|_| _|

```

```bash

go run . "Hello world!" thinkertoy

```

Result:

```*

o  o     o o                             o    o o
|  |     | |                             |    | |
O--O o-o | | o-o       o   o   o o-o o-o |  o-O o
|  | |-' | | | |        \ / \ /  | | |   | |  |
o  o o-o o o o-o         o   o   o-o o   o  o-o O

```

OR YOU CAN START TEST FILE

```*

bash test.sh

```

### Audit

```*

https://github.com/01-edu/public/blob/master/subjects/ascii-art/fs/audit.md

```

## Authors

-Samuel Uzoagba <https://01.kood.tech/git/suzoagba/Ascii-Art-Fs.git>
