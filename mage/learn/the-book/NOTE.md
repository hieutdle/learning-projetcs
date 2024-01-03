## Immutable versus `const`

`const` is not for variables; it's for constant values which may not be stored anywhere; they're effectively an alias for a literal value.

Constants can not be **shadowed**.

`let` defines **scoped** function variable bindings. Not global scope like `const`.

## `match` versus `if/else`

Clarity and destructuring and makes exhaustiveness (examining every possible path) more obvious.

```rust
let light_value = match traffic_light { // I only have to read this line to know what we are branching on
    Red => 1,
    Yellow => 2,
    Green => 3,
    // _ => panic!("We don't need this line because Rust will check to make sure all cases is used")
};

let light_value2 = if traffic_light == Red { // I have to read way further down to know the logic of this if conditional
    1
} else if traffic_light == Yellow {
    2
} else { // I have to read up to here to know if the case statement is only about traffic_light. Also, did it handle all cases? What if TrafficLight added a new state?
    3
};
/*
} else if traffic_light == Green { // we can't just use this to assert exhaustiveness, as then we break the expression type check.
    3
} else {
    panic!("We can have a catch all else at the end, but then it means it is a runtime error")
}
*/
```


## `let mut` versus **shadowing**

**Shadowing** can change the data type, `let mut` cannot

**Shadowing** only applies to the innermost scope. **Mutable variables** can allow you to propagate data from a scope to the outside.

**Shadowing** a variable does not shorten the lifespan of the variable you shadowed, so it's less memory efficient, and forces you to allocate new memory rather than reusing what you already allocated.

`mut` allows you to do stuff like passing a variable into a function and modifying it in place or reusing a scratch buffer across multiple iterations of a loop.

**Borrowing a reference:**

```rust
let a = 5;
println!("a = {}", a);
let c = &a;
let a = 10;
println!("a = {}", a);
println!("c = {}", c);

a = 5
a = 10
c = 5
```

So the original value of `a` remains on the stack, and the second `let` binding is allocating a new spot on the stack and pointing `a` to it. This is safe because it's not mutating the value of `a`, it's changing what `a` points to. The original value of `a` won't be dropped until the end of the block, so the lifetime is valid.

If you try to do this with `b`, however: 

```rust
let mut b = 5;
println!("b = {}", b);
let mut d = &b;
b = 10;
println!("b = {}", b);
println!("d = {}", d);

4 |     let mut d = &b;
|                 -- `b` is borrowed here
5 |     b = 10;
|     ^^^^^^ `b` is assigned to here but it was already borrowed
6 |     println!("b = {}", b);
7 |     println!("d = {}", d);
|                        - borrow later used here
```

## `i8` versus `u8`

`i8` : -128 to 127

`u8` : 0 to 255

