
# thread management

This would be a service layer specifically for managing threads.  Its actions include:

- create thread
- delete thread (disable*)
- edit thread
- get latest threads (by date created)
- get most active threads (recent activity, most posts)
- get threads by member
- add tag
- remove tag

_It is also a possibility that thread management could be combined with the post management service due to how closely these relate.  We could also find a data design that doesn't need them to be separate._

Because threads are no longer connected to boards, we don't need to worry about moving a thread from place to place.  Simply adding and removing tags fixes that.  _We may also decide that we only care about tags on posts, and can ignore tags by thread._  In my initial design, I believe a tag applied to a post in a thread should also be applied to the thread, since the thread is a group of posts that represents all of them.


## data design considerations

The representation of a thread is very basic.  It consists of a `title`, a `created on` date, `author id`, and optionally `tags` and `update on` fields.

If we chose to merge threads with posts, we could design posts to connect to another posts id and treat it as a thread, and add a `title` field optionally to that post.  This would likely hurt performance and complicate design though.

_Is it unlikely that posts will be pulled directly without a thread, and that we will want to know the name of the thread from a post?_  If so, we may want to store a copy of the thread name in the post, but then we have the "eventual consistency" problem if the thread name gets changed.

Similar to other components, such as chat and posts, we may want to store a copy of some member data, such as the display name, so we don't need to pull the member profile for every thread being loaded.  While this can create the problem of "eventual consistency" it would be on a minor bit of information that should rarely create confusion.
