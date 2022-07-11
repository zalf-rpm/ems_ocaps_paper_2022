# Code examples for 2022 EMS object capability paper

This repository contains some working and complete code examples for the (as of yet) unpublished EMS paper on object capabilities in our ZALF model and simulation infrastructure. These examples are not necessarily supposed to present the most secure and best way to design an API, rather to support the paper and show working examples. We try to add different implementations in different languages and maybe subsequently add more complete and closer to the real world examples. The currently incomplete real infrastructure code can be found https://github.com/zalf-rpm/mas-infrastructure.

## Use & Compilation
Different languages have different needs. Except for Python all languages need the schema files in the capnp folder to be processed by the Capn'n Proto schema compiler to generate the actual languages code. The following sections will try to list what's needed to run and/or compile the according languages code.

The different implementations should be interchangeable. 

### Python

__requirements__

- Python3
- install pycapnp via pip

* run Python example

    python revokable_forwarder_example_aio.py a=py b=py c=py m=py

### C#

__requirements__
- dotnet 6.0

### Go

__requirements__

### C++

__requirements__

- cmake
- VCPKG
- git
- Python3 (for running the starter script)

__installation__

* checkout this repository and dependencies

    git clone https://github.com/zalf-rpm/ems_ocaps_paper_2022.git
    git clone https://github.com/Microsoft/vcpkg.git

* build vcpkg

    cd vcpkg
    ./bootstrap-vcpkg.sh

* build Cap'n Proto

    ./vcpkg install capnproto:x64-linux

* build cpp code

    cd ems_ocaps_paper_2022
    sh init_cmake_linux_debug.sh
    cd _cmake_debug
    make

* run cpp example

    python revokable_forwarder_example_aio.py a=cpp b=cpp c=cpp m=cpp


## Examples

1. classic Alice, Bob and Carol interaction with Alice sends to Bob a revokable capability to Carol


