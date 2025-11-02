# mambu-assessment-test

## Brief description
This is an application writen in Go that has two main strategies:
 - A REST API that allows transactions to be published into a database and File storage by a XML file, following the provided XSD schema.
 - A File System Watcher that monitors a given directory for new CSV files received from the banking institution.

    This watcher will read the CSV for the received transactions, and update the transaction's status accordingly if present in DB 

## General instructions
1. Consider running the application or unit tests by make file. Run any playbooks by using UNIX based terminal

## Observed issues

C compiler "gcc" not found: exec: "gcc": executable file not found in %PATH%

**Solution:** some C compiler is needed to make sqlite work. On Windows, install [TDM-GCC](https://jmeubank.github.io/tdm-gcc/) and make sure to add it to your PATH.

## Running the application
1) Run a go mod install to install dependencies
2) Use make run to start the application. The server will start on port 8080 by default.

## Enhancements
1) Better error handling, and deep analysis if DB and files are properly closing their connections and buffers.
2) Units tests, unfortunately due time restrictions I was not able to implement them. However the code is developed around interfaces which makes gomock or other mocking approaches to be easily included. 