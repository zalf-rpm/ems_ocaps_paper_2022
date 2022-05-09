# Code examples for 2022 EMS object capability paper

This repository contains the some working and complete code examples for the (as of yet) unpublished EMS paper on object capabilities in our ZALF model and simulation infrastructure. These examples are not necessarily supposed to present the most secure and best way to design an API, rather to support the paper and show working examples. We try to add different implementations in different languages and maybe subsequently add more complete and closer to the real world examples. The currently incomplete real infrastructure code can be found https://github.com/zalf-rpm/mas-infrastructure.

## Use & Compilation
Different languages have different needs. Except for Python all languages need the schema files in the capnp folder to be processed by the Capn'n Proto schema compiler to generate the actual languages code. The following sections will try to list what's needed to run and/or compile the according languages code.

The different implementations should be interchangeable. 

### Python

requirements
- Python3
- install pycapnp via pip

### C#

requirements
- dotnet 6.0

### Go

requirements

### C++

requirements
- cmake


## Examples

1. classic Alice, Bob and Carol interaction
2. additional to 1. Alice sends to Bob a revokable capability to Carol


