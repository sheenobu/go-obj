# NOTE: this is a WIP copy vendored for obj-renderer

# go-gamekit

go-gamekit is a simple game utility library built on top of SDL2.

The design of go-gamekit follows a reactive approach. Events are handled via
channels that are wrapped as subscribable data types:

```
mouse.Position.Subscribe()                  // Mouse position changes
button.Clicked.Subscribe()                  // UI button click changes (clicked/unclicked)
window.Size.Subscribe()                     // Window Resizing
windowManager.CurrentWindowID.Subscribe()   // Focused window changes
```

We can use `Get` to get the values synchronously:

```
windowID := windowManager.CurrentWindowID.Get()
```

In the example below, `mouse.Position` has a Subscribe method that has a channel. This
channel emits an `Int32Pair`, where we've decided that L is X and R is Y. Once the message comes in, we
move the `MyObject` to follow:


```
var o *MyObject

sub := mouse.Position.Subscribe()
defer sub.Close()

for {
	select {
		case coords := <-sub.C:
			o.Move(coords.L, coords.R)
		case <-ctx.Done():
			return
	}
}
```

This general structure has lots of power. We can define our game objects as a series of goroutines which
listen on channels and update their internal state approprietely.

We can also make pieces of code very generic:

```
// Moveable is an interface for an entity which can be moved around the screen
type Moveable interface {

	// Move moves the entity to the given x and y position
	Move(x, y int32)
}

// Follow moves the given Moveable everytime the subscribed coordinates change
func Follow(o Moveable, coords *pair.RxInt32PairSubscriber ) {
	for pos := range coords.C {
		o.Move(pos.L, pos.R)
	}
}
```

In this code, we do not care about the concrete type of `Moveable`;
nor do we care if `coords` represents our mouse X and Y positions,
values coming from a game server for a sprites location, or anything else imaginable.

## Features

The following features are done or planned:

 * Multi-Window Management
 * Reactive Data Types
 * TODO: Scriptable console (needs migrating to new codebase)
 * TODO: Scripting support via Javascript/otto
 * TODO: Multiple game loop implementations
 * TODO: High level graphics API (maybe)
 * TODO: Networking

## See Also

 * [rxgen](https://github.com/sheenobu/rxgen)
 * [go-sdl2](https://github.com/veandco/go-sdl2)
