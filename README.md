# Streaming Demo

This project demonstrates how to fetch and stream data from a database using Go. It includes two main functionalities: streaming data in chunks and fetching all data at once.

## Project Structure

- `main.go`: The entry point of the application.
- `fetch_data/streaming-data.go`: Contains the function to stream data in chunks.
- `fetch_data/non-streaming-data.go`: Contains the function to fetch all data at once.
- `database`: Contains database initialization and configuration.
- `models`: Contains the data models used in the project.

## Endpoints

- `/all_records`: Streams all records from the database in chunks.
- `/all_records_non_streaming`: Fetches all records from the database in a single response.

## Setup

1. **Install Go**: Make sure Go is installed on your system. You can download it from the [official Go website](https://golang.org/dl/).

2. **Clone the repository**:
   ```sh
   git clone <repository-url>
   cd <repository-directory>