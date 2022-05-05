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

import time
import subprocess as sp

#alice_process = sp.Popen(["csharp/bin/Debug/net6.0/revokable_forwarder_example", "actor=alice"])
alice_process = sp.Popen(["python", "python/revokable_forwarder_example.py", "actor=alice", "use_asyncio=true"])

bob_process = sp.Popen(["csharp/bin/Debug/net6.0/revokable_forwarder_example", "actor=bob"])
#bob_process = sp.Popen(["python", "python/revokable_forwarder_example.py", "actor=bob", "use_asyncio=true"])

#carol_process = sp.Popen(["csharp/bin/Debug/net6.0/revokable_forwarder_example", "actor=carol"])
carol_process = sp.Popen(["python", "python/revokable_forwarder_example.py", "actor=carol", "use_asyncio=true"])
time.sleep(2)

main_process = sp.Popen(["csharp/bin/Debug/net6.0/revokable_forwarder_example", "actor=main"])
#main_process = sp.Popen(["python", "python/revokable_forwarder_example.py", "actor=main", "use_asyncio=true"])
main_process.wait()

alice_process.terminate()
bob_process.terminate()
carol_process.terminate()
print("finished")






