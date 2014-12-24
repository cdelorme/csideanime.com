
# design concepts

everything in this document is a WIP, `@TODO`

- create a detailed mongodb schema and model our expected data
- seek UX assistance with interaction design


Untested Concepts:

- client side cache using lscache for API results
- use of UTC timezones exclusively on server-side, and whether there are any potential problems with this approach

First release frontend developer perspective:

- static html intro page
- lscache, mithriljs, and custom js code
- login/registration module
- member module
    - list members
    - display profiles

We should attempt to tie all callbacks into one endpoint and design a standardized format to distribute the response to the appropriate system(s).

We can move the lscache implementation to the second release if trying to get it to play nicely becomes a problem.  One of the main concerns is sensible cache-invalidation.  If the cache is constantly being invalidated and in large chunks, then caching the content serves little value.  Additionally inconsistencies between browsers could create a problem.


First releast backend developer perspective:

- member struct /w oauth & profile
- oauth setup & connection
- member system
    - logins
    - listing
    - creation
    - fetching profile(s)
- permissions via ACL
- tags module
- threads & post structs
- chat structs

Do we want to separate oauth, profiles, and core member structs so when "list" is requested we can fulfill it?  Or do we want to enable json omitted fields?

Do we want a single communications module for threads and posts, or individual modules for each?  Is there any coupling that we can expect to be a problem between them?


