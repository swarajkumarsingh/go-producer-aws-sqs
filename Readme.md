# <h1 align="center">Dynamic AWS SQS Producer for Golang</h1>
___
<p align="center">
<a href="Java url">
    <img alt="Java" src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" />
</a>
<a href="License url" >
        <img alt="BSD Clause 3" src="https://img.shields.io/badge/License-MIT-yellow.svg"/>
    </a>
</p>

---

### Overview
This repository contains a dynamic AWS SQS producer implemented in Golang, suitable for handling various types of data. It is Dockerized for easy deployment and usage.


## Table of Contents

- [Features](#features)
- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
   - [Environment Variables](#environment-variables)
- [Contributing](#contributing)
- [Contact](#contact)
- [License](#license)

## Technologies <a name="technologies"></a>

- [Go](https://golang.org/)
- [Docker](https://www.docker.com/)
- [AWS SES](https://aws.amazon.com/)
- [AWS SQS](https://aws.amazon.com/)

## Features

- Dynamic: Supports various types of data for sending to AWS SQS.
- AWS SQS Integration: Seamlessly integrates with AWS SQS for message queueing.
- Golang Implementation: Utilizes the Go programming language for efficient and concurrent processing.
- Dockerized: Packaged into a Docker container for portability and ease of deployment.
- Easy Debug

### Prerequisites
Before running this service, ensure you have the following prerequisites installed:

- Docker: Install Docker
- AWS Account: Obtain AWS credentials and configure them locally.

## Environment Variables

Table bellow shows the obligatory environment variables for mariadb container. You should set them based on what was also set for backend container.

Environment variable  | Default value | Optional
--- | --- | ---
STAGE | dev | `NO`
SQS_URL |  | `YES`
AWS_REGION |  | `YES`
AWS_ACCESS_KEY |  | `YES`
AWS_SECRET_ACCESS_KEY |  | `YES`

## Getting started
   ```
   First of all, correctly configure the Golang development environment on your machine, see https://go.dev/doc/install
   
   - Clone this repository:
   $ https://github.com/swarajkumarsingh/go-producer-aws-sqs.git

   - Enter in directory:
   $ cd go-producer-aws-sqs

   - For install dependencies(optional):
   $ make run SQS_URL= AWS_REGION= AWS_ACCESS_KEY= AWS_SECRET_ACCESS_KEY=
   ```
---

### Contributing
Contributions are welcome! If you'd like to contribute to this project, please follow these guidelines:
```
1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes and test thoroughly.
4. Commit your changes with clear commit messages.
5. Create a pull request against the main branch.
```

## Contact

- Swaraj Singh <br> <br>
  <a  href="https://www.linkedin.com/in/swarajkumarsingh/" target="_blank"><img alt="LinkedIn" src="https://img.shields.io/badge/linkedin%20-%230077B5.svg?&style=for-the-badge&logo=linkedin&logoColor=white" /></a>
  <a href="sswaraj169@gmail.com"><img  alt="Gmail" src="https://img.shields.io/badge/Gmail-D14836?style=for-the-badge&logo=gmail&logoColor=white" />

  feel free to contact me!

## License

> You can check out the full license [here](https://github.com/swarajkumarsingh/go-producer-aws-sqs/blob/main/LICENSE)

This project is licensed under the terms of the **MIT** license

### Acknowledgments
Special thanks to AWS for providing powerful cloud services and Golang community for building robust frameworks and libraries.
