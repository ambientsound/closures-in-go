# Closures in Go

## What are closures?

- Wikipedia mentions: "a closure is a technique for implementing lexically scoped name binding in a language with first-class functions."

## What does that mean?

- We may write functions with access to the current scope, put the function in a variable, and pass it to a different part of the application that normally doesn't have access.

## Why bother?

- Create methods that have asynchronous operation.
- Send local objects to remote places. Examples are e.g. database handles or logging instances.

## Repository examples are in pkg/

## Real world examples

- [Naiserator](https://github.com/nais/naiserator) is a Kubernetes operator written by NAV.
It manages Kubernetes resources using a _create, update or delete_ mechanism implemented with closures.
