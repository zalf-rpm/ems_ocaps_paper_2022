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
import sys

PATH_TO_REPO = Path(os.path.realpath(__file__)).parent.parent
if str(PATH_TO_REPO) not in sys.path:
    sys.path.insert(1, str(PATH_TO_REPO))
schema = capnp.load(str(PATH_TO_REPO / "capnp" / "revokable_forwarder.capnp"))

import helper.capnp_async_helpers as async_helpers
import helper.common as common

#------------------------------------------------------------------------------

class Bob(schema.Bob.Server):

    def __init__(self):
        self.carol = None
        self.count = 0

    def act_context(self, context): # act @0 (msg :Text);
        print("@Bob::act | msg:", context.params.msg)
        if self.carol:
            print("@Bob::act | sending act(<Bobs ACT message to Carol>) to Carol")
            self.count += 1
            return self.carol.act("<Bobs " + str(self.count) + ". ACT message to Carol>").then(lambda _: ())

    def foo_context(self, context): # foo @1 (carol :Carol);
        print("@Bob::foo")
        self.carol = context.params.carol

#------------------------------------------------------------------------------

if __name__ == '__main__':

    config = {
        "port": "9992",
        "use_asyncio": True,
    }
    common.update_config(config, sys.argv, print_config=False)

    print("@bob.py")
    if config["use_asyncio"]:
        asyncio.run(async_helpers.serve_forever(None, config["port"], Bob()))
    else: 
        server = capnp.TwoPartyServer("*:"+config["port"], bootstrap=Bob())
        capnp.wait_forever()

