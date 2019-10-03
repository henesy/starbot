# Star Bot

Star bot provides a star-board service on a per-guild basis.

A server can configure their channels to be monitored (by default: all) and which channel is used as the starboard. 

A message will be copied and re-posted into the starboard channel when the number of reactions for the registered reaction emoji reaches the tolerance for posting (by default: 10). 

Any number of emoji to channel mappings can be created per guild. 

By default the bot can only be controlled by the guild owner. 

Star bot does not require elevated privileges, but must be able to read the contents of channels to count star reactions. 

Star bot does not retroactively delete messages posted if the original message drops below the required count for a channel. 

### Dependencies

* Go

### Usage

```
cd $GOPATH/src/github/henesy/starbot
go build
./starbot -t 'Bot BOT_AUTH_TOKEN'
```
### Commands

Register a channel named `#starboard` and the reaction emoji `:star:` as the starboard:

	s~register #starboard :star: 

De-register a channel named `#starboard` from counting a `:star:` reaction emoji:

	s~deregister #starboard :star:

Set a channel to passively scan in the guild (rather than the default of all):

	s~scan #art

Set all channels to be scanned (if able to be read):

	s~scan all

Set the count for an emoji `:star:` in a channel `#starboard` to require `2` reactions:

	s~count #starboard :star: 2

Configure the backlog scanning to be count-based for `1000` messages:

	s~backlog count 1000

Configure the backlog scanning to be time-based for `2 days`:

	s~backlog time 2d

Register someone named `henesy` to be able to control the bot in the guild:

	s~admin add @henesy

Remove someone named `henesy` from being able to control the bot in the guild:

	s~admin remove @henesy

Show current registrations configuration states and scanned channels (if not all):

	s~list

### Documentation

Star Bot invite: 
