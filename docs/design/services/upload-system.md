
# upload system

Member uploads is a feature I'd like to offer, and if possible without restrictions.

Member content would be identified by a sha256 hash, renamed accordingly with metadata in the system containing the name at upload-time.  The final product would be uploaded to S3 and a CloudFront url would be supplied for access.

With the has identification, it would be easy to prevent duplicate uploads.  Ideally the system should notify them after an upload if a duplicate was found.  It should also run the sha256 and upload-to-s3 operations in the background to reduce time consumed for a member.

The actions this would enable:

- add media (aka upload)
- delete media (disable)
- really, delete media (for administrators and moderators, things we "cannot" host on S3, but could be a problem since paths are shared if two users upload the same content, so we'd have to run a "delete all" by url)
- rename media (metadata name only)

If need be we can implement size and file-type restrictions, but I'd much rather just allow for all types.  We may also implement a "per-user" limit

If possibly I'd like to identify common types and enable the system to intelligently provide click-to-add-to-post type functionality at the client-side.


## data design considerations

We would need each record to have a `name` and a `url`.  Optionally they would have `tags` and `disabled`.
