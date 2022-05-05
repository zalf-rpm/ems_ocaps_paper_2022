@0xff44938a6ce76f11;

using Cxx = import "c++.capnp";
$Cxx.namespace("ems_ocaps_paper::schema");

using Go = import "go.capnp";
$Go.package("climate");
$Go.import("github.com/zalf-rpm/EMS_ocaps_paper_2022_code/go");

interface Actor {
    do @0 (msg :Text);
}

interface Alice extends(Actor) {
    set @0 (bob :Bob, carol :Carol);
} 

interface Bob extends(Actor) {
    foo @0 (carol :Carol);
}

interface Carol extends(Actor) {
}

interface Forwarder extends(Carol) {
    # forward Carol messages
}

interface Revoker extends(Forwarder) {
    # revoke forwarder messages
    revoke @0 ();
}

