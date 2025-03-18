/*


Panic
As we've seen, the proper way to handle errors in Go is to make use of the error interface. Pass errors up the call stack, treating them as normal values:

func enrichUser(userID string) (User, error) {
    user, err := getUser(userID)
    if err != nil {
        // fmt.Errorf is GOATed: it wraps an error with additional context
        return User{}, fmt.Errorf("failed to get user: %w", err)
    }
    return user, nil
}

However, there is another way to deal with errors in Go: the panic function. When a function calls panic, the program crashes and prints a stack trace.

As a general rule, do not use panic!

The panic function yeets control out of the current function and up the call stack until it reaches a function that defers a recover. If no function calls recover, the goroutine (often the entire program) crashes.

func enrichUser(userID string) User {
    user, err := getUser(userID)
    if err != nil {
        panic(err)
    }
    return user
}

func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("recovered from panic:", r)
        }
    }()

    // this panics, but the defer/recover block catches it
    // a truly astonishingly bad way to handle errors
    enrichUser("123")
}

Sometimes new Go developers look at panic/recover, and think, "This is like try/catch! I like this!" Don't be like them.

I use error values for all "normal" error handling, and if I have a truly unrecoverable error, I use log.Fatal to print a message and exit the program.

*/



/*

Q.Panic and recover should be used instead of errors...

when your fingers are tired and you don't care anymore
when you want to make your code more readable
when you want more control over the error handling
basically never = answer


Q.If you want your program to cleanly exit in an unrecoverable way, which is a good alternative to panic?

log.Fatal() = answer
kill()
fmt.Println()
boos.Exit()