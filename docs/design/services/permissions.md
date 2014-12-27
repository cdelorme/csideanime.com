
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

Because golang is stateful, we will need to reload the groups and permissions for every request.  Therefore, we should find an efficient way of storing groups and actions.  Similarly, how we store actions and groups will also effect how easily we can create a management page.  _The management page is slightly less of a priority since it won't be loaded per page per user._

