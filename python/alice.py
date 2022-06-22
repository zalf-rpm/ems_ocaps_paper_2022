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
from forwarder import Forwarder
from revoker import Revoker

#------------------------------------------------------------------------------

class PlainAlice(schema.Alice.Server):

    def __init__(self):
        self.bob = None
        self.carol = None

    def act_context(self, context): # act @0 (msg :Text);
        print("@Alice::act | msg:", context.params.msg)
        print("@Alice::act | sending foo(carol) to Bob")
        if self.bob and self.carol:
            return self.bob.foo(self.carol).then(lambda _: ())

    def setBobAndCarol_context(self, context): # setBobAndCarol @0 (bob :Bob, carol :Carol);
        print("@Alice::setBobAndCarol")
        self.bob = context.params.bob
        self.carol = context.params.carol

    def revokeCarol(self, **kwargs): # revokeCarol @1 ();
        pass

class Alice(schema.Alice.Server):

    def __init__(self):
        self.bob = None
        self.carol = None
        self.forwarder = schema.Forwarder._new_client(Forwarder())
        #self.forwarder = Forwarder()
        self.revoker = schema.Revoker._new_client(Revoker())
        #self.revoker = Revoker()

    def act_context(self, context): # act @0 (msg :Text);
        print("@Alice::act | msg:", context.params.msg)
        if self.bob and self.carol and self.forwarder:
            print("@Alice::act | sending foo(forwarder) to Bob")
            return self.bob.foo(self.forwarder).then(lambda _: ())

    def setBobAndCarol_context(self, context): # setBobAndCarol @0 (bob :Bob, carol :Carol);
        print("@Alice::setBobAndCarol")
        self.bob = context.params.bob
        self.carol = context.params.carol
        return capnp.join_promises([
            self.forwarder.setActor(self.revoker), 
            self.revoker.setActor(self.carol)
        ]).then(lambda _: ())

    def revokeCarol_context(self, context): # revokeCarol @1 ();
        print("@Alice::revokeCarol")
        return self.revoker.revoke().then(lambda _: ())

#------------------------------------------------------------------------------

if __name__ == '__main__':

    config = {
        "port": "9991",
        "use_asyncio": True,
    }
    common.update_config(config, sys.argv, print_config=False)

    print("@alice.py")
    if config["use_asyncio"]:
        asyncio.run(async_helpers.serve_forever(None, config["port"], Alice()))
    else: 
        server = capnp.TwoPartyServer("*:"+config["port"], bootstrap=Alice())
        capnp.wait_forever()
