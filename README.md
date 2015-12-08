# projecter
Interact with a project's related webservices (project management, time tracking, revision control, team chat) from the command line.

## What does it meaaannn?

You create a `.projector` file with your login credentials to the services you want to use, and put it in your `~/`.

Then you create a `.projector` with the specifics for the project (repo locations, slack channels, whatever) in a directory.

`projector <command> [<args>]` does stuff.

## Available commands, current and planned
- [x] `projector init` — Clones project code repositories
- [ ] `projecter status` — Print config as Projecter sees it (merged `~/.projecter` and `./.projecter`—WARNING: may print semi-sensitive data, like API keys)

## Example configs

`~/.projecter`
```
# The only working service so far is GitHub, and projecter just assumes you have ssh-agent managing your keys, so 
# nothing here so far.
```


`{project_root_path}/.projecter`
```yaml
# Enable 
use:
  - github_source

# repos will be checked out to project_root/location_key
# So... make sure the key isn't something stupid.
github_source:
  locations:
    projecter: git@github.com:cremalab/projecter.git
```
