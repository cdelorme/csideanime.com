
# independent logic

Some of ours systems should be independent of the core, but coupled intelligently.

- karma handles adding or removing karma points to users under various conditions
- stats handles custom logic for any stats related information, such as producing lists of highest and lowest counts, as well as handling incrementing of the values

While it may be tempting to implement an event-system into the code, this tends to lead to confusion during maintenance, so I'd rather have these as separate components that are used by the main-line code.

