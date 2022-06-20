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
        self.revoker = None

    def do(self, **kwargs): # do @0 ();
        pass

    def set_context(self, context): # set @0 (r :Revoker);
        print("@Forwarder::set")
        self.revoker = context.params.r

class Revoker(schema.Revoker.Server):

    def __init__(self):
        self.carol = None

    def do(self, **kwargs): # do @0 ();
        pass

    def set_context(self, context): # set @0 (carol :Carol);
        print("@Revoker::set")
        self.carol = context.params.carol




