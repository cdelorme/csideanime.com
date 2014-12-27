
# tag management

This is a key element of the new system, and will be used to structure and organize everything on the new sytem.  Starting with threads and posts, it will eventually extend to media and review systems, blogs, etc.

The design is very simple, we have two sets of data:

- tags
- tag groups

A tag is a simple bit of text that represents metadata about the related entity.  The tag groups are metadata about tags.

The tag system _may be_ a simplified version of the search system, without full-text indexing.  Ideally the search system should be separate and connected to a search cache such as elasticsearch.

This system would need to support the following actions:

- create tag
- delete tag
- update tag
- add tag group
- remove tag group
- update tag group
- associate tag with tag group
- unassociate tag with tag group

_We still also need to consider how tag groups work, and what to optionally append to them.  This will determine what additional actions, if any, will be needed._


## data design considerations

Our object design should be as simple as possible, which could technically mean just `name`.

However, handling group association may be a struggle.  We could store groups as collections of `name` and `tags`, or we could add `groups` to the tags collection.  We will be facing a similar problem to the permissions system, where we technically need a way to manage all tags and groups.

Another option is to use single-parent tagging, where tags can serve as metadata for themselves.  This would allow us to store all the data in one collection, optionally.  If a tag only has a `name` that's fine, but it can also have `groups` or `tags` depending on how we choose to manage the association for best indexing results.

_Some consideration is still required before we decide on this._


## usage example

Let's say we have a thread discussing an anime, for example "To Aru Kagaku No Railgun".

We could create a "To Aru Kagaku No Railgun" tag and mark it, so if another user is looking for discussions about the railgun series and maybe we never said the word "railgun" in the discussion, it could be found.  We could also add more basic tags like "anime", or "manga" to apply as well.

The more tags applied, the easier it becomes to find and access that thread or post.

The groups allow us to classify tags.  For example, we could say that "anime" and "manga" are "categories".  We may also say that "To Aru Kagaku No Railgun" is an "anime" and also a "manga", by creating separate groups for "manga" and "anime" (_not the tags we used earlier_).

We may use groups to change how tags are displayed, or to adjust auto-completion and search results.
