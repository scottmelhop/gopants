go_binary(
    name="bin",
    environment="golang"
)

go_package()

oci_layer(name="bin-layer", packages=[":bin"])

oci_image_build(
    name="image",
    layers=[":bin-layer"],
    base=["//:static-debian12-debug"],
)