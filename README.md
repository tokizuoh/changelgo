# changelgo

A command-line tool that converts the commit history of specified files, such as `CHANGELOG.md`, from GitHub repositories into an RSS feed.

## How to use

### 1. Fork

Fork this repository.

### 2. Edit `items.yml`

```
items:
  - link: https://github.com/swiftlang/swift/blob/main/CHANGELOG.md
  - link: https://github.com/swiftlang/swift-package-manager/blob/main/CHANGELOG.md
```

GitHub Actions regularly generates the RSS feed under the `./generated` directory.

### 3. Subscribe to RSS feed


```
https://raw.githubusercontent.com/tokizuoh/changelgo/refs/heads/main/generated/swiftlang-swift-rss.xml
```

Register the raw file of the generated content as an RSS feed, using something like Feedly.
