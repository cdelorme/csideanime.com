
# definitions

_This page contains works we will use, with a basic definition and example to describe them in the context of this project._

- `controller-level-conditions`
    - basic conditions specific to routing
    - **ex.** no id provided to a singular path (generally invalid params = 400 response)
- `first-pass-acl`
    - route based access control
    - **ex.** guests cannot delete posts, unauthenticated users should be turned away at the door
- `second-pass-acl`
    - action based access control
    - **ex.** an admin can delete other members posts
- `third-pass-acl`
    - data based access control
    - **ex.** a member can delete their own posts
- `application-level-security-config`
    - server level configuration for the application, not available through the admin

