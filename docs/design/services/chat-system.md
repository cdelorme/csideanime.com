
# chat system

This is a very popular component, and I would like to keep it going forward.  However, I would also like to trim down the complication around it to just:

- name
- message
- deleted

_Deciding whether to use the `_id` to parse the timestamp, or generate and store it as duplicate data._  The `deleted` state would allow us to choose not to display specific messages.  Hiding/removing messages that have already been displayed **should** be a goal here, since that often tends to be a concern.

Integrating custom foreground and background colors and the ability to turn them on/off as a user-setting are later features to consider.


## data design considerations

One consideration is whether to store the name or the `_id` of the member.  If we store the `_id` we need to hit two documents to gather the information, but if we store the name then we would have to update all the chat records if/when they change their display name.

My solution will likely be to ignore display name changes, and store both the `_id` along with the display name for record-keeping.  We then have the option of running a batch update to cleanup the records (eventually consistent).
