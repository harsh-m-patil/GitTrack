# GitTrack

> A command-line tool to track and display a user's GitHub information directly in the terminal.

## Demo

![Demo](./assets/demo.gif)

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Commands](#commands)
- [Contributing](#contributing)
- [License](#license)

## Features

- Retrieve and display GitHub profile information.
- View user activity, including recent commits and repositories.
- Simple and intuitive command-line interface.

## Installation

You have two options to install GitTrack: **Clone the repository** or **Install via Go**.

### Option 1: Clone the Repository

1. Clone the repository to your local machine:

   ```bash
   git clone https://github.com/harsh-m-patil/GitTrack.git
   ```

2. Navigate to the GitTrack directory:

   ```bash
   cd GitTrack
   ```

3. Build the project:

   ```bash
   go build .
   ```

### Option 2: Install via Go

You can also install GitTrack directly using the `go install` command:

```bash
go install github.com/harsh-m-patil/GitTrack@latest
```

After installation, ensure that your `GOPATH/bin` directory is in your `PATH` environment variable to run GitTrack from anywhere.

### Run the Application

You can now run GitTrack to track your GitHub profile:

```bash
gittrack profile your-github-username
```

## Usage

After installation, simply replace `your-github-username` with any GitHub username to fetch and display their profile information and activity in the terminal.

## Commands

- **Profile**: Display user profile information.
  ```bash
  ./GitTrack profile your-github-username
  ```

Feel free to extend the command list as you add more features!

---

## Contributing

Contributions are welcome! If you’d like to contribute to GitTrack, please fork the repository and submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.

---

Thanks

A special thanks to ChatGPT for providing the roast messages that added some fun to this project!
