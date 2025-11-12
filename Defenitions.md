# Definitions

[ANSI escape code](https://en.wikipedia.org/wiki/ANSI_escape_code): A standard for in-band signaling to control cursor location, color, and other options on text terminals. ANSI escape codes are used to format text in terminal emulators.


## Look up table
| Action                          | Code                | Description           |
| ------------------------------- | ------------------- | --------------------- |
| Move cursor to (row, col)       | `\033[<row>;<col>H` | Absolute position     |
| Move up N lines                 | `\033[<N>A`         | Keeps same column     |
| Move down N lines               | `\033[<N>B`         | Keeps same column     |
| Move right N columns            | `\033[<N>C`         | Keeps same row        |
| Move left N columns             | `\033[<N>D`         | Keeps same row        |
| Move to beginning of line below | `\033[<N>E`         | Jumps to next line    |
| Move to beginning of line above | `\033[<N>F`         | Jumps up lines        |
| Move to column N                | `\033[<N>G`         | Same row, column N    |
| Save cursor position            | `\033[s`            | Remember current spot |
| Restore saved position          | `\033[u`            | Jump back             |
| Hide cursor                     | `\033[?25l`         | (lowercase L)         |
| Show cursor                     | `\033[?25h`         |                       |
