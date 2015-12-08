# projecter
Interact with a project's related webservices (project management, time tracking, revision control, team chat) from the command line.

## What does it meaaannn?

You create a `.projector` file with your login credentials to the services you want to use, and put it in your `~/`.

Then you create a `.projector` with the specifics for the project (repo locations, slack channels, whatever) in a directory.

`projector <command> [<args>]` does stuff.

## Available commands, current and planned
- [x] `projector init` — Clones project code repositories
- [ ] `projecter status` — Print config as Projecter sees it (merged `~/.projecter` and `./.projecter`—WARNING: may print semi-sensitive data, like API keys)
