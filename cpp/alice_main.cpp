/* This Source Code Form is subject to the terms of the Mozilla Public
* License, v. 2.0. If a copy of the MPL was not distributed with this
* file, You can obtain one at http://mozilla.org/MPL/2.0/. */

/*
Authors:
Michael Berg <michael.berg@zalf.de>

Maintainers:
Currently maintained by the authors.

This file is part of the MONICA model.
Copyright (C) Leibniz Centre for Agricultural Landscape Research (ZALF)
*/

#include <iostream>
#include <string>
#include <vector>
#include <chrono>
#include <thread>

#include <kj/debug.h>
#include <kj/common.h>
#include <kj/main.h>
#include <kj/string.h>

#include "rpc-connections.h"
#include "revokable_forwarder.capnp.h"
#include "actors.h"

class AliceMain
{
public:
  AliceMain(kj::ProcessContext &context) : context(context), ioContext(kj::setupAsyncIo()) {}

  kj::MainBuilder::Validity setHost(kj::StringPtr name) { host = name; return true; }

  kj::MainBuilder::Validity setPort(kj::StringPtr name) { port = std::max(0, std::stoi(name.cStr())); return true; }

  kj::MainBuilder::Validity startAlice()
  {
    auto alice = kj::heap<Alice>();
    auto proms = conMan.bind(ioContext, kj::mv(alice), host, port);
    kj::NEVER_DONE.wait(ioContext.waitScope);
    return true;
  }

  kj::MainFunc getMain()
  {
    return kj::MainBuilder(context, "Alice v0.1", "")
      .addOptionWithArg({'h', "host"}, KJ_BIND_METHOD(*this, setHost),
                        "<host-IP>", "Set host IP.")
      .addOptionWithArg({'p', "port"}, KJ_BIND_METHOD(*this, setPort),
                        "<port>", "Set port.")
      .callAfterParsing(KJ_BIND_METHOD(*this, startAlice))
      .build();
  }

private:
  mas::infrastructure::common::ConnectionManager conMan;
  kj::StringPtr name{"Alice"};
  kj::StringPtr host;
  int port{0};
  kj::ProcessContext &context;
  kj::AsyncIoContext ioContext;
};

KJ_MAIN(AliceMain)
