
# member management

Member management would be a service layer, which connects to the member dao.  Models include:

- member
- member profile

_The member profile contains public fields, while the member object will contain the email(s), and password hash.  By offloading the display name to the member profile object, we can restrict what gets shown to others._

It would need to be able to:

- create a new member (eg. registration)
- delete a member (which actually disables)
- add group to member
- remove group from member
- get permissions by member
- get profile
- update/patch profile
- get profiles

Still not sure about open authentication integration, but it would likely be closely tied to the member service layer.

_We should **never** delete data from our system, we should only ever disable them._  Design considerations for this need to be noted.

Other connected components we need to consider:

- statistics
- private conversations
- uploads/media
- blogs


## data design considerations

If we organize the data in the right ways we can probably write some slick simplified code to handle separation of member from member profile logic.  My guess is that member profiles will be a sub-document.
