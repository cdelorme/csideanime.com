
# data architecture

Access control should be tiered and managed as groups and actions.

- Our API should remain as simple as possible, complexity adds potential for security holes and a bad time
- ACL (Access Control Layer) interfaces with simple lists of groups and actions
- Groups are not hierarchical and do not inherit
- Members have one or more groups (the default for any visitor being Guest)

**Independent Logic Systems:**

- Karma - handles adding or removing karma points to users under various conditions
- Stats - handles custom logic for any stats related information, such as producing lists of highest/lowest counts

**Decisions:**

- I think our tagging system should allow freehand text entry
- I think adding tags to our auto-complete list could be tied to a karma point cost
- once we have a complete list of actions, we should decide which should be made accessible via karma points

No data should ever actually be deleted, instead we mark the data as deleted, and anyone with privileges who turns on admin mode will see “deleted” records.

We have to identify what we want to exist as a forum setting. Simpler is better. If we integrate new features, they should maybe have their own pages? Our first release may not have any specific forum settings.

The ability to ban users is important. Without it bots could overrun our community. We should discuss ways to efficiently handle that situation. Right now I prevent new accounts from making posts without capta and they cannot make a new thread until they have at least 5 posts. We could potentially use karma points for creating threads, which would be a built-in way of handling the same scenario without adding a bunch of custom logic.

We should have a group called "Developer" with the same permissions as an administrator, but to help differentiate (since these are seen by users).

- Member struct /w OAuth support
- OAuth setup & connection
- Member Registration
- Member Login
- Permissions via ACL (a list of actions per group)
- Tags List Module /w return-all for type­ahead client­side
- Thread & Post Structs
- Chat Structs


### struct design

We will need the following structs:

- message (chat) - member
- settings
- profile - thread
- post

We can avoid complexity by using a basic service layer and simple lists for this data:

- settings: Settings belonging to the forum, and accessible to moderators.
- groups: Members belong to one or more groups
- actions: Permissive by name belonging to one or more groups
- tags: Simple list of known tags for auto-completion when searching or tagging

Things like `stats` and `Karma` will be integrated into members, threads, and posts.

I am still working out the details and necessity of the following data, and they may not be in our first release due to the added complexity behind them:

- private conversations or threads: require separate structs, and access logic
- scanlation/subbing groups /w websites, works, & status of works
- friend system / follower system: we may not need? Wait for members to request it?


### Permissions

We need to decide on a set of groups. The smaller and simpler the list the better.  We need to determine the available operations based on the set of defined structs and their access layers, then create a feasible permission schema, such as who can access. This includes determining which level of ACL control to implement.

Groups:

- developer - admin
- guest
- member

Actions:

- login


### mongodb schema design

This means we have to define a set of actions, then a set of groups who are able to perform said actions. These should range from Guest to Admin, but ideally we should be able to get by with a short list. Making each group its own document and storing all actions in that document is one approach. Another is storing all groups and actions inside one document. Both are reasonable, until you have a sizable number of actions and groups. Also worth noting, if a document is going to grow by 20% or more it will have to reallocate space for the document, which is inefficient and creates fragmentation that later requires compaction. So it would be wise to pick one that is as future­proofed as possible against that when considering single­documents. If you miss a group or add one in the future that could be a concern. Another question is how to assign groups. Do members have multiple groups, or just one. My thought is they should have multiple groups. There should be no “inheritance” between groups, simply a check that looks at each group available to a member.

With mongodb schema is more important than it would be with SQL, because the integrity has to be enforced exclusively by the code.

Before we can document the API specifications we also need to know the public face of these structs. It would be smarter to create the internal appearance, then map them externally.

Some data needs to remain private. We may want to send a members profile to other users, but not their credentials and username/email. In that case we should create an endpoint that allows for those bits to be omitted. This may be either two separate structs and endpoints, or it could be a single struct with two endpoints.

Individual struct data in golang can be made hidden when converted to json using marshall syntax.

All endpoints will need a service layer to manage any complexity separate from the controller level.

We may consider creating data access layers that handle acquisition, modification of, and saving of data separate from the service layer.

