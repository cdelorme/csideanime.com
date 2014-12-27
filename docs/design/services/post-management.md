
# post management

This would be a service layer specifically for managing posts.  Its actions include:

- create post
- delete post
- edit post
- add tag
- remove tag
- get posts by member
- get posts by thread
- get posts by thread and by member

This should be a very strait forward component.


## data design considerations

Posts should contain the `author id`, `thread id`, `message`, and `created on` date.  Optionally the `tags` and `updated on` fields.

With slightly more of a concern than threads and chat, we will need a bunch of member information per post, such as their name and avatar.  It would certainly improve query speed to keep a copy of those with the posts.  However, we face even more concerns than "eventual consistency" with this area, because of displaying stats and other member information.  Keeping the post-count etc up to date for a member in a copy of every post would be terrible for performance.

In this case, we may want to store a reference and somehow load member metadata from that.  The problems this creates are then directly related to how many records we are pulling.  If we shard the data and end up hitting multiple servers this also becomes a concern.
