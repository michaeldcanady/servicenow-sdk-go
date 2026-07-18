# Activity Subscriptions API overview

The Activity Subscriptions API (`actsubapi`) manages activity-stream
subscriptions: which objects a user follows, their notification preferences,
and their activity feed.

## Basic usage

Access it through the `Now()` namespace via `ActSub()`:

```go
actSub := client.Now().ActSub()

// List activities
activities, err := actSub.Activities().Get(context.Background(), nil)

// Check whether the current user is subscribed to an object
subscribed, err := actSub.Subscriptions().
    ByObjectId("{objectID}").
    IsSubscribed().
    Get(context.Background(), nil)

// Subscribe to an object
_, err = actSub.Subscriptions().
    ByObjectId("{objectID}").
    Subscribe().
    Post(context.Background(), body, nil)
```

## Available operations

- **Activities** — `Activities().Get` lists activity records.
- **Contexts** — `Contexts().Get` lists activity contexts.
- **Facets** — `Facets().ByContext(context).ByInstance(instance).Get` lists facets for a context instance.
- **Followings** — `Followings().ByFollower(follower).Get` lists what a user follows.
- **Preferences** — `Preferences().Post` creates a preference; `Preferences().ByProfileId(profileID).Get` retrieves one.
- **Subscribed objects** — `SubObjects().Get` lists subscribable objects; `Subscribers().BySubObject(subObject).Get` lists subscribers.
- **Subscriptions** — `Subscriptions().BySubscriberId(id).Get`, plus `ByObjectId(id)` with `IsSubscribed()`, `Subscribe()` (POST), and `Unsubscribe()` (DELETE).
- **User stream** — `UserStream().ByProfileId(profileID)` with `Get` and `Put`.
