# cli

This is the command line interface for canifest. To do anything, the CLI needs
the core server up and running. Eventually, we need to have a wrapper script
that can start both at once (or maybe let the CLI start the core on its own),
but for now you're going to need to start the core before you start the CLI.

Assuming you've built the core, type ./bin/core & from your $GOPATH. (Make sure
  you add the ampersand so it will run in the background)

First, install the CLI with go install github.com/canifest/cli from your $GOPATH

Then run ./bin/cli and you should see it all connect.

The only thing you can do with the CLI right now is ask for help or quit, which
should successfully bring down the core server too.
