# Contributing to SourceVault

First off, thank you for considering contributing to SourceVault! It's people 
like you that make open source such a great community.

## Developer Certificate of Origin (DCO)

We enforce the Developer Certificate of Origin (DCO) on all pull requests. 
This means that all commits must be signed off by the author.

You can sign off your commits using the `-s` flag with Git:

    git commit -s -m "feat: added a new feature"

By signing off, you agree to the terms listed in the `DCO` file in the root 
of this repository.

## How to Contribute

We use a decentralized, email-driven workflow for contributions. Because our 
GitHub repository is purely a mirror, **we do not accept Pull Requests via the 
GitHub UI.** 

Instead, we use a signed-tag pull request workflow to ensure cryptographic 
integrity:

1. **Clone the repository** and create your feature branch from `main`.
2. **Log the Task:** Add your task to the `TODO` file and create a detailed 
   Markdown file in `boards/tasks/YYYY/MM/DD/` that includes your 
   Implementation Plan.
3. **Write clear, well-documented code** that aligns with the Hexagonal 
   Architecture and CQRS patterns outlined in our documentation.
4. **Make sure your code compiles** and passes all existing tests.
5. **Sign off on all commits** (as detailed above).
6. **Create a signed tag** for your changes: 
   `git tag -s my-feature -m "Feature description"`
   *(This ensures that once your pull request is created, the code cannot be 
   silently altered).*
7. **Generate a pull request summary** using `git request-pull`:
   `git request-pull main <your-public-repo-url> my-feature`
8. **Send the resulting output** via email to `dev@gotcode.org`.

## Code of Conduct

Please note that this project is released with a Contributor Code of Conduct. 
By participating in this project you agree to abide by its terms. See the 
`CODE_OF_CONDUCT.md` file for more details.
