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

PATH_TO_REPO = Path(os.path.realpath(__file__)).parent.parent
if str(PATH_TO_REPO) not in sys.path:
    sys.path.insert(1, str(PATH_TO_REPO))
schema = capnp.load(str(PATH_TO_REPO / "capnp" / "revokable_forwarder.capnp"))

import helper.capnp_async_helpers as async_helpers
import helper.common as common

#------------------------------------------------------------------------------

class Forwarder(schema.Forwarder.Server):

    def __init__(self):
        self.actor = None

    def act_context(self, context): # act @0 (msg :Text);
        msg = context.params.msg
        print("@Forwarder::act | msg:", msg)
        if self.actor:
            print("@Forwarder::act | forwarding act(", msg, ") to attached actor")
            return self.actor.act(context.params.msg).then(lambda _: ())

    def setActor_context(self, context): # setActor @0 (a :Actor);
        print("@Forwarder::setActor")
        self.actor = context.params.a

#------------------------------------------------------------------------------

if __name__ == '__main__':

    config = {
        "port": "9994",
        "use_asyncio": True,
    }
    common.update_config(config, sys.argv, print_config=False)

    print("@forwarder.py")
    if config["use_asyncio"]:
        asyncio.run(async_helpers.serve_forever(None, config["port"], Forwarder()))
    else: 
        server = capnp.TwoPartyServer("*:"+config["port"], bootstrap=Forwarder())
        capnp.wait_forever()



