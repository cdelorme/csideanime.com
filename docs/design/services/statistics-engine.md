
# stats

Not yet sure how to integrate or exactly what to track, but there are a lot of ideas that stem from this.

One concern is how to connect.  Do we attempt to make statistics independent of each core service, or couple them tightly.  To decouple would require complex architecture or an event system.  Event systems are, traditionally, very difficult to debug.  _I would rather have stats directly coupled._

- statistics /w per-member and forum averages, fancy charts and graphs (maybe?)
    - page views, searches run, posts/threads made, top contributing members


## interesting statistics

Besides my own interests, I've asked for ideas from our members.  If you have a suggestion open a ticket.  Currently the list includes:

- pie-chart of activity by tag per user
    - the ability to group/find users with similar activity
    - the ability to identify areas that get used most and focus on features for them
- line-based histograph of recent activity such as active users, recent posts, threads, all over time
    - allows any member once logged in to easily get a picture for the last N hours of activity
