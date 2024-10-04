# wcgo - Custom Word Count CLI Tool

`wcgo` is a custom command-line tool inspired by the Unix `wc` command written in go as per `Coding challenges` by `John Cricket`. It can be used to count lines, words, bytes, and characters from a given file or standard input.

## Features

- Count **lines**, **words**, **bytes**, and **characters**.
- Supports **file input** or **standard input** (stdin).
- Provides easy-to-use command-line flags to specify the desired count.

## Installation

To build the `wcgo` CLI tool, clone the repository and build the binary:

```sh
# Clone the repository
$ git clone <repository_url>

# Navigate to the project directory
$ cd wc-go

# Build the binary
$ go build -o wcgo
```

After running the above command, you will have an executable file named `wcgo`.

## Usage

The `wcgo` tool provides several command-line options to specify what counts you want:

- `-l` : Count the number of **lines**.
- `-w` : Count the number of **words**.
- `-c` : Count the number of **bytes**.
- `-m` : Count the number of **characters**.

You can use multiple flags at once to get different counts simultaneously.

### Example Usage with File Path

If you have a file named `example.txt` and you want to use `wcgo` to count lines, words, or other metrics:

```sh
# Count lines, words, and bytes in the file
$ ./bin/wcgo -l -w -c example.txt
```

### Example Usage Without File Path (Using Stdin)

You can also use `wcgo` to read input directly from `stdin`. For example:

```sh
# Pipe text into wcgo to count lines, words, and bytes
$ echo "Hello World! This is a test." | ./bin/wcgo
```

### Flags

- **`-l`**: Print the number of **lines**.
- **`-w`**: Print the number of **words**.
- **`-c`**: Print the number of **bytes**.
- **`-m`**: Print the number of **characters**.

If no flags are provided, `wcgo` defaults to counting **lines**, **words**, and **bytes**.

### Example Output

For a file named `example.txt`:

```sh
$ ./bin/wcgo -f example.txt
10 45 256 example.txt
```

This output shows:

- **10** lines
- **45** words
- **256** bytes
- The filename (`example.txt`)

## Cross-Platform Compilation

If you need to compile `wcgo` for another platform, you can use Go's cross-compilation capabilities:

- For Linux:

  ```sh
  GOOS=linux GOARCH=amd64 go build -o wcgo
  ```

- For Windows:

  ```sh
  GOOS=windows GOARCH=amd64 go build -o wcgo.exe
  ```

- For macOS:
  ```sh
  GOOS=darwin GOARCH=amd64 go build -o wcgo
  ```

## License

This project is licensed under the MIT License. See the LICENSE file for more details.

## Contributing

Contributions are welcome! Feel free to open an issue or submit a pull request for any improvements or new features.
