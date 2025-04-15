# Shredder

Shredder is a file shredding application written in Go. 
It irreversibly overwrites files with random data, rendering them unusable while keeping the file intact.

## Features

- **Secure File Shredding:** Overwrites file contents with random data for a specified number of passes.
- **Cross-Platform Compatibility:** Supports macOS and Windows.
- **Interactive TUI:** A terminal user interface built with [tview](https://github.com/rivo/tview) for easy navigation and operation.
- **Native File Dialog Integration:**
- **Safety Checks:**  
  - Empty files (0 bytes) are detected and flagged as unnecessary for shredding.
  - A maximum pass count is enforced (default is 10, configurable in the code) to prevent excessive processing and storage decay.
- **Logging:** Visual feedback is provided via a log window in the TUI, displaying status messages and error notifications.

## How It Works

1. **Select a File:**  
   Use the "Select a File" button to open a native file dialog and choose a file to shred.
   
2. **Configure Passes:**  
   Enter the desired number of overwrite passes (secure deletion standards typically require 3 to 7 passes).
   
3. **Start Shredding:**  
   Click "Start Shredding" to begin the process.
   The application overwrites the file content in chunks with random data, logging progress and any errors encountered.
   
4. **Quit:**  
   Use the "Quit" button to exit the application or just CTRL + C.

## Usage (from source)

Clone the repository, install dependencies, and run the application using the following commands:

```bash
go mod tidy
go run cmd/main.go
```

## System Requirements

- **Go:** Version 1.24 or higher.
- **Operating System:** macOS or Windows.
- **Dependencies:**  
  - [tview](https://github.com/rivo/tview) for the TUI.
  - [tcell](https://github.com/gdamore/tcell/v2) for terminal cell handling.
  - [sqweek/dialog](https://github.com/sqweek/dialog) (or AppleScript based logic for macOS) for file selection.