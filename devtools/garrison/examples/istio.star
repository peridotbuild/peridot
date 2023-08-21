exec("./examples/default_flags.star")

flags(
    host = "example.com",
)

default(
    istio = True,
    namespace = args().namespace,
)

env = Environment(
    PORT = "80",
)
env.update(var("ga_env"))

service(
    name = "simple",
    image = args().image,
    ports = [
        Port(
            port = 80,
            expose = True,
            external = True,
            host = "test.resf.org",
        ),
    ],
    env = env,
)
