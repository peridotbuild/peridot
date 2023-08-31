<p align="center">
  <picture>
    <source srcset="ui/mship_gopher.png" media="(prefers-color-scheme: dark)" height="150">
    <img src="ui/mship_gopher_dark.png" alt="Mship" height="150">
  </picture>
</p>
<p align="center">Tool to archive RPM packages and attest to their authenticity</p>
<hr />

### Development

Using the taskrunner2 target is sufficient. `bazel run //tools/mothership`.

This will watch for changes in mship_admin_server, mship_server and both UIs.

The target will also start Dex IDP and Temporal if not already running.
