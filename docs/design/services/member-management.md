
# member management

Member information will be stored in a document, and be split into several sub-documentsf.  This allows us to separate things such as public and private profile information.

The basic design should include top-level storage of email, name, and password hash.


## considerations

I am not entirely certain how we can specify whether or not each section of profile information is public or private, but ideally private is a copy of public **plus** additional information.

Users can choose a display name, and change it regularly.  _While we may associate a karma cost to doing so,_ the best solution to identity theft with this system is to provide a history of their names and timestamps as part of their public profile.  _Another related concern is that since all data is being sent to and from an API, we would need a method of preventing users from changing their own username history in spite of it being part of their user profile._

Currently logging in, checking in (as an already logged in user running a new request), and things like username completion are likely to be part of the members system, **however** I am considering the creation of a separate but dependent "auth" component that accesses the same data.  This would separate flow-control concerns, and allow our base controller to handle those sorts of things.


## actions

The member actions include:

- create a member
- delete a member
- edit a member
- get member
- get profile (for specific member)
- get profiles (for all members)
- login
- checkin

Member records will never be deleted from the system, only disabled.  We can use a boolean flag to disable login and checkin methods for a disabled account.

_Any part of a member which is modified (such as profile or groups) is all part of the same `edit` method._  Open authentication will also be a component of member, making it also part of the same "edit" path.

We can use a timestamp field for banning.  No ban will be "permanent", but instead we can assign a duration, and easily check against todays timestamp to prevent access.


## model(s)

The model will look like this:

    {
        "name": "name@mail.com",
        "password": "hash",
        "profile": {
            "name": "display name",
            "joined": #
            "names": [],
            "email": "public@mail.com"
        },
        "settings": {},
        "private": {},
        "oauth": [],
        "groups": []
    }

Passwords will be stored as hashes, and logins will accept both the parent name or display name.

The `names` array is a historical record with names and timestamps when a name was adopted.

The private object is a copy of the profile object with any information the user did not wish to make public.

The `settings` object is a key-value pair of specific toggles, which may include things like a language filter or adult content.

_The profile object itself has not been fully architected and will likely grow organically._

Open authentication is not yet planned out, so it's an empty array.


## closely related systems

Decoupling these systems may be difficult because they will likely rely on eachothers data:

- karma
- statistics
- private conversations
- uploads/media
- blogs
