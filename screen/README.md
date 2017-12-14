# Screen

This package implements a screen "object" in an event-driven
manner. You can manipulate the methods on this object in order to
mutate its internal state or access it. However, any mutations would
also emit related change events. These events, if replayed by the
subscriber, could be used to rebuild the screen state and keep it in
sync.

Usage:

screen.New()


// all methods generate a change event and apply it
// they read state in order to make decisions
//
