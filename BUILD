go_mod(
    name="root",
)

oci_pull_image(
    name="static-debian12",
    repository="gcr.io/distroless/static-debian12",
    digest="5c7e2b465ac6a2a4e5f4f7f722ce43b147dabe87cb21ac6c4007ae5178a1fa58",
    architecture="amd64",
    os="linux",
)

oci_pull_image(
    name="static-debian12-debug",
    repository="gcr.io/distroless/static-debian12",
    digest="0eb1b021dc83cd103a4f5bec6d5292907b7f649137a200e178aa42d768d93ba2",
    architecture="amd64",
    os="linux",
)

docker_environment(
    name="golang", 
    image="golang@sha256:70031844b8c225351d0bb63e2c383f80db85d92ba894e3da7e13bcf80efa9a37",
    platform="linux_x86_64"
)
