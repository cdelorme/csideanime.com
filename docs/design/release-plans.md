
# release plans

Our goal is to reach an alpha release checkpoint swiftly.  By sticking to a bare-bones design we avoid unnecessary time spent on small details.

We can then identify the more critical areas of concern for a beta release.  **The beta release is likely when we will migrate from `seaofnumbers.com` to the new site.**

Our deadline for the beta release is early March 2015, which does not give us a great deal of time.  If we can hit our Alpha release by end of January it may be possible.


## alpha release

The tough part for this section is going to be identifying critical services as opposed to "desired" services.

Here is a known list:

- open-authentication (oauth /w google & facebook)



## beta release

The beta release will focus on bug-fixes from the alpha-release, and a plan to migrate data from our existing site.  The migration is expected to be the harder part here.  If we migrate posts and threads we also need to migrate user accounts, but the login schema being so different will make that very difficult.


## future releases

After the beta release, we will focus on all the missing, incomplete, or broken features.  This will likely be a fast paced process, and involve quite a bit of innovation.  Ideally our alpha->beta foundation will leave us with a sturdy design that won't end up with dozens of API breaking changes and semantic version increments as a result.
