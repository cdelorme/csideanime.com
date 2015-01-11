
# release plans

Our goal is to reach an alpha release checkpoint swiftly.  By sticking to a bare-bones design we avoid unnecessary time spent on small details.

We can then identify the more critical areas of concern for a beta release.  **The beta release is likely when we will migrate from `seaofnumbers.com` to the new site.**

Our deadline for the beta release is early March 2015, which does not give us a great deal of time.  If we can hit our Alpha release by end of January it may be possible.


## alpha release

The tough part for this section is going to be identifying critical services as opposed to "desired" services.

Here tentative outline of services for alpha release:

- [open authentication](services/open-authentication.md)
- [member management](services/member-management.md)
- [permissions system](services/permissions-system.md)
- [tag management](services/tag-management.md)
- [thread management](services/thread-management.md)
- [post management](services/post-management.md)
- [frontend](services/frontned.md)
- [javascript](services/javascript.md)

_The karma engine and statistics systems are pending, because they are key elements to what I want to build, but are also complicated and may not be possible to have built in time._


## beta release

The beta release will focus on bug-fixes from the alpha-release, and a plan to migrate data from our existing site.  The migration is expected to be the harder part here.  If we migrate posts and threads we also need to migrate user accounts, but the login schema being so different will make that very difficult.


## stable release

After the beta release all missing, incomplete, or broken features will be worked on long-term to produce a stable release.

The stable release should include all documented services in a final and complete state, which is not to say they cannot be improved or added to later.

The stable release must be bug free, which means comprehensive testing should be performed on all components prior to tagging.  Semantic versioning will be used from here forward.


## future releases

After reaching a stable release, no unstable code should be pushed to master.  All code changes will be triaged using feature-branches, allowing multiple features to be worked on independently without introducing breaking changes.
