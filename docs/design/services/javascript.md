
# javascript

Besides external libraries we will need our own custom javascript for a few things.

- dynamic banner
- client-side markdown rendering
- centralized callback mechanism
    - centralized error handler
- cache integration /w bucketing for easy removal of oldest items
    - logic for invalidation and updating or updating of individual items?
- immediate rendering component based on application state
- templates for various display pages
- pushState logic for the links and navigation on the page
- hmac sha system (something that can provide two-way validation would be nice)

Ideally lscache would also store member permissions at login, allowing us to dynamically change things such as moderator and administrator options.  Ideally the desktop and mobile browsers would have toggles to enable or disable moderator/administrator features, to keep the UI clean when not acting as administrators or moderators.

The same could be applied to regular members as well, making things like editing and deleting a post or thread display conditionally.

_The initial release will likely only have basic "always displayed" buttons._

Things that might be displayed for admins and moderators:

- delete/edit (for any post or thread)
- posters ip address
- ban hammer

Things displayed for normal users:

- delete (their own posts, but **not** threads with content from other members)
- edit (their own posts or threads)
- add/remove tags

_The key difference between administrator and moderator might end up being that moderators cannot change forum level settings and permissions.  Although ban management may be preferred in moderator hands as well._

The ability to dynamically generate just one set of controls, or to conditionally display the controls for the posts and threads in view would be ideal too.  We can add media-queries to decide when to hide or display controls for non-mobile displays.
