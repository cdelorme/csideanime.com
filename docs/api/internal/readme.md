
# interal api documentation

**_This section is to provide a full overview of our code from an internal perspective._**

### standard system layers

<table>
    <tr>
        <th>Layer</th>
        <th>Purpose</th>
    </tr>
    <tr>
        <td>Controller</td>
        <td>Handles routed connections, parses supplied arguments and handles controller-level-conditions and first-pass-acl.  Generally depends on two or more services, and responsible for replying with final results to the request object.</td>
    </tr>
    <tr>
        <td>Service</td>
        <td>Handles business logic and second-pass-acl.  Depends on the data layer and may depend directly on some models.</td>
    </tr>
    <tr>
        <td>Data Access</td>
        <td>Depends on models, connets to database, manipulates data, and handles third-pass-acl.</td>
    </tr>
    <tr>
        <td>Model</td>
        <td>Definition of the data (eg. structs), should have no dependencies.</td>
    </tr>
</table>

These are the layers that will exist in all of our system.  It gives us plenty of separation of concerns.

The core is responsible for instantiating and connecting all of these, as well as starting the application server.

We should have a controller for every defined service, and each service exists as an endpoint.  A service can represent whole structs, partial structs, or simpler sets of data.

Our goal is to keep the code on the server as simple as possible.  Complexity only adds potential for security exploits and higher probability of bugs.

