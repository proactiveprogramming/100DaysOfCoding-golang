# 01_hellomulti
Demonstrate how to setup a multi file golang project.  

Each directory is intended to be its own project unless it is a package that is included in main.

In a normal project, each diretory would be its own repository, but due to the fact that these are learning projects, we are just
placing them all into one main repository.  

All of these projects have a corresponding "helloworldproject" and "greetingsproject" that go together.
All of these projects are demonstrating how to ue a local library package with different scenarios.

The only exception is "greetingstest" which is demonstrating how to set a local library package
without having a main module. Demonstrations on how to create "Examples", "Fuzz", and "Benchmark" testing are located in their own examples.

The main purpose of these examples is to demonstrate how to setup a main project that links with
a local library package.  Anything beyond that scope is in a separate demonstration.


## Greetings.sh
This is a demonstration file for how to compile a library without any main code.

## helloworldmulti.sh
This is a demonstration file for how to compile a main module that links in local library files.

## helloworlderror
This is a simple demonstration on how to setup error check and error logging.  This is 
just a simple example to demonstrate a very basic setup.

## helloworldinit
This is a simple demonstration on how to initialize a local library package.

## helloworldmap
This is a simple demonstration on how to use a map to store initial data.

## helloworldtest
This is a simple demonstration on how to test a local library package.  This is 
just a simple example to demonstrate a very basic setup.  Example does not include
"Examples", "Fuzz", or "Benchmarks".

