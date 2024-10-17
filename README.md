# TinyDB

## Introduction

TinyDB is a simple, lightweight key-value store implemented in Go. This project serves as an introduction to a larger database project, focusing on the core functionalities of data storage and retrieval. TinyDB aims to provide a clear understanding of how a basic database operates, including operations such as adding, retrieving, deleting, and searching for entries.

## Features

- **Key-Value Storage**: Store data as key-value pairs.
- **Concurrent Access**: Thread-safe operations using mutex locks.
- **JSON Serialization**: Load and save data in JSON format.
- **Basic CRUD Operations**:
  - `Put(key, value)`: Add or update an entry.
  - `Get(key)`: Retrieve an entry by its key.
  - `Delete(key)`: Remove an entry by its key.
  - `Search(value)`: Find all keys that have a specific value.

## Installation

To use TinyDB in your Go project, clone this repository and import the package:
