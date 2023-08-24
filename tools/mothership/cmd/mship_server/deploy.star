exec("base/ci/defaults.star")

namespace("mship")

service(
    name = "mship_api",
    image = "%s/mship_api:%s" % (args().registry, args().version),
    ports = [
        Port(
            name = "grpc",
            port = 6677,
        ),
        Port(
            name = "http",
            port = 6678,
            expose = True,
            external = True,
            host = "mship.resf.org",
            path = "/api",
        ),
    ],
)
