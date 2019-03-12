
Always use plural syntax for resources in routes.

This prevents complexity and better represents the fact that you are dealing with a collection.


---

I need to reconsider the resources that actually matter, and their names.


---

I would like to do away with the concept of actions and groups, and instead simply have "permissions", which are strings that represent one or more actions available in the system.  This simple relationship means that we can avoid the complexity of group based association but still benefit from multiple possible actions, all with reduced conditional processing (eg. less overall checks in a loop when there are fewer permissions than actual actions).

This also makes it dead simple to add to a JWT token claims.  I should absolutely create an extended claims model for use both parsing and generated the JWT tokens used by my system.

It may even make sense to share permissions with feature toggles, since both determine whether an operation in code may proceed.


---

Switch naming of member to account, as they represent inanimate data not an actual individual of any sort.


---

Is there a way we can associate multiple string values with the same "tag"?  That is roughly as difficult as choosing the default value...  This is valuable if someone uses the romaji or full japanese characters representing the title of an anime with a different english title.




---

I want to reverse direction a bit, and switch to server side rendered GET routes, while all other operations will be API driven.

This enables strong SEO support, while retaining high performance and powerfully dynamic functionality.

We can use javascript to override links with single-page-application behavior, which also means users with browsers that disable javascript will still work for browsing the site, but not posting contents.

This reduces first-load latency as well, since contents will be rendered up front.

We may need to design a method of sharing the HTML templates so we don't have two different ways of rendering content.

_In the future, we might be able to fully replace our javascript code with WASM compiled from pure go assuming the loader isn't too massive._


This time I should focus on simplifying the implementation of a transparent authorization system, which I can quickly swap out with mocks as needed to expedite testing before focusing too much on oAuth implementation.

Obviously I want oauth support, so I will eventually need to add support for it.  I also want my own account system so users do not need to use third party authentication, and I have support to extend as needed.

I was interested in a tag cloud style system for labeling resources, and popularity of tags by number of references to determine what items make the top list of forums.  This is supplemental to the most active and most recent sections of the forums, all on the home page, but also existing as their own paginated search options.

We might actually just want `/threads/popular` and `/threads/latest`.  Getting these to work is an exact science rather than a complex algorithm.

_We could use a custom algorithm to study a users behavior and assemble their own dashboard of relevant threads by most commonly visited tags in threads._

The tags themselves could be expanded into full resources, for example if the tag represents a band, a game, an anime, a television show.  We could use that information to capture additional information, such as website links including official websites, myanimelist, wikipedia, gamefaqs, and more.


---


A button to click for more content rather than auto-loading on scroll.

Originally I had considered an organic structure that allows multiple paths to locate the title of a thread, which uses the various tags and methods of finding those tags.

The only parent tags would be those that are not also child tags.  Thus we would need to keep track of the total number of children and parents of each tag.

Setting child and parent tags would be done through the tag manager, rather than through.

All threads would be listed underneath every single tag in this structure, which also means finding it may result in varying depths based on parent and child values.

When listing the organic system, we would probably want to mix threads and tags, which means storing most recent modification date on tags as well (anytime a tag is added to a thread for example).

This organic structure would easily allow people to find content in the way that they themselves seem to find appropriate when labeling.


---

A fun and deceptively simple way to add security layers into calls, is to utilize functions as parameters, since they are first class citizens.

When I register routes, I can create a set of common layered security checks to standardize access.

The optimal solution may be to have roughly two functions that encapsulate access logic.

Accepting the intended function to call, and then a list of named permissions that will be accepted.

The function handles the ACL behavior, as well as error responses in a standardized way.

The reason we might have two of them is to differentiate between API requests and server side rendered content.

The end result might look like this:

	router.Handle(http.MethodGet, "/", rootHandler)
	router.Handle(http.MethodGet, "/latest", securedHTTP(latestHTTP, "admin", "user"))
	router.Handle(http.MethodGet, "/api/latest", securedAPI(latestAPI, "admin", "user"))

_The deception is that adding restricted access is a trivial change at the routing registration._

When the person accessing does not have valid permissions we can appropriately respond or redirect depending on the operation.



---



---

## options

Add root options response with 200 status, and a header "Allow" with all methods comma delineated.

	Allow: GET,POST,PUT,PATCH,DELETE,OPTIONS

For the API it may make sense to include a response body that describes the model; basically a self-documenting API.

Unfortunately, when it comes to CORS, things get complicated and shitty very fast.

In order to accept requests with `Authorization` or similar headers, you must set `Access-Control-Allow-Credentials` to `true`.

When you set `Access-Control-Allow-Credentials` to `true` you are no longer allowed to set `Access-Control-Allow-Origin` to `*`.

**Basically, if you have content that requires permissions, you cannot dynamically allow all callers, and pouring more fuel on the flames the protocol matters, so you have to set HTTP or HTTPS appropriately.**

The only good news is that you can _mostly_ automate setting `Access-Control-Allow-Headers` and `Access-Control-Allow-Methods` using the requests `Access-Control-Request-Headers` and `Access-Control-Request-Method` headers.

_It may be possible to capture the `Origin` header to set a value, or to assume `localhost`._. When specifying an actual origin rather than a wildcard, the `Vary` header is expected to be set to `Origin`?

I may try creating a dynamic catch-all CORS compatible automated OPTIONS reply and seeing if that breaks at any point...


Keep in mind that while this is awesome for being dynamic, it may prevent model responses for self-documenting APIs:

	router.Handle(http.MethodOptions, "/", DefaultOptions)
	router.Handle(http.MethodOptions, "/api*all", CORSOptions)


	func DefaultOptions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Allow", "GET,PUT,POST,PATCH,DELETE,OPTIONS")
		w.WriteHeader(http.StatusOK)
	}

	func CORSOptions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Allow", "GET,PUT,POST,PATCH,DELETE,OPTIONS")
		w.Header().Set("Access-Control-Allow-Methods", "GET,PUT,POST,PATCH,DELETE,OPTIONS")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", r.Header.Get("Access-Control-Request-Headers"))
		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		w.Header().Set("Vary", "Origin")
		w.WriteHeader(http.StatusOK)

		// @todo: parse URI to determine model
		// this would enable a dynamic self-documenting API
	}


---

Regarding login....  How dod that go again?

Rate limits, tab order, email not registered, password is common and should not be used (eg. 10,000 most common).

The sort of stuff that sounds like it should be standard.

Support for at least google oauth, and probably no others.


---

Implementation of a per-user system to track relative interests and status.

This may be tied to the extended tagging system, assuming we are able to craft it as in depth as we desire.  Or perhaps the same tags are used to link to these alternative resources?

Anime, TV Shows, Games.  The ability to track ones status somehow.

Possibly, in the future, exposing an API to update those details such that third party systems can be hooked in.


---

Regarding private messages, we may need some kind of thread system that restricts who can access.

This sort of system may also allow multi-user private conversations...

We could also choose to forgo private messages in general.


---

No longer use nginx for this project.

All assets will be served from memory using statically compiled binary.

This alone eliminates any reasons to use nginx ontop of go for server purposes.
