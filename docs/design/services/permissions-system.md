
# permissions system

Following the basic Access Control Layer design, permissions should be associated with groups and actions.

In this system, members belong to one or more groups, and groups can perform one or mor eactions.

_There is will be no option to "restrict" an action, as overrides like that create very complex requirements._

The permissions service layer will need to provide the following operations:

- create a group
- delete a group
- change group name
- add an action
- delete action
- change action name
- add an action to a group
- remove an action from a group

**Each of these could, technically, represent an `action` in the very system they apply to.**


## Notes

Actions will be represented and checked as plain text.  This means using constants per module to define its set of actions is the most sensible.  New modules will need to install or provide their list of actions to the main system, so we can assign permissions dynamically.

Permission checks can only occur in the three major layers of the system, the controller, the service, and the data access layers.

- we can stop requests before they begin based on permissions at the controller layer
- service layer can augment the behavior of a specific action based on permissions
- the data access layer may apply augmented filters to queries based on permissions


## data design considerations

For scalability we have to assume that the system will need to load the permissions per request and that they will not persist in the cache of any single server (except the database from which they are requested).

We also need to consider the most efficient storage approach, such that they can be retreived more efficiently for the most common cases.  This means we should priotize the order that makes checking permissions the fastest, which is likely to be the reverse of how the management page is displayed.

This means that the `ACL` is just another service layer, which means it too will need a data-access component to be used by the various services.  _There could be reprecussions with importing a service layer into the data access layer._

Persistence should exist via a `permissions` collection.  The schema of the documents can be `name` and an `actions` array of strings.

For management purposes we can create a special group called `all`, which contains all of the permissions.  The reason we don't just make that the same as administrator is because some new components could introduce breaking changes and need to be tested before they are fully enabled.  It also adds greater flexibility.

Creation of a new group will always create a new document.  Because the number of documents is not likely to grow by any significant degree, and we won't be modifying them as regularly as we query them, we can get away with always pulling all documents from this collection.


## states

All visitors will inherently belong to the `Guest` group.  This should give them basic access rights for the forum contents.

Members will belong to a `Member` group.  This will give them write access.

For the frontend, we may provide members with specific featuresets the ability to toggle them on or off, which will keep the interface from getting cluttered when they visit just to browse.  _Keep in mind, `Administrators` are also `Members`.

_Because there is no distinguishing factor such as boards that allows heirarchical control over a subset of forum content, we may no longer have global and non-global `Moderator` groups._  However, as we add more features to the system, we may create feature specific moderator positions (eg. `Blog Moderator`).

We will have a `Developer` group that has access to all the same settings as Administrators would.  This is important because developers will want to be able to test systems, but are likely not participating in the same activities as the administrators, so a separate title gives a distinction that is missing from most other platforms.

**Banning will not be related to the permissions system at all.**  Banned users will simply be denied access to anything and everything when they attempt to connect.  The easiest way to apply this check is to make the acquisition of the user part of the base controller.
