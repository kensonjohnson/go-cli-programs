# Command Line Todo App

Super simple todo list app, fully in the terminal.

## Running the App

If you have [Just](https://just.systems/) installed, you can `just run` to launch the app.

Otherwise, run `go run .` from the project root.

## Usage

There are 4 commands you can use once the app is running:
- `show`
- `create`
- `remove`
- `done`

`show` will provide a list of all todos including their Id.

`create` can accept any number of arguments, which will all be treated as a single string:

`create My new todo!`

`remove` and `done` need a todo's Id, or panics:

`remove 1`, `done 5`

