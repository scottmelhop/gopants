go_binary(
    name="bin",
)

go_package()

oci_layer(name="bin-layer", packages=[":bin"])

oci_image_build(
    name="image",
    layers=[":bin-layer"],
    base=["//:static-debian12-debug"],
    repository="example.com",
    tag="latest",
    # entrypoint="python3",
    # args=[f"{dirctory_no_slash}/{name}-srcs.pex"],
    # env=[
    #     f"PEX_PATH=/{dirctory_no_slash}/{name}-srcs.pex:/{dirctory_no_slash}/{name}-deps.pex",
    # ],
    # tags=tags
)