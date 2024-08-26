![image](https://github.com/user-attachments/assets/5577002a-7320-44f6-a4cd-171d471ebd53)

# Wikinow
Wikinow is an open-source project designed to provide a quick and easy way to set up a wiki for your next project. Built with Go, Wikinow allows you to maintain and serve your wiki content with minimal setup and dependencies.

## Features
**Simple Setup**: Get your wiki up and running in minutes with minimal configuration.
**Markdown Support**: Write your content in Markdown (CommonMark) for easy readability and formatting.
**Fast Performance**: Leverage the speed of Go to serve your content quickly and efficiently.
**Customizable Templates**: Use custom HTML templates to style your wiki to your liking.

## Installation
To install Wikinow, make sure you have Go installed on your system.

Clone the repository:

```bash
git clone https://github.com/lucasForato/wikinow.git
cd wikinow
```

Install with Make by running:
```bash
make install
```

## Usage
Wikinow has two basic commands that you can use:

#### `wikinow init`
This command will create a `wiki` directory and a `wikinow.yaml` file. 
The directory is where you will store all your markdown files that will be rendered as a wiki on the web for you. 
The configuration is just so you tell wikinow what port and the title of your application (for now).

#### `wikinow start`
This command will start the server with all your content within the `wiki` directory served at `http://localhost:{PORT}/wiki/{FILE_NAME}

For example:

```bash
wiki/
├── main.md  (accessible at http://localhost:{PORT}/wiki/)
├── example/
│   └── hello.md  (accessible at http://localhost:{PORT}/wiki/example/hello)
```

Configuration
Wikinow comes with a configuration file (wikinow.yaml) where you can customize the `title` and `port` of your application.

Contributing
Contributions are welcome! If you find a bug or have a feature request, please open an issue. If you’d like to contribute code, please fork the repository and submit a pull request.
