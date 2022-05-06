# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at http://mozilla.org/MPL/2.0/. */

# Authors:
# Michael Berg-Mohnicke <michael.berg-mohnicke@zalf.de>
#
# Maintainers:
# Currently maintained by the authors.
#
# This file has been created at the Institute of
# Landscape Systems Analysis at the ZALF.
# Copyright (C: Leibniz Centre for Agricultural Landscape Research (ZALF)

import capnp
import os
from pathlib import Path
import socket
import sys
import time
import uuid

PATH_TO_REPO = Path(os.path.realpath(__file__)).parent.parent.parent
if str(PATH_TO_REPO) not in sys.path:
    sys.path.insert(1, str(PATH_TO_REPO))
persistence_capnp = capnp.load(str(PATH_TO_REPO / "capnp" / "persistence.capnp")) 
common_capnp = capnp.load(str(PATH_TO_REPO / "capnp" / "common.capnp")) 

#------------------------------------------------------------------------------

def update_config(config, argv, print_config=False, allow_new_keys=False):
    if len(argv) > 1:
        for arg in argv[1:]:
            k, v = arg.split("=")
            if not allow_new_keys and k in config:
                config[k] = v.lower() == "true" if v.lower() in ["true", "false"] else v 
        if print_config:
            print(config)

#------------------------------------------------------------------------------

class Restorer(persistence_capnp.Restorer.Server):

    def __init__(self):
        self._issued_sr_tokens = {} # sr_token to capability
        self._actions = []
        self._host = socket.gethostbyname(socket.gethostname()) #socket.getfqdn() #gethostname()
        self._port = None


    @property
    def port(self):
        return self._port
    
    @port.setter
    def port(self, p):
        self._port = p


    @property
    def host(self):
        return self._host
    
    @host.setter
    def host(self, h):
        self._host = h


    def sturdy_ref(self, sr_token=None):
        if sr_token:
            return "capnp://insecure@{host}:{port}/{sr_token}".format(host=self.host, port=self.port, sr_token=sr_token)
        else:
            return "capnp://insecure@{host}:{port}".format(host=self.host, port=self.port)


    def save(self, cap, sr_token=None, create_unsave=True):
        sr_token = sr_token if sr_token else str(uuid.uuid4())
        self._issued_sr_tokens[sr_token] = cap
        if create_unsave:
            unsave_sr_token = str(uuid.uuid4())
            unsave_action = Action(lambda: [self.unsave(sr_token), self.unsave(unsave_sr_token)]) 
            self._issued_sr_tokens[unsave_sr_token] = unsave_action
        return (self.sturdy_ref(sr_token), self.sturdy_ref(unsave_sr_token) if create_unsave else None)


    def unsave(self, sr_token): 
        if sr_token in self._issued_sr_tokens:
            del self._issued_sr_tokens[sr_token]


    def restore_context(self, context): # restore @0 (srToken :Text) -> (cap :Capability);
        srt = context.params.srToken
        if srt in self._issued_sr_tokens:
            context.results.cap = self._issued_sr_tokens[srt]

#------------------------------------------------------------------------------

class Persistable(persistence_capnp.Persistent.Server):

    def __init__(self, restorer=None):
        self._restorer = restorer


    @property
    def restorer(self):
        return self._restorer

    @restorer.setter
    def restorer(self, r):
        self._restorer = r


    def save_context(self, context): # save @0 () -> (sturdyRef :Text, unsaveSR :Text);
        if self.restorer:
            sr, unsave_sr = self.restorer.save(self)
            context.results.sturdyRef = sr
            context.results.unsaveSR = unsave_sr

#------------------------------------------------------------------------------

class ConnectionManager:

    def __init__(self):
        self._connections = {}


    def connect(self, sturdy_ref, cast_as = None):

        # we assume that a sturdy ref url looks always like capnp://hash-digest-or-insecure@host:port/sturdy-ref-token
        if sturdy_ref[:8] == "capnp://":
            rest = sturdy_ref[8:]
            hash_digest, rest = rest.split("@") if "@" in rest else (None, rest)
            host, rest = rest.split(":")
            port, sr_token = rest.split("/") if "/" in rest else (rest, None)

            host_port = "{}:{}".format(host, port)
            if host_port in self._connections:
                bootstrap_cap = self._connections[host_port]
            else:
                bootstrap_cap = capnp.TwoPartyClient(host_port).bootstrap()
                self._connections[host_port] = bootstrap_cap

            if sr_token:
                restorer = bootstrap_cap.cast_as(persistence_capnp.Restorer)
                dyn_obj_reader = restorer.restore(sr_token).wait().cap
                return dyn_obj_reader.as_interface(cast_as) if cast_as else dyn_obj_reader
            else:
                return bootstrap_cap.cast_as(cast_as) if cast_as else bootstrap_cap


    def try_connect(self, sturdy_ref, cast_as = None, retry_count=10, retry_secs=5, print_retry_msgs=True):
        while True:
            try:
                return self.connect(sturdy_ref, cast_as=cast_as)
            except Exception as e:
                print(e)
                if retry_count == 0:
                    if print_retry_msgs:
                        print("Couldn't connect to sturdy_ref at {}!".format(sturdy_ref))
                    return None
                retry_count -= 1
                if print_retry_msgs:
                    print("Trying to connect to {} again in {} secs!".format(sturdy_ref, retry_secs))
                time.sleep(retry_secs)
                retry_secs += 1

#------------------------------------------------------------------------------

class Action(common_capnp.Action.Server):

    def __init__(self, action, *args, exec_action_on_del=False, **kwargs):
        self._args = args
        self._kwargs = kwargs
        self._action = action
        self._already_executed = False
        self._exec_action_on_del = exec_action_on_del

    def __del__(self):
        if self._exec_action_on_del and not self._already_executed:
            self._action(*self._args, **self._kwargs)

    def do_context(self, context): # do @0 () -> ();
        self._action(*self._args, **self._kwargs)
        self._already_executed = True

#------------------------------------------------------------------------------
