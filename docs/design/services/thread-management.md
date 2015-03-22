
# thread management

This is two parts, the primary focus of our subsystem here is "threads", which have long existed in communications tools as containers for sets of sequentially related to the second half, "posts".

Threads themselves tend to have a very simple model, at times limited to just a title.  However, metadata and pre-calculated fields such as a post count, date created, date updated, author, and author name will probably be included.

Posts could be considered authored messages, which exist under a thread, and contain the message, author, authors name, date created, and possibly a date edited/updated.

In both cases, the ability to delete (or hide) content is required.  Additionally, we will be supplying a tags field, which allows us to categorize and provide a means of searching without knowing the written contents.

We also want special flags for thread `locking`, and post `stickying`.  The locking will prevent non-privileged members from posting to a locked thread.  Stickied posts will always load first in a thread.


## considerations

_One concern is how to distinguish whether a member has read a specific thread or not,_ so we can change its appearance to indicate whether they have not.  This may be a modification to a member model and not to the threads or posts.

One possible option was to store posts inside of threads as single documents.  The mongodb database provides a generous 16 megabyte per document limit, which is effectively 8000 pages of raw text.  However, there are problems with documents increasing by 20%, write locking, and requesting "parts" of a record.  If we clump the data together, a thread cannot be polled for just its title and metadata, all posts must be loaded.  If more than one person submits a new post, the write-lock forces the second to wait for the first to complete.  Because a thread could constantly grow, it would regularly increase by 20%, forcing the database to request a brand new container for the data, leading to a significant slowdown in performance (especially as it grows).  Finally, whether we return upwards of 1mb of raw text, or filter at the server-side, we are working with much larger sets of data than we need to be, which will cause performance problems.

The other concern is regarding storing external data, such as author names, for query convenience.  If a member updates their display name, it becomes a sizable task to find and update all of the rest.  Therefore, we might offload that to a later job, which results in "eventual-consistency".  The other option is to just leave old names as they were at the time of a post, which will probably help with context for readers.

One of the biggest concerns is path-control.  Since we have two separate sets of objects with similar methods, we need a path for each.  However, there are some cases where we need more.  My goal is to avoid redundancy, and complexity which makes layering (eg. `/threads/#id/post`) undesirable.  However, we want to avoid creating unexpected limitations in the API's ability to be extended when creating the paths.


## actions

These actions combine both models into a single service layer:

- create thread
- edit thread
- delete thread
- get thread
- get threads
- add post
- edit post
- delete post
- get posts by thread

_Both deleting and locking a thread count as editing, but action-permissions will need to be checked.  Adding and removing tags will also be an edit to either a thread or a post, respectively._  I have not yet decided how to implement tagging into the UI.

We probably won't need a "get thread" method for the UI anywhere, but rather for internal operations.

There are some actions which we may want to place elsewhere as they involve two or more services:

- get posts by member
- get threads by member

Members often want to see the option in their profile to find their old posts or threads.

Finally, here are some non-standard actions that probably need to be coordinated:

- get recent threads
- get most active threads

These would be by `date created` and `date updated`.  We may also use `count` as another respective measure of "most active".


## model(s)

A thread will look like this:

    {
        "title": "Some descriptive words",
        "author": #,
        "authorName": "name",
        "created": #,
        "updated": #,
        "count": #,
        "deleted": true,
        "locked": true
        "tags": [ "topics" ]
    }

The name and member id are stored, allowing us to easily print what needs to be displayed when this object is requested.

The `created` and `updated` fields are unix timestamps.

The `deleted` and `locked` flags should either not exist, or be set to true.  It should never be "false".

Indexing will depend on the queries we eventually end up running.

A post model looks like this:

    {
        "message": "...",
        "author": #,
        "authorName": "name",
        "created": #,
        "updated": #,
        "tags": [],
        "deleted": true,
        "sticky": true
    }

As with previous boolean flags, `deleted` and `sticky` should never exist and be "false" in the storage of the data model.

_We may consider placing tags only onto the threads, which would reduce complexity along with search accuracy._
