# Data Compression Tool

This project provides a simple command-line tool for compressing and decompressing files using the Huffman coding algorithm. It supports compressing text files into a `.huff` format and decompressing `.huff` files back into their original text format using a corresponding `.json` file containing the Huffman codes.

## Features

- **Compression**: Compress a text file into a `.huff` file and generate a corresponding `.json` file containing the Huffman codes.
- **Decompression**: Decompress a `.huff` file back into its original text format using the `.json` file.

## Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/your-username/data-compression.git
   cd data-compression
   ```

2. Install dependencies:
   ```sh
   go mod tidy
   ```

3. Build the project:
   ```sh
   go build -o data-compression
   ```

## Usage

The tool provides two main functionalities: compression and decompression. Use the following flags to specify the desired operation:

### Compress a File

To compress a file, use the `-c` flag followed by the input file name and the desired output file name (without extensions). For example:

```sh
./data-compression -c input.txt output
```

This will create:
- `output.huff`: The compressed file.
- `output_codes.json`: The JSON file containing the Huffman codes.

### Decompress a File

To decompress a file, use the `-e` flag followed by the `.huff` file, the `.json` file containing the Huffman codes, and the desired output file name (without extensions). For example:

```sh
./data-compression -e input.huff input_codes.json output
```

This will create:
- `output.txt`: The decompressed file.

### Example

#### Compressing a File
```sh
./data-compression -c example.txt compressed
```
Output:
- `compressed.huff`
- `compressed_codes.json`

#### Decompressing a File
```sh
./data-compression -e compressed.huff compressed_codes.json decompressed
```
Output:
- `decompressed.txt`

## Error Handling

- If the required arguments are not provided, the tool will display an error message and exit.
- Ensure that the input files exist and are accessible before running the commands.

## Project Structure

```
.
├── main.go               # Entry point of the application
├── huffman/
│   ├── encoding.go       # Huffman encoding logic
│   ├── decoding.go       # Huffman decoding logic
├── go.mod                # Go module file
├── big.txt               # Example large input file
├── small.txt             # Example small input file
├── small.huff            # Example compressed file
├── small_codes.json      # Example Huffman codes file
```

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.