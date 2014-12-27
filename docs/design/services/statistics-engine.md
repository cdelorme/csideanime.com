
# stats

Not yet sure how to integrate or exactly what to track, but there are a lot of ideas that stem from this.

One concern is how to connect.  Do we attempt to make statistics independent of each core service, or couple them tightly.  To decouple would require complex architecture or an event system.  Event systems are, traditionally, very difficult to debug.  _I would rather have stats directly coupled._

- statistics /w per-member and forum averages, fancy charts and graphs (maybe?)
    - page views, searches run, posts/threads made, top contributing members
