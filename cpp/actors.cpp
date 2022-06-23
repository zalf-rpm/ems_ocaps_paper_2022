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

#include "actors.h"

#include <iostream>
#include <string>
#include <vector>
#include <thread>

#include <kj/common.h>
#include <kj/main.h>
#include <kj/string.h>

#include "revokable_forwarder.capnp.h"
#include "common.h"

Alice::Alice() 
: _forwarder(kj::heap<Forwarder>())
, _revoker(kj::heap<Revoker>())
{}

kj::Promise<void> Alice::act(ActContext context) {
  std::cout << "@Alice::act | msg:" << context.getParams().getMsg().cStr() << std::endl;
  std::cout << "@Alice::act | sending foo(forwarder) to Bob" << std::endl;
  auto fr = _bob.fooRequest();
  fr.setCarol(_forwarder);
  return fr.send().ignoreResult();
}

kj::Promise<void> Alice::setBobAndCarol(SetBobAndCarolContext context) {
  std::cout << "@Alice::setBobAndCarol" << std::endl;
  _bob = context.getParams().getBob();
  _carol = context.getParams().getCarol();
  auto r1 = _forwarder.setActorRequest();
  r1.setA(_revoker);
  auto r2 = _revoker.setActorRequest();
  r2.setA(_carol);
  auto ps = kj::heapArrayBuilder<kj::Promise<void>>(2);
  ps.add(r1.send().ignoreResult());
  ps.add(r2.send().ignoreResult());
  return kj::joinPromises(ps.finish());
}

kj::Promise<void> Alice::revokeCarol(RevokeCarolContext context) {
  return _revoker.revokeRequest().send().ignoreResult();
}

//-----------------------------------------------------------------------------------

kj::Promise<void> Bob::act(ActContext context) {
  std::cout << "@Bob::act | msg:" << context.getParams().getMsg().cStr() << std::endl;
  std::cout << "@Bob::act | sending act(<Bobs ACT message to Carol>) to Carol" << std::endl;
  auto cr = _carol.actRequest();
  cr.setMsg("<Bobs " + std::to_string(_count++) + ". ACT message to Carol>");
  return cr.send().ignoreResult();
}

kj::Promise<void> Bob::foo(FooContext context) {
  std::cout << "@Bob::foo" << std::endl;
  _carol = context.getParams().getCarol();
  return kj::READY_NOW;
}

//-----------------------------------------------------------------------------------

kj::Promise<void> Carol::act(ActContext context) {
  std::cout << "@Carol::do | msg:" << context.getParams().getMsg().cStr() << std::endl;
  return kj::READY_NOW;
}

//-----------------------------------------------------------------------------------

kj::Promise<void> Forwarder::act(ActContext context) {
  auto msg = context.getParams().getMsg();
  std::cout << "@Forwarder::act | msg:" << msg.cStr() << std::endl;
  std::cout << "@Forwarder::act | forwarding act(" << msg.cStr() << ") to attached actor" << std::endl;
  auto ar = _actor.actRequest();
  ar.setMsg(msg);
  return ar.send().ignoreResult();
}

kj::Promise<void> Forwarder::setActor(SetActorContext context) {
  std::cout << "@Forwarder::setActor" << std::endl;
  _actor = context.getParams().getA();
  return kj::READY_NOW;
}

//-----------------------------------------------------------------------------------

kj::Promise<void> Revoker::act(ActContext context) {
  auto msg = context.getParams().getMsg();
  std::cout << "@Revoker::act | msg:" << msg.cStr() << std::endl;
  std::cout << "@Revoker::act | forwarding act(" << msg.cStr() << ") to attached actor" << std::endl;
  auto ar = _actor.actRequest();
  ar.setMsg(msg);
  return ar.send().ignoreResult();
}

kj::Promise<void> Revoker::setActor(SetActorContext context) {
  std::cout << "@Revoker::setActor" << std::endl;
  _actor = context.getParams().getA();
  return kj::READY_NOW;
}

kj::Promise<void> Revoker::revoke(RevokeContext context) {
  std::cout << "@Revoker::revoke" << std::endl;
  _actor = nullptr;
  return kj::READY_NOW;
}
