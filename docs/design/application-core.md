
# application core

Application config file located in standard path (ex. `/etc/csideanime.com` || `~/.csideanime.com` || `./.csideanime.com`, and is json format.  It should contain:

- database connection info
- application­level­security­config not available to the community

_We may (optionally) also look at environment variables for this information._

Our core should instantiate, configure, and connect (in order):

- data access layers
- service layers
- controllers
- router

By creating all our core components up front we simplify the code and its readability.

Dependencies are created and supplied up front without any abstractions or lazy­loading, which reduces code complexity and can improve performance in memory.

After connecting all of our core, we register controllers with the router.

Controllers contain the routes, and the router uses wildcard­left routing to reach them, allowing complete support for any number of clean­url parameters and additional constructs without complicating the server or proxy configuration.

