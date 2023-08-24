# devdex

devdex is a local instance of Dex IDP that is used for development. User does NOT need to have dex pre-installed. devdex ships with its own configuration that should work for most use cases in this repo.

Only ONE instance of devdex can run at a time, multiple invocations will *silently* fail.

This is done so users can run multiple services locally without having to keep devdex out of the taskrunner2 target list.
That means the second instance of devdex won't start, but first one will keep running and the service that relies on dex will not notice any difference.

### Static clients

* ```
  client_id: local
  client_secret: local
  redirect_uris:
    - http://localhost:9111/auth/oidc/callback
  ```

### Static users

*All static users have the same password: `password`*

* admin@rockylinux.org
* mustafa@rockylinux.org
