#!/usr/bin/python
# -*- coding: UTF-8

# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at http://mozilla.org/MPL/2.0/. */

# Authors:
# Michael Berg-Mohnicke <michael.berg@zalf.de>
#
# Maintainers:
# Currently maintained by the authors.
#
# This file is part of the util library used by models created at the Institute of
# Landscape Systems Analysis at the ZALF.
# Copyright (C: Leibniz Centre for Agricultural Landscape Research (ZALF)

import subprocess as sp
import sys
import time

def update_config(config, argv, print_config=False, allow_new_keys=False):
    if len(argv) > 1:
        for arg in argv[1:]:
            k, v = arg.split("=")
            if not allow_new_keys and k in config:
                config[k] = v.lower() == "true" if v.lower() in ["true", "false"] else v 
        if print_config:
            print(config)

config = {
    "a": "cpp", # cpp | c# | (go) | py
    "b": "cpp", # cpp | c# | (go) | py
    "c": "cpp", # cpp | c# | (go) | py
    "m": "cpp", # cpp | c# | (go) | py
    "py_use_asyncio": "true",
}
update_config(config, sys.argv, print_config=False)

if config["a"] == "cpp":
    alice_process = sp.Popen(["cpp/_cmake_debug/revokable_forwarder_example", "--actor=alice"])
elif config["a"] == "c#":
    alice_process = sp.Popen(["csharp/bin/Debug/net6.0/revokable_forwarder_example", "actor=alice"])
elif config["a"] == "py":
    alice_process = sp.Popen(["python", "python/alice.py", "use_asyncio="+config["py_use_asyncio"]])

if config["b"] == "cpp":
    bob_process = sp.Popen(["cpp/_cmake_debug/revokable_forwarder_example", "--actor=bob"])
elif config["b"] == "c#":
    bob_process = sp.Popen(["csharp/bin/Debug/net6.0/revokable_forwarder_example", "actor=bob"])
elif config["b"] == "py":
    bob_process = sp.Popen(["python", "python/bob.py", "use_asyncio="+config["py_use_asyncio"]])

if config["c"] == "cpp":
    carol_process = sp.Popen(["cpp/_cmake_debug/revokable_forwarder_example", "--actor=carol"])
elif config["c"] == "c#":
    carol_process = sp.Popen(["csharp/bin/Debug/net6.0/revokable_forwarder_example", "actor=carol"])
elif config["c"] == "py":
    carol_process = sp.Popen(["python", "python/carol.py", "use_asyncio="+config["py_use_asyncio"]])
time.sleep(2)

if config["m"] == "cpp":
    main_process = sp.Popen(["cpp/_cmake_debug/revokable_forwarder_example", "--actor=main"])
elif config["m"] == "c#":
    main_process = sp.Popen(["csharp/bin/Debug/net6.0/revokable_forwarder_example", "actor=main"])
elif config["m"] == "py":
    main_process = sp.Popen(["python", "python/revokable_forwarder_example.py", "actor=main", "use_asyncio="+config["py_use_asyncio"]])

main_process.wait()

alice_process.terminate()
bob_process.terminate()
carol_process.terminate()
print("finished")






