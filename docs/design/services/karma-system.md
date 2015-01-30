
# karma system

A lot of mystery around this still.

My goal is to deliver a karma system unlike any other.  It will control the direction of the forum, by connecting karma directly to forum ownership and privileges.  Karma allows voting, voting enables members to choose the authority figures, and keeps authority among active users instead of having administrators choose other admins and moderators like a dictatorship.

_Developer would be the special privilege group that needs admin features for testing, but does not act as an administrator, and is not able to be removed from power._

Currently thinking of a weighted karma system, where points are awards in a balanced and fair way based on various facts.  Information coming from statistics that promote community involvement (posting), but not in a negative way (spam is not good).  An emphasis placed on member feedback.

Will also need that balance to still (sanely) be able to provide a measure of control even during periods of low activity, such as when the system is just starting out, so that the balance is maintained regardless of activity levels, and costs of specific forum services are fair.


## objectives

Not all members will have access to all forum actions.  Some forum actions will only be available to specific groups, while karma will allow users to "purchase" specific actions.

One such action is a voting permit, which allows members to vote for moderators and administrators.  Another possible item would be campaign permit, allowing members to put themselves up for voting into specific positions.

Ideally moderator and administrator privileges would be rolling, and the community will be responsible for handling assignment.  Those with moderator and administrator privileges may also have additional options such as to increase the number of available moderator/administrator positions.  This allows the system to scale based on need and not because every member had enough karma-points to campaign.


## design considerations

We need to decide when to add or subtract from member karma.

Contributions to the system or system meta-data should generate positive karma.  For example, posting, creating a new thread, adding tags, creating new tags, and filling out profile information.

**Karma should never be negative.**  However, it can be reduced to zero in these ways:

- negative user feedback
- spending

_We may also consider depletion of karma over durations of inactivity, or whether to use member activity as an augmentation in the amount of karma earned when acting._

Members should be able to effect eachothers karma via feedback.  For this to work, we need members to provide feedback or critique eachothers posts.  Negative feedback will require supporting text.  Neither negative or positive feedback will have a great impact alone, if members agree with feedback they should specify it.  In this way we can apply exponential growth to the points.

Another consideration is whether to tie karma to the permissions system directly.  If we make karma an optional cost for all actions it bakes karma directly into the system.  This approach may even make it possible to spot and eliminate bots easily.  If posting or opening threads costs karma, first signups cannot do either before building some (perhaps by reading posts and other actions).
