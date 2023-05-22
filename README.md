# Buckle

Toy project.

## Why tagging

Version dependency without being dependent on the latest semantic (commit) version:
- semantic version: `v0.0.0-00010101000000-000000000000` (initial standard)
- tag: `v0.0.0`

## How to tag

[Git documentation](https://git-scm.com/docs/git-tag)

```bash
export TAG="v0.0.0"
export MSG="some message"
git tag -s -a $$TAG -m "$$MSG" # Create signed annotated tag with message
git push origin refs/tags/$$TAG # Push to remote
```
Tags can be deleted, too:
```bash
git tag --delete $$TAG
git push origin --delete refs/tags/$$TAG
```
