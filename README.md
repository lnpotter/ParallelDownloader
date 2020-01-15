# WARNING: This program has been created for testing and educational purposes, it is not 100% functional.

# ParallelDownloader

ParallelDownloader is a CLI tool written in Go that allows you to download multiple files concurrently. It reads URLs from a text file (`urls.txt`) and downloads each file in parallel, providing a robust and efficient way to handle large batches of downloads.

## Features

- **Concurrent Downloads:** Downloads files concurrently using Go's goroutines for maximum efficiency.
- **Error Handling & Retries:** Automatically retries failed downloads up to 3 times with a backoff mechanism.
- **Dynamic URL Loading:** Reads URLs dynamically from a text file (`urls.txt`), making it easy to manage your download tasks.

## Technologies Used

- **Go 1.14.15**: The project is built using the Go programming language version 1.14.15, leveraging its powerful concurrency model.

## How to Use

1. Clone the repository to your local machine.
2. Add the URLs you want to download to the `urls.txt` file, one per line.
3. Run the program using the following command:

   \`\`\`bash
   go run downloader.go
   \`\`\`

4. The files will be downloaded and saved in the same directory as the program.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
