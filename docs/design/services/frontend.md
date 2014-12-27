
# frontend

This document is a summary of all the things that will need to be done to create a static frontend page.

**The initial design for everything here should be done "Mobile First", which assumes a mobile device and browser for the initial design.  More complicated interfaces and dynamic controls can be implemented later.**

- we need a header, which should contain the sites name and optionally a rotating banner.

- we need navigation with links for:
    - home
    - login (if not logged in)
    - register (if not logged in)
    - account (if logged in)
    - admin (if an admin)
    - members (if logged in)
    - stats ???
    - search ???
    - tags ???

_The navigation items displayed will vary based on the member being logged in and their permissions._

The login could be a button with username and password input fields, or a separate "page" with inputs for username and password.  Personally, I would much rather have a simplified login than a separate page.  _How we choose to display open authentication options such as "Login with Google" or "Login with Facebook" is yet to be decided._

The registration page should have:

- email
- display name
- password (*2)

The `account` is their own profile, using the `member/#` route, and should provide "click-to-edit" with the following fields:

- display name (change)
- email address (add/change)
- change password
- open authentication accounts (add/remove)
- themes? (change?)
- website?
- game & contact names (steam, skype, etc)
- add/remove/change avatar

_A preview button should open their profile, using the `members/` route, in a new page._  The same page layout can be used to display other members.  Specific information for other members will not be displayed.  Guests cannot browse member profiles.

The index should show the most popular threads, with a post count, last activity date/time, and creator.  _The post count and last-activity date/time should increase without page refresh._


Viewing a thread should show all the posts.  Ideally it should use automatic scrolling to pull down new contents.  _The tough part will be making it work well with mobile._  An indicator that new content is being loaded should be supplied.  Also, "End of thread" should be displayed when no more content is available (how to determine when to display this may also be a challenge).

A button to create a new thread should be displayed at the top and bottom of the home page or any thread, and buttons to create a new post at the top and bottom of any thread.

How and where to display the search is still under consideration.

Administrative pages include:

- forum settings
- ban manager
- permission manager

If we can somehow manage to display all three administrator components on a single page, that might be especially nice, but otherwise a separate set of links, a tab-based sub-menu, or a drop-down menu might be needed.

The forum settings would have to be "server safe", so something like the default theme and manage banner logos to cycle.

The ban manager is effectively a member search filtered by "banned" status.  It allows administrators to unban members.

The permission management allows the administrator to see available groups and permissions, and add and remove groups and permissions.  However, with ties to the karma system, group management might be slightly more complicated (adding groups would require conditions to acheive, and removing would require it is either not connected to karma or no members have that group).  This at least provides a bit of fore-thought to prevent, for example, deleting moderators after a bunch of people spent karma to elect some folks.

We may also want a tags link, which will allow all members to search by and manage tags.  The page could, by default, display a tag-cloud.

A chat system will need to be integrated with the home page, though it may be a feature for a future release and later consideration.

A members page would need to include a filter, that applies to more than just the name.  For example, the ability to filter members by group, name, or status (eg `banned`).
