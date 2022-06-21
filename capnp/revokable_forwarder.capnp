@0xff44938a6ce76f11;

using Cxx = import "c++.capnp";
$Cxx.namespace("ems_ocaps_paper::schema");

using Go = import "go.capnp";
$Go.package("capnp");
$Go.import("github.com/zalf-rpm/EMS_ocaps_paper_2022_code/go/capnp");

interface Actor {
    act @0 (msg :Text);
}

interface Alice extends(Actor) {
    setBobAndCarol  @0 (bob :Bob, carol :Carol);

    revokeCarol     @1 ();
} 

interface Bob extends(Actor) {
    foo @0 (carol :Carol);
}

interface Carol extends(Actor) {
}

interface Forwarder extends(Carol) {
    # forward Carol messages

    setActor @0 (a :Actor);
    # set the Actor we want to forward to
}

interface Revoker extends(Forwarder) {
    # revoke forwarder messages

    revoke @0 ();
    # revoke any further forwarding
}

