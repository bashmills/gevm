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

## Making changes

To contribute follow these steps:

1. Clone the repository to your local machine:

   ```
   git clone https://github.com/bashmills/gevm.git
   ```

2. Create a new branch for your feature or bug fix:  

   ```
   git checkout -b feature/new-feature
   ```

3. Make your changes and commit them:  

   ```
   git commit -m "Add new feature"
   ```

4. Push your changes to the `bashmills/gevm` repository:  

   ```
   git push origin feature/new-feature
   ```

5. Open a pull request from your branch to the `main` branch of the `bashmills/gevm` repository.

6. Your pull request will be reviewed and once approved it will be merged into the `main` branch.

## Deployment

New deployments are automatically built and published via GitHub Actions whenever a new tag is pushed to the repository.

## License

This project is licensed under the [MIT License](LICENSE).
