Here's a `README.md` file for your CLI:

# My Personal CLI Assistant

A powerful and customizable command-line assistant to manage VPN connections and interact with your local LLM.  

## Prerequisites

- **Ollama**: Ensure that Ollama is installed on your local machine to use the LLM functionality.  
  Follow the instructions in the [Ollama documentation](https://ollama.com/) to set it up.

## Installation  

1. Clone the repository and navigate to its directory.
2. Build the binary using the following command:  
   ```bash
   git clone ...
   make build
   ```

## Setting Up an Alias

To simplify usage, create an alias for the CLI with the required environment variables. Replace placeholders like `[VPN HOST]` and `[LLM MODEL]` with your specific values:

```bash
alias foo="APP_NAME=Foo VPN_HOST=[VPN HOST] CISCO_BIN_DIR=[CISCO BIN DIR] LLM_HOST=[LLM HOST] LLM_MODEL=[LLM MODEL] /usr/local/bin/Foo"
```

Add this alias to your shell configuration file (`~/.bashrc`, `~/.zshrc`, etc.) for persistence:
```bash
echo 'alias foo="APP_NAME=Foo VPN_HOST=[VPN HOST] CISCO_BIN_DIR=[CISCO BIN DIR] LLM_HOST=[LLM HOST] LLM_MODEL=[LLM MODEL] /usr/local/bin/Foo"' >> ~/.bashrc
source ~/.bashrc
```

## Usage

### Basic Command Structure
```bash
foo [command]
```

### Available Commands

| Command      | Description                                  |  
|--------------|----------------------------------------------|  
| `completion` | Generate the autocompletion script for your shell. |  
| `help`       | Get help about any command.                 |  
| `llm`        | Manage requests to the local LLM.           |  
| `vpn`        | Manage the Cisco secure VPN.                |  

### Flags

| Flag          | Description                                  |  
|---------------|----------------------------------------------|  
| `-h, --help`  | Display help information for the CLI.        |  
| `-v, --version` | Display the current version of the CLI.   |  

### Examples

- **Check CLI version:**
  ```bash
  foo --version
  ```  

- **Manage VPN:**
- Connect to the VPN
  ```bash
  foo vpn c
  ```  
- Disconnect to the VPN
  ```bash
  foo vpn d
  ```  
- Show the stats of the VPN
  ```bash
  foo vpn s
  ```  

- **Interact with LLM:**
- Correct a given text
  ```bash
  foo llm c "What is the weather today?"
  ```  
- Reformulate a given text
  ```bash
  foo llm r "Announcement to my team to reach out the the HR?"
  ```  
- Answer a given question
  ```bash
  foo llm r "What's the color of the sky"
  ```  

---

### Contributing

Feel free to submit issues or contribute to the project by opening pull requests.

---

### License

This project is licensed under the MIT License.