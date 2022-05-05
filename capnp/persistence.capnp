@0x855efed3475f6b26;

#using Persistent = import "/capnp/persistent.capnp".Persistent;

using Cxx = import "c++.capnp";
$Cxx.namespace("mas::schema::persistence");

using Go = import "go.capnp";
$Go.package("persistence");
$Go.import("github.com/zalf-rpm/mas-infrastructure/capnp_schemas/gen/go/persistence");

interface Persistent {
  # simplified version of persistent.capnp::Persistent interface

  save @0 () -> (sturdyRef :Text, unsaveSR :Text);
  # create a sturdy ref to be able to restore this object and 
  # optionally return another SR refering to a Common.Action object representing the action to unsave this object
}

interface Restorer {
  # restore a capability from a sturdy ref

  restore @0 (srToken :Text) -> (cap :Capability);
  # restore from the given sturdy ref token a live capability
}

