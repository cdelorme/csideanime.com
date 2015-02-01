
# tag management

This is a key element of the new system, and will take the place of "boards" in a traditional forum.

Instead of "Most recent" or "Most Popular" threads being treated as small stubs, they will take the center stage, and their contents will be used to provide a set of homepage categories.  This means as traffic patterns change or emerge, what appears on the home page will vary.

Tags are simply metadata that represent a category of things, and in that same sense tags can also represent other tags.  The tag design should be a single collection with documents containing a name, maybe a description, and optionally parent-tags.


## considerations

This dynamic or "organic" method of displaying information will have the price of inconsistency.  To counter that problem we need to provide intuitive controls to search for tags.

While we will, at first, relate tags to the posts and/or threads, in the future they will likely be related to many other components in the system, including the reviews, media, and blog subsystems.

While we want the design to remain simple and somewhat intuitive, there are specific things to consider, such as metadata or tags representing other tags.  For example, if we have an anime title as a tag, we may also want to generically toss the "anime" tag onto it as well.  To do this we need tags to contain metadata or "parent tags".



Our object design should be as simple as possible, which could technically mean just `name`.

However, handling group association may be a struggle.  We could store groups as collections of `name` and `tags`, or we could add `groups` to the tags collection.  We will be facing a similar problem to the permissions system, where we technically need a way to manage all tags and groups.

Another option is to use single-parent tagging, where tags can serve as metadata for themselves.  This would allow us to store all the data in one collection, optionally.  If a tag only has a `name` that's fine, but it can also have `groups` or `tags` depending on how we choose to manage the association for best indexing results.

_Some consideration is still required before we decide on this._

Let's say we have a thread discussing an anime, for example "To Aru Kagaku No Railgun".

We could create a "To Aru Kagaku No Railgun" tag and mark it, so if another user is looking for discussions about the railgun series and maybe we never said the word "railgun" in the discussion, it could be found.  We could also add more basic tags like "anime", or "manga" to apply as well.

The more tags applied, the easier it becomes to find and access that thread or post.

The groups allow us to classify tags.  For example, we could say that "anime" and "manga" are "categories".  We may also say that "To Aru Kagaku No Railgun" is an "anime" and also a "manga", by creating separate groups for "manga" and "anime" (_not the tags we used earlier_).

We may use groups to change how tags are displayed, or to adjust auto-completion and search results.



## actions

- create tag
- edit tag
- delete tag
- get tags

Tags are fortunately very simple, and adding parent tags is a simple edit.

Unlike other systems, deleting tags should actually involve deletion.  I suspect we will accrue a lot of tags over time, and keeping erroneous ones, such as mispelled words, would be silly.  **However, when a tag is deleted, additional steps may need to be taken to remove them from all related subsystems, which is going to be very challenging.**  For now, we can ignore deleted tags in other subsystems.

We will supply a single "get tags" method, for navigating tag management.  However, single tags should never be accessed for anything other than editing.  Otherwise it's part of either the search engine or specific to each subsystem (eg. post search or thread search logic).


## model(s)

Tag data will look like this:

    {
        "name": "To Aru Kagaku No Railgun",
        "description": "..."
        "parents": [
            "anime",
            "scifi"
        ]
    }


## closely related systems

It will be quite difficult to decouple the tag subsystem from these others:

- search engine

If we use something like elastic-search, we will need a means of providing additional categorization.  If we support tags AND full text search, then we effectively need to distinguish between the two, which may be more difficult than just treating tag names as part of the full text.  Whether this would cause a problem or not is debatable.
