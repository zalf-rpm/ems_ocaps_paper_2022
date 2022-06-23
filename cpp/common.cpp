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

#include "common.h"

#include <iostream>
#include <fstream>
#include <string>
#include <tuple>
#include <vector>
#include <algorithm>

#include <kj/debug.h>
#include <kj/thread.h>
#include <kj/common.h>
#define KJ_MVCAP(var) var = kj::mv(var)

#include <capnp/capability.h>
#include <capnp/ez-rpc.h>
#include <capnp/message.h>
#include <capnp/schema.h>
#include <capnp/dynamic.h>
#include <capnp/list.h>
#include <capnp/rpc-twoparty.h>

#include "sole.hpp"

//#include "tools/debug.h"
//#include "tools/date.h"

using namespace std;
//using namespace Tools;
using namespace mas::rpc::common;

//-----------------------------------------------------------------------------

kj::Promise<void> Restorer::restore(RestoreContext context) {
  auto srt = context.getParams().getSrToken();
  KJ_IF_MAYBE(cap, _issuedSRTokens.find(srt)) {
    context.getResults().setCap(*cap);
  }
  return kj::READY_NOW;
}

std::string Restorer::sturdyRef(std::string srToken) const {
  if(srToken.empty()) return "capnp://insecure@" + _host + ":" + to_string(_port);
  else return "capnp://insecure@" + _host + ":" + to_string(_port) + "/" + srToken;
}

std::pair<std::string, std::string> Restorer::save(capnp::Capability::Client cap, 
 std::string srToken, bool createUnsave) {
  if(srToken.empty()) srToken = sole::uuid4().str();
  _issuedSRTokens.insert(kj::str(srToken), cap);
  string unsaveSRToken = "";
  if(createUnsave)
  {
    unsaveSRToken = sole::uuid4().str();
    auto unsaveAction = kj::heap<Action>([this, srToken, unsaveSRToken]() { unsave(srToken); unsave(unsaveSRToken); }); 
    schema::common::Action::Client unsaveActionClient = kj::mv(unsaveAction);
    _issuedSRTokens.insert(kj::str(unsaveSRToken), unsaveActionClient);
  }
  return make_pair(sturdyRef(srToken), unsaveSRToken.empty() ? "" : sturdyRef(unsaveSRToken));
}

void Restorer::unsave(std::string srToken) {
  _issuedSRTokens.erase(srToken.c_str());
}

//-----------------------------------------------------------------------------

Action::Action(std::function<void()> action, 
                           bool execActionOnDel,
                           std::string id)
  : action(kj::mv(action))
  , execActionOnDel(execActionOnDel)
  , id(id) {}

Action::~Action() noexcept(false) {
  if (execActionOnDel && !alreadyCalled)
    action();
}

kj::Promise<void> Action::do_(DoContext context) {
  action();
  alreadyCalled = true;
  return kj::READY_NOW;
}

//-----------------------------------------------------------------------------



