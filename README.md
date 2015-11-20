# bijou

A functional programming language, putting the emphasis on programming with
values.

## Project overview

Bijou is an experimental language, with a few key design goals.

  1. Fun to develop in
  2. It must be functional at its core
  3. The emphasis should be on values, not syntax
    - almost everything should be passable as a value at runtime
  4. The runtime should be a small, single file distribution
  5. Immutable throughout
  6. Concurrency should be
    - robust
    - easy to understand
    - lightweight
  7. Pattern matching should be used for expressiveness
  8. Abstractions should provide a great deal of consistency
  9. The stdlib (bundled into a single file) should be rich
  10. Modules are used to group functions together
    - import automatically constructs modules
    - modules can be dynamically created at runtime
    - provide a mechanism for object-oriented programming

Currently the language is a work in progress, but it is already in a state that
allows some experimentation.

## Language features

### Data types

The following data concrete types are currently implemented:

  1. 32-bit integers
  2. Symbols (bare keywords for variable lookup etc)
  3. Strings
  4. Linked lists
  5. nil
    - acts like a true null object
    - accepts invocations, returns nil
    - is sequenceable, is empty
  6. Booleans
  7. Functions
  8. Macros

The following abstract data types are also implemented:

  1. Sequences
    - Lists
    - Strings
    - Cons lists
    - Nil
  2. Callable
    - Functions
    - nil
    - Eventually other data types
  3. Ports
    - these support for I/O type operations
    - read/write/close

The following types will exist in future:

  1. Floating point numbers
  2. Rational numbers
  3. Aribtrary length integers (automatic)
  4. Keywords (interned strings)
  5. Regular expressions (will be Callable)
  6. Modules (e.g. `math/sin` points to `sin` in the `math` module)
  7. Vectors
  8. Maps

### Syntax

Bijou is a Lisp family language. The syntax is based on atomic values, and
data stuctures that group these values. Most literal values are implemented as
in any other language:

#### Literals

Integers.

``` lisp
42
-42
```

Strings (always double-quoted).

``` lisp
"This is a \"string\"\n"
```

The nil value.

``` lisp
nil
```

True and false.

``` lisp
true
false
```

Some data types must be quoted if given in literal form. We'll see why shortly.

Lists.

``` lisp
'(1 2 3 4) ; same as (list 1 2 3 4)
```

Symbols.

``` lisp
'some-symbol
```

Quoting values that do not need to be quoted has no effect.

``` lisp
'42
'true
'nil
'"some string"
```

#### Defining symbols

We can give values a name, by associating them with a symbol. We use the `def`
form to do this.

``` lisp
(def answer 42)
```

Once something has been defined, we can refer to it by using the symbol
unquoted.

``` lisp
answer
```

We can provide documentation for our definitions.

``` lisp
(def answer
  "The answer to life, the universe and everything"
  42)
```

Currently the docstring is simply ignored by the runtime. In the future it will
be accessible by at runtime.

#### Calling functions

We call a function by using the list syntax, unqoted. The first value in the
list is the function we want to call. The rest of the values are the arguments.

``` lisp
(* 2 3) ; 6
```

Here we call the function `*`, which is used for multiplication, as you might
expect, with the arguments `2` and `3`. It returns 6.

``` lisp
(* 2 (+ 4 7))
```

What do you think that might return?

#### Creating functions

Here's where that design goal of focusing on values comes into play. Functions
are values. We create them with the `fn` form.

``` lisp
(fn (x y) (* x y))
```

Let's break that down. The `fn` creates our function for us. The `(x y)` is the
parameter list. They must be symbols, as in the `def` form. The remaining forms
represent the function body. There are no `return` keywords. The last evaluated
expression is always the return value. This is true for all forms in bijou.

If functions are just values, we can give them a name with `def`.

``` lisp
(def squared
  "Return the square of n"
  (fn (n)
    (* n n)))

(squared 4) ; 16
```

