
# member management

Member information will be stored in a single collection, but contain multiple sub-documents for public and private profile information.  This allows us to separate what can be publicly accessible to other members from what is privately available to the member themselves.

The member name, email, and password hash will be stored at top level, while the public and privately accessible profile information will be in separate containers.  One such example is a history of display names and timestamps for the member (allowing other members to verify older messages).

We will also want to store information such as the


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

## closely related systems

Decoupling these systems may be difficult because they will likely rely on eachothers data:

- karma
- statistics
- private conversations
- uploads/media
- blogs


## model(s)

The model will look like this:

    {
        "name": "name@mail.com",
        "password": "hash",
        "profile": {
            "name": "display name",
            "names": [],
            "email": "public@mail.com"
        },
        "private": {},
        "oauth": [],
    }

Passwords will be stored as hashes, and logins will accept both the parent name or display name.

The `names` array is a historical record with names and timestamps when a name was adopted.

The private object is a copy of the profile object with any information the user did not wish to make public.

_The profile object itself has not been fully architected and will likely grow organically._

Open authentication is not yet planned out, so it's an empty array.
