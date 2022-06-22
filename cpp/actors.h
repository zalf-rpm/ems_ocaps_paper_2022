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
#include <thread>

#include <kj/common.h>
#include <kj/main.h>
#include <kj/string.h>

#include "revokable_forwarder.capnp.h"

class Forwarder;
class Revoker;
class Bob;
class Carol;

class Alice final : public ems_ocaps_paper::schema::Alice::Server
{
public:
  Alice() : _forwarder(kj::heap<Forwarder>()) {
    auto f = kj::heap<Forwarder>();
    _forwarder = kj::mv(f);
    auto r = kj::heap<Revoker>();
    _revoker = kj::mv(r);
  }

  virtual ~Alice() noexcept(false) {}

  // act @0 (msg :Text);
  kj::Promise<void> act(ActContext context) override {
    std::cout << "@Alice::act | msg:" << context.getParams().getMsg().cStr() << std::endl;
    std::cout << "@Alice::act | sending foo(forwarder) to Bob" << std::endl;
    auto fr = _bob.fooRequest();
    fr.setCarol(_forwarder);
    return fr.send().ignoreResult();
  }

  // setBobAndCarol  @0 (bob :Bob, carol :Carol);
  kj::Promise<void> setBobAndCarol(SetBobAndCarolContext context) override {
    std::cout << "@Alice::setBobAndCarol" << std::endl;
    _bob = context.getParams().getBob();
    _carol = context.getParams().getCarol();
    auto r1 = _forwarder.setActorRequest();
    r1.setA(_revoker);
    auto r2 = _revoker.setActorRequest();
    r2.setA(_carol);
    return kj::joinPromises(kj::heapArray({r1.send().ignoreResult(), r2.send().ignoreResult()})).ignoreResult();
  }

  // revokeCarol     @1 ();
  kj::Promise<void> revokeCarol(RevokeCarolContext context) override {
    return _revoker.revokeRequest().send().ignoreResult();
  }

private:
  ems_ocaps_paper::schema::Bob::Client _bob{nullptr};;
  ems_ocaps_paper::schema::Carol::Client _carol{nullptr};;
  ems_ocaps_paper::schema::Forwarder::Client _forwarder;
  ems_ocaps_paper::schema::Revoker::Client _revoker{nullptr};;
};

//-----------------------------------------------------------------------------------

class Bob final : public ems_ocaps_paper::schema::Bob::Server
{
public:
  Bob() {}

  virtual ~Bob() noexcept(false) {}

  // act @0 (msg :Text);
  kj::Promise<void> act(ActContext context) override {
    std::cout << "@Bob::act | msg:" << context.getParams().getMsg().cStr() << std::endl;
    std::cout << "@Bob::act | sending act(<Bobs ACT message to Carol>) to Carol" << std::endl;
    auto cr = _carol.actRequest();
    cr.setMsg("<Bobs " + std::to_string(_count) + ". ACT message to Carol>");
    return cr.send().ignoreResult();
  }

  // # foo @1 (carol :Carol);
  kj::Promise<void> foo(FooContext context) override {
    std::cout << "@Bob::foo" << std::endl;
    _carol = context.getParams().getCarol();
    return kj::READY_NOW;
  }

private:
  ems_ocaps_paper::schema::Actor::Client _carol{nullptr};
  int _count {1};
};

//-----------------------------------------------------------------------------------

class Carol final : public ems_ocaps_paper::schema::Carol::Server
{
public:
  Carol() {}

  virtual ~Carol() noexcept(false) {}

  // act @0 (msg :Text);
  kj::Promise<void> act(ActContext context) override {
    std::cout << "@Carol::do | msg:" << context.getParams().getMsg().cStr() << std::endl;
    return kj::READY_NOW;
  }

private:
};

//-----------------------------------------------------------------------------------

class Forwarder final : public ems_ocaps_paper::schema::Forwarder::Server
{
public:
  Forwarder() {}

  virtual ~Forwarder() noexcept(false) {}

  // act @0 (msg :Text);
  kj::Promise<void> act(ActContext context) override {
    auto msg = context.getParams().getMsg();
    std::cout << "@Forwarder::act | msg:" << msg.cStr() << std::endl;
    std::cout << "@Forwarder::act | forwarding act(" << msg.cStr() << ") to attached actor" << std::endl;
    auto ar = _actor.actRequest();
    ar.setMsg(msg);
    return ar.send().ignoreResult();
  }

  // setActor @0 (a :Actor);
  kj::Promise<void> setActor(SetActorContext context) override {
    std::cout << "@Forwarder::setActor" << std::endl;
    _actor = context.getParams().getA();
    return kj::READY_NOW;
  }

private:
  ems_ocaps_paper::schema::Actor::Client _actor{nullptr};
};

//-----------------------------------------------------------------------------------

class Revoker final : public ems_ocaps_paper::schema::Revoker::Server
{
public:
  Revoker() {}

  virtual ~Revoker() noexcept(false) {}

  // act @0 (msg :Text);
  kj::Promise<void> act(ActContext context) override {
    auto msg = context.getParams().getMsg();
    std::cout << "@Forwarder::act | msg:" << msg.cStr() << std::endl;
    std::cout << "@Forwarder::act | forwarding act(" << msg.cStr() << ") to attached actor" << std::endl;
    auto ar = _actor.actRequest();
    ar.setMsg(msg);
    return ar.send().ignoreResult();
  }

  // setActor @0 (a :Actor);
  kj::Promise<void> setActor(SetActorContext context) override {
    std::cout << "@Forwarder::setActor" << std::endl;
    _actor = context.getParams().getA();
    return kj::READY_NOW;
  }

  // revoke @0 ();
  kj::Promise<void> revoke(RevokeContext context) override {
    std::cout << "@Revoker::revoke" << std::endl;
    _actor = nullptr;
    return kj::READY_NOW;
  }

private:
  ems_ocaps_paper::schema::Actor::Client _actor{nullptr};
};