#### Grouping expressions together

The `do` form allows us to evaluate multiple expressions in a single form.

``` lisp
(do (this)
    (that)
    (more))
```

It returns the last evaluated expression, so in this case, whatever `(more)`
returns.

#### Conditions

Conditions are implemented with `if`.

``` lisp
(if (true?)
  (this)
  (that)
  (more))
```

The `if` form takes at a minimum, two arguments. The first, `(true?)`, as
you'd expect is the condition to evaluate. The second, `(this)` is the
expression to evaluate on success. The remaining forms `(that) (more)`
are evaluated in an implicit `do` as the else case.

Without any forms in the else position, `if` returns `nil` on false.

``` lisp
(if false
  (this)) ; nil
```

If you want to evaluate multiple forms in the success position, use `do`.

``` lisp
(if (true?)
  (do (this)
      (that)
      (more)))
```

#### Looping

There are no loop constructs in bijou. Instead, we use recursion to loop. Say
we want to implement the function `exponent`, such as `(exponent 2 3)`, giving
two to the power of three, in imperative style languages, like C, we might
write.

``` c
int exponent(int n, int e) {
  int m = 1;
  while (e-- > 0) {
    m = m * n;
  }

  return m;
}
```

In bijou we'd write this recursively:

``` lisp
(def exponent
  "Return n to the power of e"
  (fn (n e)
    (if (= e 0)
      1
      (* (exponent n (- e 1)) n))))
```

#### Tail Call Optimization

Recursive functions are a key element in functional programming. Bijou
implements proper tail call optimization, so we can recurse infinitely without
blowing the stack, provided we only use recursion in a "tail position". This
translates to the last evaluated expression being a call to the function
itself, or a call to another function, which may in turn mutually call the
current function.

Say we want to output all the integers, between `n` and `m`.

``` lisp
(def print-range
  (fn (n m)
    (println n)
    (if (< n m)
      (print-range (+ n 1) m))))

(print-range 0 100000)
```

That's 100,000 recursive calls, and that's fine.

#### Concurrency

Bijou puts concurrency as a core language feature. Wherever you want
concurrency, you use `spawn`, which starts running a function in a separate
lightweight thread. Spawn returns a `port`, which we'll look at later. The same
`port` is passed as the first argument to the function running in parallel.

``` lisp
(spawn (fn (p) (println "This happens concurrently")))
```

We can give arguments to `spawn`, which are passed to the function it calls.

``` lisp
(spawn (fn (p who)
         (println "Hello, " who "!"))
       "World")
```

Now it's time to talk about ports. Ports are values that we can read from with
`<-` and write to with `->`. When we open a file, or a network socket, we get
a port, which we can read and write to. When we use `spawn`, we also get a
port. This is how we communicate with the spawned thread. We can send and
receive any kind of value over the port: integers, strings, symbols, lists,
even functions or other ports.

Say we want to implement a score counter, which we can increment and decrement
with arbitrary values:

``` lisp
(def counter
  (fn (p n)
    (println n)
    (counter p (+ n (<- p)))))

(def score
  (spawn counter 0))

(-> score 15) ; 15
(-> score 10) ; 25
(-> score -2) ; 23
```

You might be beginning to see how by using recursion, we don't need mutability.

What's happening here? First we define our `counter` function, which accepts
a port and a number as arguments.

Next we spawn that counter function so it runs in the background. It starts by
printing `n`, then recursing with the value of `(+ n (<- p))`. What is this?

`(<- p)` takes a value from the port. We are only going to send integers here,
so it will always add some integer to `n`.

If there are no values to read on the port, `(<- p)` will block, with no
runtime cost. Since it's happening concurrently, this does not block the
program execution.

Eventually, once a value is a available on the port, the recursion happens and
the new value of `n` is printed.

We adjust our score by a few numbers by using `(-> score 15)`, for example.

Like with `<-`, if the receiving end of the port is busy, `->` will block.
