# Fundstelle000

A tool for my dad.
Do not use this. It will hurt you.
I might learn some golang here...

## Installation

```[bash]
go install -ldflags="-s -w"
```

## Usage

```[bash]
cd myStuff

# this will list what would happen
fundstelle000

# this will actually make it happen
fundstelle000 -w
```

## What?

This will search for files with two uppercase alphas followed by one or two
digits and pad them with zeroes until there are three digits.

## Features

### Exclude stuff by prefix

files with "gps" or "track" as prefix will get skipped

### Prevent multimatching

Only the first occurence of a fundstelle-like thing will get replaced.
