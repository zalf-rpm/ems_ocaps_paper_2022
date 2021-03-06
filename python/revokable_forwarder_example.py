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

import asyncio
import capnp
import os
from pathlib import Path
import socket
import sys
import time
from threading import Thread
import subprocess as sp

PATH_TO_SCRIPT_DIR = Path(os.path.realpath(__file__)).parent
PATH_TO_REPO = PATH_TO_SCRIPT_DIR.parent
if str(PATH_TO_REPO) not in sys.path:
    sys.path.insert(1, str(PATH_TO_REPO))
schema = capnp.load(str(PATH_TO_REPO / "capnp" / "revokable_forwarder.capnp"))

import helper.capnp_async_helpers as async_helpers

# for the multithreading example we need support to for eventloops in every single thread
# this removes the old single threaded one and creates a new one for the main thread with multithreading support
capnp.reset_event_loop(ignore_errors=False, threaded=True)

from alice import Alice, PlainAlice
from bob import Bob
from carol import Carol

def run_alice(s, aio=False):
    print("@run_alice")
    if aio:
        asyncio.run(async_helpers.serve_forever(None, s[2:], Alice()))
    else: 
        server = capnp.TwoPartyServer(s, bootstrap=Alice())
        capnp.wait_forever()


def run_bob(s, aio=False):
    print("@run_bob")
    if aio:
        asyncio.run(async_helpers.serve_forever(None, s[2:], Bob()))
    else:
        server = capnp.TwoPartyServer(s, bootstrap=Bob())
        capnp.wait_forever()


def run_carol(s, aio=False):
    print("@run_carol")
    if aio:
        asyncio.run(async_helpers.serve_forever(None, s[2:], Carol()))
    else:
        server = capnp.TwoPartyServer(s, bootstrap=Carol())
        capnp.wait_forever()


if __name__ == '__main__':

    config = {
        "actor": "aio",
        "use_threads": False,
        "use_asyncio": True,
    }
    if len(sys.argv) > 1:
        for arg in sys.argv[1:]:
            k, v = arg.split("=")
            config[k] = v.lower() == "true" if v.lower() in ["true", "false"] else v 
    #print(config)

    actor = config["actor"]
    threaded = actor == "aio" and config["use_threads"]
    is_asyncio = config["use_asyncio"] and not config["use_threads"]
    as1, as2 = socket.socketpair() if threaded else ("localhost:9991", "*:9991")
    bs1, bs2 = socket.socketpair() if threaded else ("localhost:9992", "*:9992")
    cs1, cs2 = socket.socketpair() if threaded else ("localhost:9993", "*:9993")

    if actor == "alice":
        run_alice(as2, is_asyncio)
    elif actor == "bob":
        run_bob(bs2, is_asyncio)
    elif actor == "carol":
        run_carol(cs2, is_asyncio)
    elif actor == "main":
        alice = capnp.TwoPartyClient(as1).bootstrap().cast_as(schema.Alice)
        bob = capnp.TwoPartyClient(bs1).bootstrap().cast_as(schema.Bob)
        carol = capnp.TwoPartyClient(cs1).bootstrap().cast_as(schema.Carol)

        print("@main | sleep for 1s")
        time.sleep(1)
        
        print("@main | sending setBobAndCarol(bob,carol) to Alice")
        alice.setBobAndCarol(bob, carol).wait()
        print("@main | sending act(msg) to Alice")
        alice.act("<mains ACT message to Alice>").wait()

        print("@main | sending act(msg) to Bob")
        bob.act("<mains 1. ACT message to Bob>").wait()
        bob.act("<mains 2. ACT message to Bob>").wait()

        print("@main | sending revokeCarol to Alice")
        alice.revokeCarol().wait()

        print("@main | sending act(msg) to Bob")
        try:
            bob.act("<mains 3. ACT message to Bob>").wait()
        except Exception:
            print("@main | Couldn't send msg: <mains 3. ACT message to Bob>")

        time.sleep(2)
        print("@main | finished")

    elif actor == "aio":
        # start Alice
        if threaded:
            alice_thread = Thread(name="Alice", target=run_alice, args=[as2], daemon=True)
            alice_thread.start()
        else:
            alice_process = sp.Popen(["python", "python/revokable_forwarder_example.py", "actor=alice", "use_asyncio="+str(is_asyncio)])

        # start Bob
        if threaded:
            bob_thread = Thread(name="Bob", target=run_bob, args=[bs2], daemon=True)
            bob_thread.start()
        else:
            bob_process = sp.Popen(["python", "python/revokable_forwarder_example.py", "actor=bob", "use_asyncio="+str(is_asyncio)])

        # start Carol
        if threaded:
            carol_thread = Thread(name="Carol", target=run_carol, args=[cs2], daemon=True)
            carol_thread.start()
        else:
            carol_process = sp.Popen(["python", "python/revokable_forwarder_example.py", "actor=carol", "use_asyncio="+str(is_asyncio)])
        

        # wait for threads/processes to have started
        time.sleep(1)
        
        # get references to Actors
        alice = capnp.TwoPartyClient(as1).bootstrap().cast_as(schema.Alice)
        bob = capnp.TwoPartyClient(bs1).bootstrap().cast_as(schema.Bob)
        carol = capnp.TwoPartyClient(cs1).bootstrap().cast_as(schema.Carol)
        
        print("@main | sending set(bob,carol) to Alice")
        alice.set(bob, carol).wait()
        print("@main | sending act(msg) to Alice")
        alice.act("<mains ACT message to Alice>").wait()

        print("@main | sending act(msg) to Bob")
        bob.act("<mains ACT message to Bob>").wait()

        time.sleep(2)
        if not threaded:
            alice_process.terminate()
            bob_process.terminate()
            carol_process.terminate()
        print("@main | finished")







