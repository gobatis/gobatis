---
sidebar_position: 1
title: Overview
---

<img src="/img/logo.png" style={{width: 300, display: 'block', margin:'0 auto'}} />

Most database operation scenarios involve simple CRUD (Create, Read, Update, Delete) operations. For these cases,
we need straightforward methods. When conducting complex queries, excessive parameter bindings become a headache; at
these times, a convenient template syntax is required. In some specific scenarios, such as paginated queries, we need
some quick methods.

Taking into account the above requirements, Gobatis was designed. It adheres to the traditional usage habits of Go ORMs,
and also draws from MyBatis's Dynamic SQL syntax. Additionally, it offers numerous other features, making system
development simpler and more efficient.

## Feature

* Simple, An engineering-oriented ORM
* Intuitive and convenient API design
* Targeted at users who prefer using native SQL
* Transaction tracing.
* Mybatis parameter syntax and Dynamic SQL syntax
* Hooks (Before/After, Insert/Update/Delete/Query/Exec)
* More rigorous query result matching mechanism
* Context, Prepared Statement Mode, Debug Mode, DryRun Mode, Loose Mode
* Logger
* Every feature comes with tests
* Developer Friendly
