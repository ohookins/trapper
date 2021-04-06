# trapper

Trap signals and complain a lot.

# TODO

- Env var to configure INT/TERM signals to trigger shutdown of application.
- INT/TERM signals shutdown timer is configurable.
- HTTP requests take a non-zero amount of time to return, to simulate
  connections in-flight.
- Health checks fail once shutdown has been triggered.