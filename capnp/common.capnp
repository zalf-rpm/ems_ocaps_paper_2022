@0x99f1c9a775a88ac9;

using Cxx = import "c++.capnp";
$Cxx.namespace("mas::schema::common");

using Go = import "go.capnp";
$Go.package("common");
$Go.import("github.com/zalf-rpm/mas-infrastructure/capnp_schemas/gen/go/common");

using Persistent = import "persistence.capnp".Persistent;

interface Action extends(Persistent) {
  # interface to an arbitrary unparameterised action object

  do @0 () -> ();
  # execute the action represented by this object
  # any parameter can be null representing optional parameters
}


