
# external api documentation

Welcome, and thank you for considering our API for your needs.  Our goal is to describe our API so you can use it without pulling your hair out.  If we are failing to do so, let us know by opening a ticket!

API Route: `https://www.csideanime.com/api/v1/`
Alternative API Route: `https://api.csideanime.com/v1/`

_We don't yet have a class 2 ssl certificate so we do not have a sub-domain at the moment._

**The following information is part of our first-release, and is a tentative outline subject to change.**  To complete our API specifications document we are still working on our internal content.

### standard api formats

<table>
    <tr>
        <th>HTTP Method</th>
        <th>Description</th>
        <th>Example</th>
    </tr>
    <tr>
        <td>GET</td>
        <td>return _all_</td>
        <td>https://www.csideanime.com/api/v1/posts</td>
    </tr>
    <tr>
        <td>GET</td>
        <td>return _one_</td>
        <td>https://www.csideanime.com/api/v1/post/#</td>
    </tr>
    <tr>
        <td>PUT</td>
        <td>create a new record</td>
        <td>https://www.csideanime.com/api/v1/post</td>
    </tr>
    <tr>
        <td>POST</td>
        <td>**replace** an existing record</td>
        <td>https://www.csideanime.com/api/v1/post/#</td>
    </tr>
    <tr>
        <td>PATCH</td>
        <td>**modify** an existing record</td>
        <td>https://www.csideanime.com/api/v1/post/#</td>
    </tr>
    <tr>
        <td>DELETE</td>
        <td>delete an existing record</td>
        <td>https://www.csideanime.com/api/v1/post/#</td>
    </tr>
</table>

**Notes:**

- all endpoints will be compatible for both singular and plural dictation
- endpoints allow/ignore suffixed slashes
- some service routes support and mayexpect clean-url parameters:
    - access a thread: `https://www.csideanime.com/api/v1/posts/thread/#`
    - get paginated posts by thread: `https://www.csideanime.com/api/v1/posts/thread/#/index/1/count/25`
    - posts by member: `https://www.csideanime.com/api/v1/posts/member/#`
    - posts by a member in a thread: `https://www.csideanime.com/api/v1/posts/thread/#/member/#`

API Services:

- [posts](#)
- [threads](#)
- [tags](#)
- [members](#)

