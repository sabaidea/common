## 🧭 Semantic Versioning Guide

This project follows **Semantic Versioning (SemVer)**:

```
vMAJOR.MINOR.PATCH
```

| Type     | Example        | When to Bump                          |
|----------|----------------|--------------------------------------|
| Patch    | `v2.0.3 → v2.0.4` | Bug fixes, internal refactoring       |
| Minor    | `v2.0.3 → v2.1.0` | Add features (non-breaking)          |
| Major    | `v2.3.1 → v3.0.0` | Breaking changes in public API       |

---

## 🛠️ Release Process

### Step 1: Check Changes

- Run `git diff` / `git log`
- Decide the type of version bump


### Step 2: Tag & Push

```bash
git add .
git commit -m "Release v2.1.0: added dynamic log formatter"
git push origin <branch>  # e.g. main or develop
git tag v2.1.0
git push origin v2.1.0
```

### Step 3: Create GitHub Release (Optional)

1. Go to https://github.com/<your-org>/<your-repo>/releases
2. Click "Draft a new release"
3. Select `v2.1.0` as the tag
4. Use the same changelog content
5. Click **Publish Release**

---

## 🚀 Optional: Automate Releases

You can automate this with tools like:

- `GitHub Actions`
- `goreleaser`
- `release.sh` script

Let us know if you want to integrate automation workflows.

---

## ✅ Go Module Compatibility

> ✅ Only tagging with `vX.Y.Z` is required for Go to pick up the release

If using a major version `v2+`, make sure your `go.mod` path ends with `/v2`, `/v3`, etc.

```go
module github.com/your-org/your-repo/v2
