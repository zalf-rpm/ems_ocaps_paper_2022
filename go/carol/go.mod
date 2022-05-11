module github.com/zalf-rpm/EMS_ocaps_paper_2022_code/go/carol

go 1.17

require (
	capnproto.org/go/capnp/v3 v3.0.0-alpha.2
	github.com/zalf-rpm/ems_ocaps_paper_2022/go/capnp v0.0.0-00010101000000-000000000000
	github.com/zalf-rpm/ems_ocaps_paper_2022/go/helper v0.0.0-00010101000000-000000000000
)

require golang.org/x/sync v0.0.0-20201020160332-67f06af15bc9 // indirect

replace (
	github.com/zalf-rpm/ems_ocaps_paper_2022/go/capnp => ../capnp
	github.com/zalf-rpm/ems_ocaps_paper_2022/go/helper => ../helper
)
