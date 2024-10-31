# Sail Host cloud

## About

This project is a simple web application that allows you to manage your Sail Host cloud.

## Installation

To install the project, you need to have Go installed on your machine.

First, clone the repository:

```bash
git clone https://github.com/sail-host/cloud.git
```

Then, install the dependencies:

```bash
go mod tidy
```

Install the web dependencies:

```bash
cd web && bun install
```

## Running the project

To run the project, you need to have Go installed on your machine.

```bash
make dev
```

Run only backend golang service:

```bash
make dev-api
```

## Project testing

```bash
make test
```

## License

This project is licensed under the MIT License. See the LICENSE file for more details.
