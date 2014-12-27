
# services

**This is a tentative outline of services, not a comprehensive list, and is subject to change and grow.**

The following list has been carefully picked apart to build a more comprehensive summary, and separate files with details created to match:

- [frontend](services/frontned.md)
- [javascript](services/javascript.md)
- [permissions system](services/permissions-system.md)
- [member management](services/member-management.md)
- [chat system](services/chat-system.md)

The following is a generic list of services, _which will need to later be broken into frontend and backend development objectives_:

- open authentication
- tag management system (add/delete/apply-to-post-or-thread)
    - tag grouping for identifying companies and authors
- thread management (create/edit/delete)
    - threads by tag by popularity/activity
- post management (create/edit/delete)
- search mechanism, auto-complete by tag, otherwise fulltext search
    - advanced search engine?
- karma system
- site-theme system
- caching & search engine (likely elastic search)
- statistics /w per-member and forum averages, fancy charts and graphs (maybe?)
    - page views, searches run, posts/threads made, top contributing members
- member media upload & management system
    - accept uploads
    - store name as metadata, and rename to sha256 hash (to avoid duplicates)
    - push to s3, provide cloudfront url
- member blogs
- content release tracking?
- watch-lists & progress-tracking-lists?  Level up schemes?
- anime/manga/game/etc management system
    - review system for anime/manga/game/etc

_We still need to detail the behavior and logic behind each of these services, and update the list if we find mistakes._


## concerns

There are a lot of areas where we haven't made conclusion decisions yet, and concerns surrounding them.

- should our tagging system should allow for freehand text entry?
- should adding tags to auto-complete list could be tied a karma-point cost to prevent abuse?
- what actions should be tied to karma, and in what ways?
- how do we introduce a fair weight system into karma?
- do our users care about fancy statistic charts?
- given the goal of an autonomous forum, is there any way to avoid scrutiny for the "developer" role not adhering to the karma system?

Without a caching engine to start with, we will likely have slow search results.  Provided we don't start out with especially high traffic this shouldn't be a concern, but it could grow to become a concern.

While I want to both support and promote user-uploaded multimedia, I do **not** want the site to become a streaming service.  Either no-videos or some kind of task for moderators will need to be implemented.

Ideally we want our admin and member profile menus to be extremely simple.  Instead of a dozen layered pages, I'd rather have one or two very specific pages.  For example, the administrative forum settings should be on a single page, and may not even be necessary in a first-release.

During an earlier design I considered making ban-management flexible instead of global, as well as making moderation specific to area.  However, with a tag based system and the ability for any user to add and remove tags, it isn't possible to implement it fairly.  Was this a heavily desired feature, or can we deal without it?  Are moderators going to be intimidated if their privileges are no-longer "per-board" but rather global?

Do we want to separate oauth, profiles, and core member structs so when "list" is requested we can fulfill it?  Or do we want to enable json omitted fields?

Do we want a single communications module for threads and posts, or individual modules for each?  Is there any coupling that we can expect to be a problem between them?


## solutions

I like freehand entry, and if a tag does not exist, it would be nice to be able to provide the option to add it.  However, the option should be more explicit than hitting enter, otherwise it would be too easy to accidentally enter mispelled tags.

A possible solution to the `developer` group concern would be using the mock api for operations.  However, because these do not actually run against the database, it limits what can be tested.  Another option is to force developers to run their tests on their local system with mock data.  _There isn't really a way to resolve the "has access to the server" concern._  As the owner of the domain name, server, and hosting, I have a biased level of access.
