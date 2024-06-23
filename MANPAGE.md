
# shikaku: The one tailed Branch Checker
A Golang CLI tool to compare a branch with the master branch, check for code changes, and validate basic coding principles using OpenAI's GPT-3.
## Installation
1. **Clone the repository:**
   ```sh
   git clone https://github.com/ray-cd/shikaku.git
   cd shikaku 
   ```
2. **Install dependencies:**

    ```sh
    go mod tidy
    ```

3. **Setup environment variables:**

    ```sh
    export "GITHUB_TOKEN=your_github_token"
    export "OPENAI_API_KEY=your_openai_api_key"
    ```

## Usage
To use the CLI tool, run the following command:
    ```sh
    go run main.go compare <repo_owner> <repo_name> <branch_name>
    ```
Replace <repo_owner>, <repo_name>, and <branch_name> with the appropriate values.

### Example

```sh
go run main.go compare yourusername yourrepository feature-branch
```



## Code Structure
- main.go: Initializes the Cobra root command and loads environment variables.
- compare.go: Contains the command to compare branches and fetch changes.
- openai.go: Interacts with OpenAI's GPT-3 to analyze code changes.
- .env: Stores your GitHub and OpenAI tokens (should not be committed to version control).
## Dependencies
- [Cobra](github.com/spf13/cobra) for building the CLI.
- [go-github](github.com/google/go-github/v46/github) for interacting with the GitHub API.
- [oauth2](golang.org/x/oauth2) for authentication with GitHub.
- [godotenv](github.com/joho/godotenv) for loading environment variables from a file.
## Contributing
Feel free to open issues or submit pull requests if you have any improvements or bug fixes.
## License
This project is licensed under the MIT License.
```
To ensure the README.md is comprehensive, replace any placeholder values like https://github.com/yourusername/branch_checker.git or yourusername with actual values specific to your repository and setup. This README.md provides a clear guide for anyone who wants to understand, set up, and use your tool. (edited) 