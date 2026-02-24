# Godot Engine Version Manager (gevm) - CONTRIBUTING

Contributions are welcome via pull requests. Not really expecting any contributions but this is still nice to have here for personal reference.

## Table of Contents
- [Prerequisites](#prerequisites)
- [Making Changes](#making-changes)
- [Deployment](#deployment)
- [License](#license)

## Prerequisites

Before you begin, ensure you have the following prerequisites installed:

- [Go v1.22.4](https://go.dev/dl/)

### Installing Go via gobrew

1. **Install gobrew:**

   Download the latest version of gobrew by following the instructions found [here](https://github.com/kevincobain2000/gobrew).

2. **Install Go:**

   Once gobrew is installed, open the terminal and run the following command to install Go version 1.22.4:

   ```
   gobrew install 1.22.4
   ```

3. **Set Go version:**

   After installing Go, you can set it as the default version:

   ```
   gobrew use 1.22.4
   ```

4. **Verify installation:**

   To verify that Go is installed correctly, you can run:

   ```
   go version
   ```

## Making Changes

To contribute follow these steps:

1. Fork the repository to your own GitHub account using the **Fork** button.

2. Clone your fork to your local machine:

   ```
   git clone https://github.com/<your-username>/gevm.git
   ```

3. Change into the project directory:

   ```
   cd gevm
   ```

4. Add the original repository as an upstream remote:

   ```
   git remote add upstream https://github.com/bashmills/gevm.git
   ```

5. Create a new branch for your feature or bug fix:

   ```
   git checkout -b feature/new-feature
   ```

6. Make your changes and commit them:

   ```
   git commit -m "Add new feature"
   ```

7. Push your changes to your fork:

   ```
   git push origin feature/new-feature
   ```

8. Deal with merge conflicts:

    1. Rebase your branch onto the latest `main` branch from upstream (preferred over merge):

       ```
       git pull --rebase upstream main
       ```

    2. Resolve conflicts and continue the rebase until success:

       ```
       git rebase --continue
       ```

    3. Force push the changes after rebasing:

       ```
       git push --force-with-lease
       ```

9. Open a pull request from your forked repository’s branch to the `main` branch of the `bashmills/gevm` repository.

10. Your pull request will be reviewed and once approved it will be merged into the `main` branch.

## Deployment

New deployments are automatically built and published via GitHub Actions whenever a new tag is pushed to the repository.

## License

This project is licensed under the [MIT License](LICENSE).
