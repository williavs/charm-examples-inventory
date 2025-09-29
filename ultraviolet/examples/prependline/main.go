package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"time"

	uv "github.com/charmbracelet/ultraviolet"
	"github.com/charmbracelet/ultraviolet/screen"
	"github.com/charmbracelet/x/ansi"
)

func main() {
	// Create a new default terminal that uses [os.Stdin] and [os.Stdout] for
	// I/O.
	t := uv.DefaultTerminal()
	t.SetLogger(log.Default())

	// Set the terminal title to "Hello World".
	t.SetTitle("Hello World")

	w, _, err := t.GetSize()
	if err != nil {
		log.Fatalf("failed to get terminal size: %v", err)
	}

	frameHeight := 1
	t.Resize(w, frameHeight)

	// We want to display the program in alternate screen mode.
	// t.EnterAltScreen()

	// Without starting the program, we cannot display anything on the screen.
	if err := t.Start(); err != nil {
		log.Fatalf("failed to start program: %v", err)
	}

	var st uv.Style
	bg := 1
	st = st.Background(ansi.BasicColor(bg)).Foreground(ansi.Black)
	display := func() {
		// This is the program display. It takes a function that receives a
		// [uv.Frame]. The frame contains the last displayed buffer, cursor
		// position, and the viewport area the program is operating on.
		// Under the hood, the program will call [uv.Screen.Display] to display the
		// frame on the screen and the implementation will depend on the screen
		// type and how it handles displaying frames.
		const hw = "Hello, World!"
		bg := uv.EmptyCell
		bg.Style = st
		screen.Fill(t, &bg)
		for i, r := range hw {
			t.SetCell(i, 0, &uv.Cell{
				Content: string(r),
				Style:   st,
				Width:   1,
			})
		}
		t.Display()
	}

	// Now input is separate from the program. Just like a uv.screen displaying
	// a channel, the channel airs the program to the uv.screen. The program is
	// not aware of the input source.
	// Here, our [uv.Screen] is a terminal and is capable of receiving input
	// events. These events can come from different sources, such as a
	// keyboard, mouse, window resize, etc. Each input source implements the
	// [uv.InputReceiver] interface, which is responsible for receiving input
	// events and sending them to a receiver channel.
	//
	//  ```go
	//  // InputReceiver is an interface for receiving input events from an input source.
	//  type InputReceiver interface {
	//  	// ReceiveEvents read input events and channel them to the given event
	//  	// channel. The listener stops when either the context is done or an error
	//  	// occurs. Caller is responsible for closing the channels.
	//  	ReceiveEvents(ctx context.Context, events chan<- Event) error
	//  }
	//  ```

	// Create a context that we can cancel anytime to stop the event loop.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	display()

	evch := make(chan uv.Event)
	go func() {
		defer close(evch)
		t.StreamEvents(ctx, evch)
	}()

	// Events returns an event channel that receives input events from the all
	// the terminal input sources. It will block until we close the event
	// channel or cancel the context.
	for ev := range evch {
		// Handle events here
		switch ev := ev.(type) {
		case uv.KeyPressEvent:
			switch {
			case ev.MatchStrings("q", "ctrl+c"):
				cancel()
			}

			st = st.Background(ansi.BasicColor(rand.Intn(16)))
		case uv.WindowSizeEvent:
			t.Resize(ev.Width, frameHeight)
			t.Erase()
		}

		t.PrependString(fmt.Sprintf("%T %v", ev, ev))

		rd := rand.Intn(8)
		st.Background(ansi.BasicColor(rd))
		display()
	}

	// Gracefully shutdown the program.
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := t.Shutdown(ctx); err != nil {
		log.Printf("shutdown: %v", err)
	}
}

func init() {
	if os.Getenv("UV_DEBUG") == "" {
		log.SetOutput(io.Discard)
		return
	}
	f, err := os.OpenFile("uv_debug.log", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o644)
	if err == nil {
		log.SetOutput(f)
		log.SetFlags(log.LstdFlags | log.Lshortfile)
	}
}
