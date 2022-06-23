/* This Source Code Form is subject to the terms of the Mozilla Public
* License, v. 2.0. If a copy of the MPL was not distributed with this
* file, You can obtain one at http://mozilla.org/MPL/2.0/. */

/*
Authors:
Michael Berg <michael.berg-mohnicke@zalf.de>

Maintainers:
Currently maintained by the authors.

This file is part of the MONICA model.
Copyright (C) Leibniz Centre for Agricultural Landscape Research (ZALF)
*/

#pragma once

#include <kj/debug.h>
#include <kj/common.h>
#include <kj/string.h>
#include <kj/vector.h>
#include <kj/map.h>

#include <capnp/rpc-twoparty.h>
#include <kj/thread.h>

#include <functional>
#include <string>
#include <vector>

//#include "model.capnp.h"
#include "common.capnp.h"
#include "persistence.capnp.h"

namespace mas {
  namespace rpc {
    namespace common {

      class Restorer final : public mas::schema::persistence::Restorer::Server
      {
        public:
        Restorer() {}

        virtual ~Restorer() noexcept(false) {}

        // restore @0 (srToken :Text) -> (cap :Capability);
        kj::Promise<void> restore(RestoreContext context) override;

        int getPort() const { return _port; }
        void setPort(int p) { _port = p; }

        std::string getHost() const { return _host; }
        void setHost(std::string h) { _host = h; }

        std::string sturdyRef(std::string srToken = "") const;

        std::pair<std::string, std::string> save(capnp::Capability::Client cap, std::string srToken = std::string(),
          bool createUnsave = true);

        void unsave(std::string srToken);

      private:
        std::string _host{ "" };
        int _port{ 0 };
        kj::HashMap<kj::String, capnp::Capability::Client> _issuedSRTokens;
        std::vector<std::function<void()>> _actions;
      };

      //-----------------------------------------------------------------------------

      class Action final : public mas::schema::common::Action::Server {
      public:
        Action(std::function<void()> action, 
                    bool execActionOnDel = false,
                    std::string id = "<-");

        virtual ~Action() noexcept(false);

        kj::Promise<void> do_(DoContext context) override;

      private:
        std::string id{ "<-" };
        std::function<void()> action;
        bool execActionOnDel{ false };
        bool alreadyCalled{ false };
      };
    } 
  }
}
