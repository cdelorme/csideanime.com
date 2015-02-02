
# permissions system

Our permission system will be based on access-control-layer logic with two components; `Groups`, and `Actions`.  The permissions will exist in three-tiers, allowing us to distinguish different logical layers where this information may be needed.

- controller layer: prevent wrongful access, eg. to `admin/` for non administrators
- service layer: prevent wrongful actions or augment response behavior by ownership or access rights
- data layer: modify query logic by access rights

To avoid calling a global permissions service, we will supply the ACL itself from the `BaseController`, alongside the member object.

Using this system, members may belong to one or more groups, and groups can perform **zero** or more actions.

**There will be no restriction of actions**, so overlapping groups simply merge and share permissions.  There are also no "individual" permissions, assigned per-member.

All actions will exist in the form of plain text, which gives us the greatest flexibility.  New modules will need to supply their list of actions to the core of the system, so we may need a registration component (depending on how abstracted we want sub-systems to be).  It may also be wise to use a set of constants to represent the data.

We will have a single `Permissions` collection.  Each document will contain a `name`.  An optional array of `actions`, an optional url field `icon`, and two optional boolean fields for `management` and `dogtag`.

The dogtag and icon features allow us to provide members with the ability (via `karma`) to create permissionless groups which allow them to select an dogtag-image to display under their username.

For scalability, we cannot cache the permissions and must poll the database or lowest level caching component for the latest states.

Query efficiency for this collection has two parts.  First, indexing the names and boolean fields.  Second is storing the data in a way that makes retreival easiest for the most common situation.

All visitors will inherently belong to the `Guest` group.  This should give them basic access rights for the forum contents.  Members will belong to a group with write access.

_Because there is no distinguishing factor such as boards that allows heirarchical control over a subset of forum content, we may no longer have global and non-global `Moderator` groups._  However, as we add more features to the system, we may create feature specific moderator positions (eg. `Blog Moderator`).

**Banning will not be related to the permissions system at all.**  Banned members will simply be denied access to login, preventing already logged in members from performing any action.  The easiest way to apply this check is to make the acquisition of the member part of the base controller.

Two factors should belong to any checks in the system.  First, whether they have the action in their list, and second is whether they own the object in question.  For example, a member cannot perform the "delete member" action, unless it is their own member record.


## considerations

We should use the `management` flag on a group labeled `All`, which will store every available action to the system, and will be used by the management page.  Still under consideration is how to handle completely dynamic actions.  There are a lot of undetermined variables, so for now we will likely use the additional document as the solution.

We may consider adjusting permissions to that actions can (optionally) have a karma cost, allowing us to create built-in restrictions.  _Integrating it directly might simplify management but would complicate the storage and retreival._

It would be nice to provide toggles on the front-end for features that complicate the UI.  Administrators, as an example, are members first.  Providing them with the same clean UI as every other member is important, and lets them continue to participate like every other member.  A toggle can allow them to take action by simply enabling the features, then later disabling them again.

I (cdelorme) carefully considered the addition of a `Developer` group with all access rights.  The concern was that developers are not necissarily acting administrators, but need privileges to test the code.  This distinction is important, but is a contradiction to the member-voting concept.

The lack of metadata associated with actions and groups may require greater complexity.  However, for now I'd rather we try to avoid complexity and simply make sure actions and groups are well-named.


## actions

Permission actions include:

- create a group
- edit groups
- get groups
- delete a group
- add an action (to a group)
- delete an action (from a group)

Standard UI shows all groups and actions in a single table, so edits for managed groups will probably involve sending all of them at once.  _This is a slight deviation from standard REST practices_, but it better supports how the data is used and interacted with.

Groups should never need to be acquired singularly, we can always assume the ACL will need all of them.

_Deleting a group simply adds a `deleted` field._


## model(s)

The model will look like this:

    {
        "name": "Administrator",
        "actions": [
                "Delete Member",
            ],
        "icon": "https://www.amazons3.com/admin-dogtag.jpg",
        "management": true,
        "dogtag": true
    }

The name is the only required field.  The dogtag and management fields should never exist with a value of false (absense assumes false).  The dogtag field indicates no actions will exist, while the management field indicates that it will be used for managing actions and should contain a "complete" list of available actions.  The icon field is optional, and should contain a URL to an image if desired.

_Specifications of the dimensions for icons is not yet determined._
