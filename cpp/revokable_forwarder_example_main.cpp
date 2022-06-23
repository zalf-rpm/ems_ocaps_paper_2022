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
#include <chrono>
#include <thread>

#include <kj/debug.h>
#include <kj/common.h>
#include <kj/main.h>
#include <kj/string.h>

#include "rpc-connections.h"
#include "revokable_forwarder.capnp.h"
#include "actors.h"

class Main
{
public:
  Main(kj::ProcessContext &context) : context(context), ioContext(kj::setupAsyncIo()) {}

  kj::MainBuilder::Validity setHost(kj::StringPtr name) { host = name; return true; }

  kj::MainBuilder::Validity setPort(kj::StringPtr name) { port = std::max(0, std::stoi(name.cStr())); return true; }

  kj::MainBuilder::Validity setActor(kj::StringPtr name) { actor = name; return true; }

  kj::MainBuilder::Validity startAlice()
  {
    capnp::Capability::Client a(nullptr);
    if(actor == "alice") { a = kj::heap<Alice>(); port = port == 0 ? 9991 : 0; }
    if(actor == "bob") { a = kj::heap<Bob>(); port = port == 0 ? 9992 : 0; }
    if(actor == "carol") { a = kj::heap<Carol>(); port = port == 0 ? 9993 : 0; }
    if(actor == "main"){
      auto alice = conMan.tryConnectB(ioContext, "capnp://insecure@localhost:9991").castAs<ems_ocaps_paper::schema::Alice>();
      auto bob = conMan.tryConnectB(ioContext, "capnp://insecure@localhost:9992").castAs<ems_ocaps_paper::schema::Bob>();
      auto carol = conMan.tryConnectB(ioContext, "capnp://insecure@localhost:9993").castAs<ems_ocaps_paper::schema::Carol>();

      std::cout << "@main | sleep for 1s" << std::endl;
      std::this_thread::sleep_for(std::chrono::milliseconds(1000));

      std::cout << "@main | sending setBobAndCarol(bob, carol) to Alice" << std::endl;
      auto sbacReq = alice.setBobAndCarolRequest();
      sbacReq.setBob(bob);
      sbacReq.setCarol(carol);
      sbacReq.send().ignoreResult();
      std::cout << "@main | sending act(msg) to Alice" << std::endl;
      auto aaReq = alice.actRequest();
      aaReq.setMsg("<mains ACT message to Alice>");
      aaReq.send().wait(ioContext.waitScope);

      std::cout << "@main | sending act(msg) to Bob" << std::endl;
      auto baReq1 = bob.actRequest();
      baReq1.setMsg("<mains 1. ACT message to Bob>");
      baReq1.send().wait(ioContext.waitScope);
      auto baReq2 = bob.actRequest();
      baReq2.setMsg("<mains 2. ACT message to Bob>");
      baReq2.send().wait(ioContext.waitScope);

      std::cout << "@main | sending revokeCarol to Alice" << std::endl;
      alice.revokeCarolRequest().send().wait(ioContext.waitScope);

      std::cout << "@main | sending act(msg) to Bob" << std::endl;
      auto baReq3 = bob.actRequest();
      baReq3.setMsg("<mains 3. ACT message to Bob>");
      try {
        baReq3.send().wait(ioContext.waitScope);
      } catch(kj::Exception e) {
        std::cout << "@main | Couldn't send msg: <mains 3. ACT message to Bob>" << std::endl;
      }

      std::this_thread::sleep_for(std::chrono::milliseconds(2000));
      std::cout << "@main | finished" << std::endl;

      return true;
    } 
    KJ_LOG(INFO, "starting", actor, host, port);
    auto proms = conMan.bind(ioContext, a, host, port);
    kj::NEVER_DONE.wait(ioContext.waitScope);
    return true;
  }

  kj::MainFunc getMain()
  {
    return kj::MainBuilder(context, "Alice v0.1", "")
      .addOptionWithArg({'h', "host"}, KJ_BIND_METHOD(*this, setHost),
                        "<host-IP *>", "Set host IP.")
      .addOptionWithArg({'p', "port"}, KJ_BIND_METHOD(*this, setPort),
                        "<port>", "Set port.")
      .addOptionWithArg({'a', "actor"}, KJ_BIND_METHOD(*this, setActor),
                        "<alice|bob|carol|main>", "Which actor to run.")
      .callAfterParsing(KJ_BIND_METHOD(*this, startAlice))
      .build();
  }

private:
  mas::infrastructure::common::ConnectionManager conMan;
  kj::StringPtr actor{"main"};
  kj::StringPtr host{"*"};
  int port{0};
  kj::ProcessContext &context;
  kj::AsyncIoContext ioContext;
};

KJ_MAIN(Main)
