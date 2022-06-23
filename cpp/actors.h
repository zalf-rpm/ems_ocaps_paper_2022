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

#pragma once

#include <kj/common.h>
#include <kj/main.h>
#include <kj/string.h>

#include "revokable_forwarder.capnp.h"

class Alice final : public ems_ocaps_paper::schema::Alice::Server
{
public:
  Alice();
  virtual ~Alice() noexcept(false) {}

  // act @0 (msg :Text);
  kj::Promise<void> act(ActContext context) override;

  // setBobAndCarol  @0 (bob :Bob, carol :Carol);
  kj::Promise<void> setBobAndCarol(SetBobAndCarolContext context) override;

  // revokeCarol     @1 ();
  kj::Promise<void> revokeCarol(RevokeCarolContext context) override;

private:
  ems_ocaps_paper::schema::Bob::Client _bob{nullptr};;
  ems_ocaps_paper::schema::Carol::Client _carol{nullptr};;
  ems_ocaps_paper::schema::Forwarder::Client _forwarder{nullptr};
  ems_ocaps_paper::schema::Revoker::Client _revoker{nullptr};
};

//-----------------------------------------------------------------------------------

class Bob final : public ems_ocaps_paper::schema::Bob::Server
{
public:
  Bob() {}
  virtual ~Bob() noexcept(false) {}

  // act @0 (msg :Text);
  kj::Promise<void> act(ActContext context) override;

  // # foo @1 (carol :Carol);
  kj::Promise<void> foo(FooContext context) override;

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
  kj::Promise<void> act(ActContext context) override;

private:
};

//-----------------------------------------------------------------------------------

class Forwarder final : public ems_ocaps_paper::schema::Forwarder::Server
{
public:
  Forwarder() {}
  virtual ~Forwarder() noexcept(false) {}

  // act @0 (msg :Text);
  kj::Promise<void> act(ActContext context) override;

  // setActor @0 (a :Actor);
  kj::Promise<void> setActor(SetActorContext context) override;

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
  kj::Promise<void> act(ActContext context) override;

  // setActor @0 (a :Actor);
  kj::Promise<void> setActor(SetActorContext context) override;

  // revoke @0 ();
  kj::Promise<void> revoke(RevokeContext context) override;

private:
  ems_ocaps_paper::schema::Actor::Client _actor{nullptr};
};
