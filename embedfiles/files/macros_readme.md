There are some macros you can use in a card.

Macro is an expression that EdoEx will interpret and replace with the result of that macro, so you don't need to type the full text.

Macros are applied in the following fields of a card:

- `description`
- `name`
- `pendulum_effect`
- `strings`

The following macros are available for now (more will be added with time)

| Name            | Type of parameter |
| --------------- | ----------------- |
| [`card`](#card) | Card id           |
| [`meta`](#meta) | Meta id           |

To use a macro just put the macro name and the parameters for that macro as following:

```
${ macro name : param1 : param2 : ... }
```

All macros support multiple parameters

# Card

Use the `card` macro the mention a card's name by it's id. This is specially useful to maintain consistency of a card's name across the entire expansion.

## Example

If you have a card with id `123` and name `Example name` you can mention it on any card by using `${card:123}` and it will be replaced with `Example name`. The effect

```
Add from your deck to your hand 1 "${card:123}".
```

turns into

```
Add from your deck to your hand 1 "Example name".
```

You can also use `card` to mention the current card by using the `self` param. The effect

```
Draw 1 card. You can only use this effect of "${card:self}" once per turn.
```

turns into

```
Draw 1 card. You can only use this effect of "Example name" once per turn.
```

If an invalid id is provided, `card` will just use the parameter itself. So

```
Send 1 "${card : some text}" from your deck to your Graveyard.
```

turns into

```
Send 1 "some text" from your deck to your Graveyard.
```

If no id is provided, `card` will put nothing.

When using multiple params, `card` concatenates the result with a comma. So

```
Destroy all "${ card : 123 : Some card : Cool card }" you control.
```

turns into

```
Destroy all "Example name, Some card, Cool card" you control.
```

Probably not that useful.


# Meta

Basically the same as [`card`](#card), but it mentions a meta (See the README on the `meta` folder).

The only differences being `meta` does not have a `self` parameter and accepts a meta's `id` (even on hexadecimal) or `alias`.
