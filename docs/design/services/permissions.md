
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

Actions are represented as plain text.  Integrating them will require code changes.  _Best practice would be to use constants to match the action text._

The design we are looking to use will require that the/an access control layer be available to data access objects, services, and controllers.

- we can stop requests before they begin based on permissions at the controller layer
- service layer can augment the behavior of a specific action based on permissions
- the data access layer may apply augmented filters to queries based on permissions


## data design considerations

How we store the data will effect how we retreive it.  Since permissions will likely need to be retreived regularly, a good schema is essential.  Of a lesser priority, it also effects how we generate the permission management page.

One consideration is that golang is stateful.  We _could_ create a new ACL layer per requested operation and load in the permissions fresh.

An alternative consideration would be to use a single shared ACL that refreshes every so often, but otherwise retains the permissions in memory.  _This could create, for a very short time, discrepancies between servers when permissions change in a scalable infrastructure._

The basic storage concept is simple.  Create a `permissions` collection, and the schema of the documents can be a string `name` and `actions` string array.

However, the problem is with managing groups and actions.  If a group is freshly made and has no actions, or an action is added but not assigned to any group, we need to persist them.  We could throw them into entirely independent collections, but for this system it might work best to create an `all_actions` and `all_groups` document, and prevent the creation of a group by those names (or something similar to those names).

With that design, any changes to actions and groups could be saved a very simple operation.
